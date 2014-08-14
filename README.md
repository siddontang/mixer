# Mixer

Mixer is a MySQL proxy powered by Go, aims to supply a simple solution for using MySQL.

## Featrues

- Supports base MySQL (select, insert, update, replace, delete).
- Splits read and write (not full test).
- MySQL HA, switchs backup automatically if main crashed (not full test).
- Base SQL Routing.

## Todo

- admin commands.
- some show command support: show databases, etc.
- some select system variable: select @@version, etc.
- Enhance Routing Rule.
- SQL validation check. 
- Statistics.
- Prepare statement.
- Lots of ......

## Install 

    cd $WORKSPACE
    git clone git@github.com:siddontang/mixer.git src/github.com/siddontang/mixer
    
    cd src/github.com/siddontang/mixer

    ./bootstrap.sh

    . ./dev.env

    make
    make test

## Keywords

### proxy

Proxy is the bridge connecting clients and the read mysql servers. 

It acts as a mysql server too, clients can communicate with it using mysql procotol.

### node

Mixer uses node to represent the real remote mysql server. A node can have three mysql servers:

+ master, main mysql server, all write operations, read operations (if ```rw_split``` and slave not set) will be executed here.
All transactions will be executed here too.
+ master backup, if master was down, mixer can switch over to the backup mysql server. 
+ slave, if ```rw_split``` is set, any select operations will be executed here.

You can only set master for a node to use.

Notice:

+ You can use ```admin upnode``` or ```admin downnode``` commands to up or down specified mysql server.
+ If master was down,  you must use admin command to up it manually.
+ You must set the mysql replication for yourself, mixer do not care it, maybe will support it later.

### schema

Schema likes mysql database, if a client executes ```use db``` command, ```db``` must be exists in schema.

A schema contains one or more nodes, if a client use the specified schema, any command will be only routed to the node which belongs to the schema to be executed.

### rule

You must set some roules for a schema to let mixer decide how to route the sql to different node to be executed.

Mixer uses ```table + key``` to route. Duplicate rule for a table is not allowed.

Rule has three types: default, hash and range.

A schema must have a default rule with only one node assigned. Any sql which can not be routed will use the default rule.

For hash and range routing you can see the example below.

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
hash algorithm: value % len(nodes)

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
range algorithm: node key start <= value < node key stop

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

## Limitation

### Select

+ join not support
+ sub select not support

### Insert

+ insert select not support
+ multi insert values to diff nodes not support
+ insert on duplicate update set can not set the routing key

### Replace

+ multi replace values to diff nodes not support

### Set

+ set autocommit support
+ set name charset support

## Feedback

Email: siddontang@gmail.com


