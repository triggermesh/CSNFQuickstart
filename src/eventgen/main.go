package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

type Event struct {
	EventType string `json:"eventType"`
	EventData string `json:"eventData"`
}

func main() {
	if os.Args[1] == "--help" || os.Args[1] == "-h" || os.Args[1] == "help" {
		displayHelp()
		return
	}

	var selectedEvents = []Event{}
	switch os.Args[1] {
	case "Azure":
		selectedEvents = AzureEventChain
	case "Oracle":
		selectedEvents = OracleEventChain
	case "Aquasec":
		selectedEvents = AquasecEventChain
	default:
		selectedEvents = AzureEventChain
	}

	// Gather the multiple URLs
	urls := []string{}
	urls = append(urls, os.Args[2:]...)

	fmt.Println("selected events: ", selectedEvents[0].EventType)
	fmt.Println("Sending events to:", urls)

	fmt.Println("CSNF event gen started...")
	for {
		time.Sleep(2 * time.Second)

		randomness := rand.Intn(len(selectedEvents))
		for _, url := range urls {
			req, err := http.NewRequest("POST", url, strings.NewReader(selectedEvents[randomness].EventData))
			if err != nil {
				log.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Ce-Type", selectedEvents[randomness].EventType)
			req.Header.Set("Ce-Id", "536808d3-88be-4077-9d7a-a3f162705f79")
			req.Header.Set("Ce-Specversion", "1.0")
			req.Header.Set("Ce-Source", "CSNFEventGenerator")

			//send event
			_, err = http.DefaultClient.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("{%v} Event %v sent to: %v \n", time.Now(), selectedEvents[randomness].EventType, url)
		}

	}

}

func displayHelp() {
	fmt.Println("Usage: eventgen <event type> <url>")
	fmt.Println("Example: eventgen Azure http://localhost:8081/")
	fmt.Println("Supported event types: Azure, Oracle, Aquasec")
	fmt.Println("You may also pass multiple URLs, separated by spaces")
	fmt.Println("Example: eventgen Azure http://localhost:8081/ http://localhost:8082/")
}
