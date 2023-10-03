# EggRoll 🐣🛼

A high-level, opinionated, lambda-based framework for Cartesi Rollups in Go.

![build](https://github.com/gligneul/eggroll/actions/workflows/go.yml/badge.svg)

## Requirements

EggRoll is built on top of the [Cartesi Rollups](https://docs.cartesi.io/cartesi-rollups/) infrastructure version 1.0.
To use EggRoll, you also need [sunodo](https://github.com/sunodo/sunodo/) version 0.8.

## Quick Look

EggRoll divides the DApp into two parts: contract and client.
The contract runs in the blockchain inside a Cartesi VM and relies on the Cartesi Rollups API.
The client runs off-chain and communicates with the contract using the Cartesi Reader Node APIs and the Ethereum API.
EggRoll provides abstractions for both sides of the DApp.

Let's look at a simple example: a DApp that keeps a text box in the blockchain.
In EggRoll, you should first define the types shared between the DApp contract and the client.
The structs `InputAppend` and `InputClear` are types sent from the client to the contract.
The struct `State` is the contract's internal state that will also be available to the client for reading.

```go
type InputAppend struct {
	Value string
}

type InputClear struct {
}

type State struct {
	TextBox string
}
```

Then, you should use the `eggroll.Contract` struct to build the DApp contract.
`Register` registers a handler function for each input struct type.
Each handler receives the mutable state and the respective input struct.
After registering all functions, call `Roll` to run the contract.

```go
func clearHandler(env *eggroll.Env, state *State, _ *InputClear) error {
	env.Logln("received input clear")
	state.TextBox = ""
	return nil
}

func appendHandler(env *eggroll.Env, state *State, input *InputAppend) error {
	env.Logf("received input append with '%v'\n", input.Value)
	state.TextBox += input.Value
	return nil
}

func main() {
	contract := eggroll.NewContract[State]()
	eggroll.Register(contract, clearHandler)
	eggroll.Register(contract, appendHandler)
	contract.Roll()
}
```

Finally, you can interact with the DApp contract from the client using the `eggroll.Client` struct.
`Send` sends inputs to the contract through the Ethereum API.
`WaitFor` waits for the given number of inputs to be processed by the contract.
`State` reads the contract state from the Cartesi reader node.
`Logs` reads the logs generated by the contract with the `env.Log` functions.

```go
func main() {
	ctx := context.Background()
	client := eggroll.NewClient[State]()

	inputs := []any{
		&InputClear{},
		&InputAppend{Value: "egg"},
		&InputAppend{Value: "roll"},
	}
	for _, input := range inputs {
		log.Printf("Sending input %#v\n", input)
		Check(client.Send(ctx, input))
	}

	log.Println("Waiting for inputs to be processed")
	Check(client.WaitFor(ctx, 3))

	state := Must(client.State(ctx))
	log.Printf("Text box: '%v'\n", state.TextBox)

	logs := Must(client.Logs(ctx))
	log.Println("Logs:")
	for _, msg := range logs {
		log.Print(">", msg)
	}
}
```

To run this example, check the README in `./examples/textbox`.
