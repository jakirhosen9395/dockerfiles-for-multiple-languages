# Build
FROM maven:3.9-eclipse-temurin-17 AS build
WORKDIR /src
COPY pom.xml .
RUN mvn -q -e -B -DskipTests dependency:go-offline
COPY . .
RUN mvn -q -e -B -DskipTests package

# Run
FROM eclipse-temurin:17-jre
WORKDIR /app
COPY --from=build /src/target/*.jar app.jar
EXPOSE 8080
CMD ["java","-jar","/app/app.jar"]
