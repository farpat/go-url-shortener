.DEFAULT_GOAL := help

.PHONY: help test run update-certificates install

-include .env

# ANSI color codes
BOLD := \033[1m
RESET := \033[0m
GREEN := \033[32m
RED := \033[31m
BLUE := \033[34m

help: ## Show this help message
	@awk 'BEGIN {FS = ":.*##"; } /^[a-zA-Z_-]+:.*?##/ { printf "$(BOLD)$(BLUE)%-7s$(RESET) %s\n", $$1, $$2 }' $(MAKEFILE_LIST) | sort

test: ## Run tests
	@go test ./tests/... $(if $(FILTER),-run $(FILTER))

run: ## Run the application
	@echo "🚀  Server starting on https://localhost:$(PORT)$(RESET)"
	@PORT=$(PORT) air -log.silent=true

update-certificates: ## Generate SSL certificates
	@mkcert -key-file certs/key.pem -cert-file certs/cert.pem localhost 127.0.0.1
	@mkcert -install > /dev/null 2>&1
	@echo "✅ $(GREEN)SSL certificates updated$(RESET)!"

install: ## Install dependencies
	@go mod download
	@echo "✅ $(GREEN)Dependencies installed$(RESET)!"
