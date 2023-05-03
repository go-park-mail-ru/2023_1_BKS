FROM node:18.14.2-alpine3.17

RUN apk add git
WORKDIR /node/src/
RUN git clone https://github.com/frontend-park-mail-ru/2023-1_Yula
WORKDIR /node/src/2023-1_Yula

RUN npm install
# If you are building your code for production
# RUN npm ci --omit=dev

# Bundle app source
WORKDIR /node/src/2023-1_Yula/server 

EXPOSE 8080
CMD [ "node", "index.js" ]