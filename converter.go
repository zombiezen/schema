// Copyright 2012 The Gorilla Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package schema

import (
	"errors"
	"reflect"
	"strconv"
)

type Converter func(string) reflect.Value

type ErrorConverter func(string) (reflect.Value, error)

var errInvalidValue = errors.New("")

func wrapConverter(c Converter) ErrorConverter {
	return func(value string) (reflect.Value, error) {
		v := c(value)
		if !v.IsValid() {
			return v, errInvalidValue
		}
		return v, nil
	}
}

var (
	invalidValue = reflect.Value{}
	boolType     = reflect.TypeOf(false)
	float32Type  = reflect.TypeOf(float32(0))
	float64Type  = reflect.TypeOf(float64(0))
	intType      = reflect.TypeOf(int(0))
	int8Type     = reflect.TypeOf(int8(0))
	int16Type    = reflect.TypeOf(int16(0))
	int32Type    = reflect.TypeOf(int32(0))
	int64Type    = reflect.TypeOf(int64(0))
	stringType   = reflect.TypeOf("")
	uintType     = reflect.TypeOf(uint(0))
	uint8Type    = reflect.TypeOf(uint8(0))
	uint16Type   = reflect.TypeOf(uint16(0))
	uint32Type   = reflect.TypeOf(uint32(0))
	uint64Type   = reflect.TypeOf(uint64(0))
)

// Default converters for basic types.
var converters = map[reflect.Type]ErrorConverter{
	boolType:    convertBool,
	float32Type: convertFloat32,
	float64Type: convertFloat64,
	intType:     convertInt,
	int8Type:    convertInt8,
	int16Type:   convertInt16,
	int32Type:   convertInt32,
	int64Type:   convertInt64,
	stringType:  convertString,
	uintType:    convertUint,
	uint8Type:   convertUint8,
	uint16Type:  convertUint16,
	uint32Type:  convertUint32,
	uint64Type:  convertUint64,
}

func convertBool(value string) (reflect.Value, error) {
	v, err := strconv.ParseBool(value)
	if err != nil {
		return invalidValue, err
	}
	return reflect.ValueOf(v), nil
}

func convertFloat32(value string) (reflect.Value, error) {
	v, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return invalidValue, err
	}
	return reflect.ValueOf(float32(v)), nil
}

func convertFloat64(value string) (reflect.Value, error) {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return invalidValue, err
	}
	return reflect.ValueOf(v), nil
}

func convertInt(value string) (reflect.Value, error) {
	v, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return invalidValue, err
	}
	return reflect.ValueOf(int(v)), nil
}

func convertInt8(value string) (reflect.Value, error) {
	v, err := strconv.ParseInt(value, 10, 8)
	if err != nil {
		return invalidValue, err
	}
	return reflect.ValueOf(int8(v)), nil
}

func convertInt16(value string) (reflect.Value, error) {
	v, err := strconv.ParseInt(value, 10, 16)
	if err != nil {
		return invalidValue, err
	}
	return reflect.ValueOf(int16(v)), nil
}

func convertInt32(value string) (reflect.Value, error) {
	v, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return invalidValue, err
	}
	return reflect.ValueOf(int32(v)), nil
}

func convertInt64(value string) (reflect.Value, error) {
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return invalidValue, err
	}
	return reflect.ValueOf(v), nil
}

func convertString(value string) (reflect.Value, error) {
	return reflect.ValueOf(value), nil
}

func convertUint(value string) (reflect.Value, error) {
	v, err := strconv.ParseUint(value, 10, 0)
	if err != nil {
		return invalidValue, err
	}
	return reflect.ValueOf(uint(v)), nil
}

func convertUint8(value string) (reflect.Value, error) {
	v, err := strconv.ParseUint(value, 10, 8)
	if err != nil {
		return invalidValue, err
	}
	return reflect.ValueOf(uint8(v)), nil
}

func convertUint16(value string) (reflect.Value, error) {
	v, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		return invalidValue, err
	}
	return reflect.ValueOf(uint16(v)), nil
}

func convertUint32(value string) (reflect.Value, error) {
	v, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return invalidValue, err
	}
	return reflect.ValueOf(uint32(v)), nil
}

func convertUint64(value string) (reflect.Value, error) {
	v, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return invalidValue, err
	}
	return reflect.ValueOf(v), nil
}
