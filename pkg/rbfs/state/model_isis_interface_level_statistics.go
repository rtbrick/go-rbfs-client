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

type IsisInterfaceLevelStatistics struct {
	// Number of IS-IS neighbors.
	NeighborCount       int32                                `json:"neighbor_count,omitempty"`
	LspSuccessCount     int32                                `json:"lsp_success_count,omitempty"`
	LspFailcount        int32                                `json:"lsp_fail:count,omitempty"`
	LspPurgeCount       int32                                `json:"lsp_purge_count,omitempty"`
	CsnpSuccessCount    int32                                `json:"csnp_success_count,omitempty"`
	CsnpFailCount       int32                                `json:"csnp_fail_count,omitempty"`
	PsnpSuccessCount    int32                                `json:"psnp_success_count,omitempty"`
	PsnpFailCount       int32                                `json:"psnp_fail_count,omitempty"`
	CsnpAuthFailCount   int32                                `json:"csnp_auth_fail_count,omitempty"`
	PsnpAuthFailCount   int32                                `json:"psnp_auth_fail_count,omitempty"`
	P2pIihAuthFailCount int32                                `json:"p2p_iih_auth_fail_count,omitempty"`
	Ingress             *IsisInterfaceLevelStatisticsIngress `json:"ingress,omitempty"`
	Egress              *IsisInterfaceLevelStatisticsIngress `json:"egress,omitempty"`
}
