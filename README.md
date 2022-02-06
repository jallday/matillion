# Matillion

A simple http service that allows you to fetch and rate Star Wars movies.

### Setup 

This service runs on a golang backend with a Postgres Database. All of the services are dockerized
for ease of use. To get going run the command:

```base
make start
```

This will setup everything and start the service for you. You can you the command `make help` to help find 
any help commands this service uses. 

### Documentation
This service uses go-swagger to build and serve the docs so please go here https://goswagger.io for information about downloading. You can run the command `serve-docs` to build and serve the documenation locally.