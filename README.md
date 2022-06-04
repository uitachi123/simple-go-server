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
Home:
```
http://localhost:8080/
```
response:
```
Welcome to new server!
```

API:
```
localhost:8080/echo/testdata
```
response:
```
testdata!
```

API:
```
localhost:8080/users
```
response:
```
[
      {
            email: "alice@test.com",
            name: "Alice",
      },
      {
            email: "bob@test.com",
            name: "Bob",
      },
]
```