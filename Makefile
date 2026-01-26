build-go-protoc:
	docker build --tag go-protoc --file protoc.Dockerfile .

protoc-docker-version:
	docker run --rm go-protoc \
		sh -c "which protoc && which protoc-gen-go && which protoc-gen-go-grpc && protoc --version && protoc-gen-go-grpc -version"

protoc-compile-docker:
	docker run --rm --volume $(CURDIR):/workdir --workdir /workdir go-protoc protoc \
		--go_out=internal \
		--go-grpc_out=internal \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		proto/random-number-gen/v1/*.proto

protoc-version:
	which protoc && which protoc-gen-go && which protoc-gen-go-grpc && protoc --version && protoc-gen-go-grpc -version

install-protoc-gen-go:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

brew-install-protoc:
	brew install protobuf

protoc-compile:
	protoc
		--go_out=internal \
		--go-grpc_out=internal \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
        proto/random-number-gen/v1/*.proto

