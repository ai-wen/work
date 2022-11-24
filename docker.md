#/bin/sh

# Docker 的旧版本被称为 docker，docker.io 或 docker-engine 。如果已安装，请卸载它们：
apt-get remove docker docker-engine docker.io containerd runc

# 当前称为 Docker Engine-Community 软件包 docker-ce
apt-get update

apt-get install apt-transport-https ca-certificates curl gnupg-agent software-properties-common

# 设置仓库,从仓库安装和更新 Docker 
curl -fsSL https://mirrors.ustc.edu.cn/docker-ce/linux/ubuntu/gpg | sudo apt-key add -

# 添加 Docker 的官方 GPG 密钥
# 9DC8 5822 9FC7 DD38 854A E2D8 8D81 803C 0EBF CD88 通过搜索指纹的后8个字符，验证您现在是否拥有带有指纹的密钥。
apt-key fingerprint 0EBFCD88

add-apt-repository   "deb [arch=amd64] https://mirrors.ustc.edu.cn/docker-ce/linux/ubuntu/ $(lsb_release -cs) stable"

apt-get update

# 安装最新版本的 Docker Engine-Community 和 containerd
apt-get install docker-ce docker-ce-cli containerd.io

# 安装特定版本的 Docker Engine-Community，请在仓库中列出可用版本，然后选择一种安装。
apt-cache madison docker-ce
#docker-ce | 5:20.10.7~3-0~ubuntu-xenial | https://mirrors.ustc.edu.cn/docker-ce/linux/ubuntu xenial/stable amd64 Packages
#docker-ce | 5:20.10.7~3-0~ubuntu-xenial | https://download.docker.com/linux/ubuntu xenial/stable amd64 Packages
#docker-ce | 5:20.10.6~3-0~ubuntu-xenial | https://mirrors.ustc.edu.cn/docker-ce/linux/ubuntu xenial/stable amd64 Packages
#docker-ce | 5:20.10.6~3-0~ubuntu-xenial | https://download.docker.com/linux/ubuntu xenial/stable amd64 Packages
apt-get install docker-ce=5:20.10.7~3-0~ubuntu-xenial docker-ce-cli=5:20.10.7~3-0~ubuntu-xenial containerd.io


# 测试 Docker 是否安装成功
docker run hello-world


# 删除安装包,删除镜像、容器、配置文件
apt-get purge docker-ce
rm -rf /var/lib/docker



#Docker 镜像加速
#科大镜像：https://docker.mirrors.ustc.edu.cn/
#网易：https://hub-mirror.c.163.com/
#七牛云加速器：https://reg-mirror.qiniu.com

# Ubuntu14.04 对于使用 upstart 的系统而言，编辑 /etc/default/docker 文件，在其中的 DOCKER_OPTS 中配置加速器地址：
# DOCKER_OPTS="--registry-mirror=https://registry.docker-cn.com"
# 重新启动服务:service docker restart

#Ubuntu16.04+ 对于使用 systemd 的系统，请在 /etc/docker/daemon.json 中写入如下内容（如果文件不存在请新建该文件）：
{"registry-mirrors":["https://reg-mirror.qiniu.com/"]}
#重新启动服务：
systemctl daemon-reload
systemctl restart docker










#------------------------------镜像----------------------------------------------------------
# 列出本地主机上的镜像
docker images
#REPOSITORY    TAG       IMAGE ID       CREATED         SIZE
#hello-world   latest    feb5d9fea6a5   14 months ago   13.3kB
#ubuntu        14.04     90d5884b1ee0   5 days ago      188 MB
#ubuntu        15.10     4e3b13c8a266   4 weeks ago     136.3 MB
#镜像的仓库源   镜像的标签 镜像ID	镜像创建时间	镜像大小
#同一仓库源可以有多个 TAG，代表这个仓库源的不同个版本，如 ubuntu 仓库源里，有 15.10、14.04 等多个不同的版本
#使用版本为15.10的ubuntu系统镜像来运行容器时，命令： 
docker run -t -i ubuntu:15.10 /bin/bash
# -i: 交互式操作
#-t: 终端。
#ubuntu:15.10: 这是指用 ubuntu 15.10 版本镜像为基础来启动容器。
#/bin/bash：放在镜像名后的是命令，这里我们希望有个交互式 Shell，因此用的是 /bin/bash

# 当我们在本地主机上使用一个不存在的镜像时 Docker 就会自动下载这个镜像。如果我们想预先下载这个镜像，我们可以使用 docker pull 命令来下载它。
# eg:docker pull ubuntu
# eg:docker pull ubuntu:13.10

#查找镜像
# https://hub.docker.com/ 网站搜索镜像
#命令来搜索镜像
docker search mysql
#NAME                            DESCRIPTION                                     STARS     OFFICIAL   AUTOMATED
#mysql                           MySQL is a widely used, open-source relation…   13504     [OK]           
#bitnami/mysql                   Bitnami MySQL Docker Image                      79                   [OK]
#镜像仓库源的名称		镜像的描述					类似 Github 里面的 star	  是否 docker 官方发布 自动构建

#拖取镜像
docker pull mysql
docker run mysql

# 删除镜像
docker rmi hello-world


#当我们从 docker 镜像仓库中下载的镜像不能满足我们的需求时，我们可以通过以下两种方式对镜像进行更改。
#1、从已经创建的容器中更新镜像，并且提交这个镜像

#更新镜像之前，我们需要使用镜像来创建一个容器
docker run -t -i ubuntu:16.04 /bin/bash
#在运行的容器内使用 apt-get update 命令进行更新
#输入 exit 命令来退出这个容器
#通过命令 docker commit 来提交容器副本
docker commit -m="has update" -a="runoob" e218edb10161 runoob/ubuntu:v2
#-m: 提交的描述信息
#-a: 指定镜像作者
#e218edb10161：容器 ID
#runoob/ubuntu:v2: 指定要创建的目标镜像名

#使用新镜像启动一个容器
docker run -t -i runoob/ubuntu:v2 /bin/bash  


#2、使用 Dockerfile 指令来创建一个新的镜像
# 需要创建一个 Dockerfile 文件，其中包含一组指令来告诉 Docker 如何构建我们的镜像
FROM    centos:6.7					
MAINTAINER      Fisher "fisher@sudops.com"

RUN     /bin/echo 'root:123456' |chpasswd		
RUN     useradd runoob
RUN     /bin/echo 'runoob:123456' |chpasswd
RUN     /bin/echo -e "LANG=\"en_US.UTF-8\"" >/etc/default/local
EXPOSE  22
EXPOSE  80
CMD     /usr/sbin/sshd -D

#FROM，指定使用哪个镜像源
#RUN 指令告诉docker 在镜像内执行命令
#RUN ["可执行文件", "参数1", "参数2"]
# CMD 在docker run 时运行。如果 Dockerfile 中如果存在多个 CMD 指令，仅最后一个生效.
# RUN 是在 docker build。
# EXPOSE 仅仅只是声明端口。帮助镜像使用者理解这个镜像服务的守护端口，以方便配置映射。


#每一个指令都会在镜像上创建一个新的层，所以过多无意义的层，会造成镜像膨胀过大。每一个指令的前缀都必须是大写的
FROM centos
RUN yum -y install wget
RUN wget -O redis.tar.gz "http://download.redis.io/releases/redis-5.0.3.tar.gz"
RUN tar -xvf redis.tar.gz
#以上执行会创建 3 层镜像。可简化为以下格式,以 && 符号连接命令，这样执行后，只会创建 1 层镜像。
FROM centos
RUN yum -y install wget \
    && wget -O redis.tar.gz "http://download.redis.io/releases/redis-5.0.3.tar.gz" \
    && tar -xvf redis.tar.gz


#docker build 命令来构建一个镜像
docker build -t runoob/centos:6.7 .
# -t ：指定要创建的目标镜像名
# . ：Dockerfile 文件所在目录，可以指定Dockerfile 的绝对路径

#设置镜像标签,旧标签的镜像仍然存在
docker tag id runoob/centos:dev






















#------------------------------容器----------------------------------------------------------
#获取镜像
docker pull ubuntu
# 启动容器
docker run -it --name kmipserver ubuntu /bin/bash
#-i: 交互式操作。
#-t: 终端。
#--name 
#ubuntu: ubuntu 镜像。
#/bin/bash：放在镜像名后的是命令，这里我们希望有个交互式 Shell，因此用的是 /bin/bash

#退出终端，直接输入 exit:

# -d 指定容器后台运行
docker run -itd --name ubuntu-test ubuntu /bin/bash
#加了 -d 参数默认不会进入容器，想要进入容器需要使用指令 
docker attach id
docker exec -it id/name /bin/bash  # 此命令会退出容器终端，但不会导致容器的停止

# 查看所有容器历史记录
docker ps  -a
# 查看运行容器
docker ps
# CONTAINER ID   IMAGE     COMMAND   		CREATED      STATUS    PORTS     				NAMES
# 容器 ID      使用的镜像  启动容器时运行的命令   容器的创建时间  容器状态  容器的端口信息和使用的连接类型（tcp\udp）  自动分配的容器名称
# 状态有7种：
# created（已创建）
# restarting（重启中）
# running 或 Up（运行中）
# removing（迁移中）
# paused（暂停）
# exited（停止）
# dead（死亡）

# 停止容器
docker stop id/name
#启动一个已停止的容器
docker start id/name
#重启一个已停止的容器
docker restart id/name

#删除容器
docker rm -f id/name

#清理掉所有处于终止状态的容器。
docker container prune
docker rm -f $(docker ps -a | awk '{print $1}')

#导出本地某个容器快照
docker export id > ubuntu.tar
#导入本地容器快照为镜像
cat ubuntu.tar | docker import - test/ubuntu:v1
#导入URL容器快照为镜像
docker import http://example.com/exampleimage.tgz example/imagerepo




#------------------------------运行一个 web 应用----------------------------------------------------------
docker pull training/webapp  # 载入镜像
docker run -d -P training/webapp python app.py
# -d:让容器在后台运行。
# -P:将容器内部使用的网络端口随机映射到我们使用的主机上。
# -P 大写 :是容器内部端口随机映射到主机的端口。
# -p 小写: 是容器内部端口绑定到指定的主机端口。

docker ps
#CONTAINER ID        IMAGE               COMMAND             ...        PORTS                 
#d3d5e39ed9d3        training/webapp     "python app.py"     ...        0.0.0.0:32769->5000/tcp
#Docker 开放了 5000 端口（默认 Python Flask 端口）映射到主机端口 32769 上。
#这时我们可以通过浏览器访问WEB应用 127.0.0.1:32769

#-p 容器内部的 5000 端口映射到我们本地主机的 5000 端口上
docker run -d -p 5000:5000 training/webapp python app.py
#CONTAINER ID        IMAGE                             PORTS                     
#bf08b7f2cd89        training/webapp     ...        0.0.0.0:5000->5000/tcp

docker run -d -p 127.0.0.1:5000:5000/udp training/webapp python app.py
docker run -d -p 127.0.0.1:5001:5000 training/webapp python app.py
#CONTAINER ID        IMAGE               COMMAND           ...     PORTS                                NAMES
#6779686f06f6        training/webapp     "python app.py"   ...   5000/tcp, 127.0.0.1:5000->5000/udp   drunk_visvesvaraya
#95c6ceef88ca        training/webapp     "python app.py"   ...  5000/tcp, 127.0.0.1:5001->5000/tcp   adoring_stonebraker
#通过访问 127.0.0.1:5001 来访问容器的 5000 端口

#查看到容器的端口映射
docker ps
docker port id/name

#查看容器内部的标准输出
docker logs id/name


#查看容器内部运行的进程
docker top id/name

#查看 Docker 的底层信息
docker inspect id/name

#查询最后一次创建的容器
docker ps -l 



#------------------------------Docker Compose----------------------------------------------------------
#Compose 是用于定义和运行多容器 Docker 应用程序的工具。通过 Compose，您可以使用 YML 文件来配置应用程序需要的所有服务。
#使用一个命令，就可以从 YML 文件配置中创建并启动所有服务。

#------------------------------Docker Machine----------------------------------------------------------
#Docker Machine 是一种可以让您在虚拟主机上安装 Docker 的工具，并可以使用 docker-machine 命令来管理主机。
#Docker Machine 也可以集中管理所有的 docker 主机，比如快速的给 100 台服务器安装上 docker。


docker pull mysql:latest
docker pull mysql:5.7
docker pull ubuntu:18.04
docker run -itd --name mysql-test -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql:5.7
docker run -itd --name mysql-test -p 3316:3306 -e MYSQL_ROOT_PASSWORD=123456 mysql:5.7

#如果指定映射的端口不是服务默认的端口，则无法访问服务，-p 3316:3316 例如 mysql  3306
docker run -itd --name mysql-test -p 3316:3316 -e MYSQL_ROOT_PASSWORD=123456 mysql:5.7


容器生命周期管理
run
start/stop/restart
kill
rm
pause/unpause
create
exec
容器操作
ps
inspect
top
attach
events
logs
wait
export
port
stats
容器rootfs命令
commit
cp
diff

镜像仓库
login
pull
push
search
本地镜像管理
images
rmi
tag
build
history
save
load
import
info|version
info
version











# Dockerfile
```
FROM  ubuntu:16.04
MAINTAINER      lkl "liukanglu@longmai.com.cn"

#COPY  etc/*  /opt/etc/
COPY  KMIPServer  /opt/
COPY  libssl.so.1.0.0  /lib/x86_64-linux-gnu/
COPY  libcrypto.so.1.0.0  /lib/x86_64-linux-gnu/
COPY  libmysqlclient.so.20  /lib/x86_64-linux-gnu/

RUN  chmod a+x /opt/KMIPServer
RUN  chmod a+x /lib/x86_64-linux-gnu/libssl.so.1.0.0
RUN  chmod a+x /lib/x86_64-linux-gnu/libcrypto.so.1.0.0
RUN  chmod a+x /lib/x86_64-linux-gnu/libmysqlclient.so.20

EXPOSE  443

#RUN  echo "hello world" 
#RUN  docker build 运行
#CMD     /bin/echo "hello world"
#CMD     /usr/sbin/sshd -D
#CMD  在docker run 时运行,如果存在多个 CMD 指令，仅最后一个生效。




#/bin/bash

#docker build -t lmkmip:v1 .
#docker build -t lmkmip:v1 -f /home/longmai/Desktop/KMIP/Dockerfile

#docker run -tid --name kmip80 -p 5555:80 lmkmip:v1  /bin/bash
#docker run -tid --name kmip443 -p 5556:443 lmkmip:v1  /bin/bash
#docker exec -it kmip80 /bin/bash

#docker stop kmip
#docker rm -f kmip
#docker rmi lmkmip:v1

#./KMIPServer -f 80 -i 192.168.0.161 -P 3316 -n seckms -u root -p root

```

# 远程部署
编辑docker的宿主机文件/lib/systemd/system/docker.service
修改以ExecStart开头的行
ExecStart=/usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock -H tcp://0.0.0.0:2375
systemctl daemon-reload
service docker restart

http://192.168.0.161:2375/version

### docker context create remote ‐‐docker "host=ssh://lm@192.168.0.161:2375"
### docker context create remote --description "192.168.0.161" --docker "host=tcp://192.168.0.161:2375,ca=~/ca-file,cert=~/cert-file,key=~/key-file"
docker context create remote --description "192.168.0.161" --docker "host=tcp://192.168.0.161:2375"

docker context ls

docker context use remote

docker ps