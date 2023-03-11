# 64-bit
GOOS=linux GOARCH=amd64 go build -o bin/go-post-amd64-linux main.go
GOOS=darwin GOARCH=amd64 go build -o bin/go-post-amd64-darwin main.go
GOOS=windows GOARCH=amd64 go build -o bin/go-post-amd64.exe main.go

# 32-bit
GOOS=linux GOARCH=386 go build -o bin/go-post-386-linux main.go
GOOS=windows GOARCH=386 go build -o bin/go-post-386.exe main.go