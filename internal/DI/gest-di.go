package DI

import (
	"fmt"
	"reflect"
)

var errType = reflect.TypeFor[error]()

type Container struct {
	constructors map[reflect.Type]reflect.Value
	instances    map[reflect.Type]reflect.Value
	resolving    map[reflect.Type]bool
}

func New() *Container {
	return &Container{
		constructors: map[reflect.Type]reflect.Value{},
		instances:    map[reflect.Type]reflect.Value{},
		resolving:    map[reflect.Type]bool{},
	}
}

func (c *Container) Provide(constructor any) error {
	value := reflect.ValueOf(constructor)
	typ := value.Type()
	if typ.Kind() != reflect.Func || typ.NumOut() == 0 || typ.NumOut() > 2 || (typ.NumOut() == 2 && typ.Out(1) != errType) {
		return fmt.Errorf("gest: invalid provider %s", typ)
	}

	out := typ.Out(0)
	if _, dup := c.constructors[out]; dup {
		return fmt.Errorf("gest: provider %s is already defined", out)
	}

	c.constructors[out] = value

	return nil
}

func (c *Container) Resolve(typ reflect.Type) (reflect.Value, error) {
	if value, ok := c.instances[typ]; ok {
		return value, nil
	}

	constructor, ok := c.constructors[typ]
	if !ok {
		return reflect.Value{}, fmt.Errorf("gest: provider %s is not defined", typ)
	}
	if c.resolving[typ] {
		return reflect.Value{}, fmt.Errorf("gest: dependency cycle at %s", typ)
	}

	c.resolving[typ] = true
	defer delete(c.resolving, typ)

	value, err := c.Invoke(constructor.Interface())
	if err != nil {
		return reflect.Value{}, err
	}

	c.instances[typ] = value
	return value, nil
}

func (c *Container) Invoke(fun any) (reflect.Value, error) {
	fn := reflect.ValueOf(fun)
	fnType := fn.Type()

	args := make([]reflect.Value, fnType.NumIn())
	for i := range args {
		arg, err := c.Resolve(fnType.In(i))
		if err != nil {
			return reflect.Value{}, err
		}

		args[i] = arg
	}

	out := fn.Call(args)
	if len(out) == 2 && !out[1].IsNil() {
		return reflect.Value{}, out[1].Interface().(error)
	}

	return out[0], nil
}
