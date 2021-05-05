package calculator

import (
	log "github.com/sirupsen/logrus"
	"pricingengine"
	"pricingengine/config"
	"pricingengine/helper"
)

//AllCalc custom struct for calculate all
type AllCalc struct {
}

//NewAllCalcImpl give a new instance of the struct
func NewAllCalcImpl() *AllCalc {
	return &AllCalc{}
}

//AllCalculator interface of the CalculateAll method
type AllCalculator interface {
	CalculateAll(input *pricingengine.GeneratePricingRequest) (pricingengine.PriceRange, error)
}

//CalculateAll calculates the final price to pay
func (i AllCalc) CalculateAll(input *pricingengine.GeneratePricingRequest) (pricingengine.PriceRange, error) {
	//Calculate the drivers age
	age, err := NewDriverAgeImpl().CalculateAge(input.Dob)
	if err != nil {
		return nil, err
	}
	log.Infof("age calculated is %v", age)
	license, err := NewLicenseCalcImpl().CalculateLicense(input.LicenseSince)
	log.Printf("license factor is %v", license)
	if err != nil {
		return nil, err
	}
	insurance, err := NewInsuranceCalcImpl().CalculateInsurance(input.InsuranceGroup)
	log.Printf("insurance factor calculated is %v", insurance)
	if err != nil {
		return nil, err
	}
	agePrice, err := NewDriverAgeImpl().CalculateAgeBasedPrice(age)
	log.Printf("age factor calculated is %v", agePrice)
	if err != nil {
		return nil, err
	}
	baseRates := config.NewBaseRateConfigImpl().GetBaseRateConfig()
	factorsMultiplied := agePrice * license * insurance
	log.Printf("all factors calculated is %v", factorsMultiplied)
	result := make(pricingengine.PriceRange)
	for j := range baseRates {
		result[j] = helper.RoundFloatToNearest(float64(baseRates[j]) * factorsMultiplied / 100)
	}

	return result, nil
}
