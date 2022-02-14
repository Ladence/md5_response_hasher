# md5_response_hasher

## Summary

This repo contains code for CLI utility to hash responses from HTTP requests for provided URLs.
Utility can handle multiple request at the same time. Number of concurrent requests can be providen by `paralel` option (by default == 10).

## Usage

`./myhttp [-parallel] urls ...`

## Build

This repo supports Makefile which is used for building and testing

So for building you can just enter

`make build`

and for test

`make test`
