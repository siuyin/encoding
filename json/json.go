package json

import (
	"encoding/json"
	"log"
)

// Log writes out a log line with json encoding.
func Log(i interface{}) {
	js, err := json.Marshal(i)
	if err != nil {
		log.Printf("ERROR: json marshal: %v", err)
	}
	log.Printf("%s", js)
}
