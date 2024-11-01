@echo off
set GOOS=windows
set GOARCH=amd64
go build -ldflags -H=windowsgui -o forge-gui.exe
if %errorlevel% equ 0 (
    echo Build successful!
    echo Output: forge-gui.exe
) else (
    echo Build failed!
)
pause 