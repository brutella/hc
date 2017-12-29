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

package scrypt

import (
	"golang.org/x/crypto/scrypt"
	"errors"
)

const maxInt = int(^uint(0) >> 1)

// NewScrypt returns a new key derivation function that uses scrypt to do
// the derivation. The returned key will be 32 bytes in size.
// If N, r, or p are invalid nil and an error are returned.
// Seegolang.org/x/crypto/scrypt#Key for details on proper values for
// N, r, and p.
func NewScrypt(N, r, p int) (func(salt, password []byte) []byte, error) {
	// The following two checks where copied directly from the scrypt implementation.
	if N <= 1 || N&(N-1) != 0 {
		return nil, errors.New("scrypt: N must be > 1 and a power of 2")
	}
	if uint64(r)*uint64(p) >= 1<<30 || r > maxInt/128/p || r > maxInt/256 || N > maxInt/128/r {
		return nil, errors.New("scrypt: parameters are too large")
	}

	return func(salt, password []byte) []byte {
		key, _ := scrypt.Key(password, salt, N, r, p, 32)
		return key
	}, nil
}
