## GraphQL Server with Golang (for practice)

### How to start

1. `make dev`

`http://localhost:8080` will be served.

## How to use

```
query {
  todos {
    id
    text
    done
    user {
      name
    }
  }
}
```

Please query the following.

```
mutation {
  createTodo(input: {
    text: "実用Go言語最後まで"
    userId: "1"
  }){
    id
    text
    done
    user {
      id
      name
    }
  }
}
```
