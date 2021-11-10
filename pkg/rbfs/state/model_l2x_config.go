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

// L2X settings.
type L2xConfig struct {
	// The L2X nexthop IPv4 or IPv6 address.
	IngressNexthop string `json:"ingress_nexthop,omitempty"`
	// The routing instance of the nexthop address.
	IngressLookupInstance string `json:"ingress_lookup_instance,omitempty"`
	// The L2X service MPLS label for ingress traffic (sent label).
	IngressServiceLabel int32 `json:"ingress_service_label,omitempty"`
	// The L2X service MPLS label for egress traffic (receive label).
	EgressServiceLabel int32 `json:"egress_service_label,omitempty"`
}
