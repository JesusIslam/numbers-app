FROM node:10.16.3 AS builder

LABEL maintainer="faldy.findraddy@gmail.com"

ADD . /app
WORKDIR /app

#RUN npm ci --only=production
RUN npm ci
#CMD ["node", "node_modules/react-scripts/scripts/start.js"]
CMD bash -c "npm start"

# there was "xdg-open ENOENT" error
#FROM gcr.io/distroless/nodejs
#COPY --from=builder /app /app
#WORKDIR /app
#USER 1000
#CMD ["node_modules/react-scripts/scripts/start.js"]
