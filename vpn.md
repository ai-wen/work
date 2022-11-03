1、可用域名和域名证书 

2、域名解析到当前云服务器ip

3、安装trojan
curl -O https://raw.githubusercontent.com/ai-wen/work/main/trojan_mult.sh && chmod +x trojan_mult.sh && ./trojan_mult.sh

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
