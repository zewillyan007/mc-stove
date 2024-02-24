package util

import (
	"strings"
	"unicode"
)

func StandardizeString(s string) string {
	str := "Esta é uma string com acentos e caracteres especiais: #$%¨&*()_+"
	str = RemoveSpecialChars(str)
	str = RemoveAcentos(str)

	return str
}

func RemoveSpecialChars(s string) string {
	specialChars := "çÇ!@#$%¨&*()_+={}[]/?;:.,|"
	r := strings.NewReplacer(specialChars, "")
	return r.Replace(s)
}

func RemoveAcentos(s string) string {
	t := unicode.SpecialCase{
		unicode.CaseRange{
			Lo: 0x00C0,
			Hi: 0x00FF,
			Delta: [unicode.MaxCase]rune{
				0x0000, // CaseUpper
				0x0000, // CaseLower
				0x0000, // CaseTitle
			},
		},
	}
	return strings.ToLowerSpecial(t, s)
}
