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

type IsisInstanceInterfaces struct {
	// The IS-IS instance name.
	InstanceName string `json:"instance_name,omitempty"`
	// The IS-IS host name.  The host_name attribute has been deprecated. Use the hostname attribute instead.
	HostName string `json:"host_name,omitempty"`
	// The IS-IS hostname.
	Hostname string `json:"hostname,omitempty"`
	// The IS-IS system ID.
	SystemId string `json:"system_id,omitempty"`
	// IS-IS areas seen by this switch.
	Areas []string `json:"areas,omitempty"`
	// Status of the IS-IS overload flag. The value is true if the overload flag is set and the switch runs out of resources and might not compute the LSP in time.
	Overload bool `json:"overload,omitempty"`
	// Segment routing settings.
	SegmentRouting *AllOfIsisInstanceInterfacesSegmentRouting `json:"segment_routing,omitempty"`
	Interfaces     *IsisInstanceInterfacesInterfaces          `json:"interfaces,omitempty"`
}
