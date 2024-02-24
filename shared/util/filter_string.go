package util

import "regexp"

type FilterString struct{}

func NewFilterString() *FilterString {
	return &FilterString{}
}

func (*FilterString) FilterStringRemoveCaracterSpecial(s string) string {

	tmp := s
	reg := regexp.MustCompile(`[^a-zA-Z0-9-]+`)
	filtered := reg.ReplaceAllString(tmp, "")

	return filtered
}

func (*FilterString) OnlyNumbers(s string) string {

	tmp := s
	reg := regexp.MustCompile(`[^0-9-]+`)
	filtered := reg.ReplaceAllString(tmp, "")

	return filtered
}
