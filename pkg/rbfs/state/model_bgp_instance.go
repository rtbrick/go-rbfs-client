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

// BGP instance details.
type BgpInstance struct {
	InstanceName string `json:"instance_name,omitempty"`
	// The AS number.
	Asn int64 `json:"asn,omitempty"`
	// Unique router ID
	RouterId string `json:"router_id,omitempty"`
	// The BGP cluster ID this router is member of.
	ClusterId string `json:"cluster_id,omitempty"`
	// The BGP host name.  The host_name has been deprecated. Use the hostname attribute instead.
	HostName string `json:"host_name,omitempty"`
	// The BGP hostname.
	Hostname string `json:"hostname,omitempty"`
	// The BGP domain name.
	DomainName string `json:"domain_name,omitempty"`
	// BGP version.
	BgpVersion int `json:"bgp_version,omitempty"`
	// The IP ToS (type of service) value set in all outgoing BGP packets.
	BgpIpTos       int                         `json:"bgp_ip_tos,omitempty"`
	Capabilities   *BgpInstanceCapabilities    `json:"capabilities,omitempty"`
	RouteSelection *BgpInstanceRouteSelection  `json:"route_selection,omitempty"`
	Timers         *BgpInstanceTimers          `json:"timers,omitempty"`
	Peerings       *BgpInstancePeeringsSummary `json:"peerings,omitempty"`
}
