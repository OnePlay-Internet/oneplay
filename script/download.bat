powershell Invoke-WebRequest -Uri https://gstreamer.freedesktop.org/data/pkg/windows/1.20.3/msvc/gstreamer-1.0-msvc-x86_64-1.20.3-merge-modules.zip -OutFile package/gstreamer.zip 
cd package && powershell Expand-Archive gstreamer.zip -DestinationPath . 
robocopy ./projects/repos/cerbero.git/1.20 msm
del /Q /S projects
cd ..