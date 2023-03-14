FROM nginx:1.22.1-alpine

RUN mkdir -p /etc/www/front-site \
 && mkdir -p /etc/www/front-admin \
 && mkdir -p /etc/nginx/keys

WORKDIR /etc/nginx/keys

COPY . .

COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
EXPOSE 443

CMD ["nginx", "-g", "daemon off;"]

