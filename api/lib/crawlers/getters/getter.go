package getters

import (
	"fmt"
	"math/rand"
	"time"
	"weather102/api/lib/models"

	utils "github.com/Hokutosei/hokutoseiUtils"
)

// GetWeather get weather
func GetWeather() []map[string]interface{} {
	weatherChans := make(chan map[string]interface{})
	cities := models.Cities()

	// MAIN loop through cities
	for _, city := range cities {
		go func(ID string) {
			sleep := time.Duration(rand.Intn(1e3)) * time.Millisecond
			utils.Info(fmt.Sprintf("sleeping for... %v", sleep))
			time.Sleep(sleep)

			cityID := fmt.Sprintf("%v", ID)
			weatherChans <- GetCityWeather(cityID)
		}(city)
	}
	defer close(weatherChans)

	results := make([]map[string]interface{}, len(cities))
	for i := 0; i < len(cities); i++ {
		x := <-weatherChans
		results[i] = x
	}

	return results
}
