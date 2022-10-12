// inch to m cast
package lenconv

import "fmt"

type Foot float64 // 0.3048m
type Metre float64

const (
	OneFootToMetre = 0.3048
)

func (f Foot) String() string  { return fmt.Sprintf("%gF", f) }
func (m Metre) String() string { return fmt.Sprintf("%gM", m) }

func FToM(f Foot) Metre { return Metre(f * OneFootToMetre) }
func MToF(m Metre) Foot { return Foot(m / OneFootToMetre) }
