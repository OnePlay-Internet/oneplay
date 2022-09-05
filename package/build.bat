robocopy sunshine-util\build  package\bin screencoder.exe
robocopy sunshine-util\build  package\lib libscreencoderlib.a
robocopy sunshine-util\dll    package\bin libstdc++-6.dll
robocopy sunshine-util\build\pre-compiled\windows\lib package\ffmpeg 
robocopy sunshine-util\capture\windows\directx package/directx

cd daemon 
go build -o daemon.exe
cd ..

cd sunshine-util
go build  -o screencoder.exe  examples/go_webrtc/main.go
cd ..

cd webrtc-proxy
go build  -o webrtc-proxy.exe cmd/agent/main.go
cd ..


cd DevSim
dotnet build .
cd ..

robocopy daemon package/bin daemon.exe
robocopy sunshine-util package/bin screencoder.exe
robocopy webrtc-proxy package/bin webrtc-proxy.exe
robocopy DevSim/bin/Debug/net6.0 package/bin 


del ".\daemon\daemon.exe"
del ".\sunshine-util\screencoder.exe"
del ".\webrtc-proxy\webrtc-proxy.exe"