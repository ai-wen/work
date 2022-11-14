cd %cd%

del main.syso
del RandomCheckTool.exe 

windres -i rc/main.rc -O coff -o main.syso
go build -o RandomCheckTool.exe  -ldflags="-H windowsgui"
