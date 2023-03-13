FROM node:18

RUN curl -f https://get.pnpm.io/v6.16.js | node - add --global pnpm@7

WORKDIR /app

COPY package.json pnpm-lock.yaml ./

RUN pnpm install --frozen-lockfile --prod

COPY . ./

# Build the application
ENV NEXT_PUBLIC_API_URL=http://localhost:8080/

RUN pnpm build

EXPOSE 3000

# Run the application
CMD [ "pnpm", "start" ]
