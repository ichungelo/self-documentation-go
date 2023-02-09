/*
This Package is used to do math operaion
*/
package selfdocumentationgo


type Calculate struct {
	Num1 float64
	Num2 float64
}

func (c Calculate) Addition() float64 {
	return c.Num1 + c.Num2
}

func (c Calculate) Subtraction() float64 {
	return c.Num1 - c.Num2
}

func (c Calculate) Multiplication() float64 {
	return c.Num1 * c.Num2
}

func (c Calculate) Division() float64 {
	return c.Num1 / c.Num2
}
