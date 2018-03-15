package output

import "fmt"

const (
	// Byte multiplicator for byte
	Byte = 1 << (iota * 10)

	// KiByte multiplicator for kib
	KiByte

	// MiByte multiplicator for mib
	MiByte

	// GiByte multiplicator for gib
	GiByte

	// TiByte multiplicator for tib
	TiByte

	// PiByte multiplicator for pib
	PiByte

	// EiByte multiplicator for eib
	EiByte
)

const (
	// IByte multiplicator for bytes
	IByte = 1

	// KByte multiplicator for kilo byte
	KByte = IByte * 1000

	// MByte multiplicator for mega byte
	MByte = KByte * 1000

	// GByte multiplicator for giga byte
	GByte = MByte * 1000

	// TByte multiplicator for terra byte
	TByte = GByte * 1000

	// PByte multiplicator for peta byte
	PByte = TByte * 1000

	// EByte multiplicator for exa byte
	EByte = PByte * 1000
)

//
// AsBytes formats a given uint64 using a byte unit
//
func AsBytes(bytes uint64, decimalPlaces byte) string {
	tpl := fmt.Sprintf("%%.%df", decimalPlaces)

	switch {
	case bytes > EByte:
		return fmt.Sprintf(tpl, float64(bytes)/EByte) + " EB"
	case bytes > PByte:
		return fmt.Sprintf(tpl, float64(bytes)/PByte) + " PB"
	case bytes > TByte:
		return fmt.Sprintf(tpl, float64(bytes)/TByte) + " TB"
	case bytes > GByte:
		return fmt.Sprintf(tpl, float64(bytes)/GByte) + " GB"
	case bytes > MByte:
		return fmt.Sprintf(tpl, float64(bytes)/MByte) + " MB"
	case bytes > KByte:
		return fmt.Sprintf(tpl, float64(bytes)/KByte) + " kB"
	default:
		return fmt.Sprintf("%d", bytes) + " Byte"
	}
}
