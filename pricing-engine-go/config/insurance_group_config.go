package config

//InsuranceGroup custom struct for insurance calculation
type InsuranceGroup struct {
}

//NewInsuranceGroupImpl gives a new instance of InsuranceGroup
func NewInsuranceGroupImpl() *InsuranceGroup {
	return &InsuranceGroup{}
}

//InsuranceGroupInterface interface for GetInsuranceGroupConfig method
type InsuranceGroupInterface interface {
	GetBaseRateConfig() map[int]int
}

//GetInsuranceGroupConfig method for getting the configuration for insurance
func (b InsuranceGroup) GetInsuranceGroupConfig() map[int]float64 {
	insuranceGroupConfig := make(map[int]float64)
	var i int
	for i = 1; i <= 8; i++ {
		insuranceGroupConfig[i] = 1.000
	}
	for i = 9; i <= 16; i++ {
		insuranceGroupConfig[i] = 1.073
	}
	for i = 17; i <= 35; i++ {
		insuranceGroupConfig[i] = 1.120
	}
	return insuranceGroupConfig
}
