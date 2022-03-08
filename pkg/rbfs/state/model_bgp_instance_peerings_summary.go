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

// A summary of the BGP instance peerings including the number of peerings per state and a reference to each peering. The reference allows obtaining the BPG peering details.
type BgpInstancePeeringsSummary struct {
	// Number of peerings in active state.
	ActiveCount int `json:"active_count,omitempty"`
	// Number of peerings in idle state.
	IdleCount int `json:"idle_count,omitempty"`
	// Number of peerings in connect state.
	ConnectCount int `json:"connect_count,omitempty"`
	// Number of peerings in opensent state.
	OpensentCount int `json:"opensent_count,omitempty"`
	// Number of peerings in openconfirm state.
	OpenconfirmCount int `json:"openconfirm_count,omitempty"`
	// Number of peerings in established state.
	EstablishedCount int `json:"established_count,omitempty"`
	// Array of BGP peering references.
	Peerings []BgpPeeringRef `json:"peerings,omitempty"`
}
