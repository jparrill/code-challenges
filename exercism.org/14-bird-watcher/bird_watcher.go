package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var total int
    for _, v := range birdsPerDay {
        total += v
    }
    
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    startingDayWeek := (week - 1) * 7
    givenWeek := birdsPerDay[startingDayWeek:startingDayWeek+7]
    var total int

    for _,v := range givenWeek {
        total += v
    }

    return total
	
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {

    for i,v := range birdsPerDay {
        if i == 0 {
            birdsPerDay[i] = v + 1
        } else if i%2 == 0 {
        	birdsPerDay[i] = v + 1
        } else {
        	birdsPerDay[i] = v
        }
    }

    return birdsPerDay
}
