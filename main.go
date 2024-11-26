package main

import "time"

// TODO: Add type field and add it to the csv file

// struct field to add labels
type Task struct {
	name      string
	id        int
	Completed bool
	Due       time.Time
}
