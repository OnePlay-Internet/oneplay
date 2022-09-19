$env:PKG_CONFIG_PATH = "C:\gstreamer\1.0\msvc_x86_64\lib\pkgconfig"

Set-Location daemon 
go build -o daemon.exe
Set-Location ..

Set-Location sunshine-util
go build  -o screencoder.exe  examples/go_webrtc/main.go
Set-Location ..

Set-Location DevSim
dotnet build . --output "bin" --self-contained true --runtime win-x64
Set-Location ..

robocopy daemon package/bin daemon.exe
robocopy sunshine-util package/bin screencoder.exe
robocopy sunshine-util/build-tools/msys2 package/bin libstdc++-6.dll
robocopy DevSim/bin package/bin 

Remove-Item "./sunshine-util/screencoder.exe"
Remove-Item "./daemon/daemon.exe"