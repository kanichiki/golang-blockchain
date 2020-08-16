# ベースとなるDockerイメージ指定
FROM golang:1.15.0-alpine
# コンテナ内に作業ディレクトリを作成
RUN mkdir /go/src/goblockchain
# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/goblockchain

RUN apk update \
  && apk add --no-cache git \
  && go get -u github.com/btcsuite/btcutil/base58 \
  && go get -u golang.org/x/crypto/ripemd160
