FROM golang:1.24 AS backend
WORKDIR /app
COPY book-club-be/ .
RUN go mod tidy
RUN go build -o server ./cmd/book-club
EXPOSE 8080
CMD ["./server"]

FROM node:22 AS frontend
WORKDIR /app
COPY book-club-fe/package*.json ./
RUN npm install
COPY book-club-fe/ .
EXPOSE 3000
CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0", "--port", "3000"]

