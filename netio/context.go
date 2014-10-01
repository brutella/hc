package netio

import(
    "github.com/brutella/hap"
    "io"
    "fmt"
)

// The server context provides access to the current session
// Every request should have it owns context
type Context struct {
    PairSetupSesssion *PairSetupServerSession
    PairVerifySesssion *PairVerifySession
    
    session Session
    // The value of this property is set as the value of `session` on
    // the next encryption
    nextSession Session
}

func NewContext() *Context {
    return &Context{session: NewPlainSession()}
}

func (c *Context) PublicKeyForAccessory(b *hap.Bridge) []byte {
    return b.PublicKey
}

func (c *Context) SecretKeyForAccessory(b *hap.Bridge) []byte {
    return b.SecretKey
}

func (s *Context) EncryptionEnabled() bool {
    if s.nextSession != nil {
        return s.nextSession.EncryptionEnabled()
    }
    
    return s.session.EncryptionEnabled()
}

func (c *Context) Encrypt(r io.Reader) (io.Reader, error) {
    return c.session.Encrypt(r)
}

func (c *Context) Decrypt(r io.Reader) (io.Reader, error) {
    if c.nextSession != nil {
        fmt.Println("Upgrading to new session")
        c.session = c.nextSession
        c.nextSession = nil
    }
    
    return c.session.Decrypt(r)
}

func (c *Context) OnSessionClosed() {
    c.SetNextSession(nil)
}

// Sets the session which should be used before the next decryption
// Until then the previous session is used to respond
//
// Discussion: On a session change, the previous session is used for
// a response. On the next request, we use the next session
func (c *Context) SetNextSession(session Session) {
    c.nextSession = session
}