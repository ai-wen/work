cd %cd%

del main.syso
del RandomCheckTool.exe 

windres -i rc/main.rc -O coff -o main.syso
::timeout /T 3 /NOBREAK

go build -ldflags="-H windowsgui"

move ui.exe RandomCheckTool.exe

  
