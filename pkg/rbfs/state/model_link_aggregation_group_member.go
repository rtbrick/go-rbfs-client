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

// A member interface of a link aggregation group.
type LinkAggregationGroupMember struct {
	IfpName string `json:"ifp_name,omitempty"`
	// The operational state of the member interface.
	OperationalState string `json:"operational_state,omitempty"`
	// The administrative state of the member interface
	AdministrativeState string `json:"administrative_state,omitempty"`
	// The configured speed limit. The maximum speed can be read from the bandwidth attribute.
	Speed string `json:"speed,omitempty"`
	// The maximum bandwidth available.
	Bandwidth string `json:"bandwidth,omitempty"`
	// The LAG member state indicates whether this LAG member is resolved or unresolved.
	LagMemberState string `json:"lag_member_state,omitempty"`
}
