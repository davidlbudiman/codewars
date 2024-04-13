// You're an elevator.

// People on each floor get on you in the order they are queued as long as you're stopped on their floor.

// Your doors are one-person wide. No one can board you when someone else is departing or vice versa.

// You must stop at each floor you pass that you can drop off and/or pick up passengers. Conversely, you don't stop if you're not changing passengers.

// You can't switch directions while holding passengers.

// People won't get on you if you're not going in the direction they want to go.

// During a stop, all passengers who can get off gets off BEFORE any new passengers could get on.

// When you're empty (at any point), you must go TOWARD the next person in the queue, taking anyone going in that direction along your path if you're capable.

// When you're empty AND you're on the same floor as the next person in the queue, you must now go in the direction they want to go.

// You must stop to open your doors, even if you haven't moved a floor. You begin each day with closed doors.

// Up to 5 people can be on you at a time.

// Given a starting_floor and a queue of people represented by from the floor they're on to the floor they want to go, return the order of stops you will take this day.

// Example
// Given this queue:

// queue := []Person{
//   {From: 3, To: 2}, // Al
//   {From: 5, To: 2}, // Betty
//   {From: 2, To: 1}, // Charles
//   {From: 2, To: 5}, // Dan
//   {From: 4, To: 3}, // Ed
// }
// then

// startingFloor := 1
// Order(startingFloor, queue)
// should return:

// []int{2, 5, 4, 3, 2, 1}
// Explanation
// You start on floor 1. Al is the first to queue for the elevator so you make your way toward him traveling UP.
// Going up a floor, you notice Charles and Dan on floor 2, however only Dan wants to go up so your first stop is 2 to pick up Dan.

// Going up to floor 3, you have now reached Al, but unfortunately for him, he wants to go down, and you can't switch directions until you drop off Dan. You promise to get Al later, so you pass him for now.

// With no one else going up on your path, your second stop is 5 to drop off Dan. Now that you're empty, you move toward Al again who is now below you. Since you begin traveling DOWN, Betty, who's also on floor 5, joins you.

// You stop on 4 to pick up Ed who is also going down.

// You stop on 3 to drop off Ed. Meanwhile, you finally pick up Al.

// You stop on 2 to drop off Al and Betty. You finally pick up Charles.

// You stop on 1 to drop off Charles, the last person to transport.

// Final Note
// The building you operate in may have ANY arbitrarily large number of stories (or basements) that people will want to travel between. For a given day, you should not expect more than 500 people to ride you.
package katas

type Person struct {
	From int
	To   int
}

func Order(level int, queue []Person) (doorOpenedAt []int) {
	lift := map[int][]*Person{}
	minFloor, maxFloor := minAndMaxFloors(queue)
	direction := 0
	if level < queue[0].From {
		direction = 1
	} else if level > queue[0].From {
		direction = -1
	} else {
		person := &queue[0]
		addPersonToLift(lift, person)
		queue = queue[1:]
		if level < person.To {
			direction = 1
		} else {
			direction = -1
		}
		doorOpenedAt = append(doorOpenedAt, person.From)
	}
	floor := level
	for {
		if len(queue) == 0 && len(lift) == 0 {
			break
		}
		if floor == maxFloor {
			direction = -1
		} else if floor == minFloor {
			direction = 1
		}
		var doesDoorOpenAt bool
		var new_queue = []Person{}
		for _, q := range queue {
			if q.From == floor {
				if (q.From < q.To && direction == 1) || (q.From > q.To && direction == -1) {
					lift = addPersonToLift(lift, &q)
					doesDoorOpenAt = true
				} else {
					new_queue = append(new_queue, q)
				}
			} else {
				new_queue = append(new_queue, q)
			}
		}
		queue = new_queue
		if _, ok := lift[floor]; ok {
			delete(lift, floor)
			doesDoorOpenAt = true
		}
		if doesDoorOpenAt {
			doorOpenedAt = append(doorOpenedAt, floor)
		}
		floor += direction
	}
	return doorOpenedAt
}

func addPersonToLift(lift map[int][]*Person, person *Person) map[int][]*Person {
	if a, ok := lift[person.To]; ok {
		a = append(a, person)
		lift[person.To] = a
	} else {
		a = []*Person{person}
		lift[person.To] = a
	}
	return lift
}

func minAndMaxFloors(queue []Person) (minFloor, maxFloor int) {
	minFloor, maxFloor = 100000, -100000
	for _, q := range queue {
		if minFloor > q.From {
			minFloor = q.From
		}
		if minFloor > q.To {
			minFloor = q.To
		}
		if maxFloor < q.From {
			maxFloor = q.From
		}
		if maxFloor < q.To {
			maxFloor = q.To
		}
	}
	return
}
