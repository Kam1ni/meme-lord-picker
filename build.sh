#!/bin/bash

# install miqt-rcc if not available
type miqt-rcc || go install github.com/mappu/miqt/cmd/miqt-rcc@latest

cd $( dirname -- "${BASH_SOURCE[0]}" )/src
miqt-rcc -Input resources.qrc -Qt6
go build .