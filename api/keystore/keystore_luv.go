package keystore

import (
	"fmt"
	"io/ioutil"

	"github.com/hellobuild/Luv-Go/utils/password"
)

//Create the governance user and store the password for verification of the whitelist
func (ks *keystore) CreateGovernanceUser(username, pw string) error {
	if username == "" {
		return errEmptyUsername
	}
	if len(username) > maxUserLen {
		return errUserMaxLength
	}

	ks.lock.Lock()
	defer ks.lock.Unlock()

	passwordHash, err := ks.getPassword(username)
	if err != nil {
		return err
	}

	if passwordHash != nil {
		return fmt.Errorf("user already exists: %s", username)
	}

	if err := password.IsValid(pw, password.OK); err != nil {
		return err
	}

	passwordHash = &password.Hash{}
	if err := passwordHash.Set(pw); err != nil {
		return err
	}

	passwordBytes, err := c.Marshal(codecVersion, passwordHash)
	if err != nil {
		return err
	}

	if err := ks.userDB.Put([]byte(username), passwordBytes); err != nil {
		return err
	}
	ks.usernameToPassword[username] = passwordHash

	verificationPassword := []byte(pw)
	err = ioutil.WriteFile("./build/db/governance_luv.txt", verificationPassword, 0600)
	if err != nil {
		return err
	}

	return nil
}

//returns true if governance user exists in the database
func (ks *keystore) GovernanceExists(username string) (bool, error) {
	ks.lock.Lock()
	defer ks.lock.Unlock()

	passwordHash, err := ks.getPassword(username)
	if err != nil {
		return false, err
	}
	if passwordHash != nil {
		return true, nil
	}

	return false, nil
}
