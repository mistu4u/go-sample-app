package config

//AgeRate custom struct for age configuration
type AgeRate struct {
}

//NewAgeRateConfigImpl gives a new instance of AgeRate
func NewAgeRateConfigImpl() *AgeRate {
	return &AgeRate{}
}

//AgeRateConfigInterface interface for getting age based configuration
type AgeRateConfigInterface interface {
	GetAgeRateConfig() map[int]float64
}

//GetAgeRateConfig gets the age based values
func (b AgeRate) GetAgeRateConfig() map[int]float64 {
	ageRateConfig := make(map[int]float64)
	ageRateConfig[18] = 1.540
	ageRateConfig[19] = 1.520
	ageRateConfig[20] = 1.440
	ageRateConfig[21] = 1.400
	ageRateConfig[22] = 1.370
	ageRateConfig[23] = 1.360
	ageRateConfig[24] = 1.360
	ageRateConfig[25] = 1.100
	ageRateConfig[26] = 1.000
	return ageRateConfig
}
