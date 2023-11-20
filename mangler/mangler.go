package mangler

// #include "cmangler.h"
import "C"

type Mangler struct {
	mangler C.Mangler
}

func NewMangler() *Mangler {
	return &Mangler{
		mangler: C.NewMangler(),
	}
}

func (m *Mangler) Close() {
	C.DestroyMangler(m.mangler)
}

func (m *Mangler) Encrypt(c byte) byte {
	return byte(C.Encrypt(m.mangler, C.uchar(c)))
}

func (m *Mangler) Decrypt(c byte) byte {
	return byte(C.Decrypt(m.mangler, C.uchar(c)))
}
