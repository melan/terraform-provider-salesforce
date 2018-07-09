package resource

import (
	"log"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nimajalali/go-force/force"
)

func Org() *schema.Resource {
	return &schema.Resource{
		Read: orgRead,

		Schema: map[string]*schema.Schema{
			"org_instance": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"org_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"asserted_user": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"user_nick_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_email": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"active": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"user_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"language": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"locale": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func orgRead(d *schema.ResourceData, m interface{}) error {
	log.Println("Reading an org")
	client := m.(*force.ForceApi)

	instanceUrl := client.GetInstanceURL()
	key := instanceUrl

	log.Printf("Found instance url for the org: %v\n", instanceUrl)
	d.Set("instance_url", instanceUrl)

	if urlObj, err := url.Parse(instanceUrl); err != nil {
		d.Partial(true)
		return err
	} else {
		host := urlObj.Host
		components := strings.Split(host, ".")
		d.Set("org_instance", components[0])
	}

	var apiResources map[string]string

	if err := client.Get("/services/data/v43.0/", nil, &apiResources); err != nil {
		return err
	} else {
		identityUrl := apiResources["identity"]
		if iurl, err := url.Parse(identityUrl); err != nil {
			return err
		} else {
			var identity map[string]interface{}

			if err := client.Get(iurl.Path, nil, &identity); err != nil {
				return err
			}

			log.Printf("Identity: path: %s, result: %v", iurl.Path, identity)

			d.Set("org_id", identity["organization_id"].(string))
			d.Set("user_id", identity["user_id"].(string))
			d.Set("asserted_user", identity["asserted_user"].(bool))
			d.Set("user_nick_name", identity["nick_name"].(string))
			d.Set("user_display_name", identity["display_name"].(string))
			d.Set("user_email", identity["email"].(string))
			d.Set("active", identity["active"].(bool))
			d.Set("user_type", identity["user_type"].(string))
			d.Set("language", identity["language"].(string))
			d.Set("locale", identity["locale"].(string))

			key = identity["organization_id"].(string)
		}
	}

	d.Partial(false)
	d.SetId(key)

	return nil
}
