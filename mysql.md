# 启用外键约束
SET FOREIGN_KEY_CHECKS = 1;

# 外键更新删除设置
## 1、父表，从表
假如表A（id, foreifn_id），表B（foreign_id），我们说表A参考了表B的主键作为其外键使用，所以B表示父表，A表是子表。

## 2、删除和更新有四种设置方式。
cascade：级联，当父表更新、删除，子表会同步更新和删除。
    删除：删除主表时自动删除从表。删除从表，主表不变。
    更新：更新主表时自动更新从表。更新从表，主表不变。

set null : 置空，当父表更新、删除的时候，子表会把外键字段变为null，所以这个时候设计表的时候该字段要允许为null，否则会出错。
    删除：删除主表时自动更新从表为NULL，删除从表，主表不变。
    更新：更新主表时自动更新从表值为NULL。更新从表，主表不变。

restrict : 父表在删除和更新记录的时候，要在子表中检查是否有有关该父表要更新和删除的记录，如果有，则不允许删除更改。
no action：和restrict一样。
什么也不选：什么也不选就会默认选择no action。
    删除：从表记录不存在时，主表才可以删除，删除从表，主表不变。
    更新：从表记录不存在时，主表菜可以更新，更新从表，主表不变。


# 事务
在 MySQL 中只有使用了 Innodb 数据库引擎的数据库或表才支持事务。事务处理可以用来维护数据库的完整性，保证成批的 SQL 语句要么全部执行，要么全部不执行。事务一般用来管理 insert,update,delete语句。

事务的开启：(不区分大小写）BEGIN;或START TRANSACTION;
事务的提交：COMMIT;
建立事务的存档点：SAVEPOINT nameofsavepoint（自定义）；
删除存档点：RELEASE SAVEPOINT nameifsavepoint（无此存档点会报错）；
事务的回滚：ROLLBACK; //回滚到begin的位置。或 ROLLBACK TO nameofsavepoint; //回滚到指定存档点。
