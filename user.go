package modlibb

import "github.com/stella-xx/modlib/v3/lib"

func Read() {
	user := lib.NewUser("Eva")
	user.Read("golang")
}