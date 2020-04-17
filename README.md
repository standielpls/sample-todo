## Todo App

## Motive 

This project will go over fundamental Go concepts, practicing idiomatic Go development 
to illustrate a typical CrowdRiff microservice.

We will follow guidelines that are presented in the following:
- https://golang.org/doc/effective_go.html
- https://github.com/golang/go/wiki/CodeReviewComments
- https://dmitri.shuralyov.com/idiomatic-go
- https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091
- https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
- https://medium.com/@benbjohnson/structuring-tests-in-go-46ddee7a25c

We can refer to these guidelines as we work through the project. The goal
is to finish the project with these practices in mind and understand why
they are important.

## Description

We are creating a Todo app that is consumed as a REST API. We are to create a 
set of endpoints that allows for the creation of a profile (a user), 
CRUD of todo lists, CRUD of todo items. 

A user can have 0 to many lists.    
A list can have 0 to many items.

We will cover HTTP handler design and database designs.

## Operations

### Create a profile 
REQUEST:
POST /v1/profile
```
{
    "name": "Kara"
}
```
RESPONSE:
```
{
    "data": {
        "id": "123"
    }
}
```

### Create a todo list
REQUEST:
POST /v1/todo
```
{
    "name": "Shopping",
    "profile_id": "123"
}
```
RESPONSE:
```
{
    "data": {
        "id": "500"
    }
}
```

### Create a todo list item
REQUEST:
POST /v1/todo/{todo_id}/item
```
{
    "name": "Milk",
    "profile_id": "123"
    "todo_id": ""
}
```
RESPONSE:
```
{
    "data": {
        "id": "900"
    }
}
```