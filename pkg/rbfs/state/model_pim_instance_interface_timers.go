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

// Configured timer values.
type PimInstanceInterfaceTimers struct {
	// The hello timer interval in seconds.
	HelloInterval int `json:"hello_interval,omitempty"`
	// The hold down interval in seconds.
	HoldDownInterval int `json:"hold_down_interval,omitempty"`
	// The Join/Prune interval in seconds.
	JoinPruneInterval int `json:"join_prune_interval,omitempty"`
	// The override interval in milliseconds.
	OverrideInterval int `json:"override_interval,omitempty"`
	// The prune delay interval in milliseconds.
	PruneDelayInterval int `json:"prune_delay_interval,omitempty"`
}
