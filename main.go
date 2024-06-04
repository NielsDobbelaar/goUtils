package main

import (
	"errors"
	"goUtils/assert"
)

type ValueSnapshot struct {
	sensorType string
	value      float64
}

type PlantSnapshot struct {
	PlantID int
	Values  ValueSnapshot
}

func main() {
	assert.Assert(true, "This is a test")
	// assert.Assert(false, "This is the second test")
	// assert.AssertWithError(false, "assertWith error test", errors.New("Error string"))
	valueSnapshot := ValueSnapshot{sensorType: "Temperature", value: 25.0}
	plantSnapshot := PlantSnapshot{PlantID: 1, Values: valueSnapshot}

	assert.LogContext(
		assert.ErrorContext{
			Name:  "plantSnapshot",
			Value: plantSnapshot})
	assert.AssertNoFatal("Error creating object", errors.New("Error with Name value"))
	assert.AssertWithErrorAndContext(false,
		"Error in creating object",
		errors.New("Error string"),
		assert.ErrorContext{Name: "plantSnapshot", Value: plantSnapshot},
		assert.ErrorContext{Name: "plantID", Value: 1})
}
