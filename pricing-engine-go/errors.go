package pricingengine

//Error is the type for errors defined in the project
type Error string

const (
	//ErrInvalidDateFormat error for invalid input date format
	ErrInvalidDateFormat = Error("date format must be yyyy-mm-dd")
	//ErrAgeNotAllowed error for disallowed age of a driver
	ErrAgeNotAllowed = Error("driver with age below 17 is not allowed")
	//ErrInvalidInsuranceGroup error for invalid insurance
	ErrInvalidInsuranceGroup = Error("insurance group is out of allowed range")
)

//Error method returns string representation of the error message
func (e Error) Error() string {
	return string(e)
}
