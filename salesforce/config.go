package salesforce

import (
	"context"
	"log"
	"net/http"

	"github.com/melan/go-force/force"
	iotSwagger "github.com/melan/terraform-provider-salesforce/salesforce/resource/iot/v43.0/swagger"
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

type Client struct {
	OAuthClient *force.OAuth
	SObjectApi  *force.Api
	IotAPI      *iotSwagger.APIClient
	Ctx         context.Context
}

func (c *Config) Client() (*Client, error) {
	log.Printf("Authenticating client with config: %v\n", c)

	oauth, err := force.CreateOAuth(
		c.ClientId,
		c.ClientSecret,
		c.Username,
		c.Password,
		c.SecurityToken,
		c.Environment,
	)

	forceApi, err := force.CreateApi(c.APIVersion, oauth)

	if err != nil {
		return nil, err
	}

	if c.TraceOn {
		forceApi.TraceOn("[ForceAPI]", &clientLogger{})
	}

	client := &http.Client{}
	iotClientConfig := iotSwagger.NewConfiguration()
	iotClientConfig.UserAgent = "terraform-provider-salesforce"
	iotClientConfig.HTTPClient = client

	iotClient := iotSwagger.NewAPIClient(iotClientConfig)

	ctx, _ := context.WithCancel(context.Background())

	return &Client{
		OAuthClient: oauth,
		SObjectApi:  forceApi,
		IotAPI:      iotClient,
		Ctx:         ctx,
	}, nil
}
