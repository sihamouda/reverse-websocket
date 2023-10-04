start:
ifdef service
	@docker compose -f docker-compose.yml -f docker-compose.dev.yml up $(service) -d --build --force-recreate $(service)
else
	@docker compose -f docker-compose.yml -f docker-compose.dev.yml down --remove-orphans
	@docker compose -f docker-compose.yml -f docker-compose.dev.yml  up -d --build --force-recreate
endif

stop:
ifdef service
	@docker compose -f docker-compose.yml -f docker-compose.dev.yml down $(service) --remove-orphans
else
	@docker compose -f docker-compose.yml -f docker-compose.dev.yml down --remove-orphans
endif