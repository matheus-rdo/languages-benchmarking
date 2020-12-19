## Build image

`docker build -t heroes-golang .`

## Start container

100 MB / 0.5 CPU  
`docker run -i --rm -p 8080:8080 --memory=100MB --cpus=0.5 -e MYSQL_HOST=heroes-mysql --network heroes_network heroes-golang`  

200 MB / 1 CPU  
`docker run -i --rm -p 8080:8080 --memory=200MB --cpus=1 -e MYSQL_HOST=heroes-mysql --network heroes_network heroes-golang`

300 MB / 2 CPU  
`docker run -i --rm -p 8080:8080 --memory=300MB --cpus=2 -e MYSQL_HOST=heroes-mysql --network heroes_network heroes-golang`  