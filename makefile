deploy:
	go build -o ./build/ ./cmd/fltr/
	mv ./build/fltr ~/bin/fltr
