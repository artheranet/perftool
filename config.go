package main

import (
	"errors"
	"github.com/artheranet/arthera-node/params"

	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/urfave/cli.v1"

	"github.com/artheranet/perftool/utils/toml"
)

var ConfigFileFlag = cli.StringFlag{
	Name:  "config",
	Usage: "TOML configuration file",
	Value: "txgen.toml",
}

type PerfConfig struct {
	ChainId int64 // chain id for sign transactions
	Payer   common.Address
	URLs    []string // WS nodes API URL
}

func DefaultConfig() *PerfConfig {
	return &PerfConfig{
		ChainId: int64(params.FakeNetworkID),
		URLs: []string{
			"ws://127.0.0.1:4500",
		},
	}
}

func OpenConfig(ctx *cli.Context) *PerfConfig {
	cfg := DefaultConfig()
	f := ctx.GlobalString(ConfigFileFlag.Name)
	err := cfg.Load(f)
	if err != nil {
		panic(err)
	}
	return cfg
}

func (cfg *PerfConfig) Load(file string) error {
	data, err := toml.ParseFile(file)
	if err != nil {
		return err
	}

	err = toml.Settings.UnmarshalTable(data, cfg)
	if err != nil {
		err = errors.New(file + ", " + err.Error())
		return err
	}

	return nil
}
