FROM node:lts-alpine as build-stage

ARG BASE_CONTENT_URL
ARG BASE_AUTH_URL
ENV BASE_CONTENT_URL "$BASE_CONTENT_URL"
ENV BASE_AUTH_URL "$BASE_AUTH_URL"

WORKDIR /app

COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build


FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
COPY --from=build-stage /app/substitute_environment_variables.sh /
RUN chmod +x /substitute_environment_variables.sh
EXPOSE 80
CMD [ "sh", "substitute_environment_variables.sh" ]
