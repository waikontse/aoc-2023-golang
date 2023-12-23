package utils

type Position Pair[int, int]

func (p Position) GetLeft() Position {
	return Position{First: p.First - 1, Second: p.Second}
}

func (p Position) GetRight() Position {
	return Position{First: p.First + 1, Second: p.Second}
}

func (p Position) GetTop() Position {
	return Position{First: p.First, Second: p.Second - 1}
}

func (p Position) GetBottom() Position {
	return Position{First: p.First, Second: p.Second + 1}
}
