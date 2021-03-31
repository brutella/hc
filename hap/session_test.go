package hap

import (
	"sync"
	"testing"
)

func TestSession_DataRace(t *testing.T) {
	session := NewSession(testConn)

	cryptographer := &fakeCryptographer{}
	pairSetupHandler := &fakeContainerHandler{}
	pairVerifyHandler := &fakePairVerifyHandler{}

	wg := sync.WaitGroup{}
	wg.Add(7)

	go func() {
		defer wg.Done()
		session.SetCryptographer(cryptographer)
	}()

	go func() {
		defer wg.Done()
		if is, want := session.Decrypter(), cryptographer; is != want {
			t.Fatalf("is = %v, want = %v", is, want)
		}
	}()

	go func() {
		defer wg.Done()
		session.Encrypter()
	}()

	go func() {
		defer wg.Done()
		session.SetPairSetupHandler(pairSetupHandler)
	}()

	go func() {
		defer wg.Done()
		if is, want := session.PairSetupHandler(), pairSetupHandler; is != want {
			t.Fatalf("is = %v, want = %v", is, want)
		}
	}()

	go func() {
		defer wg.Done()
		session.SetPairVerifyHandler(pairVerifyHandler)
	}()

	go func() {
		defer wg.Done()
		if is, want := session.PairVerifyHandler(), pairVerifyHandler; is != want {
			t.Fatalf("is = %v, want = %v", is, want)
		}
	}()

	wg.Wait()
}
