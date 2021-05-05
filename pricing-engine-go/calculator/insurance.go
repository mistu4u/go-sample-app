package calculator

import (
	log "github.com/sirupsen/logrus"
	"pricingengine"
	"pricingengine/config"
)

//InsuranceCalc custom struct for insurance operations
type InsuranceCalc struct {
}

//NewInsuranceCalcImpl gives a new instance of InsuranceCalc
func NewInsuranceCalcImpl() *InsuranceCalc {
	return &InsuranceCalc{}
}

//InsuranceCalculator interface for CalculateInsurance method
type InsuranceCalculator interface {
	CalculateInsurance(iGroup int) (float64, error)
}

//CalculateInsurance calculates the insurance factor
func (i InsuranceCalc) CalculateInsurance(iGroup int) (float64, error) {
	//Get the insurance config
	insuranceGrp := config.NewInsuranceGroupImpl().GetInsuranceGroupConfig()
	if insuranceGrp[iGroup] == 0.00 {
		return 0.00, pricingengine.ErrInvalidInsuranceGroup
	}
	log.Infof("insurance group is %v", iGroup)
	return insuranceGrp[iGroup], nil
}
