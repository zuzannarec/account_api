# Account API
Author: Zuzanna Rec, zuzanna.rec@gmail.com

The functions responsible for handling Create, Fetch and Delete operations takes context.Context as their first parameter to allow setting timeouts and/or request cancellation.

# Run tests using docker-compose

The tests can be executed using `docker-compose.yml` which is an extended version of https://github.com/form3tech-oss/interview-accountapi/blob/master/docker-compose.yml

Use the following command to run tests:

`docker-compose up --build`

# TODO
- retries with exponential backoff
- add a circuit breaker
