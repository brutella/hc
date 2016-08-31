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

package srp

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"testing"
)

var groups []string = []string{
	"openssl.1024",
	"openssl.1536",
	"openssl.2048",
	"openssl.3072",
	"openssl.4096",
	"openssl.6144",
	"openssl.8192",
	"rfc5054.1024",
	"rfc5054.1536",
	"rfc5054.2048",
	"rfc5054.3072",
	"rfc5054.4096",
	"rfc5054.6144",
	"rfc5054.8192",
}

var passwords []string = []string{
	"0",
	"a",
	"password",
	"This Is A Long Password",
	"This is a really long password a;lsdfkjauiwjenfasueifxl3847tq8374y(*&^JHG&*^$.kjbh()&*^KJG",
}

type hashFunc func() hash.Hash

var hashes []hashFunc = []hashFunc{
	sha1.New,
	sha256.New,
	sha512.New,
}

func testSRP(t *testing.T, group string, h func() hash.Hash, username, password []byte) {
	srp, err := NewSRP(group, h, nil)
	if err != nil {
		t.Fatal(err)
	}
	cs := srp.NewClientSession(username, password)
	salt, v, err := srp.ComputeVerifier(password)
	if err != nil {
		t.Fatal(err)
	}
	ss := srp.NewServerSession(username, salt, v)

	ckey, err := cs.ComputeKey(salt, ss.GetB())
	if err != nil {
		t.Fatal(err)
	}

	skey, err := ss.ComputeKey(cs.GetA())
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(ckey, skey) {
		if cs._A.Cmp(ss._A) != 0 {
			t.Logf("A isn't the same for client and server")
		}
		if cs._B.Cmp(ss._B) != 0 {
			t.Logf("B isn't the same for client and server")
		}
		if cs._u.Cmp(ss._u) != 0 {
			t.Logf("u isn't the same for client and server")
		}
		t.Fatalf("Keys don't match(%s:%d):\n    Ckey: %v\n    Skey: %v\n",
			group, h().Size(), ckey, skey)
	}

	cauth := cs.ComputeAuthenticator()
	if !ss.VerifyClientAuthenticator(cauth) {
		t.Fatal("Client Authenticator is not valid")
	}

	sauth := ss.ComputeAuthenticator(cauth)
	if !cs.VerifyServerAuthenticator(sauth) {
		t.Fatal("Server Authenticator is not valid")
	}
}

func TestSRPSimple(t *testing.T) {
	for _, g := range groups {
		for _, h := range hashes {
			for _, p := range passwords {
				testSRP(t, g, h, []byte("test"), []byte(p))
			}
		}
	}
}
