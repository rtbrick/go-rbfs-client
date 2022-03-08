/*
 * RBFS Operational State API
 *
 * This contract describes the RBFS Operational State API contract defined by RBMS, the RtBrick Management System. This API is a _consumer-driven_ API, which means that all changes to this API **must be approved** by RBMS, the consumer of this API to avoid compatibility issues.  The API is kept backwards-compatible and anyone is allowed to _use_ this API.  The consumer of the API _must_ ignore additional attributes not explained in this specification. Additional attributes are _not_ considered violating backwards compatibility. In contrary, additional attributes allow extending the API while preserving backward compatibility.
 *
 * API version: 1.0.0
 * Contact: martin@rtbrick.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package state

// A reference to an active subscriber session.
type SubscriberSessionRef struct {
	SubscriberId int `json:"subscriber_id,omitempty"`
	// String subscriber_id representation for I-JSON compliance.
	SubscriberIdStr    string `json:"subscriber_id_str,omitempty"`
	SubscriberUserName string `json:"subscriber_user_name,omitempty"`
	// The current subscriber FSM state.
	SubscriberState string                `json:"subscriber_state,omitempty"`
	AccessType      *SubscriberAccessType `json:"access_type,omitempty"`
	IfpName         string                `json:"ifp_name,omitempty"`
	InnerVlan       int                   `json:"inner_vlan,omitempty"`
	OuterVlan       int                   `json:"outer_vlan,omitempty"`
	// Client MAC address.
	ClientMac           string `json:"client_mac,omitempty"`
	AccountingSessionId string `json:"accounting_session_id,omitempty"`
	AgentRemoteId       string `json:"agent_remote_id,omitempty"`
	AgentCircuitId      string `json:"agent_circuit_id,omitempty"`
}
