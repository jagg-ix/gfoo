package gfoo

import (
	"io"
	"strconv"
)

type Argument struct {
	id string
	index int
	valType Type
	val Val
}

func AIndex(id string, index int) Argument {
	return Argument{id: id, index: index}
}

func AType(id string, valType Type) Argument {
	return Argument{id: id, index: -1, valType: valType}
}

func AVal(id string, val Val) Argument {
	return Argument{id: id, index: -1, val: val}
}

func (self Argument) Dump(out io.Writer) error {
	if self.index == -1 {
		if self.valType == nil {
			if err := self.val.Dump(out); err != nil {
				return err
			}
		} else {
			if _, err := io.WriteString(out, self.valType.Name()); err != nil {
				return err
			}
		}
	} else {
		if _, err := io.WriteString(out, strconv.Itoa(self.index)); err != nil {
			return err
		}
	}

	return nil
}
