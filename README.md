# Mixer

Mixer is a MySQL proxy with Golang, aims to supply a simple solution for using MySQL.

## Featrues

- Supports base MySQL 
- Splits read and write.
- MySQL HA, switchs backup automatically if main crashed

## Todo

- Routes SQL to multi MySQL server with specified rule.
- Statistics.

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
 
## Doc

[My Chinese Blog](http://blog.csdn.net/siddontang/article/category/2093877)

## Feedback

Email: siddontang@gmail.com


