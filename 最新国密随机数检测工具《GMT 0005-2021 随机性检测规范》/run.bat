
cd  %cd%
cd rddetector
go build .

cd ..
cd rdgen
go build .

cd ..

rdgen\rdgen.exe -s 1000 -n 1000000 -o data


rddetector\rddetector.exe -i data -o RandomnessTestReport.csv