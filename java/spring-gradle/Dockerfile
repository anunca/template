ARG JAVA_VERSION
ARG TOMCAT_VERSION

FROM $JAVA_VERSION AS dev

WORKDIR /app

EXPOSE 8080

ENTRYPOINT ["./gradlew"]
CMD [ "bootRun" ]

FROM dev AS build

COPY src .

RUN ./gradlew clean build

FROM $TOMCAT_VERSION AS prod

# COPY --from=build "/app/build/libs/*.war" ./webapps.dist/
COPY --from=build "/app/build/libs/*.war" ./webapps/

# RUN mv /usr/local/tomcat/webapps.dist/* /usr/local/tomcat/webapps/

EXPOSE 8080

CMD ["catalina.sh", "run"]
