[![CircleCI](https://circleci.com/gh/rokane/semvar-tools/tree/master.svg?style=svg&circle-token=fb912370a98970dfb454dd707e425de968979f27)]
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/gojp/goreportcard/blob/master/LICENSE)

# Semver

A golang Semantic Versioning tool that follows the [Semver Specification](https://semver.org/)

This tool can be used as a command line utility or as a golang package for use within other projects.

## Usage

As a command line utility:

```bash
A Golang implementation of the http://semver.org/ specification

Usage: 
semver [options] <version> [<version> [...]]
Prints valid versions sorted by SemVer precedence

Options:
-r --range <range>
        Print versions that match the specified range.

-i --increment [<level>]
        Increment a version by the specified level.  Level can
        be one of: major, minor, patch, premajor, preminor,
        prepatch, or prerelease.  Default level is 'patch'.
        Only one version may be specified.

--preid <identifier>
        Identifier to be used to prefix premajor, preminor,
        prepatch or prerelease version increments.

-l --loose
        Interpret versions and ranges loosely

-c --coerce
        Coerce a string into SemVer if possible
        (does not imply --loose)

Program exits successfully if any valid version satisfies
all supplied ranges, and prints all satisfying versions.

If no satisfying versions are found, then exits failure.

Versions are printed in ascending order, so supplying
multiple versions to the utility will just sort them.

```



# Design Document

The aim is to have either a standalone or set of executable programs which will help to do semantic versioning on the current git repository.

## Configuring the Application

One thought might be to have the user provide a set of configurations in the form of a json file. Think .gometalinter.json file where the user can specify the set of linters they want to parse over the current directories go files. Potentially the same concept could occur where we pass any configuration values:

```json
{
  "initial_version": 0.0.1,
  "git": {
    ...
  },
  "docker": {
    "enable": true,
    "organisation": "rokane",
    "image_name": "semvar"
  }
}
```

## Semantic Versioning Tasks

 1. Version Bumping
      * The application should be able to determine the current version from the git repository we are sitting in and perform a version bump. The version bump will increase the tag sitting alongside the git repository and push it upstream.
        * Another option would be to also bump tag on the related Docker image. 
 2. Initialize Version Tag
      * Perhaps this is something which you would like to occur when attempting to bump the program not finding any existing semantic tags.
 3. Validate Existing Versions
      * User may have an existing git repository which they want to validate the set of semantic tags currently existing to see if there are any inconsistencies amongst the tags. i.e. Missed versions etc.
  