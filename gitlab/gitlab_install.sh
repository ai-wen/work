#!/bin/bash

查看与rpm包相关的文件和其他信息     rpm -qa | grep 包名 
查询包是否被安装    rpm -q 包名
删除软件包  rpm -e 包名


# 硬件配置
内存 >=8G
CPU >= 4核

# 邮件服务器 postfix
yum install -y curl policycoreutils-python openssh-server postfix wget vim
systemctl enable sshd
systemctl start sshd
systemctl enable postfix
systemctl start postfix
firewall-cmd --add-service=ssh --permanent
firewall-cmd --add-service=http --permanent
firewall-cmd --reload
#安装防火墙
yum install firewalld systemd -y
service firewalld  start
systemctl stop firewalld.service 
systemctl disable firewalld.service
firewall-cmd --zone=public --add-port=80/tcp --permanent
firewall-cmd --reload

# 安装
# https://packages.gitlab.com/gitlab/gitlab-ce/install#bash-rpm
curl -s https://packages.gitlab.com/install/repositories/gitlab/gitlab-ce/script.rpm.sh | sudo bash
yum install -y gitlab-ce

# 设置开机自启
#/etc/systemd/system/multi-user.target.wants/gitlab-runsvdir.service
systemctl enable gitlab-runsvdir.service
# 备份

# 命令行 获取/修改超级管理员root的密码
cd /opt/gitlab/bin
gitlab-rails console -e production
输入 user=User.where(id:1).first 来查找与切换账号
     user = User.where(username:"root").first
（User.all 可以查看所有用户）
输入 user.password = 'qwer1234' 修改密码为12345678
输入再次确认 user.password_confirmation='12345678'
输入保存 user.save!
输入 exit



修改配置文件 vi /etc/gitlab/gitlab.rb
# 1、设置站点
## GitLab默认的配置文件路径是：/etc/gitlab/gitlab.rb
## 默认的站点Url配置项是： external_url 'http://gitlab.example.com'
## 配置首页地址（大约在第15行） 
external_url 'http://127.0.0.1'
#EXTERNAL_URL="http://127.0.0.1:8090" apt-get install gitlab-ce

# 2、禁用创建组权限
## GitLab默认所有的注册用户都可以创建组。但对于团队来说，通常只会给Leader相关权限。
## 虽然可以在用户管理界面取消权限，但毕竟不方便。我们可以通过配置GitLab默认禁用创建组权限。 
## 开启gitlab_rails['gitlab_default_can_create_group'] 选项，并将值设置为false
gitlab_rails['gitlab_default_can_create_group'] = false

# 3、设置git-data存储位置
git_data_dirs({ "default" => { "path" => "/gitdata" } })   #添加指定存储位置
gitlab-ctl reconfigure
gitlab-ctl stop
rsync -av /var/opt/gitlab/git-data/ /gitdata/   #同步文件内容
gitlab-ctl upgrade
gitlab-ctl start

# 4、设置发件邮箱
gitlab_rails['smtp_enable']=true
gitlab_rails['smtp_address']='smtp.ym.163.com'
['smtp_port']=25
['smtp_user_name'] = ""
['smtp_password']=""
['smtp domain']='smtp.ym.163.com'
['smtp_tls'] = false

# 4、备份
手动备份
gitlab_rails['backup_path'] = '/gitdata/gitlab_backups'
gitlab-rake gitlab:backup:create
默然的备份目录为： /var/opt/gitlab/backups
备份文件名类似： xxxxxx_gitlab_backup.tar

自动定时备份
gitlab_rails['backup_keep_time'] = 259200      # 删除注释 #， 3天*3600秒*24时=259200s
gitlab-ctl reconfigure
配置定时任务 需重启cron服务
crontab -e
# 0 2 * * * /opt/gitlab/bin/gitlab-rake gitlab:backup:create   #每天凌晨2点备份
0 15 * * * /opt/gitlab/bin/gitlab-rake gitlab:backup:create   #每天下午3点备份
crontab -l   查看定时任务
systemctl enable crond.service  #设置cron服务开机使能
systemctl restart crond     #修改后重启cron服务


0 18 * * * /opt/gitlab/bin/gitlab-rake gitlab:backup:create
15 18 * * * /gitlab/bak.sh
创建脚本备份到另一个目录/ /mnt/hgfs/ 本机等
sudo cp /gitdata/repositories /gitlab/ -rfp
sudo cp /gitdata/gitlab_backups/* /gitlab/ -fp
sudo rm /gitdata/gitlab_backups/* -f



#重新配置并启动
## gitlab-ctl reconfigure
## 完成后将会看到如下输出
## Running handlers complete
## Chef Client finished, 432/613 resources updated in 03 minutes 43 seconds
## gitlab Reconfigured!
## cat /opt/gitlab/embedded/service/gitlab-rails/VERSION 查看版本号


# 5、手动恢复
gitlab-ctl stop unicorn
gitlab-ctl stop sidekiq
gitlab-ctl status

gitlab-rake gitlab:backup:restore BACKUP=141111111
gitlab-rake gitlab:check SANITIZE=true       #check恢复数据是否OK
gitlab-ctl start


# 6、web其它设置
1：access protocols 查看是否开启ssh和http
2：Account and Limit Gravatar enabled 关闭头像服务，因其在国外，减缓了访问速度
3：Sign-up Restrictions Sign-up enabled 关闭注册接口，因为这是企业内部的访问仓库，账号由管理员分发
4：Sign in text 登录页面的信息提示 这是企业内部gitlab地址，已关闭注册接口。如果您是内部人员，请联系管理员进行账号分发！
5: Save



# gitlab-ctl 命令
## restart 重启GitLab服务
## start   如果GitLab服务停止了就启动服务，如果已启动就重启服务
## stop    停止GitLab服务
## status  查看GitLab服务状态
## reconfigure 重新配置GitLab并启动
## uninstall    卸载
## cleanse      清空 gitlab 配置，推倒重来 /etc/gitlab下的配置文件会自动备份 Your config files have been backed up to /root/gitlab-cleanse-2020-07-02T10:41.
## tail   查看日志

## check-config    检查在gitlab中是否有任何配置。在指定版本中删除的rb
## deploy-page     安装部署页面
## diff-config     将用户配置与包可用配置进行比较
## remove-accounts     删除所有用户和组
## upgrade     升级
## service-list    查看所有服务
## once    如果GitLab服务停止了就启动服务，如果已启动就不做任何操作
## show-config 查看 gitlab 配置










yum update
yum -y install ntp ntpdate  #安装ntpdate工具
ntpdate cn.pool.ntp.org   #设置系统时间与网络时间同步
hwclock --systohc          #将系统时间写入硬件时间

systemctl restart network

vi /etc/sysconfig/network-scripts/ifcfg-ens32
BOOTPROTO="static" #dhcp改为static     这个是改静态用的.实际不改. 自己记录用 
ONBOOT="yes"       #开机启用本配置      改这里就OK
IPADDR=192.168.0.106 #静态IP     根据路由器实际分配的IP进行设置    
NETMASK=255.255.255.0 #子网掩码
GATEWAY=192.168.0.1 #默认网关      根据路由器网关设置
DNS1= 8.8.8.8
DNS2=8.8.8.4

TYPE=Ethernet
PROXY_METHOD=none
BROWSER_ONLY=no
BOOTPROTO=static
DEFROUTE=yes
IPV4_FAILURE_FATAL=no
IPV6INIT=yes
IPV6_AUTOCONF=yes
IPV6_DEFROUTE=yes
IPV6_FAILURE_FATAL=no
IPV6_ADDR_GEN_MODE=stable-privacy
NAME=ens33
UUID=bfac8541-73ed-48d9-819d-7ac31793c566
DEVICE=ens33
ONBOOT=yes
IPADDR=192.168.0.106
NETMASK=255.255.255.0
GATEWAY=192.168.0.1
DNS1= 8.8.8.8
DNS2=8.8.8.4