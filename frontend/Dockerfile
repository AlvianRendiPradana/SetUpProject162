# Base build for production and staging
FROM node:20-alpine AS builder

WORKDIR /app

# Copy dependency info first
COPY package.json package-lock.json* ./
RUN npm install

# Copy source
COPY . .

# Build step
RUN npm run build

# Development stage
FROM node:20-alpine AS development

WORKDIR /app
# Intentionally we do not copy source code here because it will be mounted as a volume
# But we copy package.json if it exists to be able to install
COPY package.json package-lock.json* ./
RUN npm install
CMD ["npm", "run", "dev"]

# Production and Staging server stage
FROM nginx:alpine AS production

# Copy the built assets from the builder stage
COPY --from=builder /app/dist /usr/share/nginx/html

# Expose port (default 80)
EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
