FROM openjdk:26-ea-25-jdk-oraclelinux8
 
WORKDIR /userservice
 
COPY .mvn/ .mvn
COPY mvnw pom.xml ./
RUN ./mvnw dependency:go-offline
 
COPY src ./src
 
CMD ["./mvnw", "spring-boot:run"]