FROM node:18

RUN curl -f https://get.pnpm.io/v6.16.js | node - add --global pnpm@7

COPY package.json pnpm-lock.yaml ./app/

RUN pnpm install --frozen-lockfile --prod

COPY . ./app/

# Build the application
RUN pnpm build

EXPOSE 3000

# Run the application
CMD [ "pnpm", "start" ]
