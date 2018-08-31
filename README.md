[![CircleCI](https://circleci.com/gh/rokane/semvar-tools/tree/master.svg?style=svg&circle-token=fb912370a98970dfb454dd707e425de968979f27)]
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/gojp/goreportcard/blob/master/LICENSE)

# Semvar Tools

Repository contains a collection of resources, which intend to simlpify the semantic versioning of projects.

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
  