@echo off
echo.
echo Refreshing PATH from registry
:: Get System PATH
for /f "tokens=3*" %%A in ('reg query "HKLM\SYSTEM\CurrentControlSet\Control\Session Manager\Environment" /v Path') do set syspath=%%A%%B

:: Get User Path
for /f "tokens=3*" %%A in ('reg query "HKCU\Environment" /v Path') do set userpath=%%A%%B

:: Set Refreshed Path
set PATH=%userpath%;%syspath%
set PATH=%PATH%;c:gopath\bin
set PATH=%PATH%;c:\gopath
set PATH=%PATH%;c:\Users\circleci\go\bin\wails.exe
set PATH=%PATH%;c:\Users\circleci\go\bin\mewn.exe
set PATH=%PATH%;C:\Users\circleci\go\bin\wails.exe
set PATH=%PATH%;C:\Users\circleci\go\bin\mewn.exe
set PATH=%PATH%;c:gopath\bin
set PATH=%PATH%;c:\gopath
echo Refreshed PATH
echo %PATH%
