package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"
)

type Subsc struct {
	Index string `json:"Index"`
	Size  int    `json:"Size"`
	Stamp string `json:"Stamp"`
}

func main() {

	// Open file instead of device:
	ff, err := os.Open("mrTest.csv")
	if err != nil {
		panic(err)
	}
	defer ff.Close()

	// For map-red, we need a CSV file... for now.
	// Read file into a variable
	lines, err := csv.NewReader(ff).ReadAll()
	if err != nil {
		panic(err)
	}

	// Create the first mapper list
	lists := make(chan []Subsc)
	finalValue := make(chan []Subsc)
	var wg sync.WaitGroup
	// Mapping
	wg.Add(len(lines))

	for _, line := range lines {
		go func(secData []string) {
			defer wg.Done()
			lists <- Map(secData)
		}(line)
	}

	go Reducer(lists, finalValue)

	wg.Wait()
	close(lists)
	fmt.Println(<-finalValue)

}

func Map(secData []string) []Subsc {
	list := []Subsc{}
	size, _ := strconv.Atoi(secData[2])
	list = append(list, Subsc{
		Index: secData[1],
		Size:  size,
		Stamp: secData[2],
	})
	return list
}

func Reducer(mapList chan []Subsc, sendFinalValue chan []Subsc) {
	final := []Subsc{}
	for list := range mapList {
		for _, value := range list {
			// process here

			//  Switch Stamps here:
			if value.Size <= 0 {
				final = append(final, value)
			}
		}
	}
	sendFinalValue <- final
}
