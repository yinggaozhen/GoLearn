#!/bin/bash

BIN_DOCKER=`which docker`

REDIS_IMAGE="redis"
REDIS_NETWORK="redis-net"


## clean local redis container
{
    ${BIN_DOCKER} rm `${BIN_DOCKER} ps -a -q` -f
}

{
    ${BIN_DOCKER} network create ${REDIS_NETWORK}
}

{
    tmp_dir="/tmp/redis_conf"
    for port in `seq 7000 7005`; do
        mkdir -p ${tmp_dir}
        cp /data/github/GoLearn/mysql/stub/redis-cluster.tmpl ${tmp_dir}/${port}.conf
        sed -i '' "s/\${PORT}/${port}/g" ${tmp_dir}/${port}.conf
    done

    for port in `seq 7000 7005`; do
        docker run -d -it -p ${port}:${port} -v ${tmp_dir}:/etc/redis --name redis-${port} --net redis-net redis redis-server /etc/redis/${port}.conf
    done
}