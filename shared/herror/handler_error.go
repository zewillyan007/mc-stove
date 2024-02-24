package herror

import (
	"fmt"
	"strconv"
	"strings"
)

type ErrorItem struct {
	Clc  string      `json:"clc"`
	Clt  string      `json:"clt"`
	Cod  string      `json:"cod"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ErrorClass struct {
	Type string                `json:"type"`
	List map[string]*ErrorItem `json:"list"`
}

type ErrorDictionary struct {
	Revision    string                 `json:"revision"`
	DigitsClass int                    `json:"digitsclass"`
	DigitsError int                    `json:"digitserror"`
	ErrorClass  map[string]*ErrorClass `json:"dictionary"`
}

type HandlerError struct {
	Dictionary    *ErrorDictionary
	DictionaryApp *ErrorDictionary
}

func NewHandlerError() *HandlerError {
	return &HandlerError{
		Dictionary:    DictionaryError,
		DictionaryApp: DictionaryErrorApp,
	}
}

func (o *HandlerError) SplitInt(n int) []int {

	slc := []int{}
	for n > 0 {
		slc = append(slc, n%10)
		n = n / 10
	}
	for i, j := 0, len(slc)-1; i < j; i, j = i+1, j-1 {
		slc[i], slc[j] = slc[j], slc[i]
	}
	return slc
}

func (o *HandlerError) SliceToInt(ints []int) int {

	stringVals := make([]string, len(ints))
	for ind, val := range ints {
		stringVals[ind] = strconv.Itoa(val)
	}
	newInt, _ := strconv.Atoi(strings.Join(stringVals, ""))
	return newInt
}

func (o *HandlerError) Err(code int, dictionary *ErrorDictionary) *ErrorItem {

	if code <= 0 {
		return nil
	}

	clenth := dictionary.DigitsClass
	dlenth := dictionary.DigitsError
	digits := o.SplitInt(code)
	formatc := fmt.Sprintf("%v", clenth)
	formate := fmt.Sprintf("%v", dlenth)
	ccode := fmt.Sprintf("%0"+formatc+"d", o.SliceToInt(digits[:len(digits)-dlenth]))
	ecode := fmt.Sprintf("%0"+formate+"d", o.SliceToInt(digits[len(digits)-dlenth:]))

	class := dictionary.ErrorClass["C"+ccode]
	if class == nil {
		return nil
	}
	eitem := class.List["E"+ecode]
	if eitem == nil {
		return nil
	}

	eitem.Clc = ccode
	eitem.Clt = class.Type
	eitem.Cod = "E" + ccode + ecode
	return eitem
}

func (o *HandlerError) Error(code int) *ErrorItem {

	return o.Err(code, o.Dictionary)
}

func (o *HandlerError) ErrorApp(code int) *ErrorItem {

	return o.Err(code, o.DictionaryApp)
}
