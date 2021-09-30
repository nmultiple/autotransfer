package mail

import (
	"io"
	"log"

	"golang.org/x/text/encoding/japanese"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/charset"
	mailer "github.com/emersion/go-message/mail"
)

type Mail struct {
	Subject, Text string
	AttachmentCount int
}

func Fetch(server, user, password string) []Mail {
	var mails []Mail

	log.Println("Connecting to server...")

	// Connect to server
	c, err := client.DialTLS(server, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	// Don't forget to logout
	defer c.Logout()

	// Login
	if err := c.Login(user, password); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")

	// List mailboxes
	done := make(chan error, 1)
	mbox, err := c.Select("Circle", false)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Flags for Circle:", mbox.Flags)

	criteria := imap.NewSearchCriteria()
	criteria.WithoutFlags = []string{imap.SeenFlag}
	uids, err := c.Search(criteria)
	if err != nil {
		log.Println(err)
	}
	seqset := new(imap.SeqSet)
	seqset.AddNum(uids...)
	log.Printf("Search complete, found %d messages", len(uids))

	if len(uids) == 0 {
		log.Println("No unread messages")
	} else {
		messages := make(chan *imap.Message, 10)
		var section imap.BodySectionName
		go func() {
			done <- c.Fetch(seqset, []imap.FetchItem{section.FetchItem()}, messages)
		}()

		charset.RegisterEncoding("iso-2022-jp", japanese.ISO2022JP)

		for {
			msg, more := <-messages
			if !more {
				log.Println("All messages read.")
				break
			}

			if msg == nil {
				log.Fatal("Server didn't returned message")
			}

			r := msg.GetBody(&section)
			if r == nil {
				log.Fatal("Server didn't returned message body")
			}

			// Create a new mail reader
			mr, err := mailer.CreateReader(r)
			if err != nil {
				log.Fatal(err)
			}

			mail := new(Mail)
			if sub, err := mr.Header.Subject(); err == nil {
				mail.Subject = sub
			}

			for {
				part, err := mr.NextPart()
				if err == io.EOF {
					break
				} else if err != nil {
					log.Fatal(err)
				}

				switch part.Header.(type) {
				case *mailer.InlineHeader:
					b, _ := io.ReadAll(part.Body)
					mail.Text = string(b)
				case *mailer.AttachmentHeader:
					mail.AttachmentCount++;
				}
			}

			mails = append(mails, *mail)
		}

		if err := <-done; err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Mail fetched successfully!")

	return mails
}
