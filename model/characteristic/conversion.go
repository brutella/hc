package characteristic

import(
    "reflect"
)

func ConvertValue(unk interface{}, toType reflect.Type) interface{} {
    v := reflect.ValueOf(unk)
    v = reflect.Indirect(v)
    if !v.Type().ConvertibleTo(toType) {
        return nil
    }
    fv := v.Convert(toType)
    return fv.Interface()
}