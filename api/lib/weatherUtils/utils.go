package weatherUtils

import "time"

// ConvertKelvin convert kelvin to celsius
func ConvertKelvin(kelvin float64) int {
	return int(kelvin - 273.15)
}

// UnixToLocalTime convert unix time to local
func UnixToLocalTime(unixTime int64) time.Time {
	t := time.Unix(unixTime, 0)
	return t.Local()
}
