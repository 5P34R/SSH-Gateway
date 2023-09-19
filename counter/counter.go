package counter


type UserCount struct {
	data map[string]int
}

var GlobalUserCount = NewUserCount()

func NewUserCount() *UserCount {
	return &UserCount{
		data: make(map[string]int),
	}
}

func (u *UserCount) Increment(container string) {
	if _, ok := u.data[container]; ok {
		u.data[container]++
	} else {
		u.data[container] = 1
	}
}

func (u *UserCount) Decrement(container string) {
	if val, ok := u.data[container]; ok && val > 0 {
		u.data[container]--
	}
}

func (u *UserCount) GetCount(container string) int {
	if val, ok := u.data[container]; ok {
		return val
	}
	return 0
}

func (u *UserCount) GetCounts() map[string]int {
	return u.data
}
