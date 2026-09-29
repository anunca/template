ARG COMPOSER_IMAGE=composer:2.10.2
ARG PHP_IMAGE=php:8.5.10-fpm

FROM $COMPOSER_IMAGE AS composer

FROM $PHP_IMAGE AS base

# Build PHP extension and remove build-only headers in the same layer
RUN apt-get update -qq\
  && apt-get install -y --no-install-recommends libzip-dev\
  && docker-php-ext-install -j"$(nproc)" zip\
  && apt-get purge -y libzip-dev\
  && rm -rf /var/lib/apt/lists/*

WORKDIR /var/www/html

ARG APP_ENV
ARG APP_DEBUG
# hadolint ignore=DL3064
ARG APP_SECRET
# hadolint ignore=DL3064
ARG DATABASE_URL

FROM base AS base-builder

#Composer dependencies
RUN apt-get update -qq\
  && apt-get install -y --no-install-recommends git unzip\
  && rm -rf /var/lib/apt/lists/*

COPY --from=composer /usr/bin/composer /usr/bin/composer
ENV COMPOSER_ALLOW_SUPERUSER=1

SHELL ["/bin/bash", "-o", "pipefail", "-c"]
#symfony
RUN  curl -fsSL https://get.symfony.com/cli/installer | bash\
  && mv /root/.symfony5/bin/symfony /usr/local/bin/symfony

FROM base-builder AS dev

RUN cp $PHP_INI_DIR/php.ini-development $PHP_INI_DIR/php.ini\
&& sed -i '\
s/;opcache.enable=.*/opcache.enable=1/;\
s/;opcache.enable_cli=.*/opcache.enable_cli=1/;\
s|;date.timezone =.*|date.timezone = Europe/Paris|;\
s/memory_limit =.*/memory_limit = 128M/;\
' $PHP_INI_DIR/php.ini

# COPY src/composer.json src/composer.lock ./
# RUN --mount=type=cache,target=/tmp/composer-cache \
#     COMPOSER_CACHE_DIR=/tmp/composer-cache composer install
# COPY src .

FROM base-builder AS prod-builder
 
COPY src/composer.json src/composer.lock ./
RUN --mount=type=cache,target=/tmp/composer-cache \
    COMPOSER_CACHE_DIR=/tmp/composer-cache composer install --no-dev --no-interaction --no-progress --prefer-dist --no-autoloader --no-scripts
COPY src .
# Generate the optimized autoloader after application classes are available.
RUN composer dump-autoload --no-dev --optimize --no-interaction

RUN bin/console cache:warmup

FROM base AS prod

RUN cp $PHP_INI_DIR/php.ini-production $PHP_INI_DIR/php.ini\
&& sed -i '\
s/;opcache.enable=.*/opcache.enable=1/;\
s/;opcache.enable_cli=.*/opcache.enable_cli=1/;\
s/;opcache.memory_consumption=.*/opcache.memory_consumption=256/;\
s/;opcache.max_accelerated_files=.*/opcache.max_accelerated_files=20000/;\
s/;opcache.validate_timestamps=.*/opcache.validate_timestamps=0/;\
s/;opcache.interned_strings_buffer=.*/opcache.interned_strings_buffer=16/;\
s|;opcache.preload=.*|opcache.preload=/var/www/html/vendor/preload.php|;\
s/;opcache.preload_user=.*/opcache.preload_user=nobody/;\
' $PHP_INI_DIR/php.ini\
&& echo 'opcache.jit_buffer_size=100M' >> $PHP_INI_DIR/php.ini\
&& sed -i '\
s/;realpath_cache_size =.*/realpath_cache_size = 4096k/;\
s/;realpath_cache_ttl =.*/realpath_cache_ttl = 600/;\
s|;date.timezone =.*|date.timezone = Europe/Paris|;\
s/memory_limit =.*/memory_limit = 128M/;\
s/display_errors =.*/display_errors = Off/;\
' $PHP_INI_DIR/php.ini

COPY --from=prod-builder /var/www/html .

RUN chown -R nobody:nogroup var
