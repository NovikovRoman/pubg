# PUBG
[![Go Reference](https://pkg.go.dev/badge/github.com/NovikovRoman/pubg.svg)](https://pkg.go.dev/github.com/NovikovRoman/pubg)
[![Build Status](https://travis-ci.com/NovikovRoman/pubg.svg?branch=master)](https://travis-ci.com/NovikovRoman/pubg)
[![Go Report Card](https://goreportcard.com/badge/github.com/NovikovRoman/pubg)](https://goreportcard.com/report/github.com/NovikovRoman/pubg)

> This is a Go client library for gathering [PUBG API reference] data

Relevant for [V21.3.0](https://documentation.pubg.com/en/changelog/changelog.html)

## Table of Contents

- [PUBG](#pubg)
  - [Table of Contents](#table-of-contents)
  - [Getting Started](#getting-started)
  - [Usage Example](#usage-example)
  - [Tests](#tests)
  - [Documentation](#documentation)

## Getting Started

Download the PUBG library:

```shell
go get github.com/NovikovRoman/pubg
```

## Usage Example

```go
pubgClient := pubg.NewClient(apikey, nil)

if !pubgClient.Status() {
	log.Fatalln("PUBG API not working.")
}

tournaments, err := pubgClient.Tournaments()
if err != nil {
	if e, ok := err.(*pubg.ErrBadRequest); ok {
		log.Fatalln("Bad request", e, e.GetDetail())

	} else if e, ok := err.(*pubg.ErrUnauthorized); ok {
		log.Fatalln("Unauthorized", e, e.GetDetail())

	} else if e, ok := err.(*pubg.ErrNotFound); ok {
		log.Fatalln("Not found", e, e.GetDetail())

	} else if e, ok := err.(*pubg.ErrUnsupportedMediaType); ok {
		log.Fatalln("Unsupported media type", e, e.GetDetail())

	} else if e, ok := err.(*pubg.ErrTooManyRequest); ok {
		log.Fatalln("Too many request", e, e.GetDetail())
	}

	log.Fatalln(err)
}

for _, t := range tournaments.Data {
	fmt.Println(t.ID, t.Attributes.CreatedAt)
}
```

See the example of using telemetry in the `telemetry_test.go`.

## Tests

```shell
PAUSE=true [PROXY=…] APIKEY=… go test
```

## Documentation

See the [PUBG API reference].

[PUBG API reference]: https://documentation.pubg.com/en/introduction.html