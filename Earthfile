VERSION 0.8
FROM tinygo/tinygo:0.37.0
WORKDIR /go-workdir

signer:
    COPY . .
    RUN cd signer && tinygo build -o plugin.wasm -target wasip1 -buildmode=c-shared .
    SAVE ARTIFACT signer/plugin.wasm AS LOCAL ./signer.wasm
