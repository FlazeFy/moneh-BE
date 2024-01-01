package generator

import (
	"regexp"
	"strings"
)

func GetSlug(val string) string {
	res := strings.ReplaceAll(val, " ", "-")

	res = strings.ReplaceAll(res, "_", "-")

	regExp := regexp.MustCompile(`[!:\\\[/"\;\.\'^£$%&*()}{@#~?><>,|=+¬\]]`)
	res = regExp.ReplaceAllString(res, "")

	return res
}
