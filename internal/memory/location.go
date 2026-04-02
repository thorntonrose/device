package memory

type Location []byte

// NewLocations creates a new set of memory locations.
//
// Location Number	Size	Description
// ---------------	----	-----------
// 000 - 001			-		reserved
// 002					250	buffer 1 (transmit)
// 003					250	buffer 2 (receive)
// 004 - 019			-		reserved
// 020 - 039			120	general purpose
func NewLocations() []Location {
	locations := make([]Location, MaxLocations)
	locations = AddLocations(locations, 0, 1, MaxReservedSize)
	locations = AddLocations(locations, 2, 3, MaxBufferSize)
	locations = AddLocations(locations, 4, 19, MaxReservedSize)
	locations = AddLocations(locations, 20, 39, MaxGeneralSize)
	// ???: Set reserved locations to special value?

	return locations
}

func NewLocation(size int) Location {
	location := make(Location, 0, size)
	return location
}

func AddLocations(locations []Location, start int, end int, size int) []Location {
	for i := start; i <= end; i++ {
		locations[i] = NewLocation(size)
	}

	return locations
}

//-----------------------------------------------------------------------------

func (l *Location) Append(data byte) {
	*l = (*l)[:len(*l)+1]
	(*l)[len(*l)-1] = data
}

func (l *Location) Set(data []byte) {
	*l = (*l)[:len(data)]
	copy(*l, data)
}

func (l Location) String() string {
	return string(l)
}
