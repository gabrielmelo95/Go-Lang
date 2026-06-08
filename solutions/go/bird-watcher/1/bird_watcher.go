package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	totalBirdCount := 0
	for i := 0; i < len(birdsPerDay); i++ {
		totalBirdCount += birdsPerDay[i]
	}
	return totalBirdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	weekBirds := birdsPerDay[7*(week-1) : 7*(week)]
	return TotalBirdCount(weekBirds)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay)/2; i++ {
		birdsPerDay[2*i] = birdsPerDay[2*i] + 1
	}
	return birdsPerDay
}
