# Phonebook

Сборка:
go build Book.go getCountry.go handler.go postgres.go

Параметы запуска:  Book.exe --help

Usage of Book.exe:

  -d    Name Database
  
  -p    port (default "5432")
  
  -pass DB password
  
  -s    serverDB (default "localhost")
  
  -u    Database user

запуск с параметрами:
Book.exe -s localhost -p 5432 -d postgres -u dbuser -pass dbpassword

Призапуске в БД создается таблица Phonebook и выполняется загрузка данных.

