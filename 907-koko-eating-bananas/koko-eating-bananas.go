// Returns the total hours needed if Koko eats
// at a speed of 'guess' bananas per hour.
func fun(piles []int, guess int) int {
    h := 0

    // Calculate hours required for each pile
    for _, v := range piles {

        // Add complete hours
        h += v / guess

        // If there are remaining bananas,
        // one extra hour is needed
        if v%guess != 0 {
            h++
        }
    }

    return h
}

func minEatingSpeed(piles []int, h int) int {

    // Minimum possible eating speed is 1 banana/hour
    low := 1

    // Maximum possible eating speed is the largest pile
    // (finish the largest pile in one hour)
    high := slices.Max(piles)

    // Stores the minimum valid eating speed found so far
    res := -1

    // Binary search on eating speed
    for low <= high {

        // Candidate eating speed
        guess := (low + high) / 2

        // Hours needed at this speed
        minhours := fun(piles, guess)

        // Too slow: need a higher eating speed
        if minhours > h {
            low = guess + 1
        } else {

            // Valid speed found
            res = guess

            // Try to find a smaller valid speed
            high = guess - 1
        }
    }

    return res
}