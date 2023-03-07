## geth-performance-test

```shell
git clone https://github.com/bradyjoestar/geth-perfomance-test.git
```

```shell
#initial accounts: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
#initial accounts key: 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
```

#### build op
```shell
./0_buildop.sh
```

#### build mantle
```shell
./0_buildmantle.sh
```

#### start op
```shell
./0_startOP.sh
```

#### start mantle
```shell
./0_startMantle.sh
```

#### tips
mantle and op can't run concurrently.


#### init bit accounts
```shell
go test -v -run TestInitBitAccount chain_init_test.go
```

#### test
```shell
go run bit_stress.go
```