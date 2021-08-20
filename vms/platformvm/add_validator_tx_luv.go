package platformvm

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/hellobuild/Luv-Go/snow"
	"github.com/hellobuild/Luv-Go/utils/constants"
)

var errBlacklistedStaker = errors.New("this nodeID is not in the white list of validators")

//Check governance whitelist of nodeIDs to approve the new validator
//Skip governance if not setup on this node
func (tx *UnsignedAddValidatorTx) VerifyValidatorsWhitelist(ctx *snow.Context) error {
	GovernanceUsername := constants.GovernanceUsername

	GovernanceExists, err := ctx.Keystore.GovernanceExists(GovernanceUsername)
	if err != nil {
		return err
	}
	if !GovernanceExists {
		return nil
	}

	fmt.Println("whitelist verification")

	Password, err := ioutil.ReadFile("./build/db/governance_luv.txt")
	if err != nil {
		return err
	}
	db, err := ctx.Keystore.GetDatabase(GovernanceUsername, string(Password))
	if err != nil {
		return fmt.Errorf("problem retrieving user %s: %w", GovernanceUsername, err)
	}

	defer db.Close()

	user := user{db: db}
	txNodeID := tx.Validator.NodeID
	whitelistNodeID, err := user.controlsAddress(txNodeID)
	if err != nil {
		return err
	}
	if whitelistNodeID {
		return nil
	}

	return errBlacklistedStaker

}
