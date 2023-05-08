@ECHO OFF

if "%1"=="hide" goto work
start mshta vbscript:createobject("wscript.shell").run("""%~0"" hide",0)(window.close)&&exit
:work

F:\VirtualBox\VBoxManage.exe startvm gitlab --type headless

set HOUR=%time:~0,2%
set MINI=%time:~3,2%
:loop

if %HOUR% GEQ 18 (
    if %MINI% GEQ 30 (
        F:\VirtualBox\VBoxManage.exe controlvm gitlab poweroff
        exit
    )
)
TIMEOUT /T 60
goto :loop

exit


rem F:\VirtualBox\VBoxManage.exe startvm gitlab --type headless % 开机，且无界面
rem F:\VirtualBox\VBoxManage.exe startvm gitlab --type gui % 开机，带界面
rem F:\VirtualBox\VBoxManage.exe controlvm gitlab pause % 暂停
rem F:\VirtualBox\VBoxManage.exe controlvm gitlab savestate % 休眠
rem F:\VirtualBox\VBoxManage.exe controlvm gitlab poweroff % 强制关机
rem F:\VirtualBox\VBoxManage.exe controlvm gitlab resume % 从暂停状态恢复
rem F:\VirtualBox\VBoxManage.exe controlvm gitlab acpipowerbutton % 按下电源键
rem F:\VirtualBox\VBoxManage.exe controlvm gitlab acpisleepbutton % 按下睡眠键

rem 将VMware文件（.vmx）转换为Virtual Box文件（.ovf）
rem D:\SW\VMware Workstation\OVFTool
rem ovftool.exe input.vmx output.ovf
rem ovftool.exe "D:\VMOS\centos7.61810\CentOS 7.6 64 位.vmx" "D:\VMOS\centos7"\centos7.6.ovf
rem 打开Oracle VM VirtualBox => 导入

