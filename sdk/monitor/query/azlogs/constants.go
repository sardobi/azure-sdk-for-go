//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azlogs

// LogsColumnType - The data type of this column.
type LogsColumnType string

const (
	LogsColumnTypeBool     LogsColumnType = "bool"
	LogsColumnTypeDatetime LogsColumnType = "datetime"
	LogsColumnTypeDecimal  LogsColumnType = "decimal"
	LogsColumnTypeDynamic  LogsColumnType = "dynamic"
	LogsColumnTypeGUID     LogsColumnType = "guid"
	LogsColumnTypeInt      LogsColumnType = "int"
	LogsColumnTypeLong     LogsColumnType = "long"
	LogsColumnTypeReal     LogsColumnType = "real"
	LogsColumnTypeString   LogsColumnType = "string"
	LogsColumnTypeTimespan LogsColumnType = "timespan"
)

// PossibleLogsColumnTypeValues returns the possible values for the LogsColumnType const type.
func PossibleLogsColumnTypeValues() []LogsColumnType {
	return []LogsColumnType{
		LogsColumnTypeBool,
		LogsColumnTypeDatetime,
		LogsColumnTypeDecimal,
		LogsColumnTypeDynamic,
		LogsColumnTypeGUID,
		LogsColumnTypeInt,
		LogsColumnTypeLong,
		LogsColumnTypeReal,
		LogsColumnTypeString,
		LogsColumnTypeTimespan,
	}
}