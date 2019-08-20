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

In peroccess

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
