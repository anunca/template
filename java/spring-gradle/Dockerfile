ARG JAVA_IMAGE=eclipse-temurin:11-jdk
ARG TOMCAT_IMAGE=tomcat:9.0-jdk11-temurin

FROM $JAVA_IMAGE AS base

WORKDIR /app

EXPOSE 8080

FROM base AS dev

COPY src .

ENTRYPOINT ["./gradlew"]
CMD [ "bootRun" ]

FROM base AS prod-builder

COPY src .

RUN ./gradlew clean build

FROM $TOMCAT_IMAGE AS prod

WORKDIR /usr/local/tomcat

# COPY --from=build "/app/build/libs/*.war" ./webapps.dist/
COPY --from=prod-builder "/app/build/libs/*.war" ./webapps/

# RUN mv /usr/local/tomcat/webapps.dist/* /usr/local/tomcat/webapps/

EXPOSE 8080

CMD ["catalina.sh", "run"]
