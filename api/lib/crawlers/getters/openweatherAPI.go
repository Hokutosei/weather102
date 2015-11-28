package getters

import (
	"fmt"
	"time"
	"weather102/api/lib/weatherUtils"

	utils "github.com/Hokutosei/hokutoseiUtils"
)

var (
	appID = "c455dfa704c974e99ad8918d2757476e"
)

// GetCityWeather api URL
func GetCityWeather(cityID string) map[string]interface{} {
	start := time.Now()

	// HTTP get to url
	var result map[string]interface{}
	if err := HTTPGetter(APIURL(cityID, appID), &result); err != nil {
		utils.Error(fmt.Sprintf("error HTTPGetter %v", err))
		return map[string]interface{}{
			"city": cityID,
		}
	}
	finish := time.Since(start)

	// return if cod == 404
	if result["cod"] == "404" {
		return map[string]interface{}{
			"city":    cityID,
			"cod":     result["cod"],
			"message": result["message"],
		}
	}

	// return if cod == 200
	// push to map required display
	return map[string]interface{}{
		"city":   result["name"],
		"temp":   weatherUtils.ConvertKelvin(result["main"].(map[string]interface{})["temp"].(float64)),
		"getter": finish,
		"dt":     weatherUtils.UnixToLocalTime(int64(result["dt"].(float64))),
	}
}
