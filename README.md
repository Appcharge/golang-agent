# Golang agent template

First of all you need to set up all environment variables (.env.sample as an example).

Where to get the key
Go to the admin panel, open the integration tab, if you are using signature authentication - copy the key from the primary key field, if you are using encryption, take the key from the primary key field and the IV from the secondary key field.

### How to run manually (go1.18 required)
go run cmd/main.go

### How to build docker image
docker build -t appcharge/golang-template .

How to run the container locally
docker run -p 8000:8000 appcharge/golang-template