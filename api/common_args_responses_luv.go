package api

type GetValidatorWhitelistArgs struct {
	UserPass
	NodeId string `json:"nodeId"`
}

type AddValidatorWhitelistArgs struct {
	UserPass
	NodeID string `json:"nodeId"`
}

type GovernancePassword struct {
	Password string `json:"password"`
}
