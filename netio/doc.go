// Package netio implements the HomeKit Accessory Protocol to pair and securily communicate
// with a HomeKit client.
// 
// The intial pairing is done by making sure that the client knows the secret (bridge password). 
// Then public keys are exchanged and used to negotiate a shared secret. The shared secret is used 
// to encrypt the communication. For every new connection a new shared secret is negotiated.
//
// When a client is already paired, it can provide public keys of other clients without
// going through the whole pairing process again.
package netio