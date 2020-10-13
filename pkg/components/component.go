package components

type Parameters struct {

}

type Component interface {
	apply() int
	getParameters() Parameters
}
