FROM node:14 AS build
WORKDIR /app

COPY client/package.json client/yarn.lock ./
RUN yarn install --frozen-lockfile

COPY client ./
RUN yarn build


FROM gcr.io/distroless/nodejs:14
WORKDIR /app

ENV NODE_ENV production

COPY --from=build /app/public ./public
COPY --from=build /app/.next ./.next
COPY --from=build /app/node_modules ./node_modules
COPY --from=build /app/package.json ./package.json

CMD [ "node_modules/.bin/next", "start" ]
