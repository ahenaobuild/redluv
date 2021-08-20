package api

import (
	"fmt"

	"github.com/hellobuild/Luv-Go/utils/constants"
)

//Verify the request Username is not the same governance username
func CheckGovernanceUsername(Username string) error {
	if Username == constants.GovernanceUsername {
		return fmt.Errorf("can't use %s username with this endpoint", constants.GovernanceUsername)
	}

	return nil
}
