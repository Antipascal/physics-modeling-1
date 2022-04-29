package main

import (
	"encoding/json"
	"errors"
	"math"

	"github.com/Knetic/govaluate"
)

//For API
type RequestBody struct {
	M_r float64 `json:"rocket_mass"`
	M_f float64 `json:"fuel_mass"`
	U   string  `json:"u_function"`
	F   string  `json:"f_function"`
}

type ResponseBody struct {
	V map[json.Number]json.Number `json:"v_plot"`
	S map[json.Number]json.Number `json:"s_plot"`
	A map[json.Number]json.Number `json:"a_plot"`
}

func (rb *ResponseBody) init() {
	rb.V = make(map[json.Number]json.Number)
	rb.S = make(map[json.Number]json.Number)
	rb.A = make(map[json.Number]json.Number)
}

// For calculations
type Func struct {
	cordinates map[float64]float64
	function   string
	name       string
}

// Creates function 'name' from 'function'
func (f *Func) init(name, function string) {
	f.name = name
	f.function = function
	f.cordinates = make(map[float64]float64)
}

// Count value of function in point t
// and stores it in map
func (f *Func) newVal(t float64) (float64, error) {
	expression, err := govaluate.NewEvaluableExpression(f.function)
	if err != nil {
		return 0, errors.New(err.Error())
	}

	parameters := make(map[string]interface{}, 8)
	parameters["t"] = t
	parameters["e"] = math.E

	result, err := expression.Evaluate(parameters)
	if err != nil {
		return 0, errors.New(err.Error())
	}

	f.cordinates[t] = result.(float64)
	return f.cordinates[t], nil
}
