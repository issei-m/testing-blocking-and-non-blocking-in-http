FROM php:8.4-fpm
COPY --from=composer:latest /usr/bin/composer /usr/bin/composer
WORKDIR /opt/app
RUN <<EOF
set -e
apt-get update
apt-get install -y --no-install-recommends libzip-dev zip unzip procps
docker-php-ext-install zip
EOF
COPY composer.* .
RUN composer install --no-scripts --no-autoloader
COPY public public
RUN composer dump-autoload --optimize
RUN echo "pm = static" >> /usr/local/etc/php-fpm.d/zz-custom.conf
RUN echo "pm.max_children = 10" >> /usr/local/etc/php-fpm.d/zz-custom.conf
EXPOSE 9000
CMD ["php-fpm"]
