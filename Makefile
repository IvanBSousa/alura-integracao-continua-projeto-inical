.PHONY: lint test

lint:
	docker run --rm --network host   -v "$(CURDIR):/app"   -w /app   golangci/golangci-lint   golangci-lint run controllers/ database/ models/ routes/
test:
	docker compose exec app go test main_test.go
start:
	docker compose up -d
finish:
	docker compose down
ci: start lint test finish