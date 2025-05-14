# Start from the official Go image
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod ./ 
COPY go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .

# Tailwind build
FROM node:18 AS tailwind

WORKDIR /root
COPY tailwind.config.js package.json package-lock.json ./
RUN npm install

COPY src/input.css ./src/input.css
COPY static ./static
RUN npx tailwindcss -i ./src/input.css -o ./static/styles/output.css --minify


# Final minimal image
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/static ./static
COPY --from=builder /app/data ./data
EXPOSE 8080
CMD ["./main"]
