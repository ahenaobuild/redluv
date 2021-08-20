// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

import (
	"fmt"
	"testing"
	"time"

	"github.com/hellobuild/Luv-Go/utils/units"
)

func TestRewardLongerDurationBonus(t *testing.T) {
	shortDuration := 24 * time.Hour
	totalDuration := 365 * 24 * time.Hour
	shortBalance := units.KiloLuv
	for i := 0; i < int(totalDuration/shortDuration); i++ {
		r := reward(shortDuration, shortBalance, 359*units.MegaLuv+shortBalance, defaultMaxStakingDuration)
		shortBalance += r
	}
	r := reward(totalDuration%shortDuration, shortBalance, 359*units.MegaLuv+shortBalance, defaultMaxStakingDuration)
	shortBalance += r

	longBalance := units.KiloLuv
	longBalance += reward(totalDuration, longBalance, 359*units.MegaLuv+longBalance, defaultMaxStakingDuration)

	if shortBalance >= longBalance {
		t.Fatalf("should promote stakers to stake longer")
	}
}

func TestRewards(t *testing.T) {
	tests := []struct {
		duration       time.Duration
		stakeAmount    uint64
		existingAmount uint64
		expectedReward uint64
	}{
		// Max duration:
		{ // (720M - 360M) * (1M / 360M) * 12%
			duration:       defaultMaxStakingDuration,
			stakeAmount:    units.MegaLuv,
			existingAmount: 360 * units.MegaLuv,
			expectedReward: 120 * units.KiloLuv,
		},
		{ // (720M - 400M) * (1M / 400M) * 12%
			duration:       defaultMaxStakingDuration,
			stakeAmount:    units.MegaLuv,
			existingAmount: 400 * units.MegaLuv,
			expectedReward: 96 * units.KiloLuv,
		},
		{ // (720M - 400M) * (2M / 400M) * 12%
			duration:       defaultMaxStakingDuration,
			stakeAmount:    2 * units.MegaLuv,
			existingAmount: 400 * units.MegaLuv,
			expectedReward: 192 * units.KiloLuv,
		},
		{ // (720M - 720M) * (1M / 720M) * 12%
			duration:       defaultMaxStakingDuration,
			stakeAmount:    units.MegaLuv,
			existingAmount: SupplyCap,
			expectedReward: 0,
		},
		// Min duration:
		// (720M - 360M) * (1M / 360M) * (10% + 2% * MinimumStakingDuration / MaximumStakingDuration) * MinimumStakingDuration / MaximumStakingDuration
		{
			duration:       defaultMinStakingDuration,
			stakeAmount:    units.MegaLuv,
			existingAmount: 360 * units.MegaLuv,
			expectedReward: 274122724713,
		},
		// (720M - 360M) * (.005 / 360M) * (10% + 2% * MinimumStakingDuration / MaximumStakingDuration) * MinimumStakingDuration / MaximumStakingDuration
		{
			duration:       defaultMinStakingDuration,
			stakeAmount:    defaultMinValidatorStake,
			existingAmount: 360 * units.MegaLuv,
			expectedReward: 1370,
		},
		// (720M - 400M) * (1M / 400M) * (10% + 2% * MinimumStakingDuration / MaximumStakingDuration) * MinimumStakingDuration / MaximumStakingDuration
		{
			duration:       defaultMinStakingDuration,
			stakeAmount:    units.MegaLuv,
			existingAmount: 400 * units.MegaLuv,
			expectedReward: 219298179771,
		},
		// (720M - 400M) * (2M / 400M) * (10% + 2% * MinimumStakingDuration / MaximumStakingDuration) * MinimumStakingDuration / MaximumStakingDuration
		{
			duration:       defaultMinStakingDuration,
			stakeAmount:    2 * units.MegaLuv,
			existingAmount: 400 * units.MegaLuv,
			expectedReward: 438596359542,
		},
		// (720M - 720M) * (1M / 720M) * (10% + 2% * MinimumStakingDuration / MaximumStakingDuration) * MinimumStakingDuration / MaximumStakingDuration
		{
			duration:       defaultMinStakingDuration,
			stakeAmount:    units.MegaLuv,
			existingAmount: SupplyCap,
			expectedReward: 0,
		},
	}
	for _, test := range tests {
		name := fmt.Sprintf("reward(%s,%d,%d)==%d",
			test.duration,
			test.stakeAmount,
			test.existingAmount,
			test.expectedReward,
		)
		t.Run(name, func(t *testing.T) {
			r := reward(
				test.duration,
				test.stakeAmount,
				test.existingAmount,
				defaultMaxStakingDuration,
			)
			if r != test.expectedReward {
				t.Fatalf("expected %d; got %d", test.expectedReward, r)
			}
		})
	}
}