We have `nginx` as api gateway for our services. Each request redirected to services `auth`. `Auth` services return `HTTP` status code.  Via this status codes `nginx` can deside wich make with request: reject or redirect to target services. Status code list:

1. 401 - Unauthorized
2. 403 - Forbidden
3. 500 - Internal Server Error
4. 200 - OK

If return status `1-3` `nginx` rejecting request. Otherwise (4) `nginx` redirect request to requested servic.
