package modlibb

import "github.com/stella-xx/modlib/lib"

func Read() {
	user := lib.NewUser("Eva")
	user.Read("golang")
}