Для запуска нужен Docker Desktop

Команды:
* make dev-up-build - собирает и стартует
* make dev-up-down - останавливает и стартует
* make dev-stop - останаваливает


Запросы:
* Создание первого пользователя:
``
curl --location 'http://127.0.0.1/create' \
--header 'Content-Type: application/json' \
--data '{
"name": "Bogdan",
"age": 30
}'
``

* Создание второго пользователя:
``
curl --location 'http://127.0.0.1/create' \
--header 'Content-Type: application/json' \
--data '{
"name": "Sergey",
"age": 34
}'
``

* Связывание пользователей:
``
curl --location 'http://127.0.0.1/make_friends' \
--header 'Content-Type: application/json' \
--data '{
"source_id": 1,
"target_id": 2
}'
``

* Удаление пользователя:
``
curl --location --request DELETE 'http://127.0.0.1/user' \
--header 'Content-Type: application/json' \
--data '{
"target_id": 2
}'
``

* Обновление пользователя:
``
curl --location --request PUT 'http://127.0.0.1/1' \
--header 'Content-Type: application/json' \
--data '{
"name": "Anton",
"age": 29
}'
``

* Получение списка друзей пользователя:
``
curl --location 'http://127.0.0.1/friends/2'
``
