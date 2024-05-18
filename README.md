# REQHTTP

Terminal app for http request to API, build by stack Golang.

### USAGE

basic usage
```
reqhttp <METHOD> <URL> <DATA>
```
**GET**
for the GET method include the **"GET"** methode and the url **"http://localhost:8000"** 
example:
```
reqhttp GET https://dummyjson.com/products?limit=3
```
**POST**
for the POST method, include the **"POST"** method and the url **"http://dummyjson.com/products/add"** as well as JSON data **'{"foo": "bar"}'**
example:
```
reqhttp POST https://dummyjson.com/add '{"foo": "bar"}'
```

**Note**: Currently there are only two http request methods GET and POST. PUT, PATCH, DELETE are still under dvelopment. 

