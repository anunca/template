.PHONY: help
.DEFAULT_GOAL := help

include .env
-include .env.local

ENV ?= dev
TAG ?= ${ENV}
DOCKER_SHELL ?= sh
BUILDER ?= composer
TEMPLATE ?= symfony/website-skeleton
VERSION ?= 6.1

dc_files = docker-compose.override.yml

ifeq ($(ENV),prod)
	dc_files += docker-compose.prod.yml
else
	dc_files += docker-compose.yml
endif

dc = docker compose $(foreach f,$(dc_files),-f $(f))

export ENV
export TAG

help h: ## Show help
	@awk 'BEGIN {FS = ":.*##"} \
	/^[a-zA-Z0-9_. -]+:.*##/ { \
		split($$1,a," "); \
		printf "%-20s %s\n", a[1], $$2 \
	} \
	/^##@/ {printf "\n%s\n", substr($$0,5)}' $(MAKEFILE_LIST)

##@ Docker
config c: ## Show config
	$(dc) config

build b: ## Build images
	$(dc) build

start s: ## Start containers
	$(dc) up -d

stop st: ## Stop containers
	$(dc) down -t0

restart r: stop start ## Restart

run: ## Run shell
	$(dc) run --rm --entrypoint $(DOCKER_SHELL) app

shell sh: ## Exec shell
	$(dc) exec app $(DOCKER_SHELL)

logs l: ## Follow logs
	$(dc) logs -f

ps: ## List containers
	$(dc) ps -a

clean.all ca: stop ## Cleanup all docker
	docker system prune -fa && docker volume prune -fa

##@ Composer
composer.install: ## Install
	$(dc) run --rm app composer install

##@ Project
project: build ## Create
ifeq ($(BUILDER), composer)
	@$(dc) run --rm --entrypoint composer app create-project ${TEMPLATE}:^v${VERSION} /var/www/html  --no-interaction
else ifeq ($(BUILDER), symfony)
	@$(dc) run --rm --entrypoint symfony app new app ${TEMPLATE} --version=${VERSION} --dir=/var/www/html --no-interaction
else
	@echo unkonwn builder: $(BUILDER)
endif