## go-sql
用于快速统计数据库行数、敏感字段匹配、数据库连接情况。

## usage

```
./go-sql_darwin_amd64 -h
./go-sql_darwin_amd64 -f db.yaml -k name,user
./go-sql_darwin_amd64 -f db.yaml --minRow 10000
```

## screenshot

Mysql

![image-20211128134840757](https://nnotes.oss-cn-hangzhou.aliyuncs.com/notes/image-20211128134840757.png)

mssql

![image-20211128135110329](https://nnotes.oss-cn-hangzhou.aliyuncs.com/notes/image-20211128135110329.png)

oracle

![image-20211128135623331](https://nnotes.oss-cn-hangzhou.aliyuncs.com/notes/image-20211128135623331.png)

postgres

![image-20211128135604972](https://nnotes.oss-cn-hangzhou.aliyuncs.com/notes/image-20211128135604972.png)

## description

- `-f`指定数据库连接配置文件

```yaml
db:
  - db_type: mysql
    conn:
      host: 127.0.0.1
      port: 3306
      db_name:
      user: root
      pass: root
    sql:
  - db_type: mssql
    conn:
      host: 127.0.0.1
      port: 1433
      db_name: testDB
      user: sa
      pass: 123qweASD
    sql:
```

- 使用出现oci8报错的情况,需要下载对应操作系统的instantclient，比如`instantclient-basic-linux.x64`和`instantclient-sdk-linux.x64`，解压到同一文件夹并添加到环境变量。(https://www.oracle.com/database/technologies/instant-client/macos-intel-x86-downloads.html)
- 如果配置文件中指定了sql则执行sql语句，否则执行内置的sql语句。

