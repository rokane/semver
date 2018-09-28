package main

import (
	"testing"

	tassert "github.com/stretchr/testify/assert"
)

func TestControllerIncrementsVersion(t *testing.T) {
	assert := tassert.New(t)
	v := Version{Major: 1, Minor: 0, Patch: 0}
	controller := UtilController{
		IncrementLevel: Major,
		PreIdentifier:  "alpha.1",
		BuildMeta:      "123456789",
	}

	expectedStr := "2.0.0+123456789"
	incremented, err := controller.Increment(v, controller.IncrementLevel)
	assert.NoError(err)
	assert.Equal(v.Major+1, incremented.Major)
	assert.Equal(expectedStr, incremented.String())

	controller.IncrementLevel = Minor
	expectedStr = "1.1.0+123456789"
	incremented, err = controller.Increment(v, controller.IncrementLevel)
	assert.NoError(err)
	assert.Equal(v.Minor+1, incremented.Minor)
	assert.Equal(expectedStr, incremented.String())

	controller.IncrementLevel = Patch
	expectedStr = "1.0.1+123456789"
	incremented, err = controller.Increment(v, controller.IncrementLevel)
	assert.NoError(err)
	assert.Equal(v.Patch+1, incremented.Patch)
	assert.Equal(expectedStr, incremented.String())

	controller.IncrementLevel = PreMajor
	expectedStr = "2.0.0-alpha.1+123456789"
	incremented, err = controller.Increment(v, controller.IncrementLevel)
	assert.NoError(err)
	assert.Equal(v.Major+1, incremented.Major)
	assert.Equal(expectedStr, incremented.String())

	controller.IncrementLevel = PreMinor
	expectedStr = "1.1.0-alpha.1+123456789"
	incremented, err = controller.Increment(v, controller.IncrementLevel)
	assert.NoError(err)
	assert.Equal(v.Minor+1, incremented.Minor)
	assert.Equal(expectedStr, incremented.String())

	controller.IncrementLevel = PrePatch
	expectedStr = "1.0.1-alpha.1+123456789"
	incremented, err = controller.Increment(v, controller.IncrementLevel)
	assert.NoError(err)
	assert.Equal(v.Patch+1, incremented.Patch)
	assert.Equal(expectedStr, incremented.String())

	controller.PreIdentifier = ""
	incremented, err = controller.Increment(v, controller.IncrementLevel)
	assert.Error(err)
	controller.IncrementLevel = PreMajor
	incremented, err = controller.Increment(v, controller.IncrementLevel)
	assert.Error(err)
	controller.IncrementLevel = PreMinor
	incremented, err = controller.Increment(v, controller.IncrementLevel)
	assert.Error(err)

}
