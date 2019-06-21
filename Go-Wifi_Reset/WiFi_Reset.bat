ping 8.8.8.8
IF ERRORLEVEL 1 netsh interface set interface Wi-Fi DISABLED

netsh interface set interface Wi-Fi ENABLED