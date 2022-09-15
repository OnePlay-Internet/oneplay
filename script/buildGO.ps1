$env:PKG_CONFIG_PATH = "C:\gstreamer\1.0\msvc_x86_64\lib\pkgconfig"


Set-Location daemon 
go build -o daemon.exe
Set-Location ..

Set-Location sunshine-util
go build  -o screencoder.exe  examples/go_webrtc/main.go
Set-Location ..

Set-Location webrtc-proxy
go build  -o webrtc-proxy.exe cmd/agent/main.go
Set-Location ..


Set-Location DevSim
dotnet build .
Set-Location ..

robocopy daemon package/bin daemon.exe
robocopy sunshine-util package/bin screencoder.exe
robocopy webrtc-proxy package/bin webrtc-proxy.exe
robocopy DevSim/bin/Debug/net6.0 package/bin 

Remove-Item "./sunshine-util/screencoder.exe"
Remove-Item "./webrtc-proxy/webrtc-proxy.exe"