@echo off
cd %~dp0
if %0 == "install" goto install



:install
echo -------------------------------------- begin to make install...
::go get github.com/chainlibs/gobtclib
mklink /D %GOPATH%\src\github.com\chainlibs\gobtclib %cd%
echo -------------------------------------- finished successfully!
pause