module github.com/artheranet/perftool

go 1.15

require (
	github.com/artheranet/arthera-node v0.0.0
	github.com/artheranet/lachesis v0.0.0
	github.com/ethereum/go-ethereum v1.10.8
	github.com/naoina/toml v0.1.2-0.20170918210437-9fafd6967416
	gopkg.in/urfave/cli.v1 v1.22.1
)

replace gopkg.in/urfave/cli.v1 => github.com/urfave/cli v1.20.0

replace github.com/ethereum/go-ethereum => ../arthera-go-ethereum

replace github.com/artheranet/lachesis => ../lachesis

replace github.com/artheranet/arthera-node => ../arthera-node
