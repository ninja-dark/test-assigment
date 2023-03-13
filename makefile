server: 
	go build -o test-assigment_server cmd/server/server.go
	./test-assigment_server
client:
	go build -o test-assigment_client cmd/client/client.go
	./test-assigment_client
up: 
	docker compose up -d

down:
	docker compose down



