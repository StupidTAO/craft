package types

import (
	"github.com/DSiSc/craft/rlp"
	"io"
	"math/big"
	"sync/atomic"
)

type Transaction struct {
	Data TxData
	Hash atomic.Value
	Size atomic.Value
	From atomic.Value
}

type TxData struct {
	AccountNonce uint64   `json:"nonce"    gencodec:"required"`
	Price        *big.Int `json:"gasPrice" gencodec:"required"`
	GasLimit     uint64   `json:"gas"      gencodec:"required"`
	Recipient    *Address `json:"to"       rlp:"nil"`
	From         *Address `json:"from"     rlp:"-"`
	Amount       *big.Int `json:"value"    gencodec:"required"`
	Payload      []byte   `json:"input"    gencodec:"required"`

	// Signature values
	V *big.Int `json:"v" gencodec:"required"`
	R *big.Int `json:"r" gencodec:"required"`
	S *big.Int `json:"s" gencodec:"required"`

	// This is only used when marshaling to JSON.
	Hash *Hash `json:"hash" rlp:"-"`
}

// EncodeRLP implements rlp.Encoder
func (tx *Transaction) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, &tx.Data)
}

// DecodeRLP implements rlp.Decoder
func (tx *Transaction) DecodeRLP(s *rlp.Stream) error {
	_, size, _ := s.Kind()
	err := s.Decode(&tx.Data)
	if err == nil {
		tx.Size.Store(StorageSize(rlp.ListSize(size)))
	}

	return err
}

type ETransaction struct {
	data txdata
	hash atomic.Value
	size atomic.Value
	from atomic.Value
}

type txdata struct {
	AccountNonce uint64   `json:"nonce"    gencodec:"required"`
	Price        *big.Int `json:"gasPrice" gencodec:"required"`
	GasLimit     uint64   `json:"gas"      gencodec:"required"`
	Recipient    *Address `json:"to"       rlp:"nil"` // nil means contract creation
	Amount       *big.Int `json:"value"    gencodec:"required"`
	Payload      []byte   `json:"input"    gencodec:"required"`

	// Signature values
	V *big.Int `json:"v" gencodec:"required"`
	R *big.Int `json:"r" gencodec:"required"`
	S *big.Int `json:"s" gencodec:"required"`

	// This is only used when marshaling to JSON.
	Hash *Hash `json:"hash" rlp:"-"`
}

func (tx *ETransaction) DecodeBytes(encodedTx []byte) error {
	return rlp.DecodeBytes(encodedTx, &tx.data)
}

func (tx *ETransaction) GetTxData() TxData {
	txData := new(TxData)
	txData.AccountNonce = tx.data.AccountNonce
	txData.Price = tx.data.Price
	txData.GasLimit = tx.data.GasLimit
	txData.Recipient = tx.data.Recipient
	txData.Amount = tx.data.Amount
	txData.Payload = tx.data.Payload

	txData.V = tx.data.V
	txData.R = tx.data.R
	txData.S = tx.data.S

	return *txData
}

func (tx *ETransaction) SetTxData(txData *TxData) error {

	//res, _ := json.Marshal(tx.data)
	//json.Unmarshal(res, txData)
	txData.AccountNonce = tx.data.AccountNonce
	txData.Price = tx.data.Price
	txData.GasLimit = tx.data.GasLimit
	txData.Recipient = tx.data.Recipient
	txData.Amount = tx.data.Amount
	txData.Payload = tx.data.Payload

	txData.V = tx.data.V
	txData.R = tx.data.R
	txData.S = tx.data.S

	return nil
}
