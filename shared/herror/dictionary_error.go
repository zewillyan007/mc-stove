package herror

var DictionaryError *ErrorDictionary = &ErrorDictionary{
	Revision:    "1.0.0",
	DigitsClass: 2,
	DigitsError: 4,
	ErrorClass: map[string]*ErrorClass{
		"C01": {
			Type: "Database",
			List: map[string]*ErrorItem{
				"E0001": {Cod: "0001", Msg: "Connection Error"},
				"E0002": {Cod: "0002", Msg: "Restrict Violation"},
				"E0003": {Cod: "0003", Msg: "Unique Violation"},
				"E0004": {Cod: "0004", Msg: "Exclusion Violation"},
				"E0005": {Cod: "0005", Msg: "Not Null  Violation"},
				"E0006": {Cod: "0006", Msg: "Integrity Constraint Violation"},
				"E0007": {Cod: "0007", Msg: "Scan error"},
				"E0008": {Cod: "0008", Msg: "Value limit not informed"},
			},
		},
		"C02": {
			Type: "Input",
			List: map[string]*ErrorItem{
				"E0001": {Cod: "0001", Msg: "Empty field"},
				"E0002": {Cod: "0002", Msg: "Invalid field"},
				"E0003": {Cod: "0003", Msg: "Required field"},
			},
		},
		"C03": {
			Type: "Service",
			List: map[string]*ErrorItem{
				"E0001": {Cod: "0001", Msg: "Aws-S3 Connection Error"},
				"E0002": {Cod: "0002", Msg: "CouchBase Connection Error"},
			},
		},
		"C04": {
			Type: "Language",
			List: map[string]*ErrorItem{
				"E0001": {Cod: "0001", Msg: "Null Point Error"},
				"E0002": {Cod: "0002", Msg: "Conversion Data Error"},
			},
		},
	},
}
