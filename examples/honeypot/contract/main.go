package main

import (
	"fmt"

	"github.com/gligneul/eggroll"
	"github.com/gligneul/eggroll/wallets"

	"honeypot"
)

type HoneypotContract struct {
	eggroll.DefaultContract
}

func (c HoneypotContract) Codecs() []eggroll.Codec {
	return honeypot.Codecs()
}

func (c *HoneypotContract) Advance(env eggroll.Env) (any, error) {
	if deposit := env.Deposit(); deposit != nil {
		return c.handleDeposit(env, deposit)
	}

	return c.handleInput(env, env.DecodeInput())
}

func (c *HoneypotContract) handleDeposit(env eggroll.Env, deposit wallets.Deposit) (any, error) {
	switch deposit := env.Deposit().(type) {
	case *wallets.EtherDeposit:
		env.Logf("received deposit: %v\n", deposit)
		if env.Sender() != honeypot.Owner {
			// Transfer Ether deposits to honeypot.Owner
			env.EtherTransfer(env.Sender(), honeypot.Owner, &deposit.Value)
		}
		return c.getBalance(env), nil

	default:
		return nil, fmt.Errorf("unsupported deposit: %v", deposit)
	}
}

func (c *HoneypotContract) handleInput(env eggroll.Env, input any) (any, error) {
	if env.Sender() != honeypot.Owner {
		// Ignore inputs that are not from honeypot.Owner
		return nil, fmt.Errorf("ignoring input from %v", env.Sender())
	}

	switch input := input.(type) {
	case *honeypot.Withdraw:
		fmt.Printf(">> %#v\n", input)
		_, err := env.EtherWithdraw(honeypot.Owner, input.Value)
		if err != nil {
			return nil, err
		}
		env.Logf("withdraw %v\n", input.Value.ToBig().String())
		return c.getBalance(env), nil

	default:
		return nil, fmt.Errorf("invalid input: %v", input)
	}
}

func (c *HoneypotContract) getBalance(env eggroll.Env) *honeypot.Honeypot {
	ownerBalance := env.EtherBalanceOf(honeypot.Owner)
	return &honeypot.Honeypot{
		Balance: &ownerBalance,
	}
}

func main() {
	eggroll.Roll(&HoneypotContract{})
}
