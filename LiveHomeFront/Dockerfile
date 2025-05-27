# 构建阶段
FROM node:16 as build-stage

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

# 关键：赋予vite可执行权限
RUN chmod +x node_modules/.bin/vite

RUN npx vite build

# 生产阶段
FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]