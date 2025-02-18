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

type TracerouteHop struct {
	// The hop sequence number.
	Seq int `json:"seq,omitempty"`
	// The hop IPv4 or IPv6 addresses.
	Hop string `json:"hop,omitempty"`
	// Round-trip time (RTT) in seconds of the first packet sent by traceroute.
	Rtt1 float64 `json:"rtt_1,omitempty"`
	// Round-trip time (RTT) in seconds of the second packet sent by traceroute.
	Rtt2 float64 `json:"rtt_2,omitempty"`
	// Round-trip time (RTT) in seconds of the third packet sent by traceroute.
	Rtt3 float64 `json:"rtt_3,omitempty"`
}
