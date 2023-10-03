migrateUser:
	migrate -path UserService/app/db/mysql/migration -database "mysql://root:root@tcp(localhost:3306)/UserService?parseTime=true" -verbose down
	migrate -path UserService/app/db/mysql/migration -database "mysql://root:root@tcp(localhost:3306)/UserService?parseTime=true" -verbose up

migrateStudy:
	migrate -path StudyService/app/db/mysql/migration -database "mysql://root:root@tcp(localhost:3306)/StudyService?parseTime=true" -verbose down
	migrate -path StudyService/app/db/mysql/migration -database "mysql://root:root@tcp(localhost:3306)/StudyService?parseTime=true" -verbose up