# docker network rm roachnet
# docker network create -d bridge roachnet
docker-compose -f db-compose.yml -p gostore up -d --force-recreate 
docker exec gostore_roach1_1 ./cockroach sql --execute="CREATE USER IF NOT EXISTS customeruser; CREATE DATABASE IF NOT EXISTS customerdb; GRANT ALL ON DATABASE customerdb TO customeruser; GRANT CREATE ON DATABASE customerdb TO customeruser;" --insecure