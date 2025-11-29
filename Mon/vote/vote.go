package vote

var Votes = map[string]int{
	"cat": 0,
	"dog": 0,
}

func VoteCore(animal string) bool {
	if animal != "cat" && animal != "dog" {
		return false
	}

	Votes[animal]++
	return true
}
