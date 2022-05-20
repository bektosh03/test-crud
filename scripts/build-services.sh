#!/bin/bash
cd data-service && make build-docker && cd ..
cd post-service && make build-docker && cd ..
cd api-gateway && make build-docker && cd ..