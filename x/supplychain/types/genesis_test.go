package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/upalchowdhury/supplychain/x/supplychain/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				ConCounter: &types.ConCounter{
					IdValue: 24,
				},
				ActorscontractList: []types.Actorscontract{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				ContractCounter: &types.ContractCounter{
					DealId:  "91",
					IdValue: 65,
				},
				NewcontractList: []types.Newcontract{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated actorscontract",
			genState: &types.GenesisState{
				ActorscontractList: []types.Actorscontract{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated newcontract",
			genState: &types.GenesisState{
				NewcontractList: []types.Newcontract{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
