@echo off
echo.
echo Refreshing PATH from registry
:: Get System PATH
for /f "tokens=3*" %%A in ('reg query "HKLM\SYSTEM\CurrentControlSet\Control\Session Manager\Environment" /v Path') do set syspath=%%A%%B

:: Get User Path
for /f "tokens=3*" %%A in ('reg query "HKCU\Environment" /v Path') do set userpath=%%A%%B

:: Set Refreshed Path
set PATH=%userpath%;%syspath%
SET PATH=%PATH%;c:gopath\bin
SET PATH=%PATH%;c:\gopath
SET PATH=%PATH%;c:gopath\bin\wails.exe
SET PATH=%PATH%;c:gopath\bin\mewn.exe
echo Refreshed PATH
echo %PATH%