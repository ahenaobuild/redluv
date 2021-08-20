// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package keystore

//returns the keystore_luv.go function to verify the existence of luv governace addValidator
func (bks *blockchainKeystore) GovernanceExists(username string) (bool, error) {
	bks.ks.log.Info("Keystore: GovernanceExists called with %s", username)
	return bks.ks.GovernanceExists(username)
}
