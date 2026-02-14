package main

/* META
	#1 Связный (связанный) список
	#8 Доп. задача
	Комбинированный список из двух связанных списков одинаковой длины
	
	Сложность пространственная O(n)
	Сложность временная: O(n)

	Рефлексия к задаче #8: 
		В данном решение нам так или иначе приходится линейно проходиться
		по всему списку, так как нужно взять каждое значение узла и прибавить
		значение из узла одного списка к другому. 

		Моя реализация безопасна для продакшн-работы, так как мы не изменяем существующие списки,
		однако теоретически можно достичь пространственной сложности O(1), изменяя один из входных списков,
		не создавая при этом новый. 
*/

func CombineLists(l1, l2 LinkedList) LinkedList {
	if l1.Count() != l2.Count() {
		return LinkedList{}
	}

	result := LinkedList{}

	current := l1.head
	current2 := l2.head

	for current != nil {
		result.AddInTail(Node{
			value: current.value + current2.value,
			next:  nil,
		})

		current = current.next
		current2 = current2.next
	}

	return result
}
