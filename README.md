# Lock

## A CLI helper tool for Morpheus plugin development and download

[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/spoonboy-io/lock?style=flat-square)](https://go.dev/)
[![Go Report Card](https://goreportcard.com/badge/github.com/spoonboy-io/lock?style=flat-square)](https://goreportcard.com/report/github.com/spoonboy-io/lock)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/spoonboy-io/lock/build.yml?branch=main&style=flat-square)](https://github.com/spoonboy-io/lock/actions/workflows/build.yml)
[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/spoonboy-io/lock/unit_test.yml?branch=main&label=tests&style=flat-square)](https://github.com/spoonboy-io/lock/actions/workflows/unit_test.yml)
[![GitHub Release Date](https://img.shields.io/github/release-date/spoonboy-io/lock?style=flat-square)](https://github.com/spoonboy-io/lock/releases)
[![GitHub commits since latest release (by date)](https://img.shields.io/github/commits-since/spoonboy-io/lock/latest?style=flat-square)](https://github.com/spoonboy-io/lock/commits)
[![GitHub](https://img.shields.io/github/license/spoonboy-io/lock?label=license&style=flat-square)](LICENSE)

## About

Lock is a CLI tool to assist with Morpheus plugin development and download of prebuilt plugin JAR files.

Users can inspect and clone Morpheus plugin starter templates, tagged by Morpheus tested version or semantic version.
Users can also inspect and download prebuilt JAR plugins at a specific version and compatibility.

Lock currently uses [share.morpheusdata.com](https://share.morpheusdata.com) for information about prebuilt plugins and a [git repository here](https://github.com/spoonboy-io/lock-plugin-metadata)
to maintain starter template project information.

### Releases

You can find the [latest software here](https://github.com/spoonboy-io/lock/releases/latest).

### Features

#### Starter Templates
 - List Morpheus plugin starter templates.
 - Inspect template and its version and Morpheus compatability.
 - Clone a starter template into a new project folder ready for customisation.

#### Plugins
 - List prebuilt plugins available in JAR format.
 - Inspect plugin, its version history and compatability.
 - Download a plugin by version or latest.

### Usage

```bash
➜  Demo lock

Usage: lock COMMAND <ARGUMENT> --FLAG(S)

A CLI helper tool for Morpheus plugin development and download

General commands:	
  help                   Show this help section
  version                Show the Lock version information

Starter templates:
  templates              List available starter project templates.
                          --category    The category as a filter
                          --morpheus    The minimum version of Morpheus as a filter
  template <TEMPLATE>    Show template information and available versions, <TEMPLATE> can be name or id.
  pick <TEMPLATE>        Creates new project from a starter template, <TEMPLATE> can be name or id.
                          --name        A project folder name (default: morpheus-plugin-project)
                          --tag         The specific tag to create project from (default: head)

Precompiled plugins:
  plugins                List available Morpheus plugin JAR files.
                          --morpheus    The minimum version of Morpheus as a filter
  plugin <PLUGIN>        Show plugin information and Morpheus version compatibility, <PLUGIN> can be name or id.
  load <PLUGIN>          Download a compiled plugin, <PLUGIN> can be name or id.
                          --version     The specific version of the plugin (default: latest)
```

### Installation
Grab the tar.gz or zip archive for your OS from the [releases page](https://github.com/spoonboy-io/lock/releases/latest).
Unpack it and run:

```
./lock
```
Make lock available in your $PATH to omit the leading `./`

### License
Licensed under [Mozilla Public License 2.0](LICENSE)

### Todo/Ideas
- plugins --morpheus= flag is not yet implemented
- Java version detection
- Watcher functionality to build plugin on file save