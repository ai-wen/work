1、可用域名和域名证书 

2、域名解析到当前云服务器ip

3、安装trojan
curl -O https://raw.githubusercontent.com/ai-wen/work/main/vpp/trojan_mult.sh && chmod +x trojan_mult.sh && ./trojan_mult.sh

服务端端配置文件
/usr/src/trojan/server.conf

客户端配置文件
/usr/src/trojan/config.json

brr加速
cd /usr/src && wget -N --no-check-certificate "https://raw.githubusercontent.com/chiakge/Linux-NetSpeed/master/tcp.sh" && chmod +x tcp.sh && ./tcp.sh


reboot


apt-get -y install  nginx
/etc/nginx/nginx.conf 
/usr/share/nginx/html/ 
systemctl restart nginx.service



4、linux下使用代理
开启Trojan客户端代理服务
cd /usr/src && wget https://github.com/trojan-gfw/trojan/releases/download/v1.15.1/trojan-1.15.1-linux-amd64.tar.xz
tar xvf trojan-1.15.1-linux-amd64.tar.xz
cd /usr/src/trojan
touch config.json

```config
{
    "run_type": "client",
    "local_addr": "127.0.0.1",
    "local_port": 1080,
    "remote_addr": "xxx.xyz",
    "remote_port": 443,
    "password": [
        "xxx"
    ],
    "log_level": 1,
    "ssl": {
        "verify": true,
        "verify_hostname": true,
        "cert": "",
        "cipher_tls13":"TLS_AES_128_GCM_SHA256:TLS_CHACHA20_POLY1305_SHA256:TLS_AES_256_GCM_SHA384",
        "sni": "",
        "alpn": [
            "h2",
            "http/1.1"
        ],
        "reuse_session": true,
        "session_ticket": false,
        "curves": ""
    },
    "tcp": {
        "no_delay": true,
        "keep_alive": true,
        "fast_open": false,
        "fast_open_qlen": 20
    }
}
```

将 trojan 配置为  service 运行
/etc/systemd/system/trojan.service

cat > /etc/systemd/system/trojan.service <<-EOF
[Unit]
Description=trojan
After=network.target

[Service]
Type=simple
PIDFile=/usr/src/trojan/trojan.pid
ExecStart=/usr/src/trojan/trojan -c /usr/src/trojan/config.json -l /usr/src/trojan/trojan.log
ExecReload=/bin/kill -HUP \$MAINPID
Restart=on-failure
RestartSec=1s

[Install]
WantedBy=multi-user.target

EOF


#### 
systemctl start trojan
systemctl enable --now trojan

ps aux ‘ grep trojan ‘ grep -v grep

Linux貌似默认不支持直接使用socks代理,使用privoxy将socks5转换为http代理

yum install -y privoxy
配置privoxy

vim /etc/privoxy/config
# 末尾增加下面内容,/后面是代理服务器的地址:端口,注意最后还有个.

forward-socks5t / 127.0.0.1:1080 .
启动服务

systemctl start privoxy && systemctl enable privoxy
设置一下系统代理变量

export https_proxy=http://127.0.0.1:1080
export http_proxy=http://127.0.0.1:1080
export all_proxy=http://127.0.0.1:1080
注:1080是privoxy默认使用的端口

测试一下
curl ipfconfig.io



5、windows下使用代理
http://qv2ray.com/wp-content/uploads/2021/10/WinXray.7z

下载使用winxary
右键新增代理服务器
[
    {
        "address":"xxx.xyz",        代理服务器
        "id":"xxx",                 密码
        "network":"tcp",
        "port":443,
        "protocol":"trojan",           trojan协议
        "sni":"xxx.xyz",
        "tls":"tls"
    }
]