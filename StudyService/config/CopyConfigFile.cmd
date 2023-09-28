set services=../../services.txt
echo "%services%"
for /F "tokens=1,2,3" %%i in (%services%) do call :process %%i

:process
set service=%1
echo "service: %service%"
if not "%service%" == "UserService" call :copy %1

:copy
set container=../../%1/config/
if not exist "%container%" mkdir "%container%"
copy "config.env" "%container%config.env"
copy "config.go" "%container%config.go"
copy "ConfigKey.go" "%container%ConfigKey.go"

