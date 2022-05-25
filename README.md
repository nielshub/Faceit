#FaceIt code challenge.

Done by Niels Sanchez van den Beuken.

##Instructions to start the application on localhost
*install docker and docker-compose
*create variables.env file inside env folder with following values:

```[env]
ENVIRONMENT="LOCAL"
VERSION = "1.0.0"
JSONSCHEMAPATH = "file:///usr/local/bin/usersSchema.json"
```

Right now there are no critical values but here we could save apikey and sensitive stuff needed for real environments

\*run following command in main folder of the repo: docker-compose up --build

##An explanation of the choices taken and assumptions made during development
A repository pattern an a hexagonal arquitecture has been applied. It is a very simple microservice but with this structure it should be easy to scalate with other services and handlers.
Moreover, with the docker image and docker compose is easy to implement in different environments using K8 and AWS stuff. Multistage has been used in docker to have a more lightweighted image, this has been done as a good practice.

##Possible extensions or improvements to the service (focusing on scalability and deployment to production)
*Configure env variables for the different environments
*Add proper authentication middlewares for cybersecurity
\*Add K8 configurations to be able to scalate easy the application if high traffic or requests
