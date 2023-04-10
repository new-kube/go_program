package weiconv

const (
	OnePoundToKilogram = 0.4536
)

func PToK(p Pound) Kilogram { return Kilogram(p * OnePoundToKilogram) }

func KToP(k Kilogram) Pound { return Pound(k / OnePoundToKilogram) }
