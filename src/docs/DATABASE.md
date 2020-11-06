## container
### mysql cli

```
# db container up
$ docker-compose exec db mysql -h db -u gather -p

Enter password: gather
```

## migration
### create migration file

```
# create table
$ docker-compose run --rm golang sql-migrate new -config=config/dbconfig.yml <file name>

or

$ docker-compose exec golang sql-migrate new -config=config/dbconfig.yml <file name>
```

### migrate up

```
$ docker-compose run --rm golang sql-migrate up -config=config/dbconfig.yml

or

$ docker-compose exec golang sql-migrate up -config=config/dbconfig.yml
```

#### options

```
  -config=dbconfig.yml   Configuration file to use.
  -env="development"     Environment.
  -limit=0               Limit the number of migrations (0 = unlimited).
  -dryrun                Don't apply migrations, just print them.
```

### rollback

```
$ docker-compose run --rm golang sql-migrate down

or

$ docker-compose exec golang sql-migrate down
```

### status check

```
$ docker-compose run --rm golang sql-migrate status

or

$ docker-compose exec golang sql-migrate status
```

### redo

```
$ docker-compose run --rm golang sql-migrate redo

or

$ docker-compose exec golang sql-migrate redo
```
