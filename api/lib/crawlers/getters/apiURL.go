package getters

import "fmt"

// APIURL return constructed api url
func APIURL(cityID, appID string) string {
	//http://api.openweathermap.org/data/2.5/forecast/weather?id=6822131&APPID=c455dfa704c974e99ad8918d2757476e
	// http://api.openweathermap.org/data/2.5/weather?id=6822131&appid=c455dfa704c974e99ad8918d2757476e
	return fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?id=%v&APPID=%v", cityID, appID)
}
