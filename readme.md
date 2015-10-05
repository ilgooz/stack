# Stack [![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](http://godoc.org/github.com/ilgooz/stack)
> Stack is a lightweight not-a-framework RESTful API that allows you to easily get started with your next project and shows some of the best practices about building RESTful APIs in Go.

![](https://cdn.rawgit.com/ilgooz/stack/master/logo.jpg)

## What is included?
* Proper using of middlewares
* Learn how to organize your files & folders structure
* User & Token APIs out of the box
* Authentication via tokens
* Keep context by *http.Request with gorilla/context
* Form parsing with gorilla/schema & validation with go-playground/validator
* MongoDB with go-mgo/mgo
* Command-line configuration
* App versioning over latest git commit hash
* ...

### Usage
##### Run
  make run
##### Run via Docker Compose
  docker-compose up
##### Stop
  ^C *ctrl+c*
##### Hard Stop
  ^C^C *twice*

### Out of The Box Endpoints
* post /users
* get /users
* get /users/{id}
* get /me
* post /tokens
* get /version

### Authentication
Send your *access_token* over *X-Auth-Token* header

> For more information read the source code and do requests on endpoints

## Best practices resources about building APIs
* GopherCon 2015: Blake Caldwell - Uptime: Building Resilient Services with Go https://www.youtube.com/watch?v=PyBJQA4clfc

## Contribute
* Share your ideas by opening new issues
* Feel free to ask for feature requests
