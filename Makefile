#----------------------------------------------------------------------------
# Makefile
# use the following command to install on a Mac:
# brew install pre-commit
#----------------------------------------------------------------------------

.PHONY: help
.DEFAULT_GOAL := help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
	| awk 'BEGIN {FS = ":.*?## "}; \
	{printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

precommit-install: ## install pre-commit
	@pre-commit install

precommit-run: ## run all pre-commit hooks
	@pre-commit run -a

build: ## build the boilr executable
	@go mod tidy
	@go build
	
clean: ## run clean up
	@pre-commit clean
