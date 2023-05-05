package faker

import (
	"reflect"

	"github.com/christian-gama/nutrai-api/testutils/faker/providers"
	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
)

func InitializeProviders() {
	_ = faker.AddProvider("time_now", func(v reflect.Value) (any, error) {
		return providers.TimeNow(), nil
	})

	_ = faker.AddProvider("time_zero", func(v reflect.Value) (any, error) {
		return providers.TimeZero(), nil
	})

	_ = faker.AddProvider("uint", func(v reflect.Value) (any, error) {
		return providers.Uint(), nil
	})

	_ = faker.AddProvider("uint8", func(v reflect.Value) (any, error) {
		return providers.Uint8(), nil
	})

	_ = faker.AddProvider("uint16", func(v reflect.Value) (any, error) {
		return providers.Uint16(), nil
	})
}

func Setup() {
	options.SetRandomMapAndSliceMinSize(2)
	options.SetRandomMapAndSliceMaxSize(5)
}
