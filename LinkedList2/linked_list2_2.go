package main

import(
	"errors"
	"sort"
)

/* META
#2 Двунаправленный связанный список
#9 Метод, который "переворачивает" порядок элементов в связном списке, меняя его на противоположный.

Сложность пространственная O(1)
Сложность временная: O(n)

Рефлексия к задаче #9:
	Необходим перебор всех значений списка. Просто идем в обратную сторону, Добавляя все значение с конца в начало.
*/

func Reverse(l LinkedList2) LinkedList2 {
	reversed := LinkedList2{}

	current := l.tail

	for current != nil {
		reversed.InsertFirst(*current)
		current = current.prev
	}

	return reversed
}

/* META
#2 Двунаправленный связанный список
#10 Булев метод, который сообщает, имеются ли циклы (замкнутые на себя по кругу) внутри списка.

Сложность пространственная O(1)
Сложность временная: O(n)

Рефлексия к задаче #10:
	В среднем временная сложность O(n), так как нужен перебор всех элементов списка,
	но также может быть и O(1), если будем находить цикл при первой же итерации.
*/

func IsCycled(l LinkedList2) bool {
	current := l.head
	var prev *Node = nil
	for current != nil {
		if prev != nil && current == prev {
			return true
		}

		prev = current
		current = current.next
	}

	return false
}

/* META
#2 Двунаправленный связанный список
#11 Метод, сортирующий список

Сложность пространственная O(n)
Сложность временная: O(n*log(n))

Рефлексия к задаче #10:
	Худший случай происходит в сортировке на 82 строке, и это является теоретически лучшим сценарием, 
	поэтому из возможных оптимизаций можно попробовать только уменьшить сложность по памяти, 
	переиспользуя существующие структуры данных.
*/

type SortBy []Node

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i].value > a[j].value }

func Sort(l LinkedList2) LinkedList2 {
	unsorted := []Node{}

	current := l.head
	// O(n)
	for current != nil {
		unsorted = append(unsorted, *current)
		current = current.next
	}

	if len(unsorted) == 0 {
		return l
	}

	// O(n*log(n))
	sort.Sort(SortBy(unsorted))

	sorted := LinkedList2{}

	// O(n)
	for i := range unsorted {
		sorted.InsertFirst(unsorted[i])
	}

	return sorted
}

/* META
#2 Двунаправленный связанный список
#12 Добавьте метод, объединяющий два списка в третий.

<<< WIP >>>
*/

/* META
#2 Двунаправленный связанный список
#13 Фиктивный/пустой (dummy) узел.

Рефлексия к задаче #13:
	За счет дополнительного булева поля структура NodeDummy утяжеляется, но при этом помогает избавиться от ограждающих проверок
	для корневых случаев.
*/

type NodeDummy struct {
	prev  *NodeDummy
	next  *NodeDummy
	value int
	dummy bool
}

type LinkedList2Dummy struct {
	head *NodeDummy
	tail *NodeDummy
}

func NewLinkedList2Dummy() LinkedList2Dummy {
	newList := LinkedList2Dummy{}
	newList.Clean()
	return newList
}

func (l *LinkedList2Dummy) AddInTailDummy(item NodeDummy) {
	tailPrev := l.tail.prev

	l.tail.prev = &item
	item.next = l.tail

	item.prev = tailPrev
	tailPrev.next = &item
}

func (l *LinkedList2Dummy) CountDummy() int {
	count := 0
	current := l.head

	for current != nil {
		if !current.dummy {
			count++
		}
		current = current.next
	}

	return count
}

func (l *LinkedList2Dummy) FindDummy(n int) (NodeDummy, error) {
	current := l.head
	for current != nil {
		if current.value == n && !current.dummy {
			return *current, nil
		}
		current = current.next
	}

	return NodeDummy{value: -1, next: nil}, errors.New("node was not found")
}

func (l *LinkedList2Dummy) FindAllDummy(n int) []NodeDummy {
	var nodes []NodeDummy

	current := l.head
	for current != nil {
		if current.value == n && !current.dummy {
			nodes = append(nodes, *current)
		}
		current = current.next
	}

	return nodes
}

func (l *LinkedList2Dummy) DeleteDummy(n int, all bool) {
	_, err := l.FindDummy(n)
	if err != nil {
		return
	}

	current := l.head
	for current != nil {
		if current.value != n || current.dummy {
			current = current.next
			continue
		}

		prev := current.prev
		next := current.next

		prev.next = next
		next.prev = prev

		current = current.next

		if !all {
			return
		}
	}
}

func (l *LinkedList2Dummy) InsertDummy(after *NodeDummy, add NodeDummy) {
	current := l.head
	for current != nil {
		if current != after || current.dummy {
			current = current.next
			continue
		}

		currentNext := current.next

		current.next = &add
		add.prev = current

		add.next = currentNext
		currentNext.prev = &add

		return
	}

	l.InsertFirstDummy(add)
}

func (l *LinkedList2Dummy) InsertFirstDummy(first NodeDummy) {
	headNext := l.head.next

	l.head.next = &first
	first.prev = l.head

	first.next = headNext
	headNext.prev = &first
}

func (l *LinkedList2Dummy) Clean() {
	l.head = &NodeDummy{dummy: true}
	l.tail = &NodeDummy{dummy: true}

	l.head.next = l.tail
	l.tail.prev = l.head
}
