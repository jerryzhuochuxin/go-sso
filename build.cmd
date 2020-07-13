rd /s /q output && md output && md output\bin && md output\sdk
xcopy sdk\*.* output\sdk /s /e /c /y /h /r

cd jwtService && go build . && cd .. && move jwtService\jwtService.exe output\bin
cd registerService && go build . && cd .. && move registerService\registerService.exe output\bin
