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
