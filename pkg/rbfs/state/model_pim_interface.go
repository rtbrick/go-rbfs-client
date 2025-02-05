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

// PIM interface settings.
type PimInterface struct {
	InstanceName string `json:"instance_name,omitempty"`
	IflName      string `json:"ifl_name,omitempty"`
	PimIflState  string `json:"pim_ifl_state,omitempty"`
	// The primary IPv4 address of the PIM interface.
	Ipv4Address string `json:"ipv4_address,omitempty"`
	// The primary IPv6 address of the PIM interface.
	Ipv6Address string `json:"ipv6_address,omitempty"`
	// The PIM generation identifier.
	GenerationId int `json:"generation_id,omitempty"`
	// Total number of PIM neighbors on this interface.
	Neighbors int `json:"neighbors,omitempty"`
	// The priority of this interface in the designated router election.
	DesignatedRouterPriority int `json:"designated_router_priority,omitempty"`
	// Number of designated router elections being executed.
	DesignatedRouterElections int                               `json:"designated_router_elections,omitempty"`
	DesignatedRouter          *PimDesignatedRouter              `json:"designated_router,omitempty"`
	Statistics                *PimInstanceInterfaceStatistics   `json:"statistics,omitempty"`
	Timers                    *PimInstanceInterfaceTimers       `json:"timers,omitempty"`
	Capabilities              *PimInstanceInterfaceCapabilities `json:"capabilities,omitempty"`
}
