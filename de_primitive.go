package serde

type DummyVisitor struct {
	expect string
}

func (vi DummyVisitor) String() string {
	if vi.expect == "" {
		return "dummy"
	}

	return vi.expect
}

func NewDummyVisitor(expect string) DummyVisitor {
	return DummyVisitor{expect: expect}
}

type StringVisitor struct {
	v *string
}

func NewStringVisitor(v *string) StringVisitor {
	return StringVisitor{v: v}
}

func (vi StringVisitor) String() string {
	return "string"
}

func (vi StringVisitor) VisitString(v string) (err error) {
	*vi.v = v
	return nil
}

func (vi StringVisitor) VisitBytes(v []byte) (err error) {
	*vi.v = string(v)
	return nil
}
