run:
	docker-compose up  --remove-orphans --build



run_fetcher:
	HTTP_ADDR=:9090 go run -race cmd/fetcher/main.go
