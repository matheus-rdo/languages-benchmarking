## Build image
`docker build -t heroes-golang .`


## Start container
`docker run -d -p 8080:8080 --name heroes-golang --network heroes_network heroes-golang`