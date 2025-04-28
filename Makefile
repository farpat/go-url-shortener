.DEFAULT_GOAL := help

.PHONY: test help

# ANSI color codes
BOLD := \033[1m
RESET := \033[0m
GREEN := \033[32m
RED := \033[31m
BLUE := \033[34m

help: ## Show this help message	
	@awk 'BEGIN {FS = ":.*##"; } /^[a-zA-Z_-]+:.*?##/ { printf "$(BOLD)$(BLUE)%-7s$(RESET) %s\n", $$1, $$2 }' $(MAKEFILE_LIST) | sort

test: ## Run tests
	go test -v ./tests/... $(if $(FILTER),-run $(FILTER))
