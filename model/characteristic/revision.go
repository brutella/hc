package characteristic

type Revision struct {
	*String
}

func NewHardwareRevision(revision string) *Revision {
	return newRevision(revision, CharTypeHardwareRevision)
}

func NewSoftwareRevision(revision string) *Revision {
	return newRevision(revision, CharTypeSoftwareRevision)
}

func NewFirmwareRevision(revision string) *Revision {
	return newRevision(revision, CharTypeFirmwareRevision)
}

func newRevision(revision string, charType CharType) *Revision {
	str := NewString(revision)
	str.Type = CharType(charType)
	str.Permissions = PermsRead()

	return &Revision{str}
}

func (r *Revision) SetRevision(revision string) {
	r.SetString(revision)
}

func (r *Revision) Revision() string {
	return r.StringValue()
}
