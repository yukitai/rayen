package transpiler

import (
	"rayen/log"

	"github.com/llir/llvm/ir"
)

type Transpiler struct {
	m *ir.Module
	np NamePool
}

func NewTranspiler(m *ir.Module) Transpiler {
	return Transpiler{
		m: m,
		np: NewNamePool(),
	}
}

func (t *Transpiler) Parse() Blocks {
	blocks := NewBlocks()
	for _, fun := range t.m.Funcs {
		t.ParseFunc(&blocks, fun)
	}
	return blocks
}

func (t *Transpiler) ParseFunc(blocks *Blocks, fun *ir.Func) {
	solver := NewLifetimeSolver()
	for _, block := range fun.Blocks {
		for _, inst := range block.Insts {
			t.ParseInst(&solver, blocks, &inst)
			solver.Next()
		}
	}
}

func (t *Transpiler) ParseInst(solver *LifetimeSolver, blocks *Blocks, inst *ir.Instruction) {
	switch (*inst).(type) {
	case *ir.InstAlloca:
		
	default:
		log.Errorf("use of unsupported instruction `%v`", (*inst).LLString())
	}
}