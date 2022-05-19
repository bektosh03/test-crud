#!/bin/bash
export GOPTIVATE="github.com/bektosh03/test-crud/*"
go1.18.2 get -u github.com/bektosh03/test-crud/genprotos@latest
go1.18.2 get -u github.com/bektosh03/test-crud/common@latest
go1.18.2 mod tidy
go1.18.2 mod vendor