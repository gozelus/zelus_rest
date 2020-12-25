package rest

import (
	"fmt"
	"testing"
)

func TestHe(t *testing.T) {
	c := Context{}
	c.handlers = []func(c *Context){
		a, b, c2, d,
	}
	c.index = -1
	c.Next()
}
func a(c *Context) {
	fmt.Println("here is a")
}
func b(c *Context) {
	c.Abort()
	c.Next()
	fmt.Println("here is b")
}
func c2(c *Context) {
	fmt.Println("here is c")
}
func d(c *Context) {
	fmt.Println("here is d")
}
