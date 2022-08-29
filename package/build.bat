robocopy sunshine-util\build  package\bin screencoder.exe
robocopy sunshine-util\build  package\lib libscreencoderlib.a
robocopy sunshine-util\dll    package\bin libstdc++-6.dll
robocopy sunshine-util\build\pre-compiled\windows\lib package\ffmpeg 
robocopy sunshine-util\capture\windows\directx package/directx

cd daemon 
go build -o daemon.exe
cd ..

cd sunshine-util
go build  examples/go_webrtc
cd ..

robocopy daemon package/bin daemon.exe
robocopy sunshine-util package/bin go_webrtc.exe