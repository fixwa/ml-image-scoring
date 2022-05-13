FROM node:18-alpine3.14
WORKDIR .
ENV PATH="./node_modules/.bin:$PATH"
COPY . .
EXPOSE 3000
RUN npm run build
CMD ["npm", "start"]