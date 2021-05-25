package managementpartner

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Code enumerates the values for code.
type Code string

const (
	// BadRequest ...
	BadRequest Code = "BadRequest"
	// Conflict ...
	Conflict Code = "Conflict"
	// NotFound ...
	NotFound Code = "NotFound"
)

// PossibleCodeValues returns an array of possible values for the Code const type.
func PossibleCodeValues() []Code {
	return []Code{BadRequest, Conflict, NotFound}
}

// State enumerates the values for state.
type State string

const (
	// Active ...
	Active State = "Active"
	// Deleted ...
	Deleted State = "Deleted"
)

// PossibleStateValues returns an array of possible values for the State const type.
func PossibleStateValues() []State {
	return []State{Active, Deleted}
}