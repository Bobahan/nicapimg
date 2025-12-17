.PHONY: run
run:
	clear && go run ./cmd/product-img/main.go

.DEFAULT_GOAL := run