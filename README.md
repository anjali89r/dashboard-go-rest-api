## Build Setup
Install latest version of Go


## install dependencies
```
go get github.com/gorilla/mux
go get github.com/gorilla/handlers
```
## Build solution
```
cd to parent folder
type: go build

```

### Run the API
After go build, binary will be created in the folder, to run the api solution type,
```
./go-rest-api

```

API willbe running at localhost:8000, incase if the port is used or not available please update the port in function main @
```
log.Fatal(http.ListenAndServe(":8000", handlers.CORS( allowedOrigins,allowedMethods)(router)))
```
### API End Points Available & Sample Requests
```
   POST /statscards
   Request:
   {
      "fromDate": "2018-11-20",
      "toDate": "2019-01-20"
   }

	POST /login
   {
      "email": "anjali@email.com",
      "password": "test1234"
   }
  
	POST /userchart
    Request:
   {
      "fromDate": "2018-11-20",
      "toDate": "2019-01-20"
   }
   
	POST /salesdata
    Request:
   {
      "fromDate": "2018-11-20",
      "toDate": "2019-01-20"
   }
   
	POST /emaildata
    Request:
   {
      "fromDate": "2018-11-20",
      "toDate": "2019-01-20"
   }
   ```
