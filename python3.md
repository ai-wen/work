# 多环境
- https://www.python.org/downloads/release/python-370/
- https://www.python.org/ftp/python/3.7.0/python-3.7.0.exe
- https://www.python.org/ftp/python/3.7.0/python-3.7.0-amd64.exe

- https://www.python.org/downloads/release/python-380/
- https://www.python.org/ftp/python/3.8.0/python-3.8.0-amd64.exe

- https://www.python.org/downloads/release/python-390/
- https://www.python.org/ftp/python/3.9.0/python-3.9.0-amd64.exe

下载安装 放到C:\SoftW\Python目录

创建 py.bat 内容如下：
并将py.bat 路径设置到PATH环境变量中
setx PATH "C:\SoftW\Python;%PATH%"

```shell

@echo off 

Title Python                                           
Color 0A    
echo.       
echo    1.python 3.70  
echo    2.python 3.80    
echo    3.python 3.90  
echo.
set /p n=select one: 
if "%n%"=="" cls&goto :caozuo 
if "%n%"=="1" call :1 
if "%n%"=="2" call :2 
if "%n%"=="3" call :3 
if /i "%n%"=="n" exit 

goto :eof 

:1 
set PATH=C:\SoftW\Python\Python37;C:\SoftW\Python\Python37\Scripts;%PATH%
goto :caozuo 
:2 
set PATH=C:\SoftW\Python\Python38;C:\SoftW\Python\Python38\Scripts;%PATH%
goto :caozuo 
:3 
set PATH=C:\SoftW\Python\Python39;C:\SoftW\Python\Python39\Scripts;%PATH%
goto :caozuo 

:caozuo 
python -V

```
命令行运行  py

# python3：
在ubuntu的包中，python的二代和三代版本的命名：二代：python，三代：python3

sudo apt install python3
sudo apt install python3-pip

注意：这时安装完毕后的pip真实地址是： /usr/bin/pip3 ,也就是说，要用pip3这个命令去查看是否安装成功： pip3 --version 

可以用： dpkg -L python3-pip 查看一下安装的程序文件位置。

升级pip：
python3 -m pip install --upgrade pip
但升级后，造成一个小错误：
Traceback (most recent call last):
  File "/usr/bin/pip3", line 9, in <module>
    from pip import main
ImportError: cannot import name 'main'
这是，只要修改 /usr/bin/pip3 文件：

from pip import main 修改为：
from pip._internal import main



# requirement.txt
当我们拿到一个项目时，首先要在项目运行环境安装 requirement.txt 所包含的依赖：
pip install -r requirement.txt
pip install --no-cache-dir -r requirements.txt

当我们要把环境中的依赖写入 requirement.txt 中时，可以借助 freeze 命令：
pip freeze >requirements.txt


pipreqs 可以通过扫描项目目录，帮助我们仅生成当前项目的依赖清单。
通过以下命令安装：
pip install pipreqs
运行：
pipreqs ./


#  当前目录下创建虚拟环境
- python -m venv myvnev
- 后面的myvnev代表是在当前路径下创建虚拟环境myvnev，后面跟的是路径
- 进入虚拟环境
- Linux下:
- cd myvenv/bin/
- source activate

- Windows下
- cd myvenv/Scripts
- activate

退出虚拟环境
- 直接输入命令：deactivate


进入之后查看已安装的包
pip list


# 创建备份的python环境
- python -m venv Python38Evn
- cd Python38Evn\Scripts
- copy activate.bat py.bat

- 将 Python38Evn\Scripts 设置到PATH系统变量
- 以后 cmd窗口 执行py 就可以进入这个备份的python 环境

  
# playwright 点击浏览器自动录制生成自动化代码

pip install playwright

安装浏览器驱动文件（安装过程稍微有点慢）
python -m playwright install

## codegen
python -m playwright codegen --target python -o 'my.py' -b chromium https://www.baidu.com


```python
python -m playwright codegen --help
Usage: playwright codegen [options] [url]

open page and generate code for user actions

Options:
  -o, --output <file name>        saves the generated script to a file
  --target <language>             language to generate, one of javascript, playwright-test, python, python-async,
                                  python-pytest, csharp, csharp-mstest, csharp-nunit, java (default: "python")
  --save-trace <filename>         record a trace for the session and save it to a file
  -b, --browser <browserType>     browser to use, one of cr, chromium, ff, firefox, wk, webkit (default: "chromium")
  --block-service-workers         block service workers
  --channel <channel>             Chromium distribution channel, "chrome", "chrome-beta", "msedge-dev", etc
  --color-scheme <scheme>         emulate preferred color scheme, "light" or "dark"
  --device <deviceName>           emulate device, for example  "iPhone 11"
  --geolocation <coordinates>     specify geolocation coordinates, for example "37.819722,-122.478611"
  --ignore-https-errors           ignore https errors
  --load-storage <filename>       load context storage state from the file, previously saved with --save-storage
  --lang <language>               specify language / locale, for example "en-GB"
  --proxy-server <proxy>          specify proxy server, for example "http://myproxy:3128" or "socks5://myproxy:8080"
  --proxy-bypass <bypass>         comma-separated domains to bypass proxy, for example ".com,chromium.org,.domain.com"
  --save-har <filename>           save HAR file with all network activity at the end
  --save-har-glob <glob pattern>  filter entries in the HAR by matching url against this glob pattern
  --save-storage <filename>       save context storage state at the end, for later use with --load-storage
  --timezone <time zone>          time zone to emulate, for example "Europe/Rome"
  --timeout <timeout>             timeout for Playwright actions in milliseconds, no timeout by default
  --user-agent <ua string>        specify user agent string
  --viewport-size <size>          specify browser viewport size in pixels, for example "1280, 720"
  -h, --help                      display help for command

Examples:

  $ codegen
  $ codegen --target=python
  $ codegen -b webkit https://example.com



from playwright import sync_playwright
with sync_playwright() as p:
    for browser_type in [p.chromium, p.firefox, p.webkit]:
        browser = browser_type.launch()
        page = browser.newPage()
        page.goto('https://baidu.com/')
        page.screenshot(path=f'example-{browser_type.name}.png')
        browser.close()

import asyncio
from playwright import async_playwright
async def main():
    async with async_playwright() as p:
        for browser_type in [p.chromium, p.firefox, p.webkit]:
            browser = await browser_type.launch()
            page = await browser.newPage()
            await page.goto('http://baidu.com/')
            await page.screenshot(path=f'example-{browser_type.name}.png')
            await browser.close()
asyncio.get_event_loop().run_until_complete(main())

from playwright import sync_playwright
with sync_playwright() as p:
    iphone_11 = p.devices['iPhone 11 Pro']
    browser = p.webkit.launch(headless=False)
    context = browser.newContext(
        **iphone_11,
        locale='en-US',
        geolocation={ 'longitude': 12.492507, 'latitude': 41.889938 },
        permissions=['geolocation']
    )
    page = context.newPage()
    page.goto('https://maps.google.com')
    page.click('text="Your location"')
    page.screenshot(path='colosseum-iphone.png')
    browser.close()
``
