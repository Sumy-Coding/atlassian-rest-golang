install-proto:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
protoo:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        ./proto/content.proto
#        /home/malandr/GolandProjects/confluence-rest-golang/proto/content.proto
build:
	env GOOS=target-OS GOARCH=target-architecture go build package-import-path
	env GOOS=linux GOARCH=amd64 go build -o atlas
	env GOOS=windows GOARCH=amd64 go build -o atlas.exe
build-simple:
	go build -o atlas