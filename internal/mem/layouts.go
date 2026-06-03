package mem

var Layouts = map[int]Layout{
	1: {
		{Start: 0, End: 1, Size: MaxReservedSize},
		{Start: 2, End: 3, Size: MaxBufferSize},
		{Start: 4, End: 19, Size: MaxReservedSize},
		{Start: 20, End: 39, Size: MaxGeneralSize},
	},
	5: {
		{Start: 0, End: 1, Size: MaxReservedSize},
		{Start: 2, End: 3, Size: MaxBufferSize},
		{Start: 4, End: 6, Size: MaxReservedSize},
		{Start: 7, End: 19, Size: MaxReservedSize},
		{Start: 20, End: 99, Size: MaxGeneralSize},
		{Start: 100, End: 112, Size: MaxReservedSize},
		{Start: 113, End: 949, Size: MaxGeneralSize},
		{Start: 950, End: 999, Size: MaxReservedSize},
	},
}

type Layout []Block

type Block struct {
	Start int
	End   int
	Size  int
}
