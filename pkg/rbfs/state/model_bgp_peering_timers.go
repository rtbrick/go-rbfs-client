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

// BGP timers and their current values.
type BgpPeeringTimers struct {
	// The keep alive timer value in seconds.
	KeepAliveTimer int32 `json:"keep_alive_timer,omitempty"`
	// The negotiated keep-alive time in seconds.
	KeepAliveTime int32 `json:"keep_alive_time,omitempty"`
	// The received keep alive time in seconds.
	KeepAliveTimeReceived int32 `json:"keep_alive_time_received,omitempty"`
	// The configured keep alive time seconds.
	KeepAliveTimeSent int32 `json:"keep_alive_time_sent,omitempty"`
	// The hold timer value in seconds.
	HoldTimer int32 `json:"hold_timer,omitempty"`
	// The negotiated hold time in seconds.
	HoldTime int32 `json:"hold_time,omitempty"`
	// The received hold time in seconds.
	HoldTimeReceived int32 `json:"hold_time_received,omitempty"`
	// The sent hold time in seconds.
	HoldTimeSent int32 `json:"hold_time_sent,omitempty"`
	// The connect timer value in seconds.
	ConnectTimer int32 `json:"connect_timer,omitempty"`
	// The negotiated connect time in seconds.
	ConnectTime int32 `json:"connect_time,omitempty"`
	// The sent connect time in seconds.
	ConnectTimeSent int32 `json:"connect_time_sent,omitempty"`
	// The received connect time in seconds.
	ConnectTimeReceived int32 `json:"connect_time_received,omitempty"`
}