module github.com/regen-network/keystone2

go 1.16

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/99designs/keyring => github.com/cosmos/keyring v1.1.7-0.20210622111912-ef00f8ac3d76

replace github.com/regen-network/regen-ledger/orm => github.com/regen-network/regen-ledger/orm v0.0.0-20210804173213-3265a868bf83

replace github.com/regen-network/regen-ledger/types => github.com/regen-network/regen-ledger/types v0.0.0-20210804173213-3265a868bf83

replace github.com/regen-network/keystone2 => ./

replace github.com/cosmos/cosmos-sdk => ../cosmos-sdk

//replace github.com/cespare/xxhash/v2 => ../xxhash

require (
	github.com/cosmos/cosmos-sdk v0.46.0-alpha3
	github.com/frumioj/crypto11 v1.2.5-0.20210823151709-946ce662cc0e
	github.com/gin-gonic/gin v1.7.0 // indirect
	github.com/regen-network/keystone/keys v0.0.0-20220129212613-fb67e4f8db9f
	github.com/stretchr/testify v1.7.1
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.28.0
)
