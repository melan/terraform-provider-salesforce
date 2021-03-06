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

// An event referenced in a context.
type OrchestrationEventTypeInputRepresentation struct {

	// The key field selected to match event messages with orchestration instances. Only one key field is supported.
	KeyFields []string `json:"keyFields,omitempty"`
}
