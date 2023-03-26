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

try this.

```
query {
  repository(name: "repo1", owner: "hsaki"){
    id
    name
    createdAt
    owner {
      name
    }
    issue(number:2) {
      title
    }
  }
}

```
