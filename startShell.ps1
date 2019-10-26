Set-ExecutionPolicy RemoteSigned
choco install -y nodejs
choco install -y mingw
$env:Path = [System.Environment]::GetEnvironmentVariable("Path","Machine") + ";" + [System.Environment]::GetEnvironmentVariable("Path","User") || true
node --version || true || npm --version
go get ./...
go get github.com/wailsapp/wails/cmd/wails
wails build
ls