## GraphQL Server with Golang (for practice)

### How to start

1. `make dev`

`http://localhost:8080` will be served.

## How to use

Please query the following.

```
query {
  user(name: "hsaki") {
    id
    name
  }
}
```

Headers

```
{
  "Authorization": "UT_hsaki"
}
```
