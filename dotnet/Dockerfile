ARG DOTNET_SDK_IMAGE=mcr.microsoft.com/dotnet/sdk:6.0-alpine
ARG DOTNET_RUNTIME_IMAGE=mcr.microsoft.com/dotnet/aspnet:6.0-alpine

FROM $DOTNET_SDK_IMAGE AS base

WORKDIR /app

ENV ASPNETCORE_ENVIRONMENT='Development' DOTNET_ENVIRONMENT='Development'
EXPOSE 5000-5001

FROM base AS dev

# COPY src .

ENTRYPOINT ["dotnet"]
CMD [ "run", "--urls", "http://*:5000;https://*:5001" ]

FROM base AS prod-builder

COPY src .

RUN dotnet publish -c Release -o /app/publish

FROM $DOTNET_RUNTIME_IMAGE AS prod

WORKDIR /app

COPY --from=prod-builder /app/publish/ .

ENV ASPNETCORE_URLS=http://+:5000
EXPOSE 5000

ENTRYPOINT ["dotnet", "App.dll"]
