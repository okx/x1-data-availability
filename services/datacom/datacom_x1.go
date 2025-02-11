package datacom

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common"
)

func (d *Endpoints) isEmptyAddress(a common.Address) bool {
	emptyAddress := common.Address{}
	return bytes.Equal(a[:], emptyAddress[:])
}
