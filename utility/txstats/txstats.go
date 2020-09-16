package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	fmt.Println("Running...")
	bcInfoReq := []byte(`{
		"jsonrpc": "1.0",
		"id": 1,
		"method": "getblockchaininfo",
		"params": []
	 }`)

	body := sendRequest(bcInfoReq)
	bcInfo := BCInfoJsonResponse{}
	err := json.Unmarshal(body, &bcInfo)
	if err != nil {
		panic(err)
	}
	blockTime := uint64(40)
	totalPrvTx := 0
	totalNPrvTx := 0
	totalBlock := 0
	for i := 0; i < 8; i++ {
		id := byte(i)
		var wg sync.WaitGroup
		a := 0
		toHeight := bcInfo.Result.BestBlocks[i].Height
		fromHeight := toHeight - uint64((60*24*60*60)/blockTime)
		for h := fromHeight; h <= toHeight; h++ {
			if a >= 2000 {
				wg.Wait()
			}
			a++
			totalBlock++
			go func(shard byte, bh uint64) {
				wg.Add(1)
				body := sendRequest(createRequestBlockBody(shard, bh))
				block := BlockJsonResponse{}
				err := json.Unmarshal(body, &block)
				if err != nil {
					panic(err)
				}
				prv, nprv := getTxCountOfBlock(block.Result[0].Txs)
				totalPrvTx += prv
				totalNPrvTx += nprv
				wg.Done()
				a--
			}(id, h)
		}
		// fmt.Println(totalBlock, totalPrvTx, totalNPrvTx)
	}

	fmt.Printf("\nTotalBlock: %v \nTotal Privacy TXs: %v \nTotal NonPrivacy TXs: %v \nAverage pTXs/block: %v \nAverage nTXs/block: %v\n", totalBlock, totalPrvTx, totalNPrvTx, totalPrvTx/totalBlock, totalPrvTx/totalBlock)

}

func getTxCountOfBlock(txs []GetBlockTxResult) (prv int, nprv int) {
	prv = 0
	nprv = 0
	for _, tx := range txs {
		body := sendRequest(createRequestTx(tx.Hash))
		tx := TxJsonResponse{}
		err := json.Unmarshal(body, &tx)
		if err != nil {
			panic(err)
		}
		if tx.Result.IsPrivacy {
			prv++
		} else {
			nprv++
		}
	}
	return
}

func sendRequest(reqBody []byte) []byte {
	resp, err := http.Post("http://51.83.237.20:9338", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		// log.Println(err)
		// log.Println("retrying...")
		return sendRequest(reqBody)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// log.Println(err)
		return sendRequest(reqBody)
	}
	resp.Body.Close()
	return body
}
func createRequestBlockBody(shardID byte, height uint64) []byte {
	return []byte(fmt.Sprintf(`{
		"jsonrpc": "1.0",
		"id": 1,
		"method": "retrieveblockbyheight",
		"params": [%v,%v,"2"]
	 }`, height, shardID))
}

func createRequestTx(txHash string) []byte {
	return []byte(fmt.Sprintf(`{
		"jsonrpc": "1.0",
		"method": "gettransactionbyhash",
		"params": ["%v"],
		"id": 1
	}`, txHash))
}
