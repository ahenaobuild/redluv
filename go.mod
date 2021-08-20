module github.com/hellobuild/Luv-Go

// Changes to the minimum golang version must also be replicated in
// scripts/ansible/roles/golang_based/defaults/main.yml
// scripts/build_avalanche.sh
// scripts/local.Dockerfile
// Dockerfile
// README.md
// go.mod (here, only major.minor can be specified)
go 1.15

require (
	github.com/AppsFlyer/go-sundheit v0.2.0
	github.com/Microsoft/go-winio v0.4.14
	github.com/NYTimes/gziphandler v1.1.1
	github.com/VictoriaMetrics/fastcache v1.6.0
	github.com/btcsuite/btcutil v1.0.2
	github.com/cespare/cp v1.1.1
	github.com/davecgh/go-spew v1.1.1
	github.com/deckarep/golang-set v1.7.1
	github.com/decred/dcrd/dcrec/secp256k1/v3 v3.0.0-20200627015759-01fd2de07837
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/ethereum/go-ethereum v1.10.4
	github.com/fjl/memsize v0.0.0-20190710130421-bcb5799ab5e5
	github.com/gballet/go-libpcsclite v0.0.0-20191108122812-4678299bea08
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.1.5
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.4
	github.com/gorilla/rpc v1.2.0
	github.com/gorilla/websocket v1.4.2
	github.com/hashicorp/go-hclog v0.14.1
	github.com/hashicorp/go-plugin v1.3.0
	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d
	github.com/holiman/bloomfilter/v2 v2.0.3
	github.com/holiman/uint256 v1.2.0
	github.com/huin/goupnp v1.0.1-0.20210310174557-0ca763054c88
	github.com/jackpal/gateway v1.0.6
	github.com/jackpal/go-nat-pmp v1.0.2
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0
	github.com/kr/pretty v0.2.0 // indirect
	github.com/mattn/go-colorable v0.1.7
	github.com/mattn/go-isatty v0.0.12
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mr-tron/base58 v1.2.0
	github.com/nbutton23/zxcvbn-go v0.0.0-20180912185939-ae427f1e4c1d
	github.com/olekukonko/tablewriter v0.0.5
	github.com/prometheus/client_golang v1.7.1
	github.com/rjeczalik/notify v0.9.2
	github.com/rs/cors v1.7.0
	github.com/spaolacci/murmur3 v1.1.0
	github.com/spf13/cast v1.3.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/status-im/keycard-go v0.0.0-20200402102358-957c09536969
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/syndtr/goleveldb v1.0.1-0.20210305035536-64b5b1c73954
	github.com/tyler-smith/go-bip39 v1.1.0
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/text v0.3.5
	golang.org/x/time v0.0.0-20201208040808-7e3f01d25324
	gonum.org/v1/gonum v0.9.1
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/olebedev/go-duktape.v3 v3.0.0-20210326210528-650f7c854440
	gopkg.in/urfave/cli.v1 v1.20.0
	gotest.tools v2.2.0+incompatible
)
