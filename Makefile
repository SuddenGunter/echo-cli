build:
	@CGO_ENABLED=0 go build -installsuffix cgo -o echo-cli
test:
	go test ./...	
demo:
	./echo-cli auth -t="54686973206973206d7920626f6f6d737469636b"
	./echo-cli user create fedor
