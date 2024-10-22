package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) (ans int) {
	for _, x := range birdsPerDay {
		ans += x
	}
	return
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	week -= 1
	// panic("Please implement the BirdsInWeek() function")
	ans := 0
	for i := week * 7; i < (week+1)*7; i++ {
		ans += birdsPerDay[i]
	}
	return ans
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	// panic("Please implement the FixBirdCountLog() function")
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i] += 1
	}
	return birdsPerDay
}
