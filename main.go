package uuidchecker

import "regexp"

var re = regexp.MustCompile(`^\w{8}-\w{4}-\w{4}-\w{4}-\w{12}$`)

func IsValid(uuid string) bool {
	return re.MatchString(uuid)
}
