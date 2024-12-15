package blockchain

import (
	"sync"

	"github.com/IMHYEWON/hyewoncoin/12.p2p/db"
	"github.com/IMHYEWON/hyewoncoin/12.p2p/utils"
)

const defaultDifficulty int = 2  // default difficulty
const difficultyInterval int = 5 // 5개의 블록마다 difficulty 재조정
const blockInterval int = 2      // 2분마다 블록이 생성되도록 설정
const allowedRange int = 2       // 2분의 여유시간
type blockchain struct {
	NewestHash        string `json:"newestHash"`
	Height            int    `json:"height"`            // 블록의 번호
	CurrentDifficulty int    `json:"currentDifficulty"` // 현재 difficulty
}

var bc *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func persist(b *blockchain) {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock() {
	block := createBlock(b.NewestHash, b.Height+1, getDifficulty(b))
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	persist(b)
}

func Blocks(b *blockchain) []*Block {
	var blocks []*Block
	hashCursor := b.NewestHash

	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)

		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}

func recalculateDifficulty(b *blockchain) int {
	// 5개의 블록마다 시간을 계산 = 5 * 2분 = 10분이 걸렸는지 확인
	allBlocks := Blocks(b)

	// 가장 최근 블록 (allBlocks 함수에서는 가장 최근 hash부터 이전 블록을 차례로 찾아서 추가하기때문에 0번째 인덱스가 가장 최근 블록)
	newestBlock := allBlocks[0]
	lastRecalculatedBlock := allBlocks[difficultyInterval-1]

	actualTime := (newestBlock.Timestamp / 60) - (lastRecalculatedBlock.Timestamp / 60)
	expectedTime := difficultyInterval * blockInterval // 5 * 2분 = 10분

	if actualTime < (expectedTime - allowedRange) {
		// 블록이 너무 빠르게 생성되고 있음 (쉬움)> difficulty 높임
		return b.CurrentDifficulty + 1
	} else if actualTime > (expectedTime + allowedRange) {
		// 블록이 너무 느리게 생성되고 있음 (어려움)> difficulty 낮춤
		return b.CurrentDifficulty - 1
	}
	return b.CurrentDifficulty

}

func getDifficulty(b *blockchain) int {
	// default difficulty (블로체인이 비어있을 때)
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height%difficultyInterval == 0 {
		// 5개의 블록마다 difficulty 재조정
		// recalculate difficulty
		return recalculateDifficulty(b)
	} else {
		return b.CurrentDifficulty
	}
}

// 모든 (블록체인에 있는) Transaction을 가져오는 함수
func Txs(b *blockchain) []*Tx {
	var txs []*Tx
	for _, block := range Blocks(b) {
		txs = append(txs, block.Transactions...)
	}
	return txs
}

// transaction ID로 Transaction을 찾는 함수
func FindTx(b *blockchain, targetTxId string) *Tx {
	for _, tx := range Txs(b) {
		if tx.Id == targetTxId {
			return tx
		}
	}
	return nil
}

// 어떤 Transaction Output이 Input으로 사용되었는지 확인
func UnspentTxOutsByAddress(address string, b *blockchain) []*UTxOut {
	// input을 생성하면 각각의 input은 항상 유니크한 transaction에서 output을 가지고 옴
	var uTxOuts []*UTxOut
	creatorTxs := make(map[string]bool) // key: TxId(string), value: TxOut의 index(bool로 표시)

	// 모든 블록을 가져옴
	for _, block := range Blocks(b) {
		// 모든 트랜잭션을 가져옴
		for _, tx := range block.Transactions {
			// 1. 해당 주소에 해당하는 Transaction의 목록에서 사용된 Output을 찾기위해 Map에 마킹해둠
			// (어떻게 ? Transaction의 Input에 포함되어있다는 것은 이미 사용된 Output이라는 뜻이기에, 이 Input의 부모 Transaction-Output을 확인해서)
			// 모든 TxIn을 가져옴

			// coinbase transaction은 제외
			for _, input := range tx.TxIns {
				if input.Signature == "COINBASE" {
					break
				}

				// 이 주소에 속한 input을 찾음 -> 각 input은 부모 transaction의 output을 가지고 있음
				if FindTx(b, input.TxId).TxOuts[input.Index].Address == address {
					creatorTxs[input.TxId] = true
				}
			}

			// 2. UTXO 생성
			// 주소에 해당하는 모든 TxOut을 가져와서, 마킹된 Transaction을 제외한 나머지를 모두 추가
			for index, output := range tx.TxOuts {
				if output.Address == address {
					_, ok := creatorTxs[tx.Id]
					// input으로 사용된 output이 아닌 output이라면 아직 사용되지 않은 output
					// --> UTXO(Unspent Transaction Output)
					// txOut중, input으로 참조되지 않은(=아직 사용되지 않은) output만 가져옴
					if !ok {
						// UTXO객체를 생성해서 uTxOuts 슬라이스에 append
						uTxOut := &UTxOut{tx.Id, index, output.Amount}

						// 이 UTXO가 mempool에 있는지 확인
						if !isOnMempool(uTxOut) {
							uTxOuts = append(uTxOuts, uTxOut)
						}
					}
				}
			}
		}
	}

	return uTxOuts
}

// 주소에 해당하는 총 거래량을 계산
func BalanceByAddress(address string, b *blockchain) int {
	var amount int
	var txOuts = UnspentTxOutsByAddress(address, b)
	for _, txOut := range txOuts {
		amount += txOut.Amount
	}
	return amount
}

func BlockChain() *blockchain {
	once.Do(func() {
		bc = &blockchain{
			Height: 0,
		}
		checkpoint := db.Checkpoint()

		if checkpoint == nil {
			bc.AddBlock()
		} else {
			bc.restore(checkpoint)
		}
	})

	return bc
}
