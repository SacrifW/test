package models

type (
	Obj map[string]interface{}

	Block struct {
		JSONRPC string `bson:"jsonrpc"`
		ID      string `bson:"id"`
		Result  Result `bson:"result"`
	}

	Result struct {
		Difficulty      string `bson:"difficulty"`
		ExtraData       string `bson:"extraData"`
		GasLimit        string `bson:"gasLimit"`
		GasUsed         string `bson:"gasUsed"`
		Hash            string `bson:"hash"`
		LogsBloom       string `bson:"logsBloom"`
		Miner           string `bson:"miner"`
		MixHash         string `bson:"mixHash"`
		Nonce           string `bson:"nonce"`
		Number          string `bson:"number"`
		ParentHash      string `bson:"parentHash"`
		ReceiptsRoot    string `bson:"receiptsRoot"`
		Sha3Uncles      string `bson:"sha3Uncles"`
		Size            string `bson:"size"`
		StateRoot       string `bson:"stateRoot"`
		Timestamp       string `bson:"timestamp"`
		TotalDifficulty string `bson:"totalDifficulty"`
		Transactions    []Transaction `bson:"transactions"`
	}
	Transaction struct {
		BlockHash        string `bson:"blockHash"`
		BlockNumber      string `bson:"blockNumber"`
		From             string `bson:"from"`
		Gas              string `bson:"gas"`
		GasPrice         string `bson:"gasPrice"`
		Hash             string `bson:"hash"`
		Input            string `bson:"input"`
		Nonce            string `bson:"nonce"`
		To               string `bson:"to"`
		TransactionIndex string `bson:"transactionIndex"`
		Value            string `bson:"value"`
		V                string `bson:"v"`
		R                string `bson:"r"`
		S                string `bson:"s"`
	}
)