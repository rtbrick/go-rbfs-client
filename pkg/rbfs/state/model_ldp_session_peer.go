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

// The peer of this LSR session.
type LdpSessionPeer struct {
	// The LSR identifier.
	LsrId   string   `json:"lsr_id,omitempty"`
	LdpRole *LdpRole `json:"ldp_role,omitempty"`
	// The peer IPv4 address.
	Ipv4Address string `json:"ipv4_address,omitempty"`
	// The peer IPv6 address.
	Ipv6Address string `json:"ipv6_address,omitempty"`
}
