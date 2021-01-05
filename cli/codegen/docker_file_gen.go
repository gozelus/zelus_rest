package codegen

import (
	"html/template"
	"io"
)

var dockerFileTpl = `
FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/zero

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY ./etc /app/etc
RUN go build -ldflags="-s -w" -o /app/{{ .AppName }} ./cmd/{{ .AppName }}.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/{{ .AppName }} /app/{{ .AppName }}
COPY --from=builder /app/etc /app/etc

CMD ["./{{ .AppName }}"]
`

type DockerFileGenner struct {
	AppName string
}

func NewDockerFileGenner(appName string) *DockerFileGenner {
	return &DockerFileGenner{AppName: appName}
}
func (d *DockerFileGenner) GenCode(writer io.Writer) error {
	var t *template.Template
	var err error
	if t, err = template.New("docker file gen").Parse(dockerFileTpl); err != nil {
		return err
	}
	return t.Execute(writer, d)
}
