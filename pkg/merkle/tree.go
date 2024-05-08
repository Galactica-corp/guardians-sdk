// Copyright © 2024 Galactica Network
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package merkle

import (
	"context"
	"fmt"
	"math/big"

	merkleProofService "github.com/galactica-corp/merkle-proof-service/gen/galactica/merkle"
	"github.com/holiman/uint256"
	"github.com/iden3/go-iden3-crypto/ff"
	"golang.org/x/crypto/sha3"
	"google.golang.org/grpc"
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

var EmptyLeafValue = new(uint256.Int).Mod(
	uint256.MustFromBig(new(big.Int).SetBytes(makeSeedForEmptyLeaf())),
	uint256.MustFromBig(ff.Modulus()),
)

func makeSeedForEmptyLeaf() []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte("Galactica"))
	return hash.Sum(nil)
}

func GetProof(ctx context.Context, client merkleProofService.QueryClient, registryAddress string, leaf string) (Proof, error) {
	proofResp, err := client.Proof(ctx, &merkleProofService.QueryProofRequest{
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

func GetEmptyIndex(ctx context.Context, client merkleProofService.QueryClient, registryAddress string) (uint32, error) {
	resp, err := client.GetEmptyIndex(ctx, &merkleProofService.GetEmptyIndexRequest{Registry: registryAddress})
	if err != nil {
		return 0, fmt.Errorf("get empty leaf index: %w", err)
	}

	return resp.Index, nil
}

func toSDKProof(proof *merkleProofService.Proof) (Proof, error) {
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

func ConnectToMerkleProofService(ctx context.Context, merkleProofServiceHost string) (merkleProofService.QueryClient, error) {
	conn, err := grpc.DialContext(ctx, merkleProofServiceHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("connect to Merkle Proof Service: %w", err)
	}

	return merkleProofService.NewQueryClient(conn), nil
}
