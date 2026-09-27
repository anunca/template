ARG MAVEN_IMAGE
ARG TOMCAT_IMAGE

FROM $MAVEN_IMAGE AS dev

WORKDIR /app

COPY src .

RUN mvn package

ENV PORT 5000
EXPOSE $PORT

CMD [ "sh", "-c", "mvn -Dserver.port=${PORT} spring-boot:run" ]

FROM dev AS build

COPY src .

RUN mvn package

FROM $TOMCAT_IMAGE AS prod

WORKDIR /usr/local/tomcat

# COPY --from=build "/app/target/*.war" ./webapps.dist/
COPY --from=build "/app/target/demo-0.0.1-SNAPSHOT.war" ./webapps/

# RUN mv /usr/local/tomcat/webapps.dist/* /usr/local/tomcat/webapps/

EXPOSE 8080

CMD ["catalina.sh", "run"]
