package main

/* META
#7. Упорядоченный список
#6 Регулярная задача
Переделайте функцию поиска с учётом признака упорядоченности и возможности раннего прерывания поиска,
если найден заведомо больший или меньший элемент, нежели искомый.

Рефлексия к задаче:
	Сложность операции в худшем случае не изменилась, также O(n), но в лучшем случае получаем теперь O(1).
*/

func (l *OrderedList[T]) FindWithInterruption(n T) (Node[T], error) {
	if l.head == nil {
		return Node[T]{}, errors.New("not found")
	}

	current := l.head
	for current != nil {
		if current.value == n {
			return *current, nil
		}

		if l._ascending {
			if current.value > n {
				return Node[T]{}, errors.New("not found")
			}
		} else {
			if current.value < n {
				return Node[T]{}, errors.New("not found")
			}
		}

		current = current.next
	}

	return Node[T]{}, errors.New("not found")
}

/* META
#7. Упорядоченный список
#8 Дополнительная задача
Добавьте метод удаления всех дубликатов из упорядоченного списка.

Рефлексия к задаче:
	Сложность алгоритма O(n), необходимо полностью пройтись по списку.
*/

func (l *OrderedList[T]) RemoveDuplicates() {
	current := l.head
	for current != nil && current.next != nil {
		if current.value == current.next.value {
			dup := current.next
			current.next = dup.next
			if dup.next != nil {
				dup.next.prev = current
			} else {
				l.tail = current
			}
		} else {
			current = current.next
		}
	}
}

/* META
#7. Упорядоченный список
#9 Напишите алгоритм слияния двух упорядоченных списков в один, сохраняя порядок элементов. Подумайте, как это сделать наиболее эффективно.

Рефлексия к задаче:
	Сложность O(n), так как имеем два указателя и поочерёдно добавляем меньший элемент в хвост результирующего списка.
*/

func MergeLists[T constraints.Ordered](l1, l2 *OrderedList[T]) *OrderedList[T] {
	result := New[T](l1._ascending)

	current1 := l1.head
	current2 := l2.head

	for current1 != nil && current2 != nil {
		compared := l1.Compare(current1.value, current2.value)
		first := (l1._ascending && compared <= 0) || (!l1._ascending && compared >= 0)

		if first {
			appendTail(result, current1.value)
			current1 = current1.next
		} else {
			appendTail(result, current2.value)
			current2 = current2.next
		}
	}

	for current1 != nil {
		appendTail(result, current1.value)
		current1 = current1.next
	}

	for current2 != nil {
		appendTail(result, current2.value)
		current2 = current2.next
	}

	return result
}

func appendTail[T constraints.Ordered](result *OrderedList[T], val T) {
	node := &Node[T]{value: val, prev: result.tail}
	if result.tail != nil {
		result.tail.next = node
	} else {
		result.head = node
	}
	result.tail = node
}

/* META
#7. Упорядоченный список
#10 Напишите метод проверки наличия заданного упорядоченного под-списка (параметр метода) в текущем списке.

Рефлексия к задаче:
	O(n) — для каждого совпадения первого элемента подсписка проверяем все остальные.
	Раннее прерывание: если текущий элемент уже больше (меньше) первого элемента подсписка — останавливаемся.
*/

func (l *OrderedList[T]) ContainsSublist(sub *OrderedList[T]) bool {
	if sub.head == nil {
		return true
	}

	current := l.head
	for current != nil {
		cmp := l.Compare(current.value, sub.head.value)

		if (l._ascending && cmp > 0) || (!l._ascending && cmp < 0) {
			return false
		}

		if cmp == 0 {
			lCurrent := current
			subCurrent := sub.head
			for subCurrent != nil && lCurrent != nil && lCurrent.value == subCurrent.value {
				lCurrent = lCurrent.next
				subCurrent = subCurrent.next
			}
			if subCurrent == nil {
				return true
			}
		}

		current = current.next
	}

	return false
}

/* META
#7. Упорядоченный список
#11 Добавьте метод, который находит наиболее часто встречающееся значение в списке.

Рефлексия к задаче:
	O(n) — один проход. Дубликаты в упорядоченном списке всегда стоят рядом,
	поэтому достаточно сравнивать соседей.
*/

func (l *OrderedList[T]) MostFrequent() (T, error) {
	if l.head == nil {
		return *new(T), errors.New("empty list")
	}

	bestValue := l.head.value
	bestCount := 1
	currentCount := 1

	current := l.head.next
	for current != nil {
		if current.value == current.prev.value {
			currentCount++
			if currentCount > bestCount {
				bestCount = currentCount
				bestValue = current.value
			}
		} else {
			currentCount = 1
		}
		current = current.next
	}

	return bestValue, nil
}

/* META
#7. Упорядоченный список
#12 Добавьте в упорядоченный список возможность найти индекс элемента (параметр) в списке, которая должна работать за O(log N).

Рефлексия к задаче:
	WIP
*/