module github.com/kogisin/arbitrage-bot

go 1.15

require (
	github.com/cosmos/cosmos-sdk v0.42.9
	github.com/cosmos/go-bip39 v1.0.0
	github.com/osmosis-labs/osmosis v1.0.3
	github.com/pelletier/go-toml v1.8.1
	github.com/rs/zerolog v1.21.0
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.11
	google.golang.org/grpc v1.37.0
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/osmosis-labs/cosmos-sdk v0.42.10-0.20210806040506-92afdc8963ca
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)
