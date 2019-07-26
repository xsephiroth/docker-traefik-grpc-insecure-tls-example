#!/bin/bash

docker run --rm -it \
    -v ${PWD}:/defs \
    namely/protoc-all \
    -f traefikgrpc.proto \
    -l go \
    -o /defs
