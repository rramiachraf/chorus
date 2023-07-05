FROM alpine:latest

RUN apk add --update build-base go vips-dev git nodejs yarn xdg-user-dirs

RUN addgroup -S gg && adduser -S uu -G gg

USER uu

RUN xdg-user-dirs-update

RUN mkdir /home/uu/chorus
WORKDIR /home/uu/chorus

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY --chown=uu:gg . .

WORKDIR /home/uu/chorus/view
RUN yarn install

WORKDIR /home/uu/chorus
RUN make build-linux

VOLUME /music

EXPOSE 3000

CMD ["sh", "-c", "./dist/$(ls -t dist | head -n 1) /music"]
