# Cannonical 

Cannoncial is a tempalate for a go grpc server, with structured logging and prom metrics out of the box. 
A similar http chi server is available on the `chi` branch.


Vist https://github.com/smantic/cannonical and click Use this template to use it as a template.


```
$ ./cannonical serve -h
Usage of serve:
  -address string
    	address to run the server on (default "localhost")
  -debugport string
    	port for http server serving prom metrics and pprof to run on (default "8081")
  -port string
    	port to run the server on (default "8080")
```


set up the service with a grpc gateway with https://github.com/grpc-ecosystem/grpc-gateway
> protoc-gen-grpc-gateway


rename the module with: 
```
go mod edit -module <NEW_NAME>
find . -type f -name '*.go' -exec sed -i -e 's,<NEW_NAME>,<OLD_NAME>,g' {} \;
```
