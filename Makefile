#!/bin/bash

generate-proto:
	buf generate

verify-proto:
	buf lint

format-proto:
	buf format -w

build:
	rm -rf ./gen
	buf format -w
	buf lint
	buf generate
