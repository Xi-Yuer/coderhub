[mysqld]
# 开启 binlog
log-bin=mysql-bin
# 设置 server id
server-id=1
# 设置 binlog 格式为 ROW
binlog-format=ROW
# 设置需要复制的数据库
binlog-do-db=coderhub 