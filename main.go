package main

import (
	"fmt"

	api_master "github.com/aler9/goroslib/apimaster"
)

func main() {
	host := "localhost"
	port := 11311
	callerID := "/client"
	client := api_master.NewClient(host, port, callerID)
	res, err := client.GetSystemState()
	if err != nil {
		panic(err)
	}

	for _, entry := range res.State.PublishedTopics {
		fmt.Println(entry.Name)
	}

	for _, entry := range res.State.SubscribedTopics {
		fmt.Println(entry.Name)
	}
}
