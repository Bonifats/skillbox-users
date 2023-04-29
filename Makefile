# -- Settings ------------------------------------------------------

compose_local = -f docker-compose.yaml

# -- Dev -----------------------------------------------------------

dev-stop:
	docker-compose $(compose_local) stop

dev-up-down:
	docker-compose $(compose_local) stop
	docker-compose $(compose_local) up -d

dev-up-build:
	docker-compose $(compose_local) up --build -d
