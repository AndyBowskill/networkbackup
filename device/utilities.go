package device

import (
	"fmt"
	"log"
	"time"
)

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func currentDateTimeAsString() string {
	currentTime := time.Now()
	return fmt.Sprintf("%v-%v-%v-%v-%v", currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Minute())
}
