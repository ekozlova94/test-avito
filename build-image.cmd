set GOOS=linux
call go mod download
call go build -o bin/test-avito cmd/test/main.go
call docker build -t docker.io/elkozlova/test-avito:latest .
call docker push docker.io/elkozlova/test-avito:latest
