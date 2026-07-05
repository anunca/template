ARG COMPOSER_IMAGE=composer:2.9.7
ARG PHP_IMAGE=php:8.3.31-apache

FROM $COMPOSER_IMAGE AS composer

FROM $PHP_IMAGE AS base

#PHP extensions
RUN docker-php-ext-enable opcache

COPY etc/000-default.conf /etc/apache2/sites-available/000-default.conf

RUN a2enmod rewrite
 
WORKDIR /var/www/html

FROM base AS dev

#Composer dependencies
RUN apt update -qq\
  && apt install -y git unzip\
  && rm -rf /var/lib/apt/lists/*
 
COPY --from=composer /usr/bin/composer /usr/bin/composer
ENV COMPOSER_ALLOW_SUPERUSER=1

FROM dev AS build
 
COPY src .

RUN composer install --no-dev --optimize --classmap-authoritative

FROM base AS prod
 
COPY --from=build /var/www/html .

RUN chown -R www-data:www-data app/cache
