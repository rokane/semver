[![Build Status](https://travis-ci.com/rokane/semver.svg?token=quzM1nwjxcjsYV2iMeew&branch=master)](https://travis-ci.com/rokane/semver)
[![Coverage Status](https://coveralls.io/repos/github/rokane/semver/badge.svg?branch=setup-ci)](https://coveralls.io/github/rokane/semver?branch=setup-ci)
[![Go Report Card](https://goreportcard.com/badge/github.com/rokane/semver)](https://goreportcard.com/report/github.com/rokane/semver)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/gojp/goreportcard/blob/master/LICENSE)

# Semver

A golang Semantic Versioning tool that follows the [Semver Specification](https://semver.org/)

This tool can be used as a command line utility or as a golang package for use within other projects.

## Installation

```bash
go get github.com/rokane/semver
```

## Usage

As a command line utility:

```bash
A Golang implementation of the http://semver.org/ specification

Usage: 
semver [options] <version> [<version> [...]]

Options:
-r <range>
        Print versions that match the specified range.

-i <level>
        Increment a version by the specified level. Level can
        be one of: major, minor, patch, premajor, preminor,
        prepatch, or prerelease. Only one version may be specified.

-preid <identifier>
        Identifier to be used to prefix premajor, preminor,
        prepatch or prerelease version increments.

-build <buildmeta>
        Metadata to be used to attached to the incremented version.
```

As a golang package:
```go
import "github/rokane/semver"

version := Version{
        Major: 1,
        Minor: 0,
        Patch: 0,
        PreID: "preidentifier"
        BuildMeta: "metadata",
}

// Increment 
version.IncrementMajor()
version.IncrementMinor()
version.IncrementPatch()

// Output to String
version.String()
```