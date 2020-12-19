# benchmark-with-quarkus project

This project uses Quarkus, the Supersonic Subatomic Java Framework.

If you want to learn more about Quarkus, please visit its website: https://quarkus.io/ .

## Running the application in prod mode (for benchmarking)

### Steps to run on Native Image:
Package the application:
```shell script

./mvnw package -Pnative

```
Build the image using:
```shell script

docker build -f src/main/docker/Dockerfile.native -t quarkus/benchmark-with-quarkus .

```
**Run container**

100 MB / 0.5 CPU

```shell script

docker run -i --rm -p 8080:8080 --memory=100MB --cpus=0.5 --network heroes_network quarkus/benchmark-with-quarkus

```
200 MB / 1 CPU
```shell script

docker run -i --rm -p 8080:8080 --memory=200MB --cpus=1 --network heroes_network quarkus/benchmark-with-quarkus

```
300 MB / 2 CPU
```shell script

docker run -i --rm -p 8080:8080 --memory=300MB --cpus=2 --network heroes_network quarkus/benchmark-with-quarkus

```
---

### Steps to run on JVM:
Package the application:
```shell script

./mvnw package -DskipTests -Dquarkus.package.type=fast-jar

```
Build the image using:
```shell script

docker build -f src/main/docker/Dockerfile.fast-jar -t quarkus/benchmark-with-quarkus-fast-jar .

```
**Run container**

100 MB / 0.5 CPU

```shell script

docker run -i --rm -p 8080:8080 --memory=100MB --cpus=0.5 --network heroes_network quarkus/benchmark-with-quarkus-fast-jar

```
200 MB / 1 CPU
```shell script

docker run -i --rm -p 8080:8080 --memory=200MB --cpus=1 --network heroes_network quarkus/benchmark-with-quarkus-fast-jar

```
300 MB / 2 CPU
```shell script

docker run -i --rm -p 8080:8080 --memory=300MB --cpus=2 --network heroes_network quarkus/benchmark-with-quarkus-fast-jar

```

## Running the application in dev mode

You can run your application in dev mode that enables live coding using:
```shell script
./mvnw compile quarkus:dev
```

## Packaging and running the application

The application can be packaged using:
```shell script
./mvnw package
```
It produces the `benchmark-with-quarkus-1.0.0-SNAPSHOT-runner.jar` file in the `/target` directory.
Be aware that it’s not an _über-jar_ as the dependencies are copied into the `target/lib` directory.

If you want to build an _über-jar_, execute the following command:
```shell script
./mvnw package -Dquarkus.package.type=uber-jar
```

The application is now runnable using `java -jar target/benchmark-with-quarkus-1.0.0-SNAPSHOT-runner.jar`.

## Creating a native executable

You can create a native executable using: 
```shell script
./mvnw package -Pnative
```

Or, if you don't have GraalVM installed, you can run the native executable build in a container using: 
```shell script
./mvnw package -Pnative -Dquarkus.native.container-build=true
```

You can then execute your native executable with: `./target/benchmark-with-quarkus-1.0.0-SNAPSHOT-runner`

If you want to learn more about building native executables, please consult https://quarkus.io/guides/maven-tooling.html.

# RESTEasy JAX-RS

<p>A Hello World RESTEasy resource</p>

Guide: https://quarkus.io/guides/rest-json
