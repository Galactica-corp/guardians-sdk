// Package merkle provides functionality for creating and manipulating Merkle trees,
// with a focus on efficient and secure verification of data integrity.
//
// The package offers tools for constructing Merkle trees, managing tree nodes,
// and generating and verifying Merkle proofs. It includes constants for default
// tree depth and empty leaf values. Key features include the ability to create
// new empty trees, set leaf values, retrieve proofs, and calculate the root hash
// of the tree.
//
// Additionally, the package provides a hash function for big integers and utility
// functions for navigating the Merkle tree structure, such as obtaining parent and
// sibling indices and determining whether a node is a right child.
//
// Note: The package assumes the use of 256-bit integers for node values.
package merkle
