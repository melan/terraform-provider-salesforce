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

// Orchestration properties.
type OrchestrationPropertiesRepresentation struct {

	// Set to true to create a new orchestration instance for an incoming event with a matching key after an error occurs. By default (false), incoming events are ignored after an error occurs in the matching orchestration instance.
	RestartOnError bool `json:"restartOnError,omitempty"`

	// Set to true to ignore incoming events with a matching key for an exited orchestration instance. By default (false), incoming events with a matching key delete an exited orchestration instance and create a new instance.
	RunOnce bool `json:"runOnce,omitempty"`
}
