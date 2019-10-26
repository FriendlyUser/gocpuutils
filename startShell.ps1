choco install -y nodejs
$env:Path = [System.Environment]::GetEnvironmentVariable("Path","Machine") + ";" + [System.Environment]::GetEnvironmentVariable("Path","User")
node --version || true || npm --version
go get ./...
go get github.com/wailsapp/wails/cmd/wails
wails build
ls