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

type OrchestrationTrackerRepresentation struct {

	// 15/18 character Salesforce identifier.
	Id string `json:"id,omitempty"`

	// Resource URL for this orchestration tracker.
	Url string `json:"url,omitempty"`

	OrchestrationUrl string `json:"orchestrationUrl,omitempty"`

	InstanceKey *OrchestrationInstanceKey `json:"instanceKey,omitempty"`

	StartDate string `json:"startDate,omitempty"`

	ExpirationDate string `json:"expirationDate,omitempty"`

	Terminated bool `json:"terminated,omitempty"`

	TerminatedDate string `json:"terminatedDate,omitempty"`

	TerminationReason string `json:"terminationReason,omitempty"`

	CreatedDate string `json:"createdDate,omitempty"`

	CreatedBy *UserReference `json:"createdBy,omitempty"`

	LastModifiedDate string `json:"lastModifiedDate,omitempty"`

	LastModifiedBy *UserReference `json:"lastModifiedBy,omitempty"`
}