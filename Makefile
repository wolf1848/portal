VERSION         := 1.0.0
CURRENT_DIR     := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
CURRENT_DATE     := $(shell date +'%y.%m.%d_%H:%M:%S')
DB_DUMP_FILE    := db.sql
DB_CONTAINER    := taxi_database
PHP_CONTAINER   := taxi_php


.SILENT: ;               # не выводить исполняемые команды, даже без использования @
.ONESHELL: ;             # исполнять рецепты в том же самом shell
.NOTPARALLEL: ;          # ждать завершения исполнения каждого рецепта перед запуском следующего
.EXPORT_ALL_VARIABLES: ; # отправить все переменные в shell

ifneq (,$(wildcard ./.env))
    include .env
    export
endif


test:
	echo $(CURRENT_DIR)/docker/$(ENV)/docker-compose.yml

copyenv:
	if ! test -f ./.env; then
	  cp ".env.example" ".env" && \
      echo "Variables copied"
	fi

checkenv:
	if ! test -f ./.env; then
	  echo "Not variables file" && \
      exit 1
	fi

build:
	docker compose -f $(CURRENT_DIR)/docker/$(ENV)/docker-compose.yml build --no-cache

up:
	docker compose -f $(CURRENT_DIR)/docker/$(ENV)/docker-compose.yml up -d

run:
	docker compose -f $(CURRENT_DIR)/docker/$(ENV)/docker-compose.yml up

down:
	docker compose -f $(CURRENT_DIR)/docker/$(ENV)/docker-compose.yml down

composer:
	docker exec -it $(PHP_CONTAINER) composer install

frontinit:
	docker exec -it $(PHP_CONTAINER) npm --prefix=./portal install

frontrun:
	docker exec -it $(PHP_CONTAINER) npm --prefix=./portal run serve

dbconfig:
	echo "[client]" > ./docker/mariadb/.root.cnf; \
	echo "user=$(DB_USERNAME)" >> ./docker/mariadb/.root.cnf; \
	echo "password=$(DB_PASSWORD)" >> ./docker/mariadb/.root.cnf; \
	echo "host=localhost" >> ./docker/mariadb/.root.cnf; \
	docker cp ./docker/mariadb/.root.cnf $(DB_CONTAINER):/.root.cnf
	rm ./docker/mariadb/.root.cnf

dbbackup:
	docker exec -it $(DB_CONTAINER) /usr/bin/mariadb-dump --defaults-extra-file=/.root.cnf portal > ./backup/$(CURRENT_DATE)_$(DB_DUMP_FILE)

dbrestore:
	cat ./backup/$(DB_DUMP_FILE) | docker exec -i $(DB_CONTAINER) /usr/bin/mariadb --defaults-extra-file=/.root.cnf $(DB_DATABASE)

backdbrestore:
	cat ./backup/$(DB_DUMP_FILE) | docker exec -i $(DB_CONTAINER) /usr/bin/mariadb --defaults-extra-file=/.root.cnf portal_back

buildfront:
	rm -rf ./public/css && \
	rm -rf ./public/js  && \
	rm -rf ./resources/views/welcome.blade.php && \
	docker exec -it $(PHP_CONTAINER) npm --prefix ./portal run build

certbottest:
	docker compose -f $(CURRENT_DIR)/docker/$(ENV)/docker-compose.yml run --rm taxi_certbot certonly --webroot --webroot-path /var/www/certbot/ -d new.taxi-portal.ru --dry-run

certbot:
	docker compose -f $(CURRENT_DIR)/docker/$(ENV)/docker-compose.yml run --rm taxi_certbot certonly --webroot --webroot-path /var/www/certbot/ -d new.taxi-portal.ru

certbot_update:
	docker compose run --rm certbot renew

tunel:
	ssh -N -f -L 3306:localhost:3306 root@taxi-portal.ru