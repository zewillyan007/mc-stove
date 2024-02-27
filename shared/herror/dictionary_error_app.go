package herror

var DictionaryErrorApp *ErrorDictionary = &ErrorDictionary{
	Revision:    "1.0.0",
	DigitsClass: 2,
	DigitsError: 4,
	ErrorClass: map[string]*ErrorClass{
		"C01": {
			Type: "Plantation",
			List: map[string]*ErrorItem{
				"E0001": {Cod: "0001", Msg: "Length cannot be null"},
				"E0002": {Cod: "0002", Msg: "Width cannot be null"},
				"E0003": {Cod: "0003", Msg: "Height cannot be null"},
				"E0004": {Cod: "0004", Msg: "Number cannot be null"},
				"E0005": {Cod: "0005", Msg: "Species cannot be null"},
			},
		},
		"C02": {
			Type: "Lottery",
			List: map[string]*ErrorItem{},
		},
		"C03": {
			Type: "Network",
			List: map[string]*ErrorItem{},
		},
		"C04": {
			Type: "Report",
			List: map[string]*ErrorItem{},
		},
		"C05": {
			Type: "Dispatch",
			List: map[string]*ErrorItem{
				"E0001": {Cod: "0001", Msg: ""},
			},
		},
		"C06": {
			Type: "Database",
			List: map[string]*ErrorItem{
				"E0001": {Cod: "0001", Msg: "Instruction error"},
			},
		},
	},
}
