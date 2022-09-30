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

当我们要把环境中的依赖写入 requirement.txt 中时，可以借助 freeze 命令：
pip freeze >requirements.txt


pipreqs 可以通过扫描项目目录，帮助我们仅生成当前项目的依赖清单。
通过以下命令安装：
pip install pipreqs
运行：
pipreqs ./


#  当前目录下创建虚拟环境
python -m venv myvnev
后面的myvnev代表是在当前路径下创建虚拟环境myvnev，后面跟的是路径
进入虚拟环境
Linux下:
cd myvenv/bin/
source activate

Windows下
cd myvenv/Scripts
activate

退出虚拟环境
直接输入命令：deactivate


进入之后查看已安装的包
pip list


# 创建备份的python环境
python -m venv Python38Evn
cd Python38Evn\Scripts
copy activate.bat py.bat

将 Python38Evn\Scripts 设置到PATH系统变量
以后 cmd窗口 执行py 就可以进入这个备份的python 环境