package main

// Release represents a level in a Version
type Release uint8

// Levels ranked in order of precedence
const (
	Major Release = iota // Major Version (X.y.z)
	Minor                // Minor Version (x.Y.z)
	Patch                // Patch Version (x.y.Z)
	PreMajor
	PreMinor
	PrePatch
)

// Version holds details about the version of a single tag.
type Version struct {
	Major     int
	Minor     int
	Patch     int
	PreID     string
	BuildMeta string
}
