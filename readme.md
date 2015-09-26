# Stack
> Stack is a lightweight not-a-framework RESTful API that allows you to easily get started with your next project and shows to you some of the best practices about building RESTful APIs in Go.

![](https://cdn.rawgit.com/ilgooz/stack/master/logo.jpg)

## What is included?
* Proper using of middlewares
* Learn how to organize your files & folders structure
* User & Token APIs out of the box
* Authentication via tokens
* Keep context by *http.Request via gorilla/context
* MongoDB via go-mgo/mgo
* Command-line configuration for the app
* Versioning over latest git commit hash

### Usage
* to start the server simply: *make run*
* to stop simply: *^C^C* (twice)

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
