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

// Next hop details.
type NextHops struct {
	// IP address of the next hop.
	NexthopIpAddress string `json:"nexthop_ip_address,omitempty"`
	// MAC addess of the next hop.
	NexthopMacAddress string `json:"nexthop_mac_address,omitempty"`
	// Logical exit interface.
	ExitIfl string `json:"exit_ifl,omitempty"`
	// Type of the next hop.
	NexthopType string `json:"nexthop_type,omitempty"`
	// MLPS label operation.
	NexthopAction string `json:"nexthop_action,omitempty"`
	// MPLS labels for the next hop.
	MplsLabelStack []string `json:"mpls_label_stack,omitempty"`
	// Routing instance in which the NH is resolved.
	LookupInstance string `json:"lookup_instance,omitempty"`
	// Address family of the next hop.
	LookupAfi string `json:"lookup_afi,omitempty"`
	// Sub-address family of the next hop.
	LookupSafi string `json:"lookup_safi,omitempty"`
}
