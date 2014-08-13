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

## Hash Sharding Example

```
table: mixer_test_shard_hash

Node: node2, node3
node2 mysql: 127.0.0.1:3307
node3 mysql: 127.0.0.1:3308

mixer-proxy: 127.0.0.1:4000

proxy> mysql -uroot -h127.0.0.1 -P4000 -p -Dmixer
node2> mysql -uroot -h127.0.0.1 -P3307 -p -Dmixer
node3> mysql -uroot -h127.0.0.1 -P3307 -p -Dmixer

proxy> insert into mixer_test_shard_hash (id, str) values (0, "a");
node2> select str from mixer_test_shard_hash where id = 0;
+------+
| str  |
+------+
| a    |
+------+

proxy> insert into mixer_test_shard_hash (id, str) values (1, "b");
node3> select str from mixer_test_shard_hash where id = 1;
+------+
| str  |
+------+
| b    |
+------+

proxy> select str from mixer_test_shard_hash where id in (0, 1);
+------+
| str  |
+------+
| a    |
| b    |
+------+

proxy> select str from mixer_test_shard_hash where id = 0 or id = 1;
+------+
| str  |
+------+
| a    |
| b    |
+------+

proxy> select str from mixer_test_shard_hash where id = 0 and id = 1;
Empty set
```


## Range Sharding Example

```
table: mixer_test_shard_range

Node: node2, node3
node2 range: (-inf, 10000)
node3 range: [10000, +inf)
node2 mysql: 127.0.0.1:3307
node3 mysql: 127.0.0.1:3308

mixer-proxy: 127.0.0.1:4000

proxy> mysql -uroot -h127.0.0.1 -P4000 -p -Dmixer
node2> mysql -uroot -h127.0.0.1 -P3307 -p -Dmixer
node3> mysql -uroot -h127.0.0.1 -P3307 -p -Dmixer

proxy> insert into mixer_test_shard_range (id, str) values (0, "a");
node2> select str from mixer_test_shard_range where id = 0;
+------+
| str  |
+------+
| a    |
+------+

proxy> insert into mixer_test_shard_range (id, str) values (10000, "b");
node3> select str from mixer_test_shard_range where id = 10000;
+------+
| str  |
+------+
| b    |
+------+

proxy> select str from mixer_test_shard_range where id in (0, 10000);
+------+
| str  |
+------+
| a    |
| b    |
+------+

proxy> select str from mixer_test_shard_range where id = 0 or id = 10000;
+------+
| str  |
+------+
| a    |
| b    |
+------+

proxy> select str from mixer_test_shard_range where id = 0 and id = 10000;
Empty set

proxy> select str from mixer_test_shard_range where id > 100;
+------+
| str  |
+------+
| b    |
+------+

proxy> select str from mixer_test_shard_range where id < 100;
+------+
| str  |
+------+
| a    |
+------+

proxy> select str from mixer_test_shard_range where id >=0 and id < 100000;
+------+
| str  |
+------+
| a    |
| b    |
+------+
```

## Feedback

Email: siddontang@gmail.com


