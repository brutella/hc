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

// Package srp provides an implementation of SRP-6a as detailed
// at: http://srp.stanford.edu/design.html
package srp

import (
	"crypto/rand"
	"crypto/subtle"
	"fmt"
	"hash"
	"io"
	"math/big"
)

const (
	DefaultSaltLength = 20
	DefaultABSize     = 256
)

type KeyDerivationFunc func(salt, password []byte) []byte

type HashFunc func() hash.Hash

// SRP contains values that must be the the same for both the client and server.
// SaltLength and ABSize are defaulted by NewSRP but can be changed after an SRP
// instance is created.
// Instances of SRP are safe for concurrent use.
type SRP struct {
	SaltLength        int  // The size of the salt in bytes
	ABSize            uint // The size of a and b in bits
	HashFunc          HashFunc
	KeyDerivationFunc KeyDerivationFunc
	Group             *SRPGroup
	_k                *big.Int
}

// ClientSession represents the client side of an SRP authentication session.
// ClientSession instances cannot be reused.
// Instances of ClientSession are NOT safe for concurrent use.
type ClientSession struct {
	SRP      *SRP
	username []byte
	salt     []byte
	password []byte
	_a       *big.Int
	_A       *big.Int
	_B       *big.Int
	_u       *big.Int
	key      []byte
	_M       []byte
}

// ServerSession represents the client side of an SRP authentication session.
// ServerSession instances cannot be reused.
// Instances of ServerSession are NOT safe for concurrent use.
type ServerSession struct {
	SRP      *SRP
	username []byte
	salt     []byte
	verifier []byte
	_v       *big.Int
	_b       *big.Int
	_A       *big.Int
	_B       *big.Int
	_u       *big.Int
	key      []byte
}

// NewSRP creates a new SRP context that will use the specified group and hash
// functions. If the KeyDeivationFunction is nil, then the HashFunc will be
// used instead.
// The set of supported groups are:
// 		rfc5054.1024
//		rfc5054.1536
//		rfc5054.2048
//		rfc5054.3072
//		rfc5054.4096
//		rfc5054.6144
//		rfc5054.8192
// 		stanford.1024
//		stanford.1536
//		stanford.2048
//		stanford.3072
//		stanford.4096
//		stanford.6144
//		stanford.8192
// The rfc5054 groups are from RFC5054
// The stanford groups where extracted from the stanford patch to OpenSSL.
func NewSRP(group string, h HashFunc, kd KeyDerivationFunc) (*SRP, error) {
	srp := new(SRP)
	srp.SaltLength = DefaultSaltLength
	srp.ABSize = DefaultABSize
	srp.HashFunc = h
	grp, ok := srp_groups[group]
	if !ok {
		return nil, fmt.Errorf("Invalid Group: %s", group)
	}
	srp.Group = grp

	srp.compute_k()

	if kd == nil {
		kd = func(salt, password []byte) []byte {
			h := srp.HashFunc()
			h.Write(salt)
			h.Write(password)
			return h.Sum(nil)
		}
	}
	srp.KeyDerivationFunc = kd

	return srp, nil
}

// utility function that hashes the provided data with the provided hashfunction,
// returning the digest.
func quickHash(hf HashFunc, data []byte) []byte {
	h := hf()
	h.Write(data)
	return h.Sum(nil)
}

// ComputeVerifier generates a random salt and computes the verifier value that
// is associated with the user on the server.
func (s *SRP) ComputeVerifier(password []byte) (salt []byte, verifier []byte, err error) {
	//  x = H(s, p)               (s is chosen randomly)
	salt = make([]byte, s.SaltLength)
	n, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, nil, err
	}
	if n != len(salt) {
		return nil, nil, fmt.Errorf("Expected %d random bytes but only got %d bytes", s.SaltLength, n)
	}

	//  v = g^x                   (computes password verifier)
	x := new(big.Int).SetBytes(s.KeyDerivationFunc(salt, password))
	v := new(big.Int).Exp(s.Group.Generator, x, s.Group.Prime)

	return salt, v.Bytes(), nil
}

// NewClientSession creates a new ClientSession.
func (s *SRP) NewClientSession(username, password []byte) *ClientSession {
	cs := new(ClientSession)
	cs.SRP = s
	cs.username = username
	cs.password = password
	cs._a = s.gen_rand_ab()

	// g^a
	cs._A = new(big.Int).Exp(cs.SRP.Group.Generator, cs._a, cs.SRP.Group.Prime)
	return cs
}

// NewServerSession creates a new ServerSession.
func (s *SRP) NewServerSession(username, salt, verifier []byte) *ServerSession {
	ss := new(ServerSession)
	ss.SRP = s
	ss.username = username
	ss.salt = salt
	ss.verifier = verifier
	ss._b = s.gen_rand_ab()
	ss._v = new(big.Int).SetBytes(verifier)

	// B = (kv + g^b) mod n (blinding)
	kv := new(big.Int).Mul(ss.SRP._k, ss._v)
	kvgb := new(big.Int).Add(kv, new(big.Int).Exp(ss.SRP.Group.Generator, ss._b, ss.SRP.Group.Prime))
	ss._B = new(big.Int).Mod(kvgb, ss.SRP.Group.Prime)
	return ss
}

// GetA returns the bytes of the A value that need to be given to the server.
func (cs *ClientSession) GetA() []byte {
	return cs._A.Bytes()
}

// SetB sets the value of B that was returned by the server
func (cs *ClientSession) setB(B []byte) error {
	cs._B = new(big.Int).SetBytes(B)
	// B == 0
	if cs._B.BitLen() == 0 {
		return fmt.Errorf("B == 0")
	}
	// B >= modulus
	if cs._B.Cmp(cs.SRP.Group.Prime) >= 0 {
		return fmt.Errorf("B >= modulus")
	}
	if !cs.SRP.is_AB_valid(cs._B) {
		return fmt.Errorf("B%%N == 0")
	}
	cs._u = cs.SRP.compute_u(cs._A, cs._B)
	if cs._u.BitLen() == 0 {
		return fmt.Errorf("H(A, B) == 0")
	}
	return nil
}

// ComputeKey computes the session key given the salt and the value of B.
func (cs *ClientSession) ComputeKey(salt, B []byte) ([]byte, error) {
	cs.salt = salt

	err := cs.setB(B)
	if err != nil {
		return nil, err
	}

	// x = H(s, p)                 (user enters password)
	x := new(big.Int).SetBytes(cs.SRP.KeyDerivationFunc(cs.salt, cs.password))

	// S = (B - kg^x) ^ (a + ux)   (computes session key)
	// t1 = g^x
	t1 := new(big.Int).Exp(cs.SRP.Group.Generator, x, cs.SRP.Group.Prime)
	// unblind verifier
	t1.Sub(cs.SRP.Group.Prime, t1)
	t1.Mul(cs.SRP._k, t1)
	t1.Add(t1, cs._B)
	t1.Mod(t1, cs.SRP.Group.Prime)

	// t2 = ux
	t2 := new(big.Int).Mul(cs._u, x)
	// t2 = a + ux
	t2.Add(cs._a, t2)

	// t1 = (B - kg^x) ^ (a + ux)
	t3 := new(big.Int).Exp(t1, t2, cs.SRP.Group.Prime)
	// K = H(S)
	cs.key = quickHash(cs.SRP.HashFunc, t3.Bytes())

	return cs.key, nil
}

// GetKey returns the previously computed key
func (cs *ClientSession) GetKey() []byte {
	return cs.key
}

func computeClientAuthenticator(hf HashFunc, grp *SRPGroup, username, salt, A, B, K []byte) []byte {
	//M = H(H(N) xor H(g), H(I), s, A, B, K)

	// H(N) xor H(g)
	hn := new(big.Int).SetBytes(quickHash(hf, grp.Prime.Bytes()))
	hg := new(big.Int).SetBytes(quickHash(hf, grp.Generator.Bytes()))
	hng := hn.Xor(hn, hg)

	hi := quickHash(hf, []byte(username))

	h := hf()
	h.Write(hng.Bytes())
	h.Write(hi)
	h.Write(salt)
	h.Write(A)
	h.Write(B)
	h.Write(K)
	return h.Sum(nil)
}

func computeServerAuthenticator(hf HashFunc, A, M, K []byte) []byte {
	h := hf()
	h.Write(A)
	h.Write(M)
	h.Write(K)
	return h.Sum(nil)
}

// ComputeAuthenticator computes an authenticator that is to be passed to the
// server for validation
func (cs *ClientSession) ComputeAuthenticator() []byte {
	cs._M = computeClientAuthenticator(cs.SRP.HashFunc, cs.SRP.Group, cs.username, cs.salt, cs._A.Bytes(), cs._B.Bytes(), cs.key)
	return cs._M
}

// VerifyServerAuthenticator returns true if the authenticator returned by the
// server is valid
func (cs *ClientSession) VerifyServerAuthenticator(sauth []byte) bool {
	sa := computeServerAuthenticator(cs.SRP.HashFunc, cs._A.Bytes(), cs._M, cs.key)
	return subtle.ConstantTimeCompare(sa, sauth) == 1
}

// Return the bytes for the value of B.
func (ss *ServerSession) GetB() []byte {
	return ss._B.Bytes()
}

func (ss *ServerSession) setA(A []byte) error {
	ss._A = new(big.Int).SetBytes(A)
	if !ss.SRP.is_AB_valid(ss._A) {
		return fmt.Errorf("A%%N == 0")
	}
	ss._u = ss.SRP.compute_u(ss._A, ss._B)
	if ss._u.BitLen() == 0 {
		return fmt.Errorf("H(A, B) == 0")
	}
	return nil
}

// ComputeKey computes the session key given the value of A.
func (ss *ServerSession) ComputeKey(A []byte) ([]byte, error) {
	err := ss.setA(A)
	if err != nil {
		return nil, err
	}

	// S = (Av^u) mod N
	S := new(big.Int).Exp(ss._v, ss._u, ss.SRP.Group.Prime)
	S.Mul(ss._A, S).Mod(S, ss.SRP.Group.Prime)

	// Reject A*v^u == 0,1 (mod N)
	one := big.NewInt(1)
	if S.Cmp(one) <= 0 {
		return nil, fmt.Errorf("Av^u) mod N <= 0")
	}

	// Reject A*v^u == -1 (mod N)
	t1 := new(big.Int).Add(S, one)
	if t1.BitLen() == 0 {
		return nil, fmt.Errorf("Av^u) mod N == -1")
	}

	// S = (S ^ b) mod N              (computes session key)
	S.Exp(S, ss._b, ss.SRP.Group.Prime)
	// K = H(S)
	ss.key = quickHash(ss.SRP.HashFunc, S.Bytes())
	return ss.key, nil
}

// ComputeAuthenticator computes an authenticator to be passed to the client.
func (ss *ServerSession) ComputeAuthenticator(cauth []byte) []byte {
	return computeServerAuthenticator(ss.SRP.HashFunc, ss._A.Bytes(), cauth, ss.key)
}

// VerifyClientAuthenticator returns true if the client authenticator
// is valid.
func (ss *ServerSession) VerifyClientAuthenticator(cauth []byte) bool {
	M := computeClientAuthenticator(ss.SRP.HashFunc, ss.SRP.Group, ss.username, ss.salt, ss._A.Bytes(), ss._B.Bytes(), ss.key)
	return subtle.ConstantTimeCompare(M, cauth) == 1
}

func (s *SRP) pad(n *big.Int) []byte {
	nbytes := n.Bytes()
	if len(nbytes) < s.Group.Size/8 {
		bytes := make([]byte, s.Group.Size/8)
		copy(bytes[len(bytes)-len(nbytes):], nbytes)
		return bytes
	} else {
		return nbytes
	}
}

func (s *SRP) compute_u(A, B *big.Int) *big.Int {
	// u = H(A, B) where A and B are padded to the same size as N
	h := s.HashFunc()
	h.Write(s.pad(A))
	h.Write(s.pad(B))
	return new(big.Int).SetBytes(h.Sum(nil))
}

func (s *SRP) compute_k() {
	// H(N | PAD(g))
	h := s.HashFunc()
	h.Write(s.Group.Prime.Bytes())
	h.Write(s.pad(s.Group.Generator))
	s._k = new(big.Int).SetBytes(h.Sum(nil))
}

func (s *SRP) gen_rand_ab() *big.Int {
	max := new(big.Int).Lsh(big.NewInt(1), s.ABSize)
	r, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}
	return r
}

func (s *SRP) is_AB_valid(AB *big.Int) bool {
	ABmodN := new(big.Int).Mod(AB, s.Group.Prime)
	return ABmodN.BitLen() != 0
}
