FROM golang:1.16 as build

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-extldflags "-static"' -o load-balancer ./cmd/load-balancer

FROM ubuntu:20.10 as lb-builder

WORKDIR /
RUN apt-get update -y --fix-missing
RUN apt-get install -y wget xz-utils
RUN wget https://github.com/Nordix/nfqueue-loadbalancer/releases/download/0.4.0/nfqlb-0.4.0.tar.xz
RUN tar -xvf /nfqlb-0.4.0.tar.xz
RUN mv nfqlb-0.4.0 nfqlb

FROM ubuntu:20.10 as grpc-health-probe

RUN apt-get update -y
RUN apt-get install -y wget
WORKDIR /
RUN wget https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.4.2/grpc_health_probe-linux-amd64
RUN mv grpc_health_probe-linux-amd64 /bin/grpc_health_probe
RUN chmod +x /bin/grpc_health_probe

FROM ubuntu:20.10

RUN apt-get update -y
run apt-get install -y libmnl-dev
run apt-get install -y libnetfilter-queue-dev
RUN apt-get install -y iproute2 tcpdump iptables net-tools iputils-ping ipvsadm netcat

WORKDIR /root/
COPY --from=build /app/load-balancer .
COPY --from=lb-builder /nfqlb /nfqlb
COPY --from=grpc-health-probe /bin/grpc_health_probe /bin/

ENV PATH="/nfqlb/bin:${PATH}"

CMD ["./load-balancer"]
