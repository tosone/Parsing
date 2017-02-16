#!/bin/bash

# submodule update
git submodule init
git submodule update
git submodule foreach git pull origin master
