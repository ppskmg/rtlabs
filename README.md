#DB

https://github.com/golang-migrate/migrate
https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md



```
createdb reastapi_dev
createdb restapi_test
```
## Create / delete local current migrations
```
migrate -database "postgresql://localhost:5432/restapi_dev?sslmode=disable" -path ../go/src/rtlabs/migrations up
migrate -database "postgresql://localhost:5432/restapi_dev?sslmode=disable" -path ../go/src/rtlabs/migrations down
```

## Create new table / migration

```
migrate create -ext sql -dir ../go/src/boopcar-back/migrations -seq create_NAME_table
```


# Tree
    cmd
    configs
    internal
      app
        apiserver
          auth
          routers.file
        model
          models.file
        store
    migrations
    pkg
    makefile
