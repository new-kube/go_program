package weiconv

import "fmt"

// bang to kg

type Pound float64    // 英镑
type Kilogram float64 // 千克

func (p Pound) String() string    { return fmt.Sprintf("%gP", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%gKg", k) }
