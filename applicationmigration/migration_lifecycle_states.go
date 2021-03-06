// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Application Migration Service API
//
// API for the Application Migration service. Use this API to migrate applications from Oracle Cloud Infrastructure - Classic to Oracle Cloud Infrastructure.
//

package applicationmigration

// MigrationLifecycleStatesEnum Enum with underlying type: string
type MigrationLifecycleStatesEnum string

// Set of constants representing the allowable values for MigrationLifecycleStatesEnum
const (
	MigrationLifecycleStatesCreating  MigrationLifecycleStatesEnum = "CREATING"
	MigrationLifecycleStatesActive    MigrationLifecycleStatesEnum = "ACTIVE"
	MigrationLifecycleStatesInactive  MigrationLifecycleStatesEnum = "INACTIVE"
	MigrationLifecycleStatesUpdating  MigrationLifecycleStatesEnum = "UPDATING"
	MigrationLifecycleStatesSucceeded MigrationLifecycleStatesEnum = "SUCCEEDED"
	MigrationLifecycleStatesDeleting  MigrationLifecycleStatesEnum = "DELETING"
	MigrationLifecycleStatesDeleted   MigrationLifecycleStatesEnum = "DELETED"
)

var mappingMigrationLifecycleStates = map[string]MigrationLifecycleStatesEnum{
	"CREATING":  MigrationLifecycleStatesCreating,
	"ACTIVE":    MigrationLifecycleStatesActive,
	"INACTIVE":  MigrationLifecycleStatesInactive,
	"UPDATING":  MigrationLifecycleStatesUpdating,
	"SUCCEEDED": MigrationLifecycleStatesSucceeded,
	"DELETING":  MigrationLifecycleStatesDeleting,
	"DELETED":   MigrationLifecycleStatesDeleted,
}

// GetMigrationLifecycleStatesEnumValues Enumerates the set of values for MigrationLifecycleStatesEnum
func GetMigrationLifecycleStatesEnumValues() []MigrationLifecycleStatesEnum {
	values := make([]MigrationLifecycleStatesEnum, 0)
	for _, v := range mappingMigrationLifecycleStates {
		values = append(values, v)
	}
	return values
}
