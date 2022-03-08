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

// The ephemeral configuration of a A10-NSP L2X service. This configuration is lost if the switch reboots.
type A10nspConfig struct {
	LagInterfaceName string `json:"lag_interface_name,omitempty"`
	// The S-VLAN ID.
	SVlan int `json:"s_vlan,omitempty"`
	// The S-VLAN ethertype, which is either 802.1q (0x8100) or 802.1ad (0x88A8).
	SVlanEthertype string `json:"s_vlan_ethertype,omitempty"`
	// The ANP-VLAN ID.
	AnpVlan int        `json:"anp_vlan,omitempty"`
	L2x     *L2xConfig `json:"l2x,omitempty"`
}
