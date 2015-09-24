package main

import (
	"fmt"
	"net/http"
	"net/smtp"

	"github.com/BurntSushi/toml"
)

type config struct {
	smtpServer string
	smtpPort   string
	smtpUser   string
	smtpPass   string
	fromEmail  string
	toEmails   []string
}

// Index provides usage information.
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Usage info goes here...")
}

// EmailNotify sends am email to a prespecified recipient based on URL parameters
func EmailNotify(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I'll let him know")
	if err := sendEmail(r.FormValue("from"), r.FormValue("message")); err != nil {
		fmt.Fprintln(w, "Error sending email")
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusTeapot)
	} else {

		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
	//if err := json.NewEncoder(w).Encode(todos); err != nil {
	//	panic(err)
	//}
}

///////////////////////////////////////
// Private helper functions
///////////////////////////////////////

func sendEmail(fromApp string, message string) error {
	c, err := loadConfig()
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth("", c.smtpUser, c.smtpPass, c.smtpServer)

	// connect, auth, set sender and recipient, and send:
	return smtp.SendMail(c.smtpServer+":"+c.smtpPort, auth, c.fromEmail, c.toEmails, []byte("Notification from "+fromApp+": "+message))

}

func loadConfig() (config, error) {
	var c config
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		return c, err
	}

	return c, nil
}
