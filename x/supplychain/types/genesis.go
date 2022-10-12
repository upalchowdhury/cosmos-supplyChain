package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ConCounter:         nil,
		ActorscontractList: []Actorscontract{},
		ContractCounter:    nil,
		NewcontractList:    []Newcontract{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in actorscontract
	actorscontractIndexMap := make(map[string]struct{})

	for _, elem := range gs.ActorscontractList {
		index := string(ActorscontractKey(elem.Index))
		if _, ok := actorscontractIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for actorscontract")
		}
		actorscontractIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in newcontract
	newcontractIndexMap := make(map[string]struct{})

	for _, elem := range gs.NewcontractList {
		index := string(NewcontractKey(elem.Index))
		if _, ok := newcontractIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for newcontract")
		}
		newcontractIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
