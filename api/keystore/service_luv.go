package keystore

import (
	"net/http"

	"github.com/hellobuild/Luv-Go/api"
	"github.com/hellobuild/Luv-Go/utils/constants"
)

//Create a governance user to store the local whitelist of validators
func (s *service) CreateGovernanceUser(_ *http.Request, args *api.GovernancePassword, reply *api.SuccessResponse) error {
	GovernanceUsername := constants.GovernanceUsername
	s.ks.log.Info("Keystore: CreateGovernanceUser called with %.*s", maxUserLen, GovernanceUsername)

	reply.Success = true
	return s.ks.CreateGovernanceUser(GovernanceUsername, args.Password)
}
