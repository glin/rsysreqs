# rsysreqs

[![Build Status](https://travis-ci.org/glin/rsysreqs.svg?branch=master)](https://travis-ci.org/glin/rsysreqs)

## Quick Start

```sh
# Run tests and build binaries
make

# Run tests
make test
```

### Usage

```sh
$ rsysreqs
Usage of rsysreqs:
  -arch string
        architecture
  -d string
        use rules from this directory
  -dist string
        distribution
  -os string
        operating system
  -release string
        release
  -s string
        system requirements
```

```sh
$ rsysreqs-server
Usage of rsysreqs-server:
  -d string
        use rules from this directory
```

## API

### `GET /packages`

Find system packages

#### Parameters

|Name|Type|Required/Optional|Description|
|----|----|-----------------|-----------|
|sysreqs|string|required|system requirements|
|os|string|required|operating system|
|dist|string|required|distribution|
|release|string|optional|release|
|arch|string|optional|architecture|
