package config

//BaseRate custom struct for base price configuration
type BaseRate struct {
}

//NewBaseRateConfigImpl gives a new instance of BaseRate
func NewBaseRateConfigImpl() *BaseRate {
	return &BaseRate{}
}

//BaseConfigInterface interface for GetBaseRateConfig method
type BaseConfigInterface interface {
	GetBaseRateConfig() map[int]int
}

//GetBaseRateConfig method to get the configuration for base price
func (b BaseRate) GetBaseRateConfig() map[int]int {
	BaseRateConfig := make(map[int]int)
	BaseRateConfig[1800] = 273
	BaseRateConfig[3600] = 493
	BaseRateConfig[7200] = 755
	BaseRateConfig[10800] = 998
	BaseRateConfig[21600] = 1242
	BaseRateConfig[43200] = 2033
	BaseRateConfig[86400] = 2211
	BaseRateConfig[172800] = 3249
	BaseRateConfig[259200] = 4419
	BaseRateConfig[345600] = 5204
	return BaseRateConfig
}
