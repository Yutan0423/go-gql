dev:
	@cd docker && docker-compose up --build

db_exec:
	@cd docker && docker-compose exec mysql /bin/bash
	# mysql -u user -ppassword

down:
	@cd docker && docker-compose down --remove-orphans

gqlgen:
	@cd app && gqlgen

sqlc_generate:
	@cd app && sqlc generate

# db起動中に実行
migrate:
	@cd app && dbmate up

tidy:
	@cd app && go mod tidy