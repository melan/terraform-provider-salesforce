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

// An orchestration rule.
type OrchestrationRuleRepresentation struct {

	// A unique identifier for the rule.
	Id int32 `json:"id,omitempty"`

	// A rule description.
	Description string `json:"description,omitempty"`

	// Specifies the condition to be evaluated. The Salesforce Expression Language is supported.
	Condition string `json:"condition,omitempty"`

	// A list of actions to execute, such as assigning a variable value or creating a Salesforce record, when a condition is met.
	Actions []OrchestrationActionRepresentation `json:"actions,omitempty"`
}