package gscript

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompiler_Compiler(t *testing.T) {
	script := `
if ( (10 +10 ) == 20 ) {
	return !(1+1!=2) 
} else {
	return 20 
}
`
	NewCompiler().Compiler(script)
}
func TestCompiler_Compiler2(t *testing.T) {
	script := `
int a=10
a++
return a
`
	compiler := NewCompiler().Compiler(script)
	fmt.Println(compiler)
	assert.Equal(t, compiler.(*LeftValue).GetValue().(int), 11)
}
