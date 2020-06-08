package stackitem

// Type represents type of the stack item.
type Type byte

// This block defines all known stack item types.
const (
	AnyT       Type = 0x00
	PointerT   Type = 0x10
	BooleanT   Type = 0x20
	IntegerT   Type = 0x21
	ByteArrayT Type = 0x28
	BufferT    Type = 0x30
	ArrayT     Type = 0x40
	StructT    Type = 0x41
	MapT       Type = 0x48
	InteropT   Type = 0x60
)

// String implements fmt.Stringer interface.
func (t Type) String() string {
	switch t {
	case AnyT:
		return "Any"
	case PointerT:
		return "Pointer"
	case BooleanT:
		return "Boolean"
	case IntegerT:
		return "Integer"
	case ByteArrayT:
		return "ByteArray"
	case BufferT:
		return "Buffer"
	case ArrayT:
		return "Array"
	case StructT:
		return "Struct"
	case MapT:
		return "Map"
	case InteropT:
		return "Interop"
	default:
		return "INVALID"
	}
}

// IsValid checks if s is a well defined stack item type.
func (t Type) IsValid() bool {
	switch t {
	case AnyT, PointerT, BooleanT, IntegerT, ByteArrayT, BufferT, ArrayT, StructT, MapT, InteropT:
		return true
	default:
		return false
	}
}
