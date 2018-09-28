package main

import (
	"fmt"
)

// UtilController is an implementation of the Utility interface, which
// holds information captured via the command line flags
type UtilController struct {
	IncrementLevel Release
	PreIdentifier  string
	BuildMeta      string
	Range          string
}

// Increment return a version which is one level higher in terms of precendence
// defined by the release version specified.
func (uc *UtilController) Increment(v Version, l Release) (Version, error) {
	v.SetMetaData(uc.BuildMeta)
	if l == PreMajor || l == PreMinor || l == PrePatch {
		if uc.PreIdentifier == "" {
			return Version{}, fmt.Errorf("Require -preid to increment PreMajor")
		}
		v.SetPreID(uc.PreIdentifier)
	}
	if l == Major || l == PreMajor {
		v.IncrementMajor()
	}
	if l == Minor || l == PreMinor {
		v.IncrementMinor()
	}
	if l == Patch || l == PrePatch {
		v.IncrementPatch()
	}
	return v, nil
}
