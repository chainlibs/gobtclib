@echo off
cd %~dp0
if %0 == "install" goto install



:install
echo -------------------------------------- begin to make install...
::go get github.com/chainlibs/gobtclib
go get -u go.uber.org/zap
go get -u github.com/gobasis/log
mkdir %GOPATH%\src\github.com\chainlibs
rmdir /s/q %GOPATH%\src\github.com\chainlibs\gobtclib
mklink /D %GOPATH%\src\github.com\chainlibs\gobtclib %cd%
echo -------------------------------------- finished successfully!
pause