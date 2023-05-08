package convert

import (
	"github.com/jinzhu/copier"
)

// ToModel converts an object to a model. It panics if an error occurs.
func ToModel[Model any, Obj any](model Model, obj Obj) Model {
	if err := copier.CopyWithOption(model, obj, copier.Option{
		DeepCopy: true, Converters: Converters,
	}); err != nil {
		panic(err)
	}

	return model
}

// FromDomain converts a model to an object. It panics if an error occurs.
func FromModel[Model any, Obj any](obj Obj, model Model) Obj {
	if err := copier.CopyWithOption(obj, model, copier.Option{
		DeepCopy: true, Converters: Converters,
	}); err != nil {
		panic(err)
	}

	return obj
}

// ToModels converts an object slice to a model slice. It panics if an error occurs.
func ToModels[Model any, Obj any](model []Model, obj []Obj) []Model {
	if err := copier.CopyWithOption(&model, &obj, copier.Option{
		DeepCopy: true, Converters: Converters,
	}); err != nil {
		panic(err)
	}

	return model
}

// FromDomains converts a model slice to an object slice. It panics if an error occurs.
func FromModels[Model any, Obj any](obj []Obj, model []Model) []Obj {
	if err := copier.CopyWithOption(&obj, &model, copier.Option{
		DeepCopy: true, Converters: Converters,
	}); err != nil {
		panic(err)
	}

	return obj
}
