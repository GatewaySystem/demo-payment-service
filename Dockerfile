FROM public.ecr.aws/docker/library/golang:1.21-alpine AS build
WORKDIR /app
# go mod download fetches modules from VCS; alpine doesn't ship git.
RUN apk add --no-cache git
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o server .

FROM public.ecr.aws/docker/library/alpine:3.19
WORKDIR /app
COPY --from=build /app/server .
EXPOSE 8082
CMD ["./server"]
