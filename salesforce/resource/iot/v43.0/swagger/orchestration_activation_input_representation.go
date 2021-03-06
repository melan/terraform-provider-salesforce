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

// Activation status and options.
type OrchestrationActivationInputRepresentation struct {

	// Use the status to deactivate an orchestration by setting the status to Inactive. No status is needed when activating an orchestration.
	Status string `json:"status,omitempty"`
}
