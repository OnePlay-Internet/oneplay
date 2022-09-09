Copy-Item  -Path .\sunshine-util\cmake\CMakeLists.txt.msvc -Destination .\sunshine-util\CmakeLists.txt
Set-Location sunshine-util

mkdir buildmsvc
Set-Location buildmsvc
cmake .. -G "Visual Studio 17 2022"

robocopy ../.. . msvc.bat


Set-Location ../..
