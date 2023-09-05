package main

//go:generate bash -c "docker run --rm -v $(pwd)/ballot:/src -v $(pwd)/ballot:/dst ethereum/solc:0.5.12 -o /dst/ --optimize --optimize-runs=2000 --bin --abi --overwrite /src/Ballot.sol"
//go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --bin=./ballot/Ballot.bin --abi=./ballot/Ballot.abi --pkg=ballot --type=Contract --out=ballot/contract.go

import (
	"context"
	"github.com/artheranet/arthera-node/genesis/fake"
	"github.com/artheranet/lachesis/hash"
	"github.com/artheranet/lachesis/inter/idx"
	"math/big"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/artheranet/perftool/ballot"
)

var ballotOptions = [][32]byte{
	ballotProposal("option 1"),
	ballotProposal("option 2"),
	ballotProposal("option 3"),
	ballotProposal("option 4"),
	ballotProposal("option 5"),
	ballotProposal("option 6"),
	ballotProposal("option 7"),
	ballotProposal("option 8"),
	ballotProposal("option 9"),
}

func ballotRandChose() int64 {
	return rand.Int63n(int64(len(ballotOptions)))
}

func (g *CallsGenerator) ballotCreateContract(admin accounts.Account) TxMaker {
	payer := g.Payer(admin)
	return func(client *ethclient.Client) (*types.Transaction, error) {
		_, tx, _, err := ballot.DeployContract(payer, client, ballotOptions)
		if err != nil {
			panic(err)
		}

		return tx, err
	}
}

func (g *CallsGenerator) ballotCountOfVoites(voiter accounts.Account, contract common.Address) TxMaker {
	payer := g.Payer(voiter, big.NewInt(100))
	return func(client *ethclient.Client) (*types.Transaction, error) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		opts := &bind.FilterOpts{
			Context: ctx,
		}
		filterer, err := ballot.NewContractFilterer(contract, client)
		if err != nil {
			panic(err)
		}
		logs, err := filterer.FilterVoiting(opts, []common.Address{contract, payer.From}, nil, nil)
		if err != nil {
			g.Log.Error("filterer.FilterVoiting()", "err", err)
			return nil, nil
		}
		defer logs.Close()

		var count int
		for ; logs.Next(); count++ {
		}
		g.Log.Info("prev voites", "count", count)

		return nil, nil
	}
}

func (g *CallsGenerator) ballotVoite(voiter accounts.Account, contract common.Address, proposal int64) TxMaker {
	payer := g.Payer(voiter, big.NewInt(100))
	return func(client *ethclient.Client) (*types.Transaction, error) {
		transactor, err := ballot.NewContractTransactor(contract, client)
		if err != nil {
			panic(err)
		}

		return transactor.Vote(payer, big.NewInt(proposal))
	}
}

func (g *CallsGenerator) ballotWinner(contract common.Address) TxMaker {
	return func(client *ethclient.Client) (*types.Transaction, error) {
		caller, err := ballot.NewContractCaller(contract, client)
		if err != nil {
			panic(err)
		}

		winner, err := caller.WinnerName(g.ReadOnly())
		g.Log.Info("The winner", "hash", winner)

		return nil, err
	}
}

var (
	addresses []common.Address
	nums      []*big.Int
)

func (g *CallsGenerator) ballotLogs(contract common.Address) TxMaker {
	if addresses == nil {
		addresses = make([]common.Address, 1000)
		for i := range addresses {
			addresses[i] = crypto.PubkeyToAddress(fake.FakeKey(idx.ValidatorID(i)).PublicKey)
		}
	}
	if nums == nil {
		nums = make([]*big.Int, 100)
		for i := range nums {
			nums[i] = big.NewInt(int64(i))
		}
	}

	return func(client *ethclient.Client) (*types.Transaction, error) {
		filterer, err := ballot.NewContractFilterer(contract, client)
		if err != nil {
			panic(err)
		}

		it, err := filterer.FilterVoiting(
			&bind.FilterOpts{},
			addresses,
			nums,
			nil,
		)
		if err != nil {
			return nil, err
		}

		count := 0
		for it.Next() {
			count++
		}
		err = it.Close()

		g.Log.Debug("Logs massive reading", "got", count)

		return nil, err
	}
}

func ballotProposal(s string) [32]byte {
	return hash.Of([]byte(s))
}
