// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	RemoteKeyRewind      int = 0
	RemoteKeyFastForward int = 1
	RemoteKeyExit        int = 10
	RemoteKeyPlayPause   int = 11
	RemoteKeyInfo        int = 15
	RemoteKeyNextTrack   int = 2
	RemoteKeyPrevTrack   int = 3
	RemoteKeyArrowUp     int = 4
	RemoteKeyArrowDown   int = 5
	RemoteKeyArrowLeft   int = 6
	RemoteKeyArrowRight  int = 7
	RemoteKeySelect      int = 8
	RemoteKeyBack        int = 9
)

const TypeRemoteKey = "E1"

type RemoteKey struct {
	*Int
}

func NewRemoteKey() *RemoteKey {
	char := NewInt(TypeRemoteKey)
	char.Format = FormatUInt8
	char.Perms = []string{PermWrite}
	char.SetMinValue(0)
	char.SetMaxValue(16)
	char.SetStepValue(1)

	return &RemoteKey{char}
}
