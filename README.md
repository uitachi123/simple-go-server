# go-server
simple http server

# features
- [x] use `ServeMux` for routing
- [x] include testing cases examples
- [x] set up in-memory database to demonstrate database access

# clean, test, build and run
```
make all
```

# server options
```
Usage of ./go-server:
  -logging string
        logging level (default "INFO")
  -port string
        listening port (default "8080")
```

# server APIs
```
localhost:8080/echo/testdata
localhost:8080/users
```