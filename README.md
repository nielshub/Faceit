# FaceIt code challenge.

Done by Niels Sanchez van den Beuken.

## Instructions to start the application on localhost

- install docker and docker-compose
- create variables.env file inside env folder (right now there are no critical values but here we could save apikey and sensitive stuff needed for real environments) with following values:

```[env]
ENVIRONMENT="LOCAL"
VERSION = "1.0.0"
JSONSCHEMAPATH = "file:///usr/local/bin/usersSchema.json"
DBURL="mongodb"
USERSCOLLECTIONNAME="users"
DATABASENAME="faceit"
```

- run following command in main folder of the repo: docker-compose up --build
- In order to check rabbitMQ connections: http://localhost:15672/ user = guest / password = guest
- In order to check mongoDB: http://localhost:8081/ no password / user is required for this project

## An explanation of the choices taken and assumptions made during development

A repository pattern an a hexagonal arquitecture has been applied. It is a very simple microservice but with this structure it should be easy to scalate with other services and handlers.
Moreover, with the docker image and docker compose is easy to implement in different environments using K8 and AWS stuff. Multistage has been used in docker to have a more lightweighted image, this has been done as a good practice.

For the pub/sub pattern a very simple approach has been taken because no requirements were defined that need to over engineer the solution. It has been done adding it to the handler once all the DB actions where correctly done and it is an ack event in the DB. The publication approach has been done with a fanout style, so the queue's name does not matter, and the queues are created with the different customers once they are created, just binding it to the proper exchange, in this case userEvents. No differenciation has been done between events (create/ update/ delete).

However, if it is required a more sensitive approach, outbox pattern can be taken, creating a outbox DB with all the events from the main database and adding a worker that reads from that database and send the info to rabbitMQ, with this approach we ensure we do not lose information if something crashes. Or a worker that reads from the change streams in mongo DB in order to send this information asynchronous to the rabbitMQ so we ensure the order of the events sent to rabbitMQ is exactly what happens in the main DB.

## Possible extensions or improvements to the service (focusing on scalability and deployment to production)

- Configure env variables for the different environments
- Add proper authentication middlewares for cybersecurity
- Add K8 configurations to be able to scalate easy the application if high traffic or requests
- More sensitive approach for pub/sub if requirements are needed
