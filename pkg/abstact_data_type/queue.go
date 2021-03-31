package abstact_data_type

type PriorityQueue interface {
	Insert(element int) 
	IncreaseKey(key int, incr int) 
	ExtractMin() (min int)
}

type Queue interface {
	Enqueue(element int)
	Dequeue()
	Empty()
}