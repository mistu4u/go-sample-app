package calculator

import (
	log "github.com/sirupsen/logrus"
	"pricingengine"
	"pricingengine/config"
	"pricingengine/helper"
	"time"
)

//LicenseCalc custom struct for license related operations
type LicenseCalc struct {
}

//NewLicenseCalcImpl gives a new instance of LicenseCalc
func NewLicenseCalcImpl() *LicenseCalc {
	return &LicenseCalc{}
}

//LicenseCalculator interface for CalculateLicense method
type LicenseCalculator interface {
	CalculateLicense(lDate string) (float64, error)
}

//CalculateLicense calculates the factor for license
func (i LicenseCalc) CalculateLicense(lDate string) (float64, error) {
	//Get the insurance config
	licenseGrp := config.NewLicenseLengthImpl().GetLicenseLengthConfig()
	license, err := time.Parse("2006-01-02", lDate)
	if err != nil {
		return -1, pricingengine.ErrInvalidDateFormat
	}
	//licenseAge := int(time.Now().Sub(license).Hours() / (24 * 365))
	licenseAge := helper.NumberOfYearsBetweenTwoDates(time.Now(), license)
	log.Infof("licenseage is %v", licenseAge)
	if licenseAge > 6 {
		return licenseGrp[6], nil
	}
	return licenseGrp[licenseAge], nil
}
