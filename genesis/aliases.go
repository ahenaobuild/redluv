// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"github.com/hellobuild/Luv-Go/ids"
	"github.com/hellobuild/Luv-Go/utils/constants"
	"github.com/hellobuild/Luv-Go/vms/avm"
	"github.com/hellobuild/Luv-Go/vms/evm"
	"github.com/hellobuild/Luv-Go/vms/nftfx"
	"github.com/hellobuild/Luv-Go/vms/platformvm"
	"github.com/hellobuild/Luv-Go/vms/propertyfx"
	"github.com/hellobuild/Luv-Go/vms/secp256k1fx"
	"github.com/hellobuild/Luv-Go/vms/timestampvm"
)

// Aliases returns the default aliases based on the network ID
func Aliases(genesisBytes []byte) (map[string][]string, map[ids.ID][]string, map[ids.ID][]string, error) {
	generalAliases := getGeneralAliases()
	chainAliases := map[ids.ID][]string{
		constants.PlatformChainID: {"P", "platform"},
	}
	vmAliases := getChainAliases()
	genesis := &platformvm.Genesis{} // TODO let's not re-create genesis to do aliasing
	if _, err := platformvm.GenesisCodec.Unmarshal(genesisBytes, genesis); err != nil {
		return nil, nil, nil, err
	}
	if err := genesis.Initialize(); err != nil {
		return nil, nil, nil, err
	}

	for _, chain := range genesis.Chains {
		uChain := chain.UnsignedTx.(*platformvm.UnsignedCreateChainTx)
		switch uChain.VMID {
		case avm.ID:
			generalAliases["bc/"+chain.ID().String()] = []string{"X", "avm", "bc/X", "bc/avm"}
			chainAliases[chain.ID()] = GetXChainAliases()
		case evm.ID:
			generalAliases["bc/"+chain.ID().String()] = []string{"C", "evm", "bc/C", "bc/evm"}
			chainAliases[chain.ID()] = GetCChainAliases()
		case timestampvm.ID:
			generalAliases["bc/"+chain.ID().String()] = []string{"bc/timestamp"}
			chainAliases[chain.ID()] = []string{"timestamp"}
		}
	}
	return generalAliases, chainAliases, vmAliases, nil
}

func GetCChainAliases() []string {
	return []string{"C", "evm"}
}

func GetXChainAliases() []string {
	return []string{"X", "avm"}
}

func getGeneralAliases() map[string][]string {
	return map[string][]string{
		"vm/" + platformvm.ID.String():             {"vm/platform"},
		"vm/" + avm.ID.String():                    {"vm/avm"},
		"vm/" + evm.ID.String():                    {"vm/evm"},
		"vm/" + timestampvm.ID.String():            {"vm/timestamp"},
		"bc/" + constants.PlatformChainID.String(): {"P", "platform", "bc/P", "bc/platform"},
	}
}

func getChainAliases() map[ids.ID][]string {
	return map[ids.ID][]string{
		platformvm.ID:  {"platform"},
		avm.ID:         {"avm"},
		evm.ID:         {"evm"},
		timestampvm.ID: {"timestamp"},
		secp256k1fx.ID: {"secp256k1fx"},
		nftfx.ID:       {"nftfx"},
		propertyfx.ID:  {"propertyfx"},
	}
}
