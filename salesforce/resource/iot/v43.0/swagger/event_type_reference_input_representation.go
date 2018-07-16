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

// Information about the event.
type EventTypeReferenceInputRepresentation struct {

	// The event type.
	Type_ string `json:"type,omitempty"`

	// The event name.
	Name string `json:"name,omitempty"`

	// A namespace prefix used for packages.
	Namespace string `json:"namespace,omitempty"`
}