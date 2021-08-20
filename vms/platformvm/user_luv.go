package platformvm

import (
	"errors"

	"github.com/hellobuild/Luv-Go/ids"
)

var (
	errNodeIDnil = errors.New("nodeID field is empty")
)

//store the nodeID in the local whitelist to vote in favor of a new validator
func (u *user) putValidatorWhitelist(nodeID string) error {
	if nodeID == "" {
		return errNodeIDnil
	}
	var NodeID ids.ShortID
	NodeID, err := ids.ShortFromString(nodeID)
	if err != nil {
		return err
	}

	controlsNodeID, err := u.controlsAddress(NodeID)
	if err != nil {
		return err
	}
	if controlsNodeID { // user already whitelist this nodeID. Do nothing.
		return nil
	}

	if err := u.db.Put(NodeID.Bytes(), NodeID.Bytes()); err != nil { //Here there is the posibility to add aditional information or state
		return err
	}

	whitelist := make([]ids.ShortID, 0)
	userHasNodeID, err := u.db.Has(addressesKey)
	if err != nil {
		return err
	}
	if userHasNodeID { // Get whitelist this user already controls, if they exist
		if whitelist, err = u.getAddresses(); err != nil {
			return err
		}
	}

	whitelist = append(whitelist, NodeID)
	Whitelist, err := Codec.Marshal(codecVersion, whitelist)
	if err != nil {
		return err
	}
	if err = u.db.Put(addressesKey, Whitelist); err != nil {
		return err
	}
	return nil
}
