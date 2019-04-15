package data

// ClaritinService is a pollen service for Zyrtec formatted data
type ClaritinService struct{}

// GetPollenReport gets the pollen report
func (s ClaritinService) GetPollenReport(zipcode string) (PollenReport, error) {
	retval := PollenReport{}

	return retval, nil
}
