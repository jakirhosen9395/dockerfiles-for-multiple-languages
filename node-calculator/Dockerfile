# syntax=docker/dockerfile:1

# -------- Builder stage --------
FROM node:20-alpine AS build
WORKDIR /app
COPY package*.json ./
RUN npm ci --only=production
COPY . .
RUN npm prune --production

# -------- Runtime stage --------
FROM node:20-alpine
WORKDIR /app

ENV NODE_ENV=production
COPY --from=build /app /app
EXPOSE 8080
CMD ["node", "server.js"]
