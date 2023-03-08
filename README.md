## geth-performance-test

```shell
git clone https://github.com/bradyjoestar/geth-perfomance-test.git
```

```shell
#initial accounts: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
#initial accounts key: 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
```

### start optimism
```shell
git clone https://github.com/bradyjoestar/optimism.git
cd optimism/ops
git checkout wb/performance-test
make build
make clean
docker-compose up -d
```

### start mantle
```shell
git clone https://github.com/mantlenetworkio/mantle.git
cd mantle/ops
git checkout wb/performance_test
make build
make clean
docker-compose up -d
```

all 10000 accounts are imported into genesis.json


#### test
```shell
go run bit_stress.go
```
