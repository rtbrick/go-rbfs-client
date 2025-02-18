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

// Downstream access line attributes.
type AccessLineInfoDownstream struct {
	// Actual downstream data rate.
	ActualRate int `json:"actual_rate,omitempty"`
	// Minimum downstream data rate.
	MinRate int `json:"min_rate,omitempty"`
	// Attainable downstream data rate.
	AttainableRate int `json:"attainable_rate,omitempty"`
	// Maximum downstream data rate.
	MaxRate int `json:"max_rate,omitempty"`
	// Minimum downstream data rate in low power state.
	MinRateLowPower int `json:"min_rate_low_power,omitempty"`
	// Maximum downstream interleaving delay.
	MaxInterleaveDelay int `json:"max_interleave_delay,omitempty"`
	// Actual downstream interleaving delay.
	ActualInterleaveDelay int `json:"actual_interleave_delay,omitempty"`
	// Downstream Expected Throughput (ETR).
	ExpectedThroughput int `json:"expected_throughput,omitempty"`
	// Downstream Attainable Expected Throughput (ATTETR).
	AttainableExpectedThroughput int `json:"attainable_expected_throughput,omitempty"`
	// Downstream Gamma Data Rate (GDR).
	GammaDataRate int `json:"gamma_data_rate,omitempty"`
	// Downstream Attainable Gamma Data Rate (ATTGDR).
	AttainableGammaDataRate int `json:"attainable_gamma_data_rate,omitempty"`
	// ONT/ONU downstream average data rate.
	OntOnuAverageRate int `json:"ont_onu_average_rate,omitempty"`
	// ONT/ONU downstream peak data rate.
	OntOnuPeakRate int `json:"ont_onu_peak_rate,omitempty"`
	// PON tree downstream maximum data rate.
	PonMaxRate int `json:"pon_max_rate,omitempty"`
}
