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

WIP

Рефлексия к задаче:
	WIP
*/

/* META
#7. Упорядоченный список
#10 Напишите метод проверки наличия заданного упорядоченного под-списка (параметр метода) в текущем списке.

WIP

Рефлексия к задаче:
	WIP
*/

/* META
#7. Упорядоченный список
#11 Добавьте метод, который находит наиболее часто встречающееся значение в списке.

WIP

Рефлексия к задаче:
	WIP
*/

/* META
#7. Упорядоченный список
#12 Добавьте в упорядоченный список возможность найти индекс элемента (параметр) в списке, которая должна работать за O(log N). Внутреннюю реализацию списка для этого можно сменить.

WIP

Рефлексия к задаче:
	WIP
*/
