# Lambda-person

POC Lambda for technical purposes

Lambda persist PERSON data inside DynamoDB and create a notification via event EventBridge

Diagram Flow

      APIGW ==> Lambda ==> DynamoDB (person_tenant)
                       ==> EventBridge (event_type {person_tenant}) <== lambda-agregation-card-person-worker

## Compile lambda

   Manually compile the function

      GOOD=linux GOARCH=amd64 go build -o ../build/main main.go

      zip -jrm ../build/main.zip ../build/main

## Endpoints

+ Get /health

+ Get /list

+ Get /person/{id}

+ Post /person

      {
         "id": "007",
         "name": "Mr Beam",
         "gender": "M"
      }

+ Post /personaddress

      {
         "person":{
            "id":"007",
            "sk":"007"
         },
         "addresses":[
            {
               "id":"ADDRESS-401",
               "sk":"ADDRESS-401",
               "street":"St Quatre",
               "street_number":4,
               "zip_code":"zip-402"
            },      {
               "id":"ADDRESS-402",
               "sk":"ADDRESS-402",
               "street":"St Quatre Due",
               "street_number":42,
               "zip_code":"zip-422"
            }
         ]
      }

+ Get /personaddress/{004}

## Event

+ event_source: lambda.person

+ event_bus_name: event-bus-person

+ event type:

      eventTypeCreated =  "personCreated"
      eventTypeUpdated = 	"personUpdated"
      eventTypeDeleted = 	"personDeleted"

+ EventPayload

      {
         "id": "9",
         "name": "eliezer junior",
         "weight": 66,
         "gender": "M"
      }

## Pipeline

Prerequisite: 

Lambda function already created

+ buildspec.yml: build the main.go and move to S3
+ buildspec-test.yml: make a go test using services_test.go
+ buildspec-update.yml: update the lambda-function using S3 build

## DynamoDB

    PERSON-100 PERSON-100 M Eliezer Antunes

    PERSON-100 ADDRESS:ADDRESS-100-2  { "sk" : { "S" : "ADDRESS-100-2" }, "street_number" : { "N" : "101" }, "id" : { "S" : "ADDRESS-100-2" }, "street" : { "S" : "St Quatre Due" }, "zip_code" : { "S" : "zip-100-2" } }
