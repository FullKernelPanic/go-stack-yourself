FROM nginx:1.27.0-bookworm

COPY ./.docker/nginx/dev/default.conf /etc/nginx/conf.d/default.conf

CMD ["nginx", "-g", "daemon off;"]