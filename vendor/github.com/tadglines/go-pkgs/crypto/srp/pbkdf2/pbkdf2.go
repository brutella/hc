// Copyright 2013 Tad Glines
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package pbkdf2

import (
	"golang.org/x/crypto/pbkdf2"
	"hash"
)

// NewPBKDF2 returns a new key derivation function that uses pbkdf2 to do
// the derivation. The returned key size will be the same as the hash size.
func NewPBKDF2(iter int, h func() hash.Hash) func(salt, password []byte) []byte {
	return func(salt, password []byte) []byte {
		hf := h()
		return pbkdf2.Key(password, salt, iter, hf.Size(), h)
	}
}
