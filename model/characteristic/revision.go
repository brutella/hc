package characteristic

type Revision struct {
	*String
}

func NewHardwareRevision(revision string) *Revision {
	return newRevision(revision, TypeHardwareRevision)
}

func NewSoftwareRevision(revision string) *Revision {
	return newRevision(revision, TypeSoftwareRevision)
}

func NewFirmwareRevision(revision string) *Revision {
	return newRevision(revision, TypeFirmwareRevision)
}

func newRevision(revision string, characteristicType CharacteristicType) *Revision {
	str := NewString(revision)
	str.Type = CharacteristicType(characteristicType)
	str.Permissions = PermsRead()

	return &Revision{str}
}

func (r *Revision) SetRevision(revision string) {
	r.SetString(revision)
}

func (r *Revision) Revision() string {
	return r.StringValue()
}
