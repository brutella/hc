// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeStreamingStatus = "120"

type StreamingStatus struct {
	*Bytes
}

func NewStreamingStatus() *StreamingStatus {
	char := NewBytes(TypeStreamingStatus)
	char.Format = FormatTLV8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue([]byte{})

	return &StreamingStatus{char}
}
