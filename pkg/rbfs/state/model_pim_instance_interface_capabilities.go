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

// PIM interface capabilities
type PimInstanceInterfaceCapabilities struct {
	// Disable join/suppress flag
	DisableJoinSuppress bool `json:"disable_join_suppress,omitempty"`
	// Use DR priority flag.
	UseDesignatedRouterPriority bool `json:"use_designated_router_priority,omitempty"`
	// Use prune delay flag.
	UseLanPruneDelay bool `json:"use_lan_prune_delay,omitempty"`
	// Use override interval flag.
	UseOverrideInterval int32 `json:"use_override_interval,omitempty"`
	// Prune delay interval in seconds.
	PruneDelayInterval int32 `json:"prune_delay_interval,omitempty"`
	// LAN prune interval in seconds.
	LanPruneInterval int32 `json:"lan_prune_interval,omitempty"`
	// LAN override interval in seconds.
	LanOverrideInterval int32 `json:"lan_override_interval,omitempty"`
	// Join/prune interval in seconds.
	JoinPruneInterval int32 `json:"join_prune_interval,omitempty"`
}