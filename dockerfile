FROM golang:1.19

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/

RUN mkdir /root/.ssh && ssh-keyscan github.com >> /root/.ssh/known_hosts

ENV GOPRIVATE github.com

RUN --mount=type=secret,id=sshKey,dst=/root/.ssh/id_ed25519 go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o sample_project presenter/api/main.go
RUN chmod +x sample_project

CMD ["/app/sample_project"]