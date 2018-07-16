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

type ContextActivationCollectionRepresentation struct {

	NextPage string `json:"nextPage"`

	TotalSize string `json:"totalSize,omitempty"`

	Url string `json:"url,omitempty"`

	Activations []ContextActivationSummaryRepresentation `json:"activations"`
}