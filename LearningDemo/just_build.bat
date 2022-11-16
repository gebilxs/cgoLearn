@echo off
del libsoe.dll
go env -w GOARCH=386
go env -w CGO_ENABLED=1
go build -ldflags "-s -w" -buildmode=c-shared -o libsoe.dll
IF %errorlevel% NEQ 0 GOTO ERROR
echo build dll success.
copy libsoe.dll c
copy libsoe.h c
cd go
run.bat
cd ..
GOTO END
:ERROR
    echo build dll failed.
:END

:: TODO 支持高阶题型