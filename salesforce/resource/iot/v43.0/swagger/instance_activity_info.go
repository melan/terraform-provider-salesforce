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

type InstanceActivityInfo struct {

	// Internal runtime code associated with the error.
	ErrorCode string `json:"errorCode,omitempty"`

	// Textual description of the error.
	ErrorMessage string `json:"errorMessage,omitempty"`

	// ID of the state from which the error state originated.
	FromStateId string `json:"fromStateId,omitempty"`

	// Event partition key.
	InstanceKey string `json:"instanceKey,omitempty"`

	// Raw json event representation associated with error.
	Event string `json:"event,omitempty"`
}
