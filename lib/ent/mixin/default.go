package mixin

import (
	"entgo.io/ent"
)

func Default() []ent.Mixin {
	return []ent.Mixin{
		DeletedTime{},
		Time{},
		ID{},
	}
}
