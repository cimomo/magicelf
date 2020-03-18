package main

import "time"

// Property defines a rental property
type Property struct {
	Date    time.Time
	Rent    int
	Address string
}

var properties = []Property{
	Property{
		time.Now(),
		3200,
		"1000 Awesome Street NE, Awesome City, WA 98007",
	},
}
