cd ../
./mvnw package -DskipTests -Dquarkus.package.type=fast-jar
docker build -f src/main/docker/Dockerfile.fast-jar -t quarkus/benchmark-with-quarkus-fast-jar .