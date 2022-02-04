# Tasks List API in Go

Based on youtube tutorial: https://www.youtube.com/watch?v=pQAV8A9KLwk

[![Golang, REST API CRUD con go
](https://img.youtube.com/vi/configuroweb/0.jpg)](https://www.youtube.com/watch?v=pQAV8A9KLwk)

--- 

In this branch:
* Add go modules
* Add mongodb driver
* Add .env file
* Handler's separation from main file
* Add Task model for operations with the database
* Add make file

#Configure database
In .env file: add your settings about your DB

# Run
Install dependencies
```
make tidy
```

To run:
```
make run
```

