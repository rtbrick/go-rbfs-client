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

// PIM neighbor settings.
type PimNeighbor struct {
	PimNeighborState string `json:"pim_neighbor_state,omitempty"`
	// Last PIM neighbor state transition.
	LastStateTransition string `json:"last_state_transition,omitempty"`
	IflName             string `json:"ifl_name,omitempty"`
	PimIflState         string `json:"pim_ifl_state,omitempty"`
	// The primary IPv4 address of the PIM interface.
	Ipv4Address string `json:"ipv4_address,omitempty"`
	// The primary IPv6 address of the PIM interface.
	Ipv6Address string `json:"ipv6_address,omitempty"`
	// The PIM interface generation ID.
	GenerationId int `json:"generation_id,omitempty"`
	// The priority of this interface in the designated router election.
	DesignatedRouterPriority int `json:"designated_router_priority,omitempty"`
	// DR elections count
	DesignatedRouterElections int                          `json:"designated_router_elections,omitempty"`
	DesignatedRouter          *PimDesignatedRouter         `json:"designated_router,omitempty"`
	Timers                    *PimInstanceNeighborTimers   `json:"timers,omitempty"`
	Neighbor                  *PimInstanceNeighborNeighbor `json:"neighbor,omitempty"`
}
