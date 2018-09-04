// Package cfg is used to hold configuration data for the user, mainly billing
// information. This is a strange implementation of a static final struct singleton
// encapsulation.
package cfg

import "sync"

type config struct {
	name       string
	email      string
	telephone  string
	address    string
	unit       string
	zipcode    string
	city       string
	state      string
	ccNumber   string
	ccExpMonth string
	ccExpYear  string
	cVV        string
}

var cfg *config
var once sync.Once

func init() {
	once.Do(func() {
		cfg = generateConfig()
	})
}

func generateConfig() *config {
	// TODO: Load from dotfile
	return &config{
		name:       "Yung Boolean",
		email:      "Yung@Boolean.com",
		telephone:  "0123456789",
		address:    "123 Boolean Rd.",
		unit:       "1",
		zipcode:    "12345",
		city:       "Brooklyn",
		state:      "NY",
		ccNumber:   "4242424242424242",
		ccExpMonth: "12",
		ccExpYear:  "2021",
		cVV:        "123",
	}
}

// If you can only get, then you can't mutate the data. I think.
func Name() string {
	return cfg.name
}

func Email() string {
	return cfg.email
}

func Telephone() string {
	return cfg.telephone
}

func Address() string {
	return cfg.address
}

func Unit() string {
	return cfg.unit
}

func Zipcode() string {
	return cfg.zipcode
}

func City() string {
	return cfg.city
}

func State() string {
	return cfg.state
}

func CCNumber() string {
	return cfg.ccNumber
}

func CCExpMonth() string {
	return cfg.ccExpMonth
}

func CCExpYear() string {
	return cfg.ccExpYear
}

func CVV() string {
	return cfg.cVV
}
