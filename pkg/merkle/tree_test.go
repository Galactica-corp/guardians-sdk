/*
 Copyright © 2023 Galactica Network

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.

 You should have received a copy of the GNU General Public License
 along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package merkle_test

import (
	"testing"

	"github.com/holiman/uint256"
	"github.com/stretchr/testify/require"

	"github.com/galactica-corp/guardians-sdk/pkg/merkle"
)

func makeTree(t *testing.T) *merkle.Tree {
	t.Helper()

	tree, _ := merkle.NewEmptyTree(3, merkle.EmptyLeafValue)
	_ = tree.SetLeaf(0, merkle.TreeNode{Value: uint256.NewInt(10)})
	_ = tree.SetLeaf(1, merkle.TreeNode{Value: uint256.NewInt(20)})
	_ = tree.SetLeaf(2, merkle.TreeNode{Value: uint256.NewInt(30)})
	_ = tree.SetLeaf(3, merkle.TreeNode{Value: uint256.NewInt(40)})

	return tree
}

func TestMakeEmptySparseTree(t *testing.T) {
	depth := 3

	tree, err := merkle.NewEmptyTree(depth, merkle.EmptyLeafValue)
	require.NoError(t, err)

	require.Len(t, tree.Nodes, 1<<depth-1)
}

func TestTree_depth8(t *testing.T) {
	depth := 8

	tree, err := merkle.NewEmptyTree(depth, merkle.EmptyLeafValue)
	require.NoError(t, err)

	require.Len(t, tree.Nodes, 1<<depth-1)

	err = tree.SetLeaf(0, merkle.TreeNode{Value: uint256.NewInt(42)})
	require.NoError(t, err)

	_, err = tree.GetProof(0)
	require.NoError(t, err)
}

func TestTree_SetLeaf(t *testing.T) {
	depth := 3

	tree, err := merkle.NewEmptyTree(depth, merkle.EmptyLeafValue)
	require.NoError(t, err)

	value := merkle.TreeNode{Value: uint256.NewInt(42)}
	index := 2

	err = tree.SetLeaf(index, value)
	require.NoError(t, err)

	leafTreeIndex := 5

	if !tree.Nodes[leafTreeIndex].Value.Eq(value.Value) {
		t.Errorf("Expected leaf value at leaf index %d: %v, got: %v", index, value, tree.Nodes[leafTreeIndex])
	}
}

func TestTree_SetLeaf_outOfRange(t *testing.T) {
	depth := 3

	tree, err := merkle.NewEmptyTree(depth, merkle.EmptyLeafValue)
	require.NoError(t, err)

	err = tree.SetLeaf(4, merkle.TreeNode{Value: uint256.NewInt(42)})
	require.Error(t, err)
}

func TestTree_GetProof(t *testing.T) {
	tree := makeTree(t)

	index := 2

	proof, err := tree.GetProof(index)
	require.NoError(t, err)

	expectedPath := []merkle.TreeNode{tree.Nodes[6], tree.Nodes[1], tree.Nodes[0]}
	require.True(t, areTreeNodeSlicesEqual(expectedPath, proof.Path), "proof paths are not equal")
	require.True(t, proof.Leaf.Value.Eq(tree.Nodes[5].Value))
}

func TestTree_GetProof_outOfRange(t *testing.T) {
	tree := makeTree(t)

	_, err := tree.GetProof(4)
	require.Error(t, err)
}

func TestTree_Root(t *testing.T) {
	tree, err := merkle.NewEmptyTree(1, merkle.EmptyLeafValue)
	require.NoError(t, err)

	root := tree.Root()
	require.True(t, merkle.EmptyLeafValue.Eq(root.Value))
}

func TestTree_Root_multipleLeaves(t *testing.T) {
	tree := makeTree(t)

	expectedValue, err := uint256.FromDecimal("17334160021514922261858520074371062620775267607575594622555131834681402379786")
	require.NoError(t, err)

	root := tree.Root()
	require.True(t, expectedValue.Eq(root.Value), "invalid merkle root")
}

func areTreeNodeSlicesEqual(a, b []merkle.TreeNode) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !a[i].Value.Eq(b[i].Value) {
			return false
		}
	}

	return true
}
