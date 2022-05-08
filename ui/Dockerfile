FROM node:16.15-alpine

ADD  ./package.json /ui/package.json
WORKDIR /ui
RUN npm i --production
ADD ./src /ui
