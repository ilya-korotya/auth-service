# Auth services

We have `nginx` as api gateway for our services. Each request redirected to services `auth`. `Auth` services return `HTTP` status code.  Via this status codes `nginx` can deside wich make with request: reject or redirect to target services. Status code list:

1. 401 - Unauthorized
2. 403 - Forbidden
3. 500 - Internal Server Error
4. 200 - OK

If return status `1-3` `nginx` rejecting request. Otherwise (4) `nginx` redirect request to requested servic.

## Run
```
make
```

## API

### Public:

1. Register new user:
   ```
   curl --header "Content-Type: application/json" --request POST --data '{"name":"Alexandr","password":"secret"}' http://localhost:8080/user/registration -v
   ```
   Body response:
   ```
   {
     "id":"faeaa9cb-7e3f-4c75-b045-de16903ebacf",
     "name":"Alexandr",
     "email":null,
     "password":"2bb80d537b1da3e38bd30361aa855686bde0eacd7162fef6a25fe97bf527a25b",
     "Token":
       {
         "ID":"7dbad679-22a8-4e73-b403-70ad32e5de3f",
         "Content":"DixLMDCtqudWxtTlbhGrlUXjfyHAzyJwhnyyPFEaGfLZYMHuknXIWQiqGEqkGFP",
         "UserID":"faeaa9cb-7e3f-4c75-b045-de16903ebacf",
         "CreatedAt":"0001-01-01T00:00:00Z",
         "UpdatedAt":"0001-01-01T00:00:00Z",
         "ExpiredAt":"2019-08-24T17:49:03.824192Z"},
         "created_at":"0001-01-01T00:00:00Z",
         "updated_at":"0001-01-01T00:00:00Z"
        }
   }
   ```

### Internal:
1. Get user by token:
   ```
   curl -H "token: your_token" http://localhsot:8081/v1/user
   ```
   This request must be run in private network. You can attach docker to available network and make request:
   ```
   docker run --network auth-service-net byrnedo/alpine-curl -H "token: secret_token_2" http://auth-service:8081/v1/user -v
   ```
   Body response:
   ```
   {
     "id":"1049a6b9-f0b2-4f69-9cee-56862fb8ad95",
     "name":"Foo",
     "email":null,
     "password":"bar",
     "created_at":"2019-08-15T10:19:21.203475Z",
     "updated_at":"2019-08-15T10:19:21.203475Z"
   }
   ```
