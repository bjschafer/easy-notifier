package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/gomail.v2"
	"gopkg.in/yaml.v2"
)

type config struct {
	SMTPServer string   `yaml:"SMTPServer"`
	SMTPPort   int      `yaml:"SMTPPort"`
	SMTPUser   string   `yaml:"SMTPUser"`
	SMTPPass   string   `yaml:"SMTPPass"`
	FromEmail  string   `yaml:"FromEmail"`
	ToEmails   []string `yaml:"ToEmails"`
}

// Index provides usage information.
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Usage info goes here...")
}

// EmailNotify sends am email to a prespecified recipient based on URL parameters
func EmailNotify(w http.ResponseWriter, r *http.Request) {
	if err := sendEmail(r.FormValue("from"), r.FormValue("message")); err != nil {
		fmt.Fprintln(w, err)
		fmt.Fprintln(w, "Error sending email")
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusTeapot)
	} else {
		fmt.Fprintln(w, "I'll let him know")
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

	// Load config
	c, err := loadConfig()
	if err != nil {
		return err
	}

	// Set headers/info
	m := gomail.NewMessage()
	m.SetHeader("From", c.FromEmail)
	m.SetHeader("To", c.ToEmails...)
	m.SetHeader("Subject", "Notification from "+fromApp)
	m.SetBody("text/html", message)

	// Send mail

	d := gomail.NewPlainDialer(c.SMTPServer, c.SMTPPort, c.SMTPUser, c.SMTPPass)

	return d.DialAndSend(m)
}

func loadConfig() (config, error) {
	var c config

	data, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		return c, err
	}

	err = yaml.Unmarshal(data, &c)
	return c, err

}
