robocopy sunshine-util\build  package\bin screencoder.exe
robocopy sunshine-util\build  package\lib libscreencoderlib.a
robocopy sunshine-util\dll    package\bin libstdc++-6.dll
robocopy sunshine-util\build\pre-compiled\windows\lib package\ffmpeg 
robocopy sunshine-util\capture\windows\directx package/directx

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


Remove-Item ".\daemon\daemon.exe"
Remove-Item ".\sunshine-util\screencoder.exe"
Remove-Item ".\webrtc-proxy\webrtc-proxy.exe"