# Phonebook

Сборка:
go build Book.go getCountry.go handler.go postgres.go

Параметы запуска:  Book.exe --help
Usage of Book.exe:
  -d string
        Name Database
  -p string
        port (default "5432")
  -pass string
        DB password
  -s string
        serverDB (default "localhost")
  -u string
        Database user


запуск с параметрами:
Book.exe -s localhost -p 5432 -d postgres -u dbuser -pass dbpassword

Призапуске в БД создается таблица Phonebook и выполняется загрузка данных.

