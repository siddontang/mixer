# Mixer

Mixer is a MySQL proxy with Golang, aims to supply a simple solution for using MySQL.

## Featrues

- Supports base MySQL 
- Splits read and write.
- MySQL HA, switchs backup automatically if main crashed
- Base SQL Routing

## Todo

- Enhance Routing Rule.
- SQL validation check. 
- Statistics.
- Prepare statement.

## Install 

    cd $WORKSPACE
    git clone git@github.com:siddontang/mixer.git src/github.com/siddontang/mixer
    
    cd src/github.com/siddontang/mixer

    ./bootstrap.sh

    . ./dev.env

    make
    make test

## Usage

Add later......
 

## Feedback

Email: siddontang@gmail.com


