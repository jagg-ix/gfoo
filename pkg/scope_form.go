package gfoo

import (
	//"fmt"
)

type ScopeForm struct {
	FormBase
	body []Form
}

func NewScopeForm(pos Pos, body []Form) *ScopeForm {
	return new(ScopeForm).Init(pos, body)
}

func (self *ScopeForm) Init(pos Pos, body []Form) *ScopeForm {
	self.FormBase.Init(pos)
	self.body = body
	return self
}

func (self *ScopeForm) Compile(in *Forms, out []Op, scope *Scope) ([]Op, error) {
	ops, err := scope.vm.Compile(self.body, scope.Clone(true), nil)

	if err != nil {
		return out, err
	}
	
	return append(out, NewScopeOp(self, ops)), nil
}

func (self *ScopeForm) Quote(scope *Scope) (Val, error) {
	scope = scope.Clone(true)
	ops, err := scope.vm.Compile(self.body, scope, nil)

	if err != nil {
		return NilVal, err
	}

	return NewVal(&TLambda, NewLambda(self.body, ops, scope)), nil
}
