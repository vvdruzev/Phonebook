# Phonebook

Сборка :

sudo docker-compose up -d --build


Настройки прокси:

    docker-compose.yml

        HTTP_PROXY: "http://user:password@proxy:port"


Поиск телефонного кода по имени страны:

curl "ip-container":8080/code/"Countryname"

Загрузка\обновление данных в БД о странах:

curl -X POST "ip-container":8080/reload


Просмотр логов: docker logs -f phonebook_book_1


Запуск тестов: go test -v
