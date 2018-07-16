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

// A reference to a context.
type ContextReferenceRepresentation struct {

	// A unique identifier for the context.
	Id string `json:"id,omitempty"`

	// The context name.
	Name string `json:"name,omitempty"`

	// A namespace prefix used for packages.
	Namespace string `json:"namespace,omitempty"`

	// Resource URL to retrieve this context.
	Url string `json:"url,omitempty"`

	// The Salesforce IoT runtime used.
	Runtime string `json:"runtime,omitempty"`
}