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
$ rsysreqs -h
Usage of rsysreqs:
  -d string
        use rules from this directory (default "../rsysreqs-db/sysreqs/")
  -s string
        system requirements (default "libXML2, curl; pkgA")

$ rsysreqs-server -h
Usage of rsysreqs-server:
  -d string
        use rules from this directory (default "../rsysreqs-db/sysreqs/")
```
