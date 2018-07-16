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

type Source struct {

	// Runtime type associated with the activity.
	Runtime string `json:"runtime,omitempty"`

	// Nominal category associated with the requested object ID.
	Category string `json:"category,omitempty"`

	// 18 character identifier or US document ID associated with activity object.
	Id string `json:"id,omitempty"`

	// The label.
	Label string `json:"label,omitempty"`

	// Name associated with the activity object.
	Name string `json:"name,omitempty"`

	// URL link to the definition associated with the requested object.
	Url string `json:"url,omitempty"`
}