ARG RUBY_VERSION

FROM $RUBY_VERSION AS base

RUN curl -sL https://deb.nodesource.com/setup_15.x | bash -\
  && curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -\
  && echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list

RUN apt update -qq\
  && apt install -y build-essential shared-mime-info\
  nodejs yarn\
  sqlite3\
  mariadb-client default-libmysqlclient-dev\
  libpq-dev\
  && rm -rf /var/lib/apt/lists/*

ARG RAILS_VERSION

RUN gem install rails -v $RAILS_VERSION

WORKDIR /app

EXPOSE 3000

ENTRYPOINT [ "bundle" ]

FROM base AS dev

CMD [ "install", "npm i", "rails server -b 0.0.0.0" ]

FROM dev AS build

COPY ./src .

FROM base AS prod

COPY --from=build /app .

RUN bundle install\
  && rm config/credentials.yml.enc\
  && EDITOR="mate --wait" rails credentials:edit

ENV RAILS_ENV=production
RUN rake assets:precompile

ENV RAILS_SERVE_STATIC_FILES=true
CMD [ "exec puma -e production -b unix:///var/run/puma.sock" ]
