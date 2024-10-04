// Copyright © 2024 Galactica Network
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import "errors"

// ErrSaltIncompatible means that for given human ID there is another zkcertificate.HolderCommitment in the registry.
// This is a way to make sure that within the same registry
// all certificates with the same zkcertificate.HolderCommitment belong to the same person, identified by human ID.
var ErrSaltIncompatible = errors.New("incompatible salt")
