package hook

import "gorm.io/gorm"

type HookFunc func(db *gorm.DB)

func Register(db *gorm.DB, hookFuncs ...HookFunc) {
	for _, hook := range hookFuncs {
		hook(db)
	}
}
