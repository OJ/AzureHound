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

import "encoding/json"

// Represents an Azure Active Directory object. The directoryObject type is the base type for many other directory entity types.
type DirectoryObject struct {
	// The unique identifier for the object.
	// Note: The value is often but not exclusively a GUID (UUID v4 variant 2)
	//
	// Key
	// Read-only
	// Supports `filter` (eq,ne,NOT,in)
	Id string `json:"id"`

	Type string `json:"@odata.type,omitempty"`
}

type DirectoryObjectList struct {
	Count    int               `json:"@odata.count,omitempty"`    // The total count of all results
	NextLink string            `json:"@odata.nextLink,omitempty"` // The URL to use for getting the next set of values.
	Value    []json.RawMessage `json:"value"`                     // A list of various Azure AD directory objects.
}
