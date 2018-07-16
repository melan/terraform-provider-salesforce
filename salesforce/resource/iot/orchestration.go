package iot

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/melan/terraform-provider-salesforce/salesforce"
	iotSwagger "github.com/melan/terraform-provider-salesforce/salesforce/resource/iot/v43.0/swagger"
)

const (
	OrchestrationField   = "orchestration"
	OrchestrationIdField = "orchestration_id"
)

func Orchestration() *schema.Resource {
	return &schema.Resource{
		Create: orchestrationCreate,
		Read:   orchestrationRead,
		//Update: orchestrationUpdate,
		Delete: orchestrationDelete,

		Schema: map[string]*schema.Schema{
			OrchestrationIdField: {
				Type:     schema.TypeString,
				Optional: true,
			},
			OrchestrationField: {
				Type:     schema.TypeString,
				Optional: true,
			},
		},

		SchemaVersion: 1,
	}
}

func orchestrationCreate(d *schema.ResourceData, m interface{}) error {
	sfdcClient := m.(*salesforce.Client)
	client := sfdcClient.IotAPI
	ctx := context.WithValue(sfdcClient.Ctx, iotSwagger.ContextAccessToken, sfdcClient.OAuthClient.AccessToken)

	orchestrationId, found := d.GetOk(OrchestrationIdField)
	if found && len(orchestrationId.(string)) > 0 {
		return fmt.Errorf("%s field is defined already", OrchestrationIdField)
	}

	orchestration, found := d.GetOk(OrchestrationField)
	if !found {
		return fmt.Errorf("%s field is undefined", OrchestrationField)
	}

	var orchestrationInputRepresentation iotSwagger.OrchestrationInputRepresentation

	if err := json.Unmarshal([]byte(orchestration.(string)), &orchestrationInputRepresentation); err != nil {
		return fmt.Errorf("can't deserialize Orchestration: %v", err)
	}

	orchResponse, _, err := client.OrchestrationApi.CreateOrchestration(ctx, orchestrationInputRepresentation)
	if err != nil {
		return err
	}

	orchId := orchResponse.Id

	jsonOrch, err := json.Marshal(orchResponse)
	if err != nil {
		return err
	}

	d.Set(OrchestrationField, jsonOrch)
	d.SetId(orchId)

	return nil
}

func orchestrationRead(d *schema.ResourceData, m interface{}) error {
	sfdcClient := m.(*salesforce.Client)
	client := sfdcClient.IotAPI
	ctx := context.WithValue(sfdcClient.Ctx, iotSwagger.ContextAccessToken, sfdcClient.OAuthClient.AccessToken)

	orchestrationId, found := d.GetOk(OrchestrationIdField)
	if !found {
		return fmt.Errorf("%s field is undefined", OrchestrationIdField)
	}

	orchestration, _, err := client.OrchestrationApi.GetOrchestrationById(ctx, orchestrationId.(string))
	if err != nil {
		return err
	}

	jsonOrch, err := json.Marshal(orchestration)
	if err != nil {
		return err
	}

	d.Set(OrchestrationField, jsonOrch)
	d.SetId(orchestrationId.(string))

	return nil
}

//func orchestrationUpdate(d *schema.ResourceData, m interface{}) error {
//	return nil
//}

func orchestrationDelete(d *schema.ResourceData, m interface{}) error {
	sfdcClient := m.(*salesforce.Client)
	client := sfdcClient.IotAPI
	ctx := context.WithValue(sfdcClient.Ctx, iotSwagger.ContextAccessToken, sfdcClient.OAuthClient.AccessToken)

	orchestrationId, found := d.GetOk(OrchestrationIdField)
	if !found {
		return fmt.Errorf("%s field is undefined", OrchestrationIdField)
	}

	_, err := client.OrchestrationApi.DeleteOrchestration(ctx, orchestrationId.(string))
	if err != nil {
		return err
	}

	return nil
}
