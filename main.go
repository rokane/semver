package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	incOpt := flag.String("i", "", "must be one of ['major', 'minor', 'patch']")
	preidOpt := flag.String("preid", "", "identifier to be used for pre increments")
	buildOpt := flag.String("build", "", "build metadata to be used for version")
	rangeOpt := flag.String("r", "", "prints the versions which match the specified range")
	flag.Parse()
	controller, err := ParseFlags(*incOpt, *preidOpt, *rangeOpt, *buildOpt)
	if err != nil {
		fmt.Printf("Error: '%s'\n", err.Error())
		os.Exit(1)
	}

	versions := make([]Version, flag.NArg())
	for i, arg := range flag.Args() {
		var err error
		if versions[i], err = ParseVerionString(arg); err != nil {
			fmt.Printf("Error: '%s'\n", err.Error())
			os.Exit(1)
		}
	}

	if *incOpt != "" && len(versions) == 1 {
		v, err := controller.Increment(versions[0], controller.IncrementLevel)
		if err != nil {
			fmt.Printf("Error: '%s'\n", err.Error())
			os.Exit(1)
		}
		fmt.Println(v.String())
	}
}

// ParseFlags validates the args captured via command line flags and stores
// in a UtilController for further processing
func ParseFlags(increment, preid, rng, build string) (UtilController, error) {
	controller := UtilController{}
	if lvl, err := ParseIncrementLevel(increment); increment != "" {
		if err != nil {
			return UtilController{}, err
		}
		controller.IncrementLevel = lvl
	}
	if preid != "" && validateMetaData(preid) {
		controller.PreIdentifier = preid
	}
	if build != "" && validateMetaData(build) {
		controller.BuildMeta = build
	}
	if err := ParseRange(rng); rng != "" && err == nil {
		controller.Range = rng
	}
	return controller, nil
}

// ParseIncrementLevel returns an error if an invalid increment value was parsed
func ParseIncrementLevel(level string) (Release, error) {
	valid := map[string]Release{"major": Major, "minor": Minor, "patch": Patch,
		"premajor": PreMajor, "preminor": PreMinor, "prepatch": PrePatch}
	if val, ok := valid[level]; ok {
		return val, nil
	}
	return 0, fmt.Errorf(
		"Increment value: '%s' not in: ['major', 'minor', 'patch', 'premajor', 'preminor', 'prepatch']", level)
}

// ParseRange -
func ParseRange(rng string) error {
	return nil
}

// ParseVerionString validates the string follows semantic versioning format
// and returns a Version object containing the information
func ParseVerionString(str string) (Version, error) {
	validRe := regexp.MustCompile("^([0-9]|[1-9][0-9]?).([0-9]|[1-9][0-9]?).([0-9]|[1-9][0-9]?)$")
	if !validRe.MatchString(str) {
		return Version{}, fmt.Errorf("Unable to parse string '%s' to version", str)
	}
	var err error
	levels := make([]int, 3)
	for i, s := range strings.Split(str, ".") {
		if levels[i], err = strconv.Atoi(s); err != nil {
			return Version{}, fmt.Errorf("Error parsing component: '%s'. Error: '%s'", s, err.Error())
		}
	}
	return Version{Major: levels[0], Minor: levels[1], Patch: levels[2]}, nil
}

// Validates the metadata contains dot separated alphanumeric chars only
func validateMetaData(s string) bool {
	validRe := regexp.MustCompile(`^([0-9A-Za-z-]*\.?)*[^\.]$`)
	return validRe.MatchString(s)
}
