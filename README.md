**Run the application**

To run the code first change directory to pricing-engine-go

`cd pricing-engine-go`

Then use the command
    
`make run`

Additionally, you may want to run the steps independently:
1) To install and verify the dependencies, run `make install`
2) To lint, run `make lint`

**Run the tests**

To run the tests, please use the command 

`make test`

Please use [Postman](https://www.postman.com/downloads/) client to test the application. 

**Sample Request Body**

```json
{
"date_of_birth": "1988-01-01",
"insurance_group": 35,
"license_held_since": "2012-08-01"
}
```

**Sample Response**

```json
{
"is_approved": true,
"price_ranges": {
"10800": 10.62,
"172800": 34.57,
"1800": 2.9,
"21600": 13.21,
"259200": 47.02,
"345600": 55.37,
"3600": 5.25,
"43200": 21.63,
"7200": 8.03,
"86400": 23.53
},
"reason_for_denial": ""
}
```

`is_approved` -> If the user is eligible to get an insurance
`price_ranges` -> Map of timeunits in seconds and amount to pay
`reason_for_denial` -> Reason of denial in case the insurance request is not approved

**Improvement points**

1) I would have liked to implement a db for this project, Postgres seems a good fit. I already added a docker config.

2) An ORM tool like `gorm` or `jorm`

3) An OpenAPI tool to use design-first approach. However, this is a small project right now,
so code-first approach works here.
   
4) Authentication/authorization mechanism; if the client of this API is another application,
API Key would be sufficient. If it is a user, implementation of authentication token is needed.

5) Since speed was mentioned, I would have used some sort of cache, preferably Redis cache with db
queries to make the system even faster. However, without a db, I think the speed is optimal at this point.
