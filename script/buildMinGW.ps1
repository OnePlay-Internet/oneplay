Copy-Item  -Path .\sunshine-util\cmake\CMakeLists.txt.mingw -Destination .\sunshine-util\CmakeLists.txt
Set-Location sunshine-util

mkdir buildmingw
Set-Location buildmingw
cmake .. -G "Ninja"
ninja
Set-Location ../../
