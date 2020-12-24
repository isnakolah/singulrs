package utils

// GetStrErr returns the string of an err or "", nil
func GetStrErr(err error) (errMessage string) {
	if err != nil {
		errMessage = err.Error()
	}
	return
}
