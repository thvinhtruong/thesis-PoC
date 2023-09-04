
migrateup:
	migrate -path app/db/mysql/migration -database "mysql://root:root@tcp(localhost:3306)/UserService?parseTime=true" -verbose up

migratedown:
	migrate -path app/db/mysql/migration -database "mysql://root:root@tcp(localhost:3306)/UserService?parseTime=true" -verbose down

sqlc:
	sqlc generate