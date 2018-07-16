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

type ContextActivationSummaryRepresentation struct {

	Id string `json:"id,omitempty"`

	// Resource URL to retrieve this context summary activation.
	Url string `json:"url,omitempty"`

	CreatedDate string `json:"createdDate,omitempty"`

	CreatedBy *UserReference `json:"createdBy,omitempty"`

	LastModifiedDate string `json:"lastModifiedDate,omitempty"`

	LastModifiedBy *UserReference `json:"lastModifiedBy,omitempty"`

	Status string `json:"status,omitempty"`
}