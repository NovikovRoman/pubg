# PUBG

> This is a Go client library for gathering [PUBG API reference] data

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
pubgClient := pubg.NewClient(apikey)

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
pause=true APIKEY=[…] go test
```

## Documentation

See the [PUBG API reference].

[PUBG API reference]: https://documentation.pubg.com/en/introduction.html