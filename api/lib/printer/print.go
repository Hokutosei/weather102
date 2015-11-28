package printer

import (
	"fmt"
	"weather102/api/lib/models"

	utils "github.com/Hokutosei/hokutoseiUtils"
)

// PrintWeatherResult print weather result
func PrintWeatherResult(weatherResponse models.WeatherResponse) {
	utils.Info(fmt.Sprintf("%v: %v", weatherResponse.Name, weatherResponse.Main.Temp))
}
