#!/bin/bash

MYSQL_BIN=`which mysql`
SETUP_SQL="/data/github/GoLearn/mysql/stub/mysql_stub.sql"

## 执行setup.sql
{
    ${MYSQL_BIN} -u root -D test < "${SETUP_SQL}"
}