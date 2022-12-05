FROM openjdk:17-jdk-alpine AS builder
ENV JAR_FILE=target/*.jar
WORKDIR /go/src/main
COPY . .
RUN wget https://mirrors.estointernet.in/apache/maven/maven-3/3.6.3/binaries/apache-maven-3.6.3-bin.tar.gz
RUN tar -xvf apache-maven-3.6.3-bin.tar.gz
RUN mv apache-maven-3.6.3 /opt/
ENV M2_HOME="/opt/apache-maven-3.6.3" 
ENV PATH="$M2_HOME/bin:$PATH"
RUN export PATH
RUN echo $PATH
RUN mvn --version
RUN mvn package
COPY ${JAR_FILE} app.jar

FROM openjdk:17-jdk-alpine
COPY --from=builder /go/src/main/app.jar /app.jar
ENTRYPOINT ["java","-jar","/app.jar"]