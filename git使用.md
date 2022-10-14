# 删除仓库
Settings/General/Advanced/Delete project

# 关闭编码格式自动转换功能
LF和CRLF其实都是换行符，但是不同的是，LF是linux和Unix系统的换行符，CRLF是window 系统的换行符。
把自动转换功能关掉即可。
git config core.autocrlf false (仅对当前git仓库有效）
git config --global core.autocrlf false (全局有效，不设置推荐全局）

# GIT
git config --global user.name "ai-wen"
git config –-global user.email "363042868@qq.com"

http://192.168.1.160/

user：liukanglu
pass：liu1989226

git clone  http://192.168.1.160/liuchao/CipherSdk.git

git remote set-url origin http://liukanglu:liu1989226@192.168.1.160/liuchao/CipherSdk.git

git remote set-url origin http://liukanglu:liu1989226@192.168.1.160/CipherServer/CipherKMIPServer.git


git remote set-url origin https://ai-wen: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@github.com/ai-wen/work.git

# 本地提交
git add .
git commit -m "First commit"

- 添加所有文件
git add .
git add --all

- 添加多个文件
git add file1 file2 file3

- 添加文件夹
git add 文件夹名

- 添加文件夹本身及文件
git config/*

git add 目的是将修改文件由工作区提交到暂存区，可以多次提交
然后commit操作，将文件从暂存区提交到版本库


# 删除
git push origin --delete <branchName>
默认分支无法删除
git push origin --delete tag <tagname>

# 分支操作
git branch -a 
git checkout dev 切换到分支
git push origin dev:dev 
git push origin dev
git pull origin dev  第一次无法pull，只能push


git push
fatal: The current branch dev has no upstream branch.
To push the current branch and set the remote as upstream, use
    git push --set-upstream origin dev

git clone -b dev http://liukanglu:liu1989226@192.168.1.160/liuchao/CipherSdk.git



## 
git clone http://192.168.1.160/liukanglu/ciphersuite.git
git remote -v           // 查看远程地址
git remote set-url origin git@192.168.1.160:liukanglu/ciphersuite.git

git remote rm origin      // 删除原有的推送地址
git remote add origin git@192.168.1.160/liukanglu/ciphersuite.git

git push
git push -u origin main















# http免密登录
git clone http://192.168.1.160/liukanglu/ciphersuite.git
git remote -v           // 查看远程地址
git remote set-url origin http://liukanglu:liu1989226@192.168.1.160/liukanglu/ciphersuite.git
或者
git remote rm origin 
git remote add origin http://liukanglu:liu1989226@192.168.1.160/liukanglu/ciphersuite.git

git push
git push -u origin main

# ssh免密登录
git config --global user.name "liukanglu"
git config –-global user.email "liukanglu@longmai.com.cn"

## 创建公钥私钥
ssh-keygen -t rsa -C "liukanglu@longmai.com.cn" 

如果之前有公钥私钥文件，那么这样操作会让你覆盖之前文件，你直接覆盖了就好了，
然后你会发现在对应位置的文件夹下生成了一个.ssh文件夹
里面有一个pub文件，这就是公钥了：
/c/Users/lm/.ssh/id_rsa
文本打开公钥文件，把里面的公钥内容复制到git仓库里面的
.ssh/authorized_keys 文件里面就可以连通本地和git仓库端了，这样配置有一个好处就是连接仓库不用单独输入密码

ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDLcp1soHqZBR5oTjotw7wpoZPWo4EkFnAubJH0rj7+HUX9WuvB8MeNAhvSL95nxXtd8IBqsau2/dF7JAv8U2DPxWCZ+TPCI4qsrnv5/DbA8GPIvA9+x1uJvyVIY7vncXNcpq5S5d+Bomaz8mMQxXyheQhOEsB8QEn02ABZxP4D+TqacPIgXnErT2GFtEXr69CbjuNVKG8/hCXFBYWWUSzDPRUNOhAvxaw04atlqhOi3oG/ZpazhOnAnL7JVU8HhgRHDqdS6Ris8E98YIftHMPxdxFvzVreWGrOLxkTCdJgtiDo8aN/acL6HAy9vEOSx1snyMDm0T8BayLyU3x5+M27 liukanglu@longmai.com.cn







# Git的基本操作
初始化操作
    $ git config -global user.name <name> #设置提交者名字
    $ git config -global user.email <email> #设置提交者邮箱
    $ git config -global core.editor <editor> #设置默认文本编辑器
    $ git config -global merge.tool <tool> #设置解决合并冲突时差异分析工具
    $ git config -list #检查已有的配置信息
创建新版本库
    $ git clone <url> #克隆远程版本库
    $ git init #初始化本地版本库
修改和提交
    $ git add . #添加所有改动过的文件
    $ git add <file> #添加指定的文件
    $ git mv <old> <new> #文件重命名
    $ git rm <file> #删除文件
    $ git rm -cached <file> #停止跟踪文件但不删除
    $ git commit -m <file> #提交指定文件
    $ git commit -m "commit message" #提交所有更新过的文件
    $ git commit -amend #修改最后一次提交
    $ git commit -C HEAD -a -amend #增补提交（不会产生新的提交历史纪录）
查看提交历史
    $ git log #查看提交历史
    $ git log -p <file> #查看指定文件的提交历史
    $ git blame <file> #以列表方式查看指定文件的提交历史
    $ gitk #查看当前分支历史纪录
    $ gitk <branch> #查看某分支历史纪录
    $ gitk --all #查看所有分支历史纪录
    $ git branch -v #每个分支最后的提交
    $ git status #查看当前状态
    $ git diff #查看变更内容
撤消操作
    $ git reset -hard HEAD #撤消工作目录中所有未提交文件的修改内容
    $ git checkout HEAD <file1> <file2> #撤消指定的未提交文件的修改内容
    $ git checkout HEAD. #撤消所有文件
    $ git revert <commit> #撤消指定的提交
分支与标签
    $ git branch #显示所有本地分支
    $ git checkout <branch/tagname> #切换到指定分支或标签
    $ git branch <new-branch> #创建新分支
    $ git branch -d <branch> #删除本地分支
    $ git tag #列出所有本地标签
    $ git tag <tagname> #基于最新提交创建标签
    $ git tag -d <tagname> #删除标签
合并与衍合
    $ git merge <branch> #合并指定分支到当前分支
    $ git rebase <branch> #衍合指定分支到当前分支
远程操作
    $ git remote -v #查看远程版本库信息
    $ git remote show <remote> #查看指定远程版本库信息
    $ git remote add <remote> <url> #添加远程版本库
    $ git fetch <remote> #从远程库获取代码
    $ git pull <remote> <branch> #下载代码及快速合并
    $ git push <remote> <branch> #上传代码及快速合并
    $ git push <remote> : <branch>/<tagname> #删除远程分支或标签
    $ git push -tags #上传所有标签