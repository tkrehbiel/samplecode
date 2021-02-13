// https://leetcode.com/problems/boats-to-save-people/

// Compute number of boats required to rescue a certain number of people.
func numRescueBoats(people []int, limit int) int {
    
    // Sort the weights of people in ascending order.
    // Assumes we can modify the input array.
    // (Otherwise a copy would need to be made.)
    sort.Ints(people)
    
    // All initialized to zero.
    var numBoats int
    var currentBoatPeople int
    var currentBoatWeight int

    // This might get into an infinite loop if every person is over the weight limit,
    // but the problem constraints indicate that will never happen.
    
    totalPeople := len(people)
    for totalPeople > 0 { // while totalPeople > 0

        // Get index of next available person.
        i := nextAvailablePerson(people, limit-currentBoatWeight)
        if i >= 0 {
            totalPeople--                   // reduce remaining people
            currentBoatPeople++             // add one to current boat
            currentBoatWeight += people[i]  // add weight to current boat
            people[i] = 0                   // mark person as used in the array
        }
        if currentBoatPeople >= 2 || currentBoatWeight >= limit || i < 0 {
            // Current boat over capacity, move to next one.
            numBoats++
            currentBoatPeople = 0
            currentBoatWeight = 0
        }
    }

    // If any leftover weight, add another boat.
    if currentBoatWeight > 0 {
        numBoats++
    }

    return numBoats
}

// Get next available person under the given weight.
func nextAvailablePerson(people []int, weight int) int {
    // Search people array from last to first
    // (highest weight to lowest weight)
    // so we get the next heaviest person available.
    for i := len(people)-1; i >= 0; i-- {
        if people[i] > 0 && people[i] <= weight {
            return i
        }
    }
    // If none available, return -1.
    return -1
}
