package hap

const (
	// MethodGET is the HTTP Get method
	MethodGET = "GET"

	// MethodPOST is the HTTP Post method
	MethodPOST = "POST"

	// MethodPUT is the HTTP Put method
	MethodPUT = "PUT"

	// MethodDEL is the HTTP Delete method
	MethodDEL = "DEL"
)

const (
	StatusSuccess                     = 0
	StatusInsufficientPrivileges      = -70401
	StatusServiceCommunicationFailure = -70402
	StatusResourceBusy                = -70403
	StatusReadOnlyCharacteristic      = -70404
	StatusWriteOnlyCharacteristic     = -70405
	StatusNotificationNotSupported    = -70406
	StatusOutOfResource               = -70407
	StatusOperationTimedOut           = -70408
	StatusResourceDoesNotExist        = -70409
	StatusInvalidValueInRequest       = -70410
)

const (
	// HTTPContentTypePairingTLV8 is the HTTP content type for pairing
	HTTPContentTypePairingTLV8 = "application/pairing+tlv8"

	// HTTPContentTypeHAPJson is the HTTP content type for json data
	HTTPContentTypeHAPJson = "application/hap+json"
)
