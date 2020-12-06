docker network create --driver=bridge heroes_network
cd database
docker build -t heroes-mysql .
docker run -d -p 3306:3306 --name heroes-mysql --network heroes_network heroes-mysql