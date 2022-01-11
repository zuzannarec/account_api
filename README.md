# Account API
Author: Zuzanna Rec, zuzanna.rec@gmail.com

This is my first project in Golang. Apart from the language documentation I used the following sources:

https://gogoapps.io/blog/passing-loggers-in-go-golang-logging-best-practices/

https://www.youtube.com/watch?v=fe8vJSIzWss

The functions responsible for handling Create, Fetch and Delete operations takes context.Context as their first parameter to allow setting timeouts and/or request cancellation.

# Run tests using docker-compose

The tests can be executed using `docker-compose.yml` which is an extended version of https://github.com/form3tech-oss/interview-accountapi/blob/master/docker-compose.yml

There's a timing issue in docker-compose - sometimes the container with tests starts before the fake API service is ready and the tests fail. I've tried to control the startup order of these services but finally I decided to simply restart tests service on failure.

Use the following command to run tests:

`docker-compose up`

# TODO

- retries with exponential backoff
- add a circuit breaker

# Notes
I noticed that the first example payload in Accounts Create section in the documentation (https://api-docs.form3.tech/api.html#organisation-accounts-create) does not contain `name` field which is mandatory. It seems that it needs to be updated as the example request fails with error: {"error_message":"validation failure list: validation failure list: validation failure list: name in body is required"}
