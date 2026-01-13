FROM node:alpine
WORKDIR /usr/src/app
COPY package*.json ./
RUN npm install
COPY . .
EXPOSE 3000 22
HEALTHCHECK CMD curl --fail http://localhost:3000 || exit 1
CMD "node","app.js"]
# ID-1768294475-f1ca24f5
