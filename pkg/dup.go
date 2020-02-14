package gfoo

type Dup struct {
	OpBase
}

func NewDup(form Form) *Dup {
	o := new(Dup)
	o.OpBase.Init(form)
	return o
}

func (self *Dup) Evaluate(stack *Slice, scope *Scope) error {
	v := stack.Peek()
	
	if v == nil {
		return scope.vm.Error(self.form.Pos(), "Nothing to dup")
	}

	stack.Push(*v)
	return nil
}
