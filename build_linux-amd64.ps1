$Env:GOOS="linux"; $Env:GOARCH="amd64";
go build -ldflags -H=windowsgui
