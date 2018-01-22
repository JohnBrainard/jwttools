
BINARY := jwttools

windows:
	mkdir -p release/windows
	GOOS=windows GOARCH=amd64 go build -o release/windows/$(BINARY).exe

linux:
	mkdir -p release/linux
	GOOS=linux GOARCH=amd64 go build -o release/linux/$(BINARY)

darwin:
	mkdir -p release/darwin
	GOOS=darwin GOARCH=amd64 go build -o release/darwin/$(BINARY)

release: windows linux darwin
	zip -r jwttools-all.zip release/

clean:
	rm -r release
