package crawlers

import (
	"fmt"
	"time"
	"weather102/api/lib/crawlers/getters"

	utils "github.com/Hokutosei/hokutoseiUtils"
)

var (
	reportDelaySec = 15
)

// WeatherCrawlerStart weather crawler start getters
// main weather crawler
func WeatherCrawlerStart(delaySec int, counter int) {
	utils.Info(fmt.Sprintf("initializing weather crawler..."))

	Initialize(&counter)

	for t := range time.Tick(time.Duration(delaySec) * time.Second) {
		fmt.Println(t)
		Initialize(&counter)
	}
}

// Initialize initialize get weathers
func Initialize(counter *int) {
	weatherResults := getters.GetWeather()
	for _, weather := range weatherResults {
		fmt.Println(weather)
	}
	*counter++
	fmt.Println("==================================== %v", *counter)
}
