package main

import (
	"testing"

	tassert "github.com/stretchr/testify/assert"
)

func TestParseVersionString(t *testing.T) {
	assert := tassert.New(t)

	validStr := "1.2.3"
	expVersion := Version{Major: 1, Minor: 2, Patch: 3}
	v, err := ParseVerionString(validStr)
	assert.NoError(err)
	assert.Equal(expVersion, v)

	invalidStr := "a.b.c"
	_, err = ParseVerionString(invalidStr)
	assert.Error(err)
}

func TestValidateMetaData(t *testing.T) {
	assert := tassert.New(t)
	assert.True(validateMetaData("alpha.123"))
	assert.True(validateMetaData("aoms-dsoj.9283429.shfj"))
	assert.False(validateMetaData("hsghgdf."))
	assert.False(validateMetaData("&sdjshb"))
	assert.False(validateMetaData("&#$%^&*"))
}

func TestParseIncrementLevel(t *testing.T) {
	assert := tassert.New(t)

	r, err := ParseIncrementLevel("major")
	assert.NoError(err)
	assert.Equal(Major, r)

	r, err = ParseIncrementLevel("minor")
	assert.NoError(err)
	assert.Equal(Minor, r)

	r, err = ParseIncrementLevel("patch")
	assert.NoError(err)
	assert.Equal(Patch, r)

	r, err = ParseIncrementLevel("premajor")
	assert.NoError(err)
	assert.Equal(PreMajor, r)

	r, err = ParseIncrementLevel("preminor")
	assert.NoError(err)
	assert.Equal(PreMinor, r)

	r, err = ParseIncrementLevel("prepatch")
	assert.NoError(err)
	assert.Equal(PrePatch, r)

	_, err = ParseIncrementLevel("something")
	assert.Error(err)
}

func TestParseFlags(t *testing.T) {
	assert := tassert.New(t)

	contr, err := ParseFlags("invalid", "alpha", "rng", "metadata")
	assert.Error(err)

	contr, err = ParseFlags("major", "alpha", "range", "meta")
	assert.NoError(err)
	assert.Equal("meta", contr.BuildMeta)
	assert.Equal(Major, contr.IncrementLevel)
	assert.Equal("alpha", contr.PreIdentifier)
}
