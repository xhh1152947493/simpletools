FROM node:20 AS builder
WORKDIR /app
COPY package*.json ./

RUN git config --global url."https://".insteadOf git://
RUN npm install --registry=https://registry.npmmirror.com
COPY . .
RUN npm run build

FROM nginx:stable-alpine AS runner
# 目录加simpletools是与vite.config.js里面的base为/simpletools相匹配的
COPY --from=builder /app/dist /usr/share/nginx/html/simpletools
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 1234
CMD ["nginx", "-g", "daemon off;"]
