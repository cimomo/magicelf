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
		time.Date(2010, time.September, 1, 0, 0, 0, 0, time.UTC),
		3200,
		"1000 Awesome Street NE, Awesome City, WA 98007",
	},
}
