title Install
@echo off
cls
:A
set /p install=">"
echo installing %install%
set GOPATH=%cd%
go install %install%
echo done
goto A