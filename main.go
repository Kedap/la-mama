package main

import "fmt"

type mama struct {
	EDAD      int
	VIVE      bool
	TIENEMAMA bool
}

func (m *mama) TieneMama() bool {
	if m.EDAD <= 100 {
		m.TIENEMAMA = true
		return true
	} else {
		m.TIENEMAMA = false
		return false
	}
}

func main() {
	jefa := mama{
		60,
		true,
		true,
	}
	fmt.Print("La mama ")
	for jefa.TIENEMAMA {
		fmt.Print("de la mama ")
		jefa.TieneMama()
	}
}
