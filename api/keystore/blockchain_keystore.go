// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package keystore

import (
	"github.com/hellobuild/Luv-Go/database"
	"github.com/hellobuild/Luv-Go/database/encdb"
	"github.com/hellobuild/Luv-Go/ids"
)

var _ BlockchainKeystore = &blockchainKeystore{}

type BlockchainKeystore interface {
	// Get a database that is able to read and write unencrypted values from the
	// underlying database.
	GetDatabase(username, password string) (*encdb.Database, error)

	// Get the underlying database that is able to read and write encrypted
	// values. This Database will not perform any encrypting or decrypting of
	// values and is not recommended to be used when implementing a VM.
	GetRawDatabase(username, password string) (database.Database, error)

	//blockchain_key_store_luv.go call to verify the excistence of governance user
	GovernanceExists(username string) (bool, error)
}

type blockchainKeystore struct {
	blockchainID ids.ID
	ks           *keystore
}

func (bks *blockchainKeystore) GetDatabase(username, password string) (*encdb.Database, error) {
	bks.ks.log.Info("Keystore: GetDatabase called with %s from %s", username, bks.blockchainID)
	return bks.ks.GetDatabase(bks.blockchainID, username, password)
}

func (bks *blockchainKeystore) GetRawDatabase(username, password string) (database.Database, error) {
	bks.ks.log.Info("Keystore: GetRawDatabase called with %s from %s", username, bks.blockchainID)
	return bks.ks.GetRawDatabase(bks.blockchainID, username, password)
}
