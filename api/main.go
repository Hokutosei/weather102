package main

import (
	"weather102/api/lib/crawlers"
)

var (
	crawlDelay  = 300
	reportDelay = 120
)

func main() {
	quit := make(chan bool)
	go crawlers.WeatherCrawlerStart(crawlDelay, 0)

	go reportGoRoutines(reportDelay)
	<-quit
}
