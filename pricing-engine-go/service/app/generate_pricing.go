package app

import (
	"context"
	"pricingengine"
	"pricingengine/calculator"
)

//App custom struct
type App struct{}

// GeneratePricing will calculate how much a 'risk' be priced or if they should
// be denied.
func (a *App) GeneratePricing(ctx context.Context, input *pricingengine.GeneratePricingRequest) (
	*pricingengine.GeneratePricingResponse, error) {
	all, err := calculator.NewAllCalcImpl().CalculateAll(input)
	emptyRange := make(pricingengine.PriceRange, 0)
	if err != nil {
		if err == pricingengine.ErrAgeNotAllowed ||
			err == pricingengine.ErrInvalidInsuranceGroup {
			return &pricingengine.GeneratePricingResponse{
				IsApproved:      false,
				PriceRanges:     emptyRange,
				ReasonForDenial: err.Error(),
			}, nil
		}
		return nil, err
	}
	return &pricingengine.GeneratePricingResponse{
		IsApproved:      true,
		PriceRanges:     all,
		ReasonForDenial: "",
	}, nil
}
