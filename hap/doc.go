// Package hap implements the HomeKit Accessory Protocol to pair and securily communicate
// with a HomeKit client. The server send data over tcp to clients (iOS devices) based on
// a http-like protocol.
//
// Pairing: A client has to know the secret pin to pair with the server. After that
// the encryption keys are negotiated and stored on the server and client side.
//
// Session: Before data is exchanged a new pair of keys is generated. The previously negoiated
// keys are verified before the new keys are generated.
package hap
