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

package zkcertificate_test

import (
	"math/big"
	"testing"

	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/stretchr/testify/require"

	"github.com/galactica-corp/guardians-sdk/pkg/zkcertificate"
)

func TestSignCertificate(t *testing.T) {
	privateKey := babyjub.NewRandPrivKey()
	contentHash := zkcertificate.Hash{}
	commitmentHash := zkcertificate.Hash{}

	signature, err := zkcertificate.SignCertificate(privateKey, contentHash, commitmentHash)
	require.NoError(t, err)
	require.NotNil(t, signature)

	isValid, err := zkcertificate.VerifySignature(privateKey.Public(), contentHash, commitmentHash, signature)
	require.NoError(t, err)
	require.True(t, isValid)

	contentHash = zkcertificate.HashFromBigInt(big.NewInt(1))
	isValid, err = zkcertificate.VerifySignature(privateKey.Public(), contentHash, commitmentHash, signature)
	require.NoError(t, err)
	require.False(t, isValid)
}
