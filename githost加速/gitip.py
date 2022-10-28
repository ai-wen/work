
import requests
from requests.packages import urllib3
import os,string

header = {
'referer':'https://ipaddress.com/',
"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36 Edg/106.0.1370.52"
}

hostfile = r"C:\Windows\System32\drivers\etc\hosts"

#hostfile = "hosts"

items = []

def get(url ,dom):
    item = []
    urllib3.disable_warnings()
    req = requests.get(url,headers=header,verify=False)
    pos = 0
    while -1 != pos:
        pos = req.text.find('https://www.ipaddress.com/ipv4/',pos)
        if -1 != pos:
            pos2 = req.text.find("\"",pos)
            if -1 != pos:
                str = req.text[pos+len('https://www.ipaddress.com/ipv4/'):pos2]                                         
                item.append(str)    
                item.append(dom)            
                items.append(item)
                pos = pos2
                break
            else:
                break

def AddDomain(ip,dom):    
    hosts=open(hostfile,"a+")
    hosts.writelines("\n"+ ip +"    "+dom)
    hosts.close()


def RMDomain(ip):
    str = ip
    hosts = open(hostfile, "r+")
    lines=hosts.readlines()
    hosts.seek(0)
    context = hosts.read()
    for line in lines:
        if line.find(str) != -1:  #遍历多次 防止用户多次添加而不能完全清除
            context = context.replace(line, "")   #获取包含域名的行内容 防止该域名为用户自行添加而无法识别
    hosts.close()

    hosts=open(hostfile, "w")
    hosts.write(context)
    hosts.close()




get('https://ipaddress.com/site/github.com','github.com')
get('https://ipaddress.com/site/assets-cdn.github.com','assets-cdn.github.com')
get('https://ipaddress.com/site/github.global.ssl.fastly.net','github.global.ssl.fastly.net')

print(items)

for i in range(0,len(items)):
    RMDomain(items[i][0])
    AddDomain(items[i][0],items[i][1])
    
os.system("ipconfig /flushdns")

os.system("notepad "+hostfile)

# pyinstaller -F -c --uac-admin gitip.py