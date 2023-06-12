# lambda-person

POC Lambda for technical purposes

Lambda persist PERSON data inside DynamoDB and create a notification via event EventBridge

Diagram Flow

      APIGW ==> Lambda ==> DynamoDB (person_tenant)
                       ==> EventBridge (agregation-card-person {person_tenant})

## Compile lambda

      GOOD=linux GOARCH=amd64 go build -o ../build/main main.go

      zip -jrm ../build/main.zip ../build/main

## Endpoints

Get /health

Get /list

Get /person/{id}

Post /person

      {
         "id": "007",
         "name": "Mr Beam",
         "gender": "M"
      }


Post /personaddress

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

Get /personaddress/{004}



