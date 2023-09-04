set services=../../services.txt
echo "%services%"
for /F "tokens=1,2,3" %%i in (%services%) do call :process %%i

:process
set service=%1
echo "service: %service%"
if not "%service%" == "MainService" call :copy %1

:copy
set container=../../%1/utils/
if not exist "%container%" mkdir "%container%"
copy "config.go" "%container%config.go"
