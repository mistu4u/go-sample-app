package calculator

import (
	"pricingengine"
	"pricingengine/config"
	"pricingengine/helper"
	"time"
)

//DriverAge custom struct for driver age based calculations
type DriverAge struct {
}

//NewDriverAgeImpl gives a new instance of DriverAge
func NewDriverAgeImpl() *DriverAge {
	return &DriverAge{}
}

//AgeCalculator interface for the methods CalculateAge and CalculateAgeBasedPrice
type AgeCalculator interface {
	CalculateAge(dob string) (int, error)
	CalculateAgeBasedPrice(age int) (float64, error)
}

//CalculateAge calculates the drivers age
func (d DriverAge) CalculateAge(dob string) (int, error) {
	dateOfBirth, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return -1, pricingengine.ErrInvalidDateFormat
	}
	//age := int(time.Now().Sub(dateOfBirth).Hours() / (24 * 365))
	age := helper.NumberOfYearsBetweenTwoDates(time.Now(), dateOfBirth)
	if age <= 17 {
		return -1, pricingengine.ErrAgeNotAllowed
	}
	return age, nil
}

//CalculateAgeBasedPrice calculates the price based on the age and factor defined
func (d DriverAge) CalculateAgeBasedPrice(age int) (float64, error) {
	var price float64
	rateConfig := config.NewAgeRateConfigImpl().GetAgeRateConfig()
	if age >= 26 {
		//Take fixed value
		price = rateConfig[26]
	} else {
		price = rateConfig[age]
	}
	return price, nil
}
