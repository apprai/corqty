FROM node:alpine as builder
WORKDIR /usr/src/app
COPY package*.json yarn.lock ./
RUN yarn install
COPY . .
RUN yarn run build

FROM nginx:alpine
COPY --from=builder /usr/src/app/dist/ /usr/share/nginx/html
