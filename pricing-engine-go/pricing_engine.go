package pricingengine

// GeneratePricingRequest is used for generate pricing requests, it holds the
// inputs that are used to provide pricing for a given user.
type GeneratePricingRequest struct {
	Dob            string `json:"date_of_birth"`
	InsuranceGroup int    `json:"insurance_group"`
	LicenseSince   string `json:"license_held_since"`
}

// GeneratePricingResponse is used for responding to the user
type GeneratePricingResponse struct {
	IsApproved      bool       `json:"is_approved"`
	PriceRanges     PriceRange `json:"price_ranges"`
	ReasonForDenial string     `json:"reason_for_denial"`
}

//PriceRange is a custom type, acts as a map to hold the time range and cost that the user has to pay combination
//It is used as part of the response object
type PriceRange map[int]float64
