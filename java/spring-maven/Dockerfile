ARG MAVEN_IMAGE
ARG TOMCAT_IMAGE

FROM $MAVEN_IMAGE AS base

WORKDIR /app

ENV PORT 5000
EXPOSE $PORT

FROM base AS dev

RUN mvn package

COPY src .

CMD [ "sh", "-c", "mvn -Dserver.port=${PORT} spring-boot:run" ]

FROM base AS prod-builder

RUN mvn package

COPY src .

FROM $TOMCAT_IMAGE AS prod

WORKDIR /usr/local/tomcat

# COPY --from=build "/app/target/*.war" ./webapps.dist/
COPY --from=prod-builder "/app/target/demo-0.0.1-SNAPSHOT.war" ./webapps/

# RUN mv /usr/local/tomcat/webapps.dist/* /usr/local/tomcat/webapps/

EXPOSE 8080

CMD ["catalina.sh", "run"]
