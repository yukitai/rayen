package main

import (
	"rayen/log"
	"rayen/transpiler"

	"github.com/kr/pretty"
	"github.com/llir/llvm/asm"
)

func main() {
	m, err := asm.ParseFile("./c/a.ll")
	if err != nil {
		log.Errorf("%+v", err)
	}
	
	t := transpiler.NewTranspiler(m)

	pretty.Println(t.Parse())
}