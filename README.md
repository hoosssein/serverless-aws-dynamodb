#Introduction

this project use AWS Api Gateway, AWS Lambda and AWS DyanmoDB to create, update and get Device Information.

The project is written with Go and uses Serverless framework to organize and deploy the code on the cloud.

##API

```
HTTP POST
URL: https://<api-gateway-url>/api/devices
Body (application/json):
{
  "id": "/devices/<id>",
  "deviceModel": "<devicemodels>",
  "name": "<name>",
  "note": "<note>"
  "serial": "<serial>"
}
```

* All fields are required.
* "id" must always start with "/devices/" and id can not be empty.
* in response, server returns:
    * the created device if nothing goes wrong.
    * error 400 if input is invalid
    * error 500 if server encounter internal error

```
HTTP GET
URL: https://<api-gateway-url>/api/devices/{id}
```

* id cannot be empty
* in response, server returns:
    * the device data if correspoinding to `/device/{id}`
    * error 404 if cannot find device data
    * error 429 if too many requests had been sent to the server
    * error 500 if server encounter internal error
    
    

## build

**Requirements** 

Go, dep (Go dependency management tool)

**build**: To build the project run following commands

```shell script
dep ensure
make 
```

**test**:To run the unit tests, run the following command

```shell script
make test
```


##deploy

**requirements**:

serverless framework, AWS Account

to deploy run the following command
```shell script
sls deploy
```
after deploy console will provide you the link to interacti with the server.