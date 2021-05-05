package calculator

import (
	"github.com/stretchr/testify/assert"
	"pricingengine"
	"testing"
)

func TestResponse(t *testing.T) {
	tests := []struct {
		name         string
		request      pricingengine.GeneratePricingRequest
		priceForADay float64
		errReason    string
	}{
		{
			name: "Happy Path",
			request: pricingengine.GeneratePricingRequest{
				Dob:            "2001-01-01",
				InsuranceGroup: 35,
				LicenseSince:   "2012-08-01",
			},
			priceForADay: 33.88,
			errReason:    "",
		},
		{
			name: "Invalid age",
			request: pricingengine.GeneratePricingRequest{
				Dob:            "2018-01-01",
				InsuranceGroup: 35,
				LicenseSince:   "2012-08-01",
			},
			priceForADay: 0,
			errReason:    pricingengine.ErrAgeNotAllowed.Error(),
		},
		{
			name: "Invalid date format",
			request: pricingengine.GeneratePricingRequest{
				Dob:            "20181-01-01",
				InsuranceGroup: 35,
				LicenseSince:   "2012-08-01",
			},
			priceForADay: 0,
			errReason:    pricingengine.ErrInvalidDateFormat.Error(),
		},
		{
			name: "Invalid insurance group",
			request: pricingengine.GeneratePricingRequest{
				Dob:            "2000-01-01",
				InsuranceGroup: 38,
				LicenseSince:   "2012-08-01",
			},
			priceForADay: 0,
			errReason:    pricingengine.ErrInvalidInsuranceGroup.Error(),
		},
	}
	for i, e := range tests {
		t.Run(e.name, func(t *testing.T) {
			value, err := NewAllCalcImpl().CalculateAll(&tests[i].request)
			if len(tests[i].errReason) == 0 {
				assert.Nil(t, err)
				assert.Equal(t, 33.88, value[86400])
			} else {
				if err != nil {
					assert.Equal(t, tests[i].errReason, err.Error())
					assert.Nil(t, value)
				} else {
					t.Error("the error reason should not be nil")
				}
			}
		})
	}

}
