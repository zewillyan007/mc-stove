package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type ParamAward struct {
	Limit     *float64 `json:"limit"`
	Inherited string   `json:"inherited"`
}

type ParamCommission struct {
	Percentage *float64 `json:"percentage"`
	Inherited  string   `json:"inherited"`
}

type ParamSale struct {
	Limit     *float64 `json:"limit"`
	Inherited string   `json:"inherited"`
}

type ParamSpecialMaxAwardPayment struct {
	Limit     *float64 `json:"limit"`
	Inherited string   `json:"inherited"`
}

type ParamsType struct {
	Award                  ParamAward                  `json:"award"`
	Commission             ParamCommission             `json:"commission"`
	Sale                   ParamSale                   `json:"sale"`
	SpecialMaxAwardPayment ParamSpecialMaxAwardPayment `json:"special_max_award_payment"`
}

func NewParamType() ParamsType {

	return ParamsType{
		Award: ParamAward{
			Limit:     new(float64),
			Inherited: "",
		},
		Commission: ParamCommission{
			Percentage: new(float64),
			Inherited:  "",
		},
		Sale: ParamSale{
			Limit:     new(float64),
			Inherited: "",
		},
		SpecialMaxAwardPayment: ParamSpecialMaxAwardPayment{
			Limit:     new(float64),
			Inherited: "",
		},
	}
}

func (o ParamsType) Value() (driver.Value, error) {

	bytes, err := json.Marshal(o)
	return string(bytes), err
}

func (o *ParamsType) Scan(src interface{}) error {

	var source []byte

	mapt := make(map[string]interface{})

	switch src.(type) {
	case []uint8:
		source = []byte(src.([]uint8))
	case string:
		source = []byte(src.(string))
	case nil:
		return nil
	default:
		return errors.New("incompatible type for string ParamsType")
	}

	err := json.Unmarshal(source, &mapt)
	if err != nil {
		return err
	}

	params := NewParamType()

	if mapt["award"] != nil && mapt["award"].(map[string]interface{})["inherited"] != nil {
		params.Award.Inherited = mapt["award"].(map[string]interface{})["inherited"].(string)
	}

	if mapt["award"] != nil && mapt["award"].(map[string]interface{})["limit"] != nil {
		*params.Award.Limit = mapt["award"].(map[string]interface{})["limit"].(float64)
	}

	if mapt["commission"] != nil && mapt["commission"].(map[string]interface{})["inherited"] != nil {
		params.Commission.Inherited = mapt["commission"].(map[string]interface{})["inherited"].(string)
	}

	if mapt["commission"] != nil && mapt["commission"].(map[string]interface{})["percentage"] != nil {
		*params.Commission.Percentage = mapt["commission"].(map[string]interface{})["percentage"].(float64)
	}

	if mapt["sale"] != nil && mapt["sale"].(map[string]interface{})["inherited"] != nil {
		params.Sale.Inherited = mapt["sale"].(map[string]interface{})["inherited"].(string)
	}

	if mapt["sale"] != nil && mapt["sale"].(map[string]interface{})["limit"] != nil {
		*params.Sale.Limit = mapt["sale"].(map[string]interface{})["limit"].(float64)
	}

	if mapt["special_max_award_payment"] != nil && mapt["special_max_award_payment"].(map[string]interface{})["inherited"] != nil {
		params.SpecialMaxAwardPayment.Inherited = mapt["special_max_award_payment"].(map[string]interface{})["inherited"].(string)
	}

	if mapt["special_max_award_payment"] != nil && mapt["special_max_award_payment"].(map[string]interface{})["limit"] != nil {
		*params.SpecialMaxAwardPayment.Limit = mapt["special_max_award_payment"].(map[string]interface{})["limit"].(float64)
	}

	*o = params

	return nil
}
