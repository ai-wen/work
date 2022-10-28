# https://cloud.tencent.com/developer/article/1930581
安装python2.7
git clone https://github.com/esrlabs/git-repo
export PATH=/c/Python27:$PATH
git config --global user.email "363042868@qq.com"
git config --global user.name "ai-wen"

git-repo/repo init -u https://aosp.tuna.tsinghua.edu.cn/platform/manifest -b android-8.0.0_r36
git-repo/repo sync

.repo\repo\project.py:278





# https://www.jianshu.com/p/a905cc193e63

1、直接下载初始化源码库
https://mirrors.tuna.tsinghua.edu.cn/aosp-monthly/aosp-latest.tar

下载完成以后会得到一个37g的以aosp-latest命名的tar文件，然后将其解压可以看到里面有一个.repo
# repo sync,然后就可以同步到最新的分支，并检出其master分支。
repo sync
# 检出repo里面的源码项目。 
repo sync -l

2、repo下载源码库
linux 下下载repo
mkdir ~/bin
PATH=~/bin:$PATH
curl https://storage.googleapis.com/git-repo-downloads/repo > ~/bin/repo
chmod a+x ~/bin/repo


Windows下安装repo
mkdir -p /f/androidsrc/.repo
# .repo是准备存放repo工具的目录，目录名称固定。
cd /f/androidsrc/.repo
git clone https://mirrors.tuna.tsinghua.edu.cn/git/git-repo repo
export PATH=$PATH:/f/androidsrc/.repo/repo

# set path=%path%;F:/androidsrc/.repo/repo
cd /f/androidsrc


checkou src
# https://mirrors.tuna.tsinghua.edu.cn/help/AOSP/
# 初始化repo，此处可以换成自己的repo地址 https://source.android.com/setup/start/build-numbers#source-code-tags-and-builds
# repo init -u https://aosp.tuna.tsinghua.edu.cn/platform/manifest -b android-10.0.0_r25 --worktree
repo init -u https://mirrors.tuna.tsinghua.edu.cn/git/AOSP/platform/manifest -b android-8.0.0_r36

# 更新git仓库，如果是第一次执行此命令则会下载git仓库，如果不是第一次执行，则会更新git仓库。
repo sync


3、放入编译文件
从github上拉下来的开源Android项目也是没有.iml等配置文件的，需要我们自己再编译一边，才能让这个项目里的个文件在AS关联起来。
那么在这里之所以不用编译，不是真的无需编译。而是有一位简书ID是difcareer的小伙伴已经帮我们把各个版本需要的AS配置文件已经编译好，我们只需要将它下载下来放到项目的根目录就ok了。https://github.com/difcareer/AndroidSourceReader
android-4.4.4_r1 
android-5.0.1_r1 
android-6.0.1_r11 
android-7.1.2_r28 
android-8.0.0_r36 
android-9.0.0_r1 
注：
如果你忘了自己源码检出的版本可以通过如下方式查看：
根据目录'/build/core/version_defaults.mk' 打开version_defaults.mk文件，然后找到PLATFORM_SDK_VERSION这个关键字的值，就可以在下表中找出相应的版本。

第一次导入的时候，可能你看到的只有那几个刚才放入的配置文件，而看不到其他的目录，这个时候需要点击File->Invalidate Cashes/Restart...，让AS重新启动编译一下项目。
如果出现循环执行任务Scanning file to index.... 不动
解决办法如下（Open module setting --> Modules --> 找到gen文件夹 --> 选择Resources）