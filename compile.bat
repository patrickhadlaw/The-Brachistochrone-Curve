title Compile
@echo off
cls
:A
set /p compile=">"
echo building %compile%
set GOPATH=%cd%
go build %compile%
echo done
goto A