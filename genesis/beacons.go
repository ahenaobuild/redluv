// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"github.com/hellobuild/Luv-Go/utils/constants"
	"github.com/hellobuild/Luv-Go/utils/sampler"
)

// getIPs returns the beacon IPs for each network
func getIPs(networkID uint32) []string {
	switch networkID {
	case constants.MainnetID:
		return nil // Remove this line, should you activate  public IP beacons in production mainnet
		/*
			return []string{
				"127.0.0.1:9651",//replace with public ip in production mainnet
			}
		*/
	case constants.FujiID:
		return []string{
			"18.188.121.35:21001",
			"3.133.83.66:21001",
			"3.15.206.239:21001",
			"18.224.140.156:21001",
			"3.133.131.39:21001",
			"18.191.29.54:21001",
			"18.224.172.110:21001",
			"18.223.211.203:21001",
			"18.216.130.143:21001",
			"18.223.184.147:21001",
			"52.15.48.84:21001",
			"18.189.194.220:21001",
			"18.223.119.104:21001",
			"3.133.155.41:21001",
			"13.58.170.174:21001",
			"3.21.245.246:21001",
			"52.15.190.149:21001",
			"18.188.95.241:21001",
			"3.12.197.248:21001",
			"3.17.39.236:21001",
		}
	default:
		return nil
	}
}

// getNodeIDs returns the beacon node IDs for each network
func getNodeIDs(networkID uint32) []string {
	switch networkID {
	case constants.MainnetID:
		return nil // Remove this line, should you deploy production mainnet nodes
		/*
			return []string{
				"NodeID-KtNSDKDkPuzUWR6goMsgB8g4HqwZcMvQj",
			}
		*/
	case constants.FujiID:
		return []string{
			"NodeID-NpagUxt6KQiwPch9Sd4osv8kD1TZnkjdk",
			"NodeID-2m38qc95mhHXtrhjyGbe7r2NhniqHHJRB",
			"NodeID-LQwRLm4cbJ7T2kxcxp4uXCU5XD8DFrE1C",
			"NodeID-hArafGhY2HFTbwaaVh1CSCUCUCiJ2Vfb",
			"NodeID-4QBwET5o8kUhvt9xArhir4d3R25CtmZho",
			"NodeID-HGZ8ae74J3odT8ESreAdCtdnvWG1J4X5n",
			"NodeID-4KXitMCoE9p2BHA6VzXtaTxLoEjNDo2Pt",
			"NodeID-JyE4P8f4cTryNV8DCz2M81bMtGhFFHexG",
			"NodeID-EzGaipqomyK9UKx9DBHV6Ky3y68hoknrF",
			"NodeID-CYKruAjwH1BmV3m37sXNuprbr7dGQuJwG",
			"NodeID-LegbVf6qaMKcsXPnLStkdc1JVktmmiDxy",
			"NodeID-FesGqwKq7z5nPFHa5iwZctHE5EZV9Lpdq",
			"NodeID-BFa1padLXBj7VHa2JYvYGzcTBPQGjPhUy",
			"NodeID-4B4rc5vdD1758JSBYL1xyvE5NHGzz6xzH",
			"NodeID-EDESh4DfZFC15i613pMtWniQ9arbBZRnL",
			"NodeID-CZmZ9xpCzkWqjAyS7L4htzh5Lg6kf1k18",
			"NodeID-CTtkcXvVdhpNp6f97LEUXPwsRD3A2ZHqP",
			"NodeID-84KbQHSDnojroCVY7vQ7u9Tx7pUonPaS",
			"NodeID-JjvzhxnLHLUQ5HjVRkvG827ivbLXPwA9u",
			"NodeID-4CWTbdvgXHY1CLXqQNAp22nJDo5nAmts6",
		}
	default:
		return nil
	}
}

// SampleBeacons returns the some beacons this node should connect to
func SampleBeacons(networkID uint32, count int) ([]string, []string) {
	ips := getIPs(networkID)
	ids := getNodeIDs(networkID)

	if numIPs := len(ips); numIPs < count {
		count = numIPs
	}

	sampledIPs := make([]string, 0, count)
	sampledIDs := make([]string, 0, count)

	s := sampler.NewUniform()
	_ = s.Initialize(uint64(len(ips)))
	indices, _ := s.Sample(count)
	for _, index := range indices {
		sampledIPs = append(sampledIPs, ips[int(index)])
		sampledIDs = append(sampledIDs, ids[int(index)])
	}

	return sampledIPs, sampledIDs
}
