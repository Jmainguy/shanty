package main

func circleOfFifths(note string, notes map[string]int) (newNote string) {

	newNoteNumber := notes[note] + 7
	if newNoteNumber > 11 {
		newNoteNumber = newNoteNumber - 12
	}

	for k, v := range notes {
		if v == newNoteNumber {
			newNote = k
		}
	}

	return
}

func scaleToNotes(root string, notes map[string]int, method string) (scale []string) {

	var newNote string
	var intervals []int
	if method == "major" {
		intervals = majorScale()
	} else if method == "naturalMinor" {
		intervals = naturalMinor()
	} else if method == "harmonicMinor" {
		intervals = harmonicMinor()
	} else if method == "melodicMinor" {
		intervals = melodicMinor()
	}
	for _, v := range intervals {
		//fmt.Println(v)
		newNoteNumber := notes[root] + v
		if newNoteNumber > 11 {
			newNoteNumber = newNoteNumber - 12
		}

		for k, v := range notes {
			if v == newNoteNumber {
				newNote = k
				root = k
			}
		}

		scale = append(scale, newNote)

	}

	return
}

func majorScale() (intervals []int) {

	// Major scale intervals are 2 + 2 + 1 + 2 + 2 + 2 + 1
	intervals = append(intervals, 2)
	intervals = append(intervals, 2)
	intervals = append(intervals, 1)
	intervals = append(intervals, 2)
	intervals = append(intervals, 2)
	intervals = append(intervals, 2)
	intervals = append(intervals, 1)

	return
}

func naturalMinor() (intervals []int) {

	// Natural Minor scale intervals are W, H, W, W, H, W, W
	intervals = append(intervals, 2)
	intervals = append(intervals, 1)
	intervals = append(intervals, 2)
	intervals = append(intervals, 2)
	intervals = append(intervals, 1)
	intervals = append(intervals, 2)
	intervals = append(intervals, 2)

	return
}

func harmonicMinor() (intervals []int) {

	// Harmonic Minor scale intervals are W, H, W, W, H, W+H, H
	intervals = append(intervals, 2)
	intervals = append(intervals, 1)
	intervals = append(intervals, 2)
	intervals = append(intervals, 2)
	intervals = append(intervals, 1)
	intervals = append(intervals, 3)
	intervals = append(intervals, 1)

	return
}

func melodicMinor() (intervals []int) {

	// Melodic Minor scale intervals are W, H, W, W, W, W, H
	intervals = append(intervals, 2)
	intervals = append(intervals, 1)
	intervals = append(intervals, 2)
	intervals = append(intervals, 2)
	intervals = append(intervals, 2)
	intervals = append(intervals, 2)
	intervals = append(intervals, 1)

	return
}
