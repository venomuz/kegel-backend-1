FROM nginx:1.22.1-alpine

RUN mkdir -p /etc/www/front-site \
 && mkdir -p /etc/www/front-admin \
 && mkdir -p /etc/nginx/keys \
 && chown -R root:root /etc/nginx/keys \
 && chmod -R 600 /etc/nginx/keys 

COPY ./fullchain.pem /etc/nginx/keys/
COPY ./privkey.pem /etc/nginx/keys/
COPY ./chain.pem /etc/nginx/keys/
COPY ./cert.pem /etc/nginx/keys/

COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
EXPOSE 443

CMD ["nginx", "-g", "daemon off;"]
