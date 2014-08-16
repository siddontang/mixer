# Mixer

Mixer is a MySQL proxy powered by Go which aims to supply a simple solution for MySQL sharding.

## Features

- Supports basic SQL statements (select, insert, update, replace, delete)
- Supports transactions
- Splits reads and writes (not fully tested)
- MySQL HA, switches to backup automatically if main crashes (not fully tested)
- Basic SQL Routing
- Supports prepared statement: `COM_STMT_PREPARE`, `COM_STMT_EXECUTE`, etc. 

## TODO

- Some admin commands
- Some show command support, i.e. ```show databases```, etc.
- Some select system variable, i.e. ```select @@version```, etc.
- Enhance routing rules
- Monitor
- SQL validation check
- Statistics
- Many other things ...

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

A proxy is the bridge connecting clients and the real MySQL servers. 

It acts as a MySQL server too, clients can communicate with it using the MySQL procotol.

### node

Mixer uses nodes to represent the real remote MySQL servers. A node can have three MySQL servers:

+ master: main MySQL server, all write operations, read operations (if ```rw_split``` and slave are not set) will be executed here.
All transactions will be executed here too.
+ master backup: if the master was down, the mixer can switch over to the backup MySQL server. (can not set)
+ slave: if ```rw_split``` is set, any select operations will be executed here. (can not set)

Notice:

+ You can use ```admin upnode``` or ```admin downnode``` commands to bring a specified MySQL server up or down.
+ If the master was down, you must use an admin command to bring it up manually.
+ You must set up MySQL replication for yourself, mixer does not do it.

### schema

Schema likes MySQL database, if a client executes ```use db``` command, ```db``` must exist in the schema.

A schema contains one or more nodes. If a client use the specified schema, any command will be only routed to the node which belongs to the schema to be executed.

### rule

You must set some rules for a schema to let the mixer decide how to route SQL statements to different nodes to be executed.

Mixer uses ```table + key``` to route. Duplicate rule for a table are not allowed. 

When SQL needs to be routed, mixer does the following steps:

+ Parse SQL and find the table operated on
+ If there are no rule for the table, mixer use the default rule
+ If a rule exists, mixer tries to route it with the specified key

Rules have three types: default, hash and range.

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

## Limitations

### Select

+ Join not supported, later only cross sharding not supported.
+ Subselects not supported, later only cross sharding not supported.
+ Cross sharding "group by" will not work ok only except the "group by" key is the routing key
+ Cross sharding "order by" only takes effect when the "order by" key exists as a select expression field
    
    ```select id from t1 order by id``` is ok.
    
    ```select str from t1 order by id``` is not ok, mixer does not known how to sort because it can not find proper data to compare with `id`

+ Limit should be used with "order by", otherwise you may receive incorrect results

### Insert

+ "insert into select" not supported, later only cross sharding not supported.
+ Multi insert values to different nodes not supported
+ "insert on duplicate key update" can not set the routing key

### Replace

+ Multi replace values to different nodes not supported

### Update

+ Update can not set the routing key

### Set

+ Set autocommit support
+ Set name charset support

### Range Rule

+ Only int64 number range supported

## Caveat

+ Mixer uses 2PC to handle write operations for multi nodes. You take the risk that data becomes corrupted if some nodes commit ok but others error. In that case, you must try to recover your data by yourself.
+ You must design your routing rule and write SQL carefully. (e.g. if your where condition contains no routing key, mixer will route the SQL to all nodes, maybe).

## Why not vitess?

Vitess is very awesome, and I use some of its code like sqlparser. Why not use vitess directly? Maybe below:

+ Vitess is too huge for me, I need a simple proxy
+ Vitess uses an RPC protocol based on BSON, I want to use the MySQL protocol
+ Most likely, something has gone wrong in my head

## Status

Mixer now is still in development and should not be used in production. 

## Feedback

Email: siddontang@gmail.com
