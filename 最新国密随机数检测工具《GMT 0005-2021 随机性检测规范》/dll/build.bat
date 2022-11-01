go build -buildmode=c-shared -o random.dll randomTest.go

@call "C:\Program Files (x86)\Microsoft Visual Studio\2019\Community\VC\Auxiliary\Build\vcvarsall.bat" x64 %*

lib.exe /def:random.def /machine:x64 /out:random64.lib