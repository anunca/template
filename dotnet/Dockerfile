ARG DOTNET_SDK_IMAGE=mcr.microsoft.com/dotnet/sdk:6.0-alpine
ARG DOTNET_RUNTIME_IMAGE=mcr.microsoft.com/dotnet/runtime:6.0-alpine

FROM $DOTNET_SDK_IMAGE AS base

WORKDIR /app

ENV ASPNETCORE_ENVIRONMENT='Development' DOTNET_ENVIRONMENT='Development'
EXPOSE 5000-5001

FROM base AS dev

ENTRYPOINT ["dotnet"]
CMD [ "run", "--urls", "http://*:5000;https://*:5001" ]

FROM base AS prod-builder

COPY src .

RUN dotnet publish -c Release

FROM $DOTNET_RUNTIME_IMAGE AS prod

WORKDIR /app

COPY --from=prod-builder /app/bin/Release/net5.0/publish/ .

ENV ASPNETCORE_URLS='http://*:5000;https://*:5001' DOTNET_URLS=='http://*:5000;https://*:5001'
EXPOSE 5000-5001

ENTRYPOINT ["dotnet", "App.dll", "--urls", "http://*:5000;https://*:5001"]
