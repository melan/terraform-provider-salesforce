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

// Information about an instance of an orchestration.
type OrchestrationInstanceDetailRepresentation struct {

	// A key that uniquely identifies an orchestration instance. The instanceKey is the value of the event key field.
	InstanceKey string `json:"instanceKey,omitempty"`

	// A unique identifier for the orchestration instance.
	Id string `json:"id,omitempty"`

	// Resource URL to retrieve this orchestration instance.
	Url string `json:"url,omitempty"`

	CreatedDate string `json:"createdDate,omitempty"`

	CreatedBy *UserReference `json:"createdBy,omitempty"`

	LastModifiedDate string `json:"lastModifiedDate,omitempty"`

	LastModifiedBy *UserReference `json:"lastModifiedBy,omitempty"`

	// A description for the orchestration instance.
	Description string `json:"description,omitempty"`

	// A list of variables that are part of the orchestration instance.
	Variables []VariableRepresentation `json:"variables,omitempty"`
}
