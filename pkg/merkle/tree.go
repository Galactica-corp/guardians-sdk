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

package merkle

import (
	"context"
	"crypto/tls"
	"fmt"
	"math/big"

	"github.com/Galactica-corp/merkle-proof-service/gen/galactica/merkle"
	"github.com/holiman/uint256"
	"github.com/iden3/go-iden3-crypto/v2/poseidon"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type TreeNode struct {
	Value *uint256.Int
}

func (n *TreeNode) UnmarshalText(text []byte) error {
	var err error
	n.Value, err = uint256.FromDecimal(string(text))

	return err
}

func (n TreeNode) MarshalText() (text []byte, err error) {
	return []byte(n.Value.Dec()), nil
}

type Proof struct {
	Leaf      TreeNode   `json:"leaf"`      // The Merkle tree node, which authenticity is proved by the Path.
	LeafIndex int        `json:"leafIndex"` // Index of the Leaf in the Merkle tree.
	Path      []TreeNode `json:"path"`
}

type Prover interface {
	Proof(
		ctx context.Context,
		in *merkle.QueryProofRequest,
		opts ...grpc.CallOption,
	) (*merkle.QueryProofResponse, error)
}

func GetProof(ctx context.Context, client Prover, registryAddress string, leaf string) (Proof, error) {
	proofResp, err := client.Proof(ctx, &merkle.QueryProofRequest{
		Registry: registryAddress,
		Leaf:     leaf,
	})
	if err != nil {
		return Proof{}, fmt.Errorf("get merkle proof: %w", err)
	}

	sdkProof, err := toSDKProof(proofResp.Proof)
	if err != nil {
		return Proof{}, fmt.Errorf("convert proof: %w", err)
	}

	return sdkProof, nil
}

type EmptyLeafProver interface {
	GetEmptyLeafProof(
		ctx context.Context,
		in *merkle.GetEmptyLeafProofRequest,
		opts ...grpc.CallOption,
	) (*merkle.GetEmptyLeafProofResponse, error)
}

func GetEmptyLeafProof(ctx context.Context, client EmptyLeafProver, registryAddress string) (uint32, Proof, error) {
	resp, err := client.GetEmptyLeafProof(ctx, &merkle.GetEmptyLeafProofRequest{Registry: registryAddress})
	if err != nil {
		return 0, Proof{}, fmt.Errorf("get empty leaf index: %w", err)
	}

	sdkProof, err := toSDKProof(resp.Proof)
	if err != nil {
		return 0, Proof{}, fmt.Errorf("convert proof: %w", err)
	}

	return resp.Proof.Index, sdkProof, nil
}

func toSDKProof(proof *merkle.Proof) (Proof, error) {
	path := make([]TreeNode, len(proof.Path))
	for i, node := range proof.Path {
		value, err := uint256.FromDecimal(node)
		if err != nil {
			return Proof{}, fmt.Errorf("convert path node value: %w", err)
		}
		path[i] = TreeNode{Value: value}
	}

	leaf, err := uint256.FromDecimal(proof.Leaf)
	if err != nil {
		return Proof{}, fmt.Errorf("convert leaf value: %w", err)
	}

	return Proof{
		Leaf:      TreeNode{Value: leaf},
		LeafIndex: int(proof.Index),
		Path:      path,
	}, nil
}

func ConnectToMerkleProofService(merkleProofServiceHost string, useTLS bool) (merkle.QueryClient, error) {
	var creds credentials.TransportCredentials
	if useTLS {
		creds = credentials.NewTLS(&tls.Config{})
	} else {
		creds = insecure.NewCredentials()
	}

	conn, err := grpc.NewClient(
		merkleProofServiceHost,
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		return nil, fmt.Errorf("connect to Merkle Proof Service: %w", err)
	}

	return merkle.NewQueryClient(conn), nil
}

func Verify(proof Proof, root TreeNode) error {
	res := proof.Leaf
	index := proof.LeafIndex

	for _, node := range proof.Path {
		var err error
		if index%2 == 0 {
			res, err = computeNodeHash(res, node)
		} else {
			res, err = computeNodeHash(node, res)
		}
		if err != nil {
			return fmt.Errorf("compute hash: %w", err)
		}

		index /= 2
	}

	resText, err := res.MarshalText()
	if err != nil {
		return fmt.Errorf("marshal proof: %w", err)
	}

	fmt.Println(string(resText))

	if !res.Value.Eq(root.Value) {
		return fmt.Errorf("root value not equal")
	}

	return nil
}

func computeNodeHash(left, right TreeNode) (TreeNode, error) {
	val, err := poseidon.Hash([]*big.Int{left.Value.ToBig(), right.Value.ToBig()})
	if err != nil {
		return TreeNode{}, err
	}

	convertedVal, isOverflow := uint256.FromBig(val)
	if isOverflow {
		return TreeNode{}, fmt.Errorf("invalid hash")
	}

	return TreeNode{Value: convertedVal}, nil
}
