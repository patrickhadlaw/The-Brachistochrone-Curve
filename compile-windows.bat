title Compile
cls
set GOPATH=%cd%
go install clog
go build host
echo done
goto A