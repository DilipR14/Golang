postgres : 
       docker run-name postgres12-p 5432:5432 POSTGRES USEll-root -POSTGRES PASSWORD-secret-d postgres:12-

mysql:

     docker run-name mysql8 -p 3306:3306 MYSOL ROOT PASSWORD secret -d mysqli

createdb:

  docker exec-it postgres12 createdb-username=root-owner-root simple bank

dropdb:

   docker exec-it postgres12 dropdb simple_bank


migrateup:

    migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple bank?side-disable--verbose



migratedown:

  migrate -path db/migration -database "postgresql://root:secret@localhost15432/simple banks node-disabla" -verbose down


sqlc:

 sqlc generate 


test :

  go test-v-cover ./...

server:

 go run main.go

PHONY:

 postgres createdb dropdb migrateup migratedown sqlc test server