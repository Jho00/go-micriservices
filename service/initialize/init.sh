docker run --name go-microservices -p 5432:5432 -e POSTGRES_PASSWORD=go-serivces -d postgres

docker exec -it go-microservices psql -U postgres