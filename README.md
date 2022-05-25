#FaceIt code challenge.

Done by Niels Sanchez van den Beuken.

##Instructions to start the application on localhost
*install docker and docker-compose
*create variables.env file inside env folder with following values:

```[env]
ENVIRONMENT="LOCAL"
VERSION = "1.0.0"
JSONSCHEMAPATH = "file://config/usersSchema.json"
```

Right now there are no critical values but here we could save apikey and sensitive stuff needed for real environments

\*run following command in main folder of the repo: docker-compose up --build

##An explanation of the choices taken and assumptions made during development
A repository pattern an a hexagonal arquitecture has been applied. It is a very simple microservice but with this structure it should be easy to scalate and use in different environments. Moreover, with the docker image and coker compose is easy to imeplement in different environments using K8 and AWS stuff.

##Possible extensions or improvements to the service (focusing on scalability and deployment to production)
*Configure env variables for the different environments
*Add proper authentication middlewares for cybersecurity
\*Add K8 configurations to be able to scalate easy the application if high traffic or requests
