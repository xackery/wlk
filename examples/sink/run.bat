@echo off
echo Building sink
go build
echo Running sink
sink.exe
exit /b 0

:error
echo Error building %lastdir%: %errorlevel%
exit /b %errorlevel%