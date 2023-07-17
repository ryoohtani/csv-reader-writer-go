FROM golang:latest

WORKDIR /app/go
# ローカルのカレントディレクトリにある全てのファイルを、コンテナの/app/goにコピーするため
COPY . .

# リナックスコマンド関連
RUN apt-get update && \
    apt-get install -y vim less sudo tree && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# 不要かもしれないがgodocを導入
RUN go install golang.org/x/tools/cmd/godoc@latest
#ここまでが最低限の内容
