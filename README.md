# rbac-example

Example Golang application with Role Based Access Control.

It is a HTTP Server, mimicking a string:string dictionary, with the following endpoints:

- `GET /v1/map/{key}`
  - Returns a string `value` with status 200 if found
  - Returns `NOT_FOUND` with status 404 if not found
- `POST /v1/map/{key}`
  - The POST body is the `value`
  - Returns the same string `value` with status 200 if set
  - Returns `UNAUTHORISED` with status 401 if not allowed

## How to Run and Test

### Running in memory

Start the server running on localhost:8080.

```bash
$ go run cmd/main.go
```

### Running with Permit.io

[Permit.io](https://permit.io) is a third party authorisation provider, that provides an easy to integrate with API for role based access control.
For our purposes, we will use this as an alternative for our in memory solution.

To get started, you'll need to grab an API key from [https://app.permit.io/settings/api-keys](https://app.permit.io/settings/api-keys) (you'll need to create a free account first!).
Store this somewhere:

```bash
$ PERMIT_API_KEY=permit_key_skkdfbljsdfudfuybdfuygoydfubydkfub
```

Set up the default data:

```bash
$ go run scripts/permit_setup.go -permit_api_key=$PERMIT_API_KEY
```

Then, start the server running on localhost:8080.

```bash
$ go run cmd/main.go -permit_api_key=$PERMIT_API_KEY
```

### Run all tests

There is a helper script to check all authorisation is set up correctly, this should be run in a separate terminal.

```bash
$ scripts/run_tests.sh
```

### Value not found

Try to get the value at key "notexist" with user "alice":

```bash
$ curl -H "User: alice" -X GET http://localhost:8080/v1/map/notexist

> NOT_FOUND
```

### Setting a value

Try to set the value to "world" at key "hello" with user "alice":

```bash
$ curl -H "User: alice" -X POST http://localhost:8080/v1/map/hello -d "world"

> world
```

### Getting a value

Try to get that value you just created at key "hello" with user "alice":

```bash
$ curl -H "User: alice" -X GET http://localhost:8080/v1/map/hello

> world
```

### Unauthenticated read

User "bob" only has the role of "writer", so they are not able to use the GET endpoint:

```bash
$ curl -H "User: bob" -X GET http://localhost:8080/v1/map/hello

> UNAUTHORISED
```

User "charli" only has the role of "reader", so they are not able to use the POST endpoint:

```bash
$ curl -H "User: charli" -X POST http://localhost:8080/v1/map/hello -d "world"

> UNAUTHORISED
```
