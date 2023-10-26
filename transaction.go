package main

type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

type TXInput struct {
	Txid      []byte
	Vout      int
	ScriptSig string
}

type TXOutput struct {
	Value        int
	ScriptPubKey string
}

func NewCoinBaseTx(to, data string) *Transaction {
	txin := TXInput{nil, -1, data}
	txout := TXOutput{
		Value:        ,
		ScriptPubKey: "",
	}
}
