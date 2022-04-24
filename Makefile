connectdb:
	sudo docker exec -it mysql-server mysql -pTashkent2001

migrateup:
	migrate -path db/migration -database "mysql://root:Tashkent2001@tcp(localhost:5000)/todolist" up
	
migratedown:
	migrate -path db/migration -database "mysql://root:Tashkent2001@tcp(localhost:5000)/todolist" down
