
tools\thrift-0.14.2.exe -r --gen go:package_prefix="credits-go/" thrift-interface-definitions\api.thrift 
cd gen-go
xcopy "*.*" "..\"  /E /I
cd ..
rd /s /q gen-go