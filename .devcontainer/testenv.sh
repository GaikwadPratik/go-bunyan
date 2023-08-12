#!/bin/bash

go get github.com/onsi/ginkgo/v2/ginkgo@latest
go install -v github.com/onsi/ginkgo/v2/ginkgo

go mod download -x
go mod tidy