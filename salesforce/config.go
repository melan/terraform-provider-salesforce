package salesforce

import (
	"log"

	"github.com/nimajalali/go-force/force"
)

type Config struct {
	APIVersion    string
	ClientId      string
	ClientSecret  string
	Username      string
	Password      string
	SecurityToken string
	Environment   string
	TraceOn       bool
}

type clientLogger struct{}

func (cl *clientLogger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (c *Config) Client() (*force.ForceApi, error) {
	log.Printf("Authenticating client with config: %v\n", c)

	forceApi, err := force.Create(
		c.APIVersion,
		c.ClientId,
		c.ClientSecret,
		c.Username,
		c.Password,
		c.SecurityToken,
		c.Environment,
	)

	if c.TraceOn {
		forceApi.TraceOn("[ForceAPI]", &clientLogger{})
	}

	return forceApi, err
}
