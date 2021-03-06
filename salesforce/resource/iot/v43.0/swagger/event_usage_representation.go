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

// Event usage information that includes the number of events processed and skipped in the last 24 hours.
type EventUsageRepresentation struct {

	// The number of events processed successfully across the whole org in the last 24 hours.
	ProcessedEventCount int32 `json:"processedEventCount,omitempty"`

	// Number of events skipped across the whole Salesforce org due to exceeded allocations in the last 24 hrs.
	RejectedEventCount int32 `json:"rejectedEventCount,omitempty"`
}
