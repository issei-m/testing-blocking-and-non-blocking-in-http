FROM ruby:3.3-slim
RUN <<EOF
set -eu
apt-get update
apt-get install -y \
  build-essential \
  libsqlite3-dev \
  curl
apt-get clean
rm -rf /var/lib/apt/lists/*
EOF
WORKDIR /opt/app
COPY Gemfile* .
RUN bundle install
COPY . .
EXPOSE 8000
CMD ["bundle", "exec", "puma", "-C", "puma.rb"]
