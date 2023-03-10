build: 
	go build -o test-assigment cmd/server/server.go
	go build -o test-assigment cmd/client/client.go
up: 
	docker compose up -d

down:
	docker compose down

run: 
	./test-assigment