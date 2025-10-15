package rank

type R uint8

const (
	Two R = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

func (r R) Compare(other R) int {
	switch {
	case r > other:
		return 1
	case r < other:
		return -1
	default:
		return 0
	}
}
