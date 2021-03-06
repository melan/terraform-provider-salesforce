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

// The definition of an orchestration including orchestration states, variables, and properties.
type OrchestrationDefinitionRepresentation struct {

	// The states defined for the orchestration.
	States []OrchestrationStateRepresentation `json:"states,omitempty"`

	// The variables defined for the orchestration.
	Variables []OrchestrationVariableRepresentation `json:"variables,omitempty"`
}
