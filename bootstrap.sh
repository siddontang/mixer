#!/bin/bash

if [ ! -f bootstrap.sh ]; then
  echo "bootstrap.sh must be run from its current directory" 1>&2
  exit 1
fi

source ./dev.env

go get github.com/siddontang/golib/log