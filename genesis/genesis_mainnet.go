cd /// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"time"

	"github.com/hellobuild/Luv-Go/utils/units"
)

var (
	mainnetGenesisConfigJSON = `{
		"networkID": 1,
		"allocations": [
			{
				"ethAddr": "0xB8CfBAd9a6617d78c4f7b415D5F4b407F19538Bc",
				"avaxAddr": "X-luv12k4qkg4sqtfcrpvjznzces7gfzq6638q26q9wd",
				"initialAmount": 600000000000000000,
				"unlockSchedule": [
					{
						"amount": 20000000000000000,
						"locktime": 1633824000
					}
				]
			}
		],
		"startTime": 1599696000,
		"initialStakeDuration": 31536000,
		"initialStakeDurationOffset": 5400,
		"initialStakedFunds": [
			"X-luv12k4qkg4sqtfcrpvjznzces7gfzq6638q26q9wd"
		],
		"initialStakers": [
			{
				"nodeID": "NodeID-KtNSDKDkPuzUWR6goMsgB8g4HqwZcMvQj",
				"rewardAddress": "X-luv12k4qkg4sqtfcrpvjznzces7gfzq6638q26q9wd",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-ELjNtAjNPosPuawqStfzXFeRYR5DMTLCE",
				"rewardAddress": "X-luv12k4qkg4sqtfcrpvjznzces7gfzq6638q26q9wd",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-NFBbbJ4qCmNaCzeW7sxErhvWqvEQMnYcN",
				"rewardAddress": "X-luv12k4qkg4sqtfcrpvjznzces7gfzq6638q26q9wd",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-GWPcbFJZFfZreETSoWjPimr846mXEKCtu",
				"rewardAddress": "X-luv12k4qkg4sqtfcrpvjznzces7gfzq6638q26q9wd",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-FmLXATAtLNHuCJ742FiWvAazebVikMW2s",
				"rewardAddress": "X-luv12k4qkg4sqtfcrpvjznzces7gfzq6638q26q9wd",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-HwCMhYEksPqxAvQfiJroHcdyxaGvasNjy",
				"rewardAddress": "X-luv12k4qkg4sqtfcrpvjznzces7gfzq6638q26q9wd",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-KUpXZ1jDPWjUsNfGDEuD7dm1fZHBsASnm",
				"rewardAddress": "X-luv12k4qkg4sqtfcrpvjznzces7gfzq6638q26q9wd",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-5qSbXCquiAyayQjLeBH4yWi4hucYEHmnt",
				"rewardAddress": "X-luv12k4qkg4sqtfcrpvjznzces7gfzq6638q26q9wd",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-KFh3qPsghJXevtxzqPqVdEdNRFufyydeX",
				"rewardAddress": "X-luv12k4qkg4sqtfcrpvjznzces7gfzq6638q26q9wd",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-AGACGtMreSp32S4ZrEM7Eqq9KccZVL3uj",
				"rewardAddress": "X-luv12k4qkg4sqtfcrpvjznzces7gfzq6638q26q9wd",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-DJWSDtUTg7VBtQpCt3rbbs56cH1v31aQm",
				"rewardAddress": "X-luv12k4qkg4sqtfcrpvjznzces7gfzq6638q26q9wd",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-2x5di7tDYKZpnUtpS2soHTyMyXP6D2aTM",
				"rewardAddress": "X-luv12k4qkg4sqtfcrpvjznzces7gfzq6638q26q9wd",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-JoTfjh9wSiGM5MDkFggRK3HUfJUw77RhT",
				"rewardAddress": "X-luv12k4qkg4sqtfcrpvjznzces7gfzq6638q26q9wd",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-9NAmuNZjNoKstaZM8qYpKqFGDNWG1pgmY",
				"rewardAddress": "X-luv12k4qkg4sqtfcrpvjznzces7gfzq6638q26q9wd",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-4XXq2tWE8aJwBL6NcKmEmy6ZWKwaSnJui",
				"rewardAddress": "X-luv12k4qkg4sqtfcrpvjznzces7gfzq6638q26q9wd",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-GbgM5Cva8iZ2fqpyC5TqA623hbC5jcY7D",
				"rewardAddress": "X-luv12k4qkg4sqtfcrpvjznzces7gfzq6638q26q9wd",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-HsTGvmpZgSSb27ahQn8L1REP9BbNtgWAJ",
				"rewardAddress": "X-luv12k4qkg4sqtfcrpvjznzces7gfzq6638q26q9wd",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-K7sf7wN9QKqBaJycq2U1Pb1xTZJBQB4nj",
				"rewardAddress": "X-luv12k4qkg4sqtfcrpvjznzces7gfzq6638q26q9wd",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-NRFVcRv9aujKTSveY6YkSpbXCfmZZa6LM",
				"rewardAddress": "X-luv12k4qkg4sqtfcrpvjznzces7gfzq6638q26q9wd",
				"delegationFee": 200000
			},
			{
				"nodeID": "NodeID-6nWf5fxk7yMdiTfczLgRv4XNa2jvexycY",
				"rewardAddress": "X-luv12k4qkg4sqtfcrpvjznzces7gfzq6638q26q9wd",
				"delegationFee": 200000
			}
			
		],
		"cChainGenesis": "{\"config\":{\"chainId\":43114,\"homesteadBlock\":0,\"daoForkBlock\":0,\"daoForkSupport\":true,\"eip150Block\":0,\"eip150Hash\":\"0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0\",\"eip155Block\":0,\"eip158Block\":0,\"byzantiumBlock\":0,\"constantinopleBlock\":0,\"petersburgBlock\":0,\"istanbulBlock\":0,\"muirGlacierBlock\":0},\"nonce\":\"0x0\",\"timestamp\":\"0x0\",\"extraData\":\"0x00\",\"gasLimit\":\"0x5f5e100\",\"difficulty\":\"0x0\",\"mixHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"coinbase\":\"0x0000000000000000000000000000000000000000\",\"alloc\":{\"0100000000000000000000000000000000000000\":{\"code\":\"0x7300000000000000000000000000000000000000003014608060405260043610603d5760003560e01c80631e010439146042578063b6510bb314606e575b600080fd5b605c60048036036020811015605657600080fd5b503560b1565b60408051918252519081900360200190f35b818015607957600080fd5b5060af60048036036080811015608e57600080fd5b506001600160a01b03813516906020810135906040810135906060013560b6565b005b30cd90565b836001600160a01b031681836108fc8690811502906040516000604051808303818888878c8acf9550505050505015801560f4573d6000803e3d6000fd5b505050505056fea26469706673582212201eebce970fe3f5cb96bf8ac6ba5f5c133fc2908ae3dcd51082cfee8f583429d064736f6c634300060a0033\",\"balance\":\"0x0\"}},\"number\":\"0x0\",\"gasUsed\":\"0x0\",\"parentHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\"}",
		"message": "From Snowflake to Avalanche. Per consensum ad astra."
	}`

	// MainnetParams are the params used for mainnet
	MainnetParams = Params{
		TxFee:                units.MilliLuv,
		CreationTxFee:        10 * units.MilliLuv,
		UptimeRequirement:    .6, // 60%
		MinValidatorStake:    2 * units.KiloLuv,
		MaxValidatorStake:    3 * units.MegaLuv,
		MinDelegatorStake:    25 * units.Luv,
		MinDelegationFee:     20000, // 2%
		MinStakeDuration:     2 * 7 * 24 * time.Hour,
		MaxStakeDuration:     365 * 24 * time.Hour,
		StakeMintingPeriod:   365 * 24 * time.Hour,
		EpochFirstTransition: time.Unix(1607626800, 0),
		EpochDuration:        6 * time.Hour,
	}
)
