## Build image
`docker build -t heroes-mysql .`

## Start container
`docker run -d -p 3306:3306 --name heroes-mysql -e MYSQL_ROOT_PASSWORD=superheroes heroes-mysql`