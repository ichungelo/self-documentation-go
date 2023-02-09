/*
This Package is used to do math operaion

# Test1

this is a test 1

# Test2

this is a test 2

# Test3

this is a test 3

# Test4

this is a test 4

# Test5

this is a test 5

*/
package calculate

/*
# Calculate Struct

this struct is used for add float number data which want to do operation
*/
type Calculate struct {
	Num1 float64
	Num2 float64
}

/*
# Addition Method

Used for do addition in calculate
*/
func (c Calculate) Addition() float64 {
	return c.Num1 + c.Num2
}

/*
# Subtraction Method

Used for do subtraction in calculate
*/
func (c Calculate) Subtraction() float64 {
	return c.Num1 - c.Num2
}

/*
# Multiplication Method

Used for do multiplication in calculate
*/
func (c Calculate) Multiplication() float64 {
	return c.Num1 * c.Num2
}

/*
# Division Method

Used for do division in calculate
*/
func (c Calculate) Division() float64 {
	return c.Num1 / c.Num2
}
