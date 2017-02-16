#!/bin/bash

DIST=dist
RELEASE=release

# remove dist
[[ -d $DIST ]] && rm -rf $DIST
[[ -d $RELEASE ]] && rm -rf $RELEASE

# build to dist
mkdir $DIST
go build -o $DIST/parsing index.go

# release
mkdir $RELEASE
cp $DIST/* $RELEASE
