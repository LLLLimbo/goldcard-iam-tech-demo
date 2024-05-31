FROM node:14 as builder

WORKDIR /usr/src/app
COPY package*.json ./

RUN npm install --save jwt-decode oidc-client-ts react react-dom react-router-dom
COPY . .
CMD ["npm", "start"]
