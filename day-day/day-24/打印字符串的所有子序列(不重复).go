// @Author: Ciusyan 11/9/23

package day_24

// AllSubsquencesNoRepeat 要求不重复，收集
func AllSubsquencesNoRepeat(str string) []string {
	if str == "" {
		return nil
	}
	chars := []byte(str)
	track := make([]byte, 0, len(chars))
	set := make(map[string]struct{})
	var dfs func(level int, chars []byte, track []byte, res *map[string]struct{})
	dfs = func(level int, chars []byte, track []byte, res *map[string]struct{}) {
		if level == len(chars) {
			(*res)[string(track)] = struct{}{}
			return
		}

		dfs(level+1, chars, track, res)
		dfs(level+1, chars, append(track, chars[level]), res)
	}

	dfs(0, chars, track, &set)

	res := make([]string, 0, 1)
	for key := range set {
		res = append(res, key)
	}

	return res
}
