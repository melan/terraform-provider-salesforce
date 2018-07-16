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

type OrchestrationLogInfoEventRepresentation struct {

	OrchestrationId string `json:"orchestrationId"`

	ActivationId string `json:"activationId"`

	InstanceKey *OrchestrationInstanceKey `json:"instanceKey"`

	Timestamp int64 `json:"timestamp"`

	Name string `json:"name"`

	Location string `json:"location"`

	Message string `json:"message"`
}
