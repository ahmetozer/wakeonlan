FROM golang:1.17 AS build
WORKDIR /app/wakeonlan
RUN apt update && apt install -y libcap2-bin
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" 
RUN setcap CAP_NET_BIND_SERVICE=+eip wakeonlan

FROM scratch
COPY --from=build  /app/wakeonlan/wakeonlan /bin/wakeonlan
LABEL org.opencontainers.image.source="https://github.com/ahmetozer/wakeonlan"
ENTRYPOINT [ "/bin/wakeonlan" ]