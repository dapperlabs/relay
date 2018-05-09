package mailgun

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/mail"

	smtp "github.com/emersion/go-smtp"
	mg "gopkg.in/mailgun/mailgun-go.v1"
)

var _ smtp.Backend = &Backend{}
var _ smtp.User = &User{}

// Backend type
type Backend struct {
	Domain     string
	privateKey string
	publicKey  string
}

// User type
type User struct {
	mailgunClient mg.Mailgun
}

// NewBackend returns new instance of backend
func NewBackend(domain, privateKey, publicKey string) (smtp.Backend, error) {
	if domain == "" || privateKey == "" || publicKey == "" {
		return nil, fmt.Errorf("domain, privateKey, publicKey must not be empty")
	}

	return &Backend{
		Domain:     domain,
		privateKey: privateKey,
		publicKey:  publicKey,
	}, nil
}

// Login is used to authenticate the user
// In relay there's no need for that at the moment
func (b *Backend) Login(username, password string) (smtp.User, error) {
	return &User{
		mailgunClient: mg.NewMailgun(b.Domain, b.privateKey, b.publicKey),
	}, nil
}

// AnonymousLogin returns anonymouse user object
func (b *Backend) AnonymousLogin() (smtp.User, error) {
	return &User{
		mailgunClient: mg.NewMailgun(b.Domain, b.privateKey, b.publicKey),
	}, nil
}

// Send will send email synchronously via Mailgun service
// TODO: Send messages via queue
func (u *User) Send(from string, to []string, r io.Reader) error {
	m, err := mail.ReadMessage(r)
	if err != nil {
		return err
	}

	mBody, err := ioutil.ReadAll(m.Body)
	if err != nil {
		return err
	}

	for _, recipient := range to {
		message := u.mailgunClient.NewMessage(from, m.Header.Get("Subject"), string(mBody), recipient)
		resp, id, err := u.mailgunClient.Send(message)
		if err != nil {
			return err
		}
		log.Printf("ID: %s Resp: %s", id, resp)
	}
	return nil
}

// Logout is calle after all operations are complete within the session
// Here in relay there's no need to implement anything special for that
func (u *User) Logout() error {
	return nil
}
