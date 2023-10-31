package protocol100r5

type ScalerPar interface {
	GetScalePar() (string, error)
}

type Massa interface {
	GetMassa() (string, error)
}
type Tare interface {
	SetTare() error
	SetTareValue(val string) error
}
type Zero interface {
	SetZero() error
}

type Commander interface {
	ScalerPar
	Massa
	Tare
	Zero
}
