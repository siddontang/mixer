# Mixer

Mixer is a MySQL proxy powered by Go, aims to supply a simple solution for using MySQL.

## Featrues

- Supports base MySQL (select, insert, update, replace, delete).
- Supports transaction.
- Splits read and write (not full test).
- MySQL HA, switchs backup automatically if main crashed (not full test).
- Base SQL Routing.

## Todo

- Some admin commands.
- Some show command support: ```show databases```, etc.
- Some select system variable: ```select @@version```, etc.
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

+ master: main mysql server, all write operations, read operations (if ```rw_split``` and slave not set) will be executed here.
All transactions will be executed here too.
+ master backup: if master was down, mixer can switch over to the backup mysql server. 
+ slave: if ```rw_split``` is set, any select operations will be executed here.

You can only set master for a node to use.

Notice:

+ You can use ```admin upnode``` or ```admin downnode``` commands to up or down specified mysql server.
+ If master was down,  you must use admin command to up it manually.
+ You must set the mysql replication for yourself, mixer does not care it.

### schema

Schema likes mysql database, if a client executes ```use db``` command, ```db``` must be exists in schema.

A schema contains one or more nodes, if a client use the specified schema, any command will be only routed to the node which belongs to the schema to be executed.

### rule

You must set some roules for a schema to let mixer decide how to route the sql to different node to be executed.

Mixer uses ```table + key``` to route. Duplicate rule for a table is not allowed. 

When a sql needs to be routed, mixer does below step:

+ Parse sql and find the operated table.
+ No rule for the table, mixer use default rule.
+ Rule exists, mixer try to route it with the specified key. 

Rule has three types: default, hash and range.

A schema must have a default rule with only one node assigned. 

For hash and range routing you can see the example below.

## Base Example

```
#start mixer
mixer-proxy -config=/etc/mixer.conf

#another shell
mysql -uroot -h127.0.0.1 -P4000 -p -Dmixer

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
hash rule:
-   db: mixer
    table: mixer_test_shard_hash
    key: id
    nodes: node2,node3
    type: hash

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
range rule:
-   db: mixer
    table: mixer_test_shard_range
    key: id
    range: 
    nodes: node2, node3
    range: -10000-
    type: range

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
+ "group by" not support
+ "order by" only takes effect when "order by" key exists in select expression field.
    
    ```select id from t1 order by id``` is ok.
    
    ```select str from t1 order by id``` is not ok, mixer does not known how to sort because it can not find proper data to compare with id.

+ limit should be used with "order by", otherwise you may receive wrong result.  

### Insert

+ insert select not support
+ multi insert values to diff nodes not support
+ insert on duplicate update set can not set the routing key

### Replace

+ multi replace values to diff nodes not support

### Update

+ update set can not set the routing key

### Set

+ set autocommit support
+ set name charset support

### Range Rule

+ only support int64 number range now

## Caveat

+ Mixer uses 2PC to handle write operations for multi nodes, you may take the risk that data corrupted if some nodes commit ok but others error. In that case, you must try to recover your data by yourself.
+ You must design your routing rule and write sql carefully. (e.g. if your where condition contains no routing key, mixer will route the sql to all nodes, maybe).

## Why not vitess?

Vitess is very awesome, and I use some of its codes like sqlparser too. Why not use vitess directly? Maybe below:

+ Vitess is too huge for me, I need a simple proxy.
+ Vitess uses bson protocol, I want to use MySQL protocol.
+ Most likely, something has gone wrong in my head.

## Status

Mixer now is still in development and can not be used in production. 

## Feedback

Email: siddontang@gmail.com


