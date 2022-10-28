
1、查询以下三个链接的DNS解析地址
github.com
assets-cdn.github.com
github.global.ssl.fastly.net

https://www.ipaddress.com/


Domain Label	github
Global Traffic Rank	25 ▾2
Estimated Visitors	40.8 Million / Day
Estimated Page Impressions	194.8 Million / Day
Domain Age	15 years and 19 days (5,498 days)
IP Address	
140.82.114.4
Web Server Location	🇺🇸 United States



What are assets-cdn.github.com DNS Records?
The DNS configuration for assets-cdn.github.com includes 4 IPv4 addresses (A) and 4 IPv6 addresses (AAAA).
Additional DNS resource records can be found via our NSLookup Tool, if necessary.
Name	Type	Data
🇺🇸 github.github.io	A	185.199.108.153
🇺🇸 github.github.io	A	185.199.109.153
🇺🇸 github.github.io	A	185.199.110.153
🇺🇸 github.github.io	A	185.199.111.153
🇺🇸 github.github.io	AAAA	2606:50c0:8000::153
🇺🇸 github.github.io	AAAA	2606:50c0:8001::153
🇺🇸 github.github.io	AAAA	2606:50c0:8002::153
🇺🇸 github.github.io	AAAA	2606:50c0:8003::153
assets-cdn.github.com	CNAME	github.github.io


What are github.global.ssl.fastly.net DNS Records?
The DNS configuration for github.global.ssl.fastly.net includes 4 IPv4 addresses (A).
Additional DNS resource records can be found via our NSLookup Tool, if necessary.
Name	Type	Data
🇺🇸 github.global.ssl.fastly.net	A	151.101.1.194
🇺🇸 github.global.ssl.fastly.net	A	151.101.65.194
🇺🇸 github.global.ssl.fastly.net	A	151.101.129.194
🇺🇸 github.global.ssl.fastly.net	A	151.101.193.194


2、修改系统Hosts文件    
路径：C:\Windows\System32\drivers\etc

140.82.114.4    github.com
185.199.109.153 assets-cdn.github.com
151.101.65.194  github.global.ssl.fastly.net

3、刷新系统DNS缓存

Windows+X 打开系统命令行（管理员身份）或powershell
运行 ipconfig /flushdns 手动刷新系统DNS缓存。
mac系统修改完hosts文件,保存并退出就可以了.不要要多一步刷新操作.
centos系统执行/etc/init.d/network restart命令 使得hosts生效