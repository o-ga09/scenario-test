# Goの公式イメージをベースにする
FROM golang:latest

# 作業ディレクトリを設定
WORKDIR /go/src/app

RUN go install github.com/zoncoen/scenarigo/cmd/scenarigo@v0.17.1

# ホストのファイルをコンテナにコピー
COPY . .

# ポートを公開
EXPOSE 8080