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

// A list of contexts. Each context representation contains context fields including the context ID and other system fields.
type ContextCollectionRepresentation struct {

	// The URL you use to get the next page of contexts.
	NextPageUrl string `json:"nextPageUrl,omitempty"`

	// A list of contexts.
	Contexts []ContextSummaryRepresentation `json:"contexts"`
}