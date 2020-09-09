# go-cqrs_event-sourcing
A go repository to train on CQRS and event sourcing patterns.

## Running project

### With docker

```sh
docker-compose up -d

docker-compose exec go /bin/sh

```

Then 
```sh
go build -o main .

./main
```

### Without docker

```sh
go build -o main .
./main
```