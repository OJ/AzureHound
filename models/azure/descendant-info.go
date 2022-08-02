// Copyright (C) 2022 The BloodHound Enterprise Team
//
// This file is part of AzureHound.
//
// AzureHound is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// AzureHound is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package azure

// The properties of the parent management group.
type DescendantParentGroupInfo struct {

	// The fully qualified ID for the parent management group.
	//
	// For example:
	// - /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
	Id string `json:"id"`
}

// DescendantInfoProperties describes the properties of the management group descendant.
type DescendantInfoProperties struct {

	// The friendly name of the management group.
	DisplayName string `json:"display_name"`

	// The properties of the parent management group.
	Parent DescendantParentGroupInfo `json:"parent"`
}

// DescendantInfo is a management group descendant.
type DescendantInfo struct {

	// The fully qualified ID for the descendant.
	//
	// For example:
	// - /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
	// - /subscriptions/0000000-0000-0000-0000-000000000000
	Id string `json:"id"`

	// The name of the descendant.
	//
	// For example:
	// - 00000000-0000-0000-0000-000000000000
	Name string `json:"name"`

	// The properties of the management group descendant.
	Properties DescendantInfoProperties `json:"properties"`

	// The type of the resource.
	//
	// For example:
	// - Microsoft.Management/managementGroups
	// - /subscriptions
	Type string `json:"type"`
}

type DescendantInfoList struct {
	NextLink string           `json:"nextLink,omitempty"` // The URL to use for getting the next set of values.
	Value    []DescendantInfo `json:"value"`              // A list of management group descendants.
}

type DescendantInfoResult struct {
	Error error
	Ok    DescendantInfo
}
