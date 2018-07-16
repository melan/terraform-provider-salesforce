package salesforce

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/melan/terraform-provider-salesforce/salesforce/resource"
	"github.com/melan/terraform-provider-salesforce/salesforce/resource/iot"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_version": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SFDX_API_VERSION", ""),
			},
			"client_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SFDX_CLIENT_ID", ""),
			},
			"client_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SFDX_CLIENT_SECRET", ""),
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SFDX_USERNAME", ""),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SFDX_PASSWORD", ""),
			},
			"security_token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SFDX_SECURITY_TOKEN", ""),
			},
			"environment": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SFDX_ENVIRONMENT", ""),
			},
			"trace_on": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"sfdx_user":              resource.User(),
			"sfdx_iot_orchestration": iot.Orchestration(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"sfdx_org": resource.Org(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		APIVersion:    d.Get("api_version").(string),
		ClientId:      d.Get("client_id").(string),
		ClientSecret:  d.Get("client_secret").(string),
		Username:      d.Get("username").(string),
		Password:      d.Get("password").(string),
		SecurityToken: d.Get("security_token").(string),
		Environment:   d.Get("environment").(string),
		TraceOn:       d.Get("trace_on").(bool),
	}
	log.Println("[INFO] Initializing SFDC clinet")
	return config.Client()
}
