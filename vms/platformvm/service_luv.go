package platformvm

import (
	"fmt"
	"net/http"

	"github.com/hellobuild/Luv-Go/api"
	"github.com/hellobuild/Luv-Go/ids"
)

//defines the response after a blacklist of staker request
type AddValidatorWhitelistReply struct {
	NodeID string `json:"nodeId"`
}

//add a node to the local list of blacklisted stakers
func (service *Service) AddValidatorWhitelist(_ *http.Request, args *api.AddValidatorWhitelistArgs, reply *AddValidatorWhitelistReply) error {
	service.vm.ctx.Log.Info("Platform: AddValidatorWhitelist called")

	db, err := service.vm.ctx.Keystore.GetDatabase(args.Username, args.Password)
	if err != nil {
		return fmt.Errorf("problem retrieving user %q: %w", args.Username, err)
	}
	defer db.Close()

	user := user{db: db}
	if addrs, _ := user.getAddresses(); len(addrs) >= maxKeystoreAddresses {
		return fmt.Errorf("keystore user has reached its limit of %d addresses", maxKeystoreAddresses)
	}

	if err := user.putValidatorWhitelist(args.NodeID); err != nil {
		return fmt.Errorf("problem saving key %w", err)
	}

	reply.NodeID = "NodeID-" + args.NodeID

	return db.Close()
}

//list all nodeIDs in the whitelist of governance user
type GetValidatorWhitelistReply struct {
	Whitelist []string `json:"whitelist"`
}

//get the list of blacklisted nodes for validators
func (service *Service) GetValidatorWhitelist(_ *http.Request, args *api.GetValidatorWhitelistArgs, reply *GetValidatorWhitelistReply) error {
	service.vm.ctx.Log.Info("Platform: GetValidatorWhitelist called")

	db, err := service.vm.ctx.Keystore.GetDatabase(args.Username, args.Password)
	if err != nil {
		return fmt.Errorf("problem retrieving user %s: %w", args.Username, err)
	}

	defer db.Close()

	user := user{db: db}
	if args.NodeId == "" {
		whitelist, err := user.getAddresses()
		if err != nil {
			return fmt.Errorf("couldn't get addresses: %w", err)
		}
		reply.Whitelist = make([]string, len(whitelist))
		for i, nodeID := range whitelist {
			reply.Whitelist[i] = nodeID.String()
		}
		return db.Close()
	}

	var NodeID ids.ShortID
	NodeID, err = ids.ShortFromString(args.NodeId)
	if err != nil {
		return err
	}
	whitelist, err := user.controlsAddress(NodeID)
	if err != nil {
		return err
	}

	reply.Whitelist = make([]string, 1)
	if whitelist {
		reply.Whitelist[0] = args.NodeId
		return db.Close()
	}

	reply.Whitelist[0] = ""
	return db.Close()
}
