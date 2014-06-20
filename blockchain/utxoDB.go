package blockchain

import (
    "github.com/oxfeeefeee/kaiju/klib/kdb"
    "github.com/oxfeeefeee/kaiju/catma"
    "github.com/oxfeeefeee/kaiju/catma/cst"
)

// All unspent tx output stored in KDB
type utxoDB struct {
    db  *kdb.KDB
}

func newUpspentDB() *utxoDB {
    db, err := kdb.New(cst.KDBCapacity, fileKDB())
    if err != nil {
        panic("Failed to create utxoDB")
    }
    return &utxoDB{
        db,
    }
}

func (u *utxoDB) addTxs(txs []*catma.Tx) {

}

func (u *utxoDB) validateInput(txi *catma.TxIn) (int64, error) {
    return 0, nil
}