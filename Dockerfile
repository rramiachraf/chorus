FROM alpine:latest AS builder

RUN apk add --update build-base go vips-dev git nodejs yarn

WORKDIR /chorus

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

WORKDIR /chorus/view
RUN yarn install

WORKDIR /chorus
RUN rm ./dist/*
RUN make build-linux
RUN mv ./dist/$(ls dist) ./dist/chorus 

########################
########################

FROM alpine:latest

RUN apk add --update xdg-user-dirs vips
RUN addgroup -S gg && adduser -S uu -G gg

USER uu
RUN xdg-user-dirs-update

WORKDIR /home/uu/chorus
COPY --from=builder /chorus/dist/chorus .

VOLUME /music

EXPOSE 3000

CMD ["./chorus", "/music"]
