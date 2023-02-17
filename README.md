```bash
docker run \
--rm \
--name postgresql-intro-gosamples \
-p 5437:5432 \
-e POSTGRES_PASSWORD=mysecretpassword \
-e POSTGRES_DB=gorm postgres
```