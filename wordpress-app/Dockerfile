FROM php:8.1-fpm-alpine

# Install necessary packages
RUN apk add --no-cache nginx


RUN apk add --no-cache vim openrc util-linux

RUN  openrc && touch /run/openrc/softlevel

RUN docker-php-ext-install mysqli

COPY . /var/www/wordpress/

# Download and extract WordPress
#RUN curl -o /tmp/wordpress.tar.gz https://wordpress.org/latest.tar.gz && \
#    tar xf /tmp/wordpress.tar.gz -C /var/www && \
#    rm /tmp/wordpress.tar.gz && \
#    chown -R nginx:nginx /var/www/wordpress

# Copy Nginx configuration files
COPY nginx.conf /etc/nginx/nginx.conf

# Expose port 80
EXPOSE 80

# Start Nginx and PHP-FPM
CMD ["sh", "-c", "nginx && php-fpm"]

