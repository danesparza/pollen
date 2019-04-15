package data

// NasacourtService is a pollen service for Zyrtec formatted data
type NasacourtService struct{}

// GetPollenReport gets the pollen report
func (s NasacourtService) GetPollenReport(zipcode string) (PollenReport, error) {
	retval := PollenReport{}

	return retval, nil
}
