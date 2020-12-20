# Programming languages benchmarks

## Set up infrastructure
To set up database, cAdvisor and docker network, execute the following command:  
```shell script

docker-compose up -d

```  

## Tools used

- Apache benchmark
- cAdvisor (Container Advisor) 

## Run stress test

`ab -c 100 -n 10000 http://localhost:8080/heroes`