FROM node:14-alpine3.13 AS build

WORKDIR /app

COPY client/package.json client/yarn.lock ./
RUN yarn install --frozen-lockfile

COPY client ./
RUN yarn build

CMD [ "yarn", "start" ]


FROM node:14-alpine3.13

WORKDIR /app

ENV NODE_ENV production

COPY --from=build /app/public ./public
COPY --from=build /app/.next ./.next
COPY --from=build /app/node_modules ./node_modules
COPY --from=build /app/package.json ./package.json

CMD [ "yarn", "start" ]
