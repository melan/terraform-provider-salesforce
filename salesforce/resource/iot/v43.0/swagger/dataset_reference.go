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

// Information about the reference object, such as the object name and type.
type DatasetReference struct {

	// The reference object type.
	Type_ string `json:"type"`

	// The reference object label that's displayed in the user interface.
	Label string `json:"label,omitempty"`

	// A unique name that's used to refer to the reference object through the API.
	Name string `json:"name"`

	// A namespace prefix used for packages when type is SalesforceObjectReference.
	Namespace string `json:"namespace,omitempty"`

	// A unique identifier for the Salesforce object when type is SalesforceObjectReference.
	Id string `json:"id,omitempty"`
}
