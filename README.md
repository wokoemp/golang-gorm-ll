## start project
1. download libraries
```bash
 go mod download
```
2. run postgres
```bash
    docker run \
    --rm \
    --name postgresql-intro-gosamples \
    -p 5437:5432 \
    -e POSTGRES_PASSWORD=mysecretpassword \
    -e POSTGRES_DB=gorm postgres
```
3. using navicat connect postgre database
4. 