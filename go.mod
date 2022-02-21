module github.com/regen-network/keystone2

go 1.16

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4

replace github.com/regen-network/regen-ledger/orm => github.com/regen-network/regen-ledger/orm v0.0.0-20210804173213-3265a868bf83

replace github.com/regen-network/regen-ledger/types => github.com/regen-network/regen-ledger/types v0.0.0-20210804173213-3265a868bf83

replace github.com/regen-network/keystone2 => ./

require (
	github.com/frumioj/crypto11 v1.2.5-0.20210823151709-946ce662cc0e // indirect
	github.com/regen-network/keystone/keys v0.0.0-20220129212613-fb67e4f8db9f // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
)
