package model

//Galinha representa uma ave
type Galinha interface {
	Cacareja() string
}

//Pato representa uma ave
type Pato interface {
	Quack() string
}

//Ave é um tipo de animal
type Ave struct {
	Nome string
}

//Cacareja é um som emitido pela galinha
func (a Ave) Cacareja() string {
	return "Cocorico...."
}

//Quack é um emitido pelo pato
func (a Ave) Quack() string {
	return "Quack quack...."
}
