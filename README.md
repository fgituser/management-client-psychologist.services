# psychologist client services

Represents a mono repository with services:
* [Client](https://github.com/fgituser/management-client-psychologist.services/tree/master/client) - client.service
* [Psychologist](https://github.com/fgituser/management-client-psychologist.services/tree/master/client) - psychologist.service
* [Operator](https://github.com/fgituser/management-client-psychologist.services/tree/master/client) - operator.service


Service monitoring is carried out by prometheus
  
### Installation
#### Docker

Install docker in containers
  ```sh
  docker-compose up
  ```

#### Building for source

To build each service separately, you need to go into the directory with the necessary service and perform:
```sh
make build
```

### URL services (default)
After the deployment of services, they will be available at the following addresses:   
psychologist-service - http://localhost:9998   
client-service - http://localhost:9999   
operator-service - http://localhost:9997   

> services interact with each other on the internal network, you can change the ports for each of them, which will not affect the communication between services

### Database services (defaul)
By default, two databases postgresql are created for the service psychologist and client 

psycholog - postgres://127.0.0.1/psychologist?sslmode=disable&user=postgres&password=postgres   

clients - postgres://clients-db/clients?sslmode=disable&user=postgres&password=postgres

> services interact with each other on the internal network, you can change the ports for each of them, which will not affect the communication between services


### Monitoring services
Monitoring is carried out by means of prometheus which is available at:   
prometheus - http://localhost:9090
> Services are pre-registered in prometheus.
The configuration file is located in the ./configs/prometheus.yml directory

### Open API Specification

* [Client specification](https://github.com/fgituser/management-client-psychologist.services/blob/master/client/api/openapi-spec/swagger.yaml) 

* [Psychologist specification](https://github.com/fgituser/management-client-psychologist.services/blob/develop/psychologist/api/openapi-spec/swagger.yaml)
* [Operator specification](https://github.com/fgituser/management-client-psychologist.services/blob/develop/operator/api/openapi-spec/swagger.yaml)
