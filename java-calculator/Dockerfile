FROM openjdk:17-jdk-slim

WORKDIR /java-calculator

COPY . .

RUN apt-get update && \
    apt-get install -y maven && \
    mvn clean install

CMD ["mvn", "spring-boot:run"]