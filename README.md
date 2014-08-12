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

## Base Example

```
#start mixer
mixer-proxy -config=/etc/mixer.conf

#another shell
mysql -P4000 -uroot

Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 158
Server version: 5.6.19 Homebrew

mysql> use mixer;
Database changed

mysql> delete from mixer_test_conn;
Query OK, 3 rows affected (0.04 sec)

mysql> insert into mixer_test_conn (id, str) values (1, "a");
Query OK, 1 row affected (0.00 sec)

mysql> insert into mixer_test_conn (id, str) values (2, "b");
Query OK, 1 row affected (0.00 sec)

mysql> select id, str from mixer_test_conn;
+----+------+
| id | str  |
+----+------+
|  1 | a    |
|  2 | b    |
+----+------+
``` 

## Feedback

Email: siddontang@gmail.com


