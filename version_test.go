package main

import (
	"testing"

	tassert "github.com/stretchr/testify/assert"
)

func TestIncrementVersion(t *testing.T) {
	assert := tassert.New(t)
	v := Version{Major: 1, Minor: 0, Patch: 0}

	old := v
	v.IncrementMajor()
	assert.Equal(old.Major+1, v.Major)
	v.IncrementMinor()
	assert.Equal(old.Minor+1, v.Minor)
	v.IncrementPatch()
	assert.Equal(old.Patch+1, v.Patch)
}

func TestVersionString(t *testing.T) {
	assert := tassert.New(t)
	v := Version{Major: 1, Minor: 0, Patch: 0}
	preid := "alpha.1"
	buildmeta := "123456789"

	v.SetMetaData(buildmeta)
	assert.Equal(buildmeta, v.BuildMeta)
	v.SetPreID(preid)
	assert.Equal(preid, v.PreID)

	expectedStr := "1.0.0-alpha.1+123456789"
	s := v.String()
	assert.Equal(expectedStr, s)
}
