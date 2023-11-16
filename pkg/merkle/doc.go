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
