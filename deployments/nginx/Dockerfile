# 1 шаг - сборки
FROM node:18.14.2-alpine3.17 AS build_stage
RUN apk add git
WORKDIR /node/src/
RUN git clone https://github.com/frontend-park-mail-ru/2023-1_Yula
WORKDIR /node/src/2023-1_Yula
RUN git checkout dz7
RUN npm i
RUN npm run build-dev

# 2 шаг
FROM nginx AS run_stage
WORKDIR /front
COPY --from=build_stage /node/src/2023-1_Yula/dist /front
COPY --from=build_stage /node/src/2023-1_Yula/server/static /front/static
EXPOSE 80
ENTRYPOINT ["nginx", "-g", "daemon off;"]

