$Env:GOOS="windows"; $Env:GOARCH="amd64";
go build -ldflags -H=windowsgui
