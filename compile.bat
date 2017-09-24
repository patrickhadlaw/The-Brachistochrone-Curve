title Compile
@echo off
cls
:A
set /p compile=">"
echo building %compile%
set GOROOT=%cd%
go build %compile%
echo done
goto A