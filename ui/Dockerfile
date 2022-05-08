FROM node:16.15-alpine

WORKDIR /ui

COPY  ./package.json ./
RUN npm i --production

COPY ./src ./

CMD [ "node", "./index.js" ]
