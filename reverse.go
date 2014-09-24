package hap

type ByteSequence []byte

func (s ByteSequence) Len() int {
    return len(s)
}

func (s ByteSequence) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s ByteSequence) Swap(i, j int) {
    element := s[i]
    s[i] = s[j]
    s[j] = element
}