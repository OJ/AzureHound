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

// The instance view of the disk.
type DiskInstanceView struct {
	// Specifies the encryption settings for the OS Disk.
	// Minimum api-version: 2015-06-15
	EncryptionSettings []DiskEncryptionSettings `json:"encryptionSettings"`

	// The disk name.
	Name string `json:"name"`

	// The resource status information.
	Statuses []InstanceViewStatus `json:"statuses"`
}
