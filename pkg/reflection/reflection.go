package reflection

import "reflect"

// FieldIterationOptions is the options for the callback function of TraverseFields.
type FieldIterationOptions struct {
	Field     reflect.Value
	Tag       reflect.StructTag
	Index     int
	FieldName string
	Value     reflect.Value
	Type      reflect.Type
}

// IterateStructFields applies a callback function to each field of a given struct.
// The callback receives information about the current field via a FieldIterationOptions instance.
func IterateStructFields(s any, callback func(opts *FieldIterationOptions)) {
	v := reflect.ValueOf(s).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := v.Type().Field(i).Tag

		callback(&FieldIterationOptions{
			Field:     field,
			Tag:       tag,
			Index:     i,
			FieldName: v.Type().Field(i).Name,
			Value:     v,
			Type:      v.Type(),
		})
	}
}

// StructHasTagValue checks if a given struct has a field with a given tag and value.
func StructHasTagValue(s any, tag string, value string) bool {
	var hasTag bool
	IterateStructFields(s, func(opts *FieldIterationOptions) {
		if opts.Tag.Get(tag) == value {
			hasTag = true
		}
	})

	return hasTag
}

// IsPointerToInterface checks if a given struct is a pointer to an interface.
func IsPointerToInterface(s any) bool {
	return reflect.TypeOf(s).Kind() == reflect.Ptr &&
		reflect.TypeOf(s).Elem().Kind() == reflect.Interface
}

// IsPointer checks if a given struct is a pointer.
func IsPointer(v any) bool {
	return reflect.ValueOf(v).Kind() == reflect.Ptr
}

// IsNil checks if a given struct is nil.
func IsNil(v any) bool {
	return reflect.ValueOf(v).IsNil()
}

// IsZero checks if a given struct is zero.
func IsZero(v any) bool {
	return reflect.ValueOf(v).IsZero()
}
