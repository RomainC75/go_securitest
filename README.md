

### backend 

sqlc generate

1- generate init migration files 
migrate create -ext sql -dir db/migration -seq init_schema
2- fill the files : 
  (create table, drop table, etc...) // watch the order depending on the key constrains

3- check pg container is running
4- create the database directly in the container (if it's not)
