# grpc-logger
Simple logging service using gRPC. Events are logged to a MongoDB database. <br>
To test run `make up_build` and send a POST request to http://localhost:8080/log <br>
JSON payload: {"name":"test_name","data":"test_data"}
