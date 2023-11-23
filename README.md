## Тестовое задание Lamoda

API для работы с товарами на складе

### Технологии:
- Go
- PostgreSQL
- Docker

### Библиотеки:

- github.com/joho/godotenv
- github.com/kelseyhightower/envconfig
- github.com/go-chi/chi/v5

### Запуск сервиса:
- git clone https://github.com/taslav8/lamoda_task
- переименовать файл cfg.txt в cfg.env
- изменить данные входа своей бд в докер файлах и переменные окружения
- make postgres
- make createdb
- make migratecreate
- make migrateup
- make migrateup
- cd cmd 
- go run main.go