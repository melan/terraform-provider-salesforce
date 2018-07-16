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

// A list of validation messages returned when saving an orchestration.
type OrchestrationValidationMessagesRepresentation struct {

	// A list of validation messages.
	ValidationMessages []OrchestrationValidationMessageRepresentation `json:"validationMessages,omitempty"`
}