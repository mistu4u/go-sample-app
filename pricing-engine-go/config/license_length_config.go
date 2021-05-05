package config

//LicenseLength custom struct for license length
type LicenseLength struct {
}

//NewLicenseLengthImpl gives a new instance of LicenseLength
func NewLicenseLengthImpl() *LicenseLength {
	return &LicenseLength{}
}

//LicenseLengthInterface interface for GetLicenseLengthConfig method
type LicenseLengthInterface interface {
	GetLicenseLengthConfig() map[int]float64
}

//GetLicenseLengthConfig gives the configuration for license length
func (b LicenseLength) GetLicenseLengthConfig() map[int]float64 {
	licenseLengthConfig := make(map[int]float64)
	var i int
	licenseLengthConfig[0] = 1.100
	licenseLengthConfig[6] = 0.950
	for i = 1; i <= 2; i++ {
		licenseLengthConfig[i] = 1.050
	}
	for i = 3; i <= 5; i++ {
		licenseLengthConfig[i] = 1.025
	}
	return licenseLengthConfig
}
