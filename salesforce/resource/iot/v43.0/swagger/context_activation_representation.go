/*
 * Salesforce IoT API
 *
 * Use the Salesforce IoT REST API to create and manage orchestrations and contexts, and retrieve usage data.
 *
 * API version: 43.0.0
 * Contact: support@salesforce.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ContextActivationRepresentation struct {

	Id string `json:"id,omitempty"`

	// Resource URL for this context activation.
	Url string `json:"url,omitempty"`

	CreatedDate string `json:"createdDate,omitempty"`

	CreatedBy *UserReference `json:"createdBy,omitempty"`

	LastModifiedDate string `json:"lastModifiedDate,omitempty"`

	LastModifiedBy *UserReference `json:"lastModifiedBy,omitempty"`

	// Resource URL for the parent context.
	ContextUrl string `json:"contextUrl,omitempty"`

	Status string `json:"status,omitempty"`

	Runtime string `json:"runtime,omitempty"`

	Definition *ContextDefinitionRepresentation `json:"definition,omitempty"`

	Result *ContextActivationResultRepresentation `json:"result,omitempty"`
}