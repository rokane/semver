package main

import (
	"fmt"
	"strconv"
)

// IncrementMajor increments the Major level of the calling Version, resetting
// the minor and patch as a result.
func (v *Version) IncrementMajor() {
	v.Major++
	v.Minor = 0
	v.Patch = 0
}

// IncrementMinor increments the Minor level of the calling Version, which
// maintains the major level, but resets the patch level as a result.
func (v *Version) IncrementMinor() {
	v.Minor++
	v.Patch = 0
}

// IncrementPatch increments the Patch level of the calling Version, which
// maintains the major and minor levels as they previously were.
func (v *Version) IncrementPatch() {
	v.Patch++
}

// SetMetaData updates the calling versions build metadata
func (v *Version) SetMetaData(buildmeta string) {
	v.BuildMeta = buildmeta
}

// SetPreID updates the calling versions preid
func (v *Version) SetPreID(preid string) {
	v.PreID = preid
}

// String returns a formatted string of the calling version.
func (v *Version) String() string {
	str := fmt.Sprintf("%s.%s.%s", strconv.Itoa(v.Major),
		strconv.Itoa(v.Minor), strconv.Itoa(v.Patch))
	if v.PreID != "" {
		str = fmt.Sprintf("%s-%s", str, v.PreID)
	}
	if v.BuildMeta != "" {
		str = fmt.Sprintf("%s+%s", str, v.BuildMeta)
	}
	return str
}
