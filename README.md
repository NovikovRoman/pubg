# PUBG

> This is a Go client library for gathering [PUBG API reference] data

## Table of Contents

- [PUBG](#pubg)
  - [Table of Contents](#table-of-contents)
  - [Getting Started](#getting-started)
  - [Usage Example](#usage-example)
  - [Documentation](#documentation)

## Getting Started

Download the PUBG library:

```shell
go get github.com/NovikovRoman/pubg
```

## Usage Example

```go
pubgClient := pubg.NewClient(apikey)

if !pubgClient.Status() {
    log.Fataln("PUBG API not working.")
}

tournaments, err := pubgClient.Tournaments()
if err != nil {
    if e, ok := err.(*ErrBadRequest); ok {
		log.Fatalln("Bad request", e, e.GetDetail())

	} else if e, ok := err.(*ErrUnauthorized); ok {
		log.Fatalln("Unauthorized", e, e.GetDetail())

	} else if e, ok := err.(*ErrNotFound); ok {
		log.Fatalln("Not found", e, e.GetDetail())

	} else if e, ok := err.(*ErrUnsupportedMediaType); ok {
		log.Fatalln("Unsupported media type", e, e.GetDetail())

	} else if e, ok := err.(*ErrTooManyRequest); ok {
		log.Fatalln("Too many request", e, e.GetDetail())
	}

    log.Fatalln(e)
}

for _, t := range tournaments.Data {
    fmt.Println(t.ID, t.Attributes.CreatedAt)
}
```

See the example of using telemetry in the `telemetry_test.go`.

## Documentation

See the [PUBG API reference].

[PUBG API reference]: https://documentation.pubg.com/en/introduction.html