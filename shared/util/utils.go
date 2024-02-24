package util

import (
	"crypto/md5"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Utils struct{}

func NewUtils() *Utils {
	return &Utils{}
}

func (*Utils) CreateHashMd5(key string) string {

	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (*Utils) FormatCurrencyCommas(n float64) string {

	//n = RoundFloat(n, 2)
	// l := log.New()
	nInt := math.Trunc(n)

	//in := strconv.FormatFloat(n, 'f', -1, 64)
	in := strconv.FormatInt(int64(nInt), 10)

	// l.Info("in %v", in)

	//in := fmt.Sprintf("%.2f", n)
	numOfDigits := len(in)
	if n < 0 {
		numOfDigits-- // First character is the - sign (not a digit)
	}
	numOfCommas := (numOfDigits - 1) / 3

	out := make([]byte, len(in)+numOfCommas)
	if n < 0 {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			//return string(out)
			break
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = ','
		}
	}

	// decimal part of n
	decPart := fmt.Sprintf("%.2f", RoundFloat(n-nInt, 2))[2:]

	decOut := fmt.Sprintf("%s.%s", out, decPart)

	// l.Info("out %v", decOut)

	return decOut
}

// RoundFloat rounds a float64 to a certain quantity of
// decimal places. E.g., 3.4371 rounded to 2 decimal places
// yields 3.44.
func RoundFloat(n float64, qtyDecimalPlaces int) float64 {
	tmp := math.Pow(10, float64(qtyDecimalPlaces))
	return math.Round(n*tmp) / tmp
}

func (*Utils) IsCSVFormat(filePath, delimiter string, numberOfColumns int) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("30124")
	}
	defer file.Close()

	// Criar um leitor de arquivos CSV.
	reader := csv.NewReader(file)
	for {
		// Ler a próxima linha do arquivo.
		record, err := reader.Read()
		if err == io.EOF {
			// O fim do arquivo foi alcançado.
			break
		}

		// Verificar se as colunas estão no formato correto.
		for _, column := range record {
			if !strings.Contains(column, delimiter) {
				return fmt.Errorf("30123")
			}
		}

		// Verificar se linha possui 3 colunas
		var countRecord int
		for _, s := range record {
			countRecord = strings.Count(s, delimiter)
		}

		if countRecord != numberOfColumns {
			return fmt.Errorf("30122")
		}

	}

	return nil
}

func (o *Utils) GetFileContentType(ouput *os.File) (string, error) {

	buf := make([]byte, 512)
	_, err := ouput.Read(buf)

	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buf)

	return contentType, nil
}
