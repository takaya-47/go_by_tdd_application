FROM golang:1.24

WORKDIR /app

# 必要な基本ツールのインストール
RUN apt-get update && apt-get install -y \
    git \
    curl

# golangci-lint のインストール (公式のバイナリインストールスクリプトを使用)
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin latest

CMD ["/bin/bash"]