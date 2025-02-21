clean:
	-rm main.exe main
	-rm screenshot.png

run:
	go run ./main.go

build:
	go build ./main.go
