package abstact_data_type

type stack interface {
	Push(element int)
	Pop()
	Empty()
}