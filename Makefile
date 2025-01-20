build:
	mkdir -p ./bin
	go build -o ./bin/demoservice cmd/server/main.go


run: build
	chmod +x ./bin/demoservice
	lsof -ti :8080 | xargs kill -9
	./bin/demoservice
