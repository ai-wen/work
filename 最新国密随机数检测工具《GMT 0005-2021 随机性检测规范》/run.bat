
cd  %cd%
cd rddetector
go build .

cd ..
cd rdgen
go build .

cd ..

rdgen\rdgen.exe -s 1000 -n 1000000 -o data

:::rddetector\rddetector.exe -i data -o RandomnessTestReport.csv

cd dll

go build -buildmode=c-shared -o random.dll randomTest.go

@call "C:\Program Files (x86)\Microsoft Visual Studio\2019\Community\VC\Auxiliary\Build\vcvarsall.bat" x64 %*

lib.exe /def:random.def /machine:x64 /out:random64.lib

cl.exe main.cpp

cd ..

dll\main.exe data >CTestReport.csv
