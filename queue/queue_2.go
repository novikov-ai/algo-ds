package main

import (
	"fmt"
)

/* META
#5. Очереди
#2 Регулярная задача
Оцените меру сложности для операций enqueue() (добавление) и dequeue() (удаление) в данной реализации.

Enqueue — Мера сложности O(1)

Dequeue() — Мера сложности O(n)

Рефлексия к задаче:
	В случае с Dequeue приходится заново копировать массив, чтобы избежать утечек памяти. При такой реализации
	O(n) — лучшее, чего можно добиться.
*/

/* META
#5. Очереди
#3 Дополнительная задача
Напишите функцию, которая "вращает" очередь по кругу на N элементов.

Рефлексия к задаче:
	Пространственная сложность O(n^2) при стандартной реализации Dequeue(),
	где каждый раз мы аллоцируем новый массив с копированием.
*/

func (q *Queue[T]) Rotate(n int) {
	if q.Size() == 0 {
		return
	}

	for range q.elements {
		if n == 0 {
			return
		}

		item, err := q.Dequeue()
		if err != nil {
			continue
		}

		q.Enqueue(item)
		n--
	}
}

/* META
#5. Очереди
#4 Дополнительная задача
Реализуйте очередь с помощью двух стеков.

Рефлексия к задаче:
	Каждый из стеков выполняет свою функцию:
		а) stackIn работает на прием элементов
		б) stackOut работает на извлечение элементов
	Таким образом, для каждой из операций получаем O(1) кроме случаев, 
	когда stackOut оказывается пустым, но это редкость, так как "размазывается" по всем операциям
*/

type QueueStacks[T any] struct {
	stackIn  Stack[T]
	stackOut Stack[T]
}

func (q *QueueStacks[T]) Size() int {
	return q.stackIn.Size() + q.stackOut.Size()
}

func (q *QueueStacks[T]) Dequeue() (T, error) {
	if q.stackOut.Size() == 0 {
		for q.stackIn.Size() > 0 {
			v, _ := q.stackIn.Pop()
			q.stackOut.Push(v)
		}
	}

	return q.stackOut.Pop()
}

func (q *QueueStacks[T]) Enqueue(itm T) {
	q.stackIn.Push(itm)
}

type Stack[T any] struct {
	values []T
	count  int
}

func (st *Stack[T]) Size() int {
	return st.count
}

func (st *Stack[T]) Peek() (T, error) {
	var result T

	if st.count == 0 {
		return result, fmt.Errorf("stack is empty")
	}

	result = st.values[st.count-1]

	return result, nil
}

// Мера сложности: O(1)
func (st *Stack[T]) Pop() (T, error) {
	var result T

	if st.count == 0 {
		return result, fmt.Errorf("stack is empty")
	}

	result = st.values[st.count-1]

	st.values = st.values[:st.count-1]
	st.count -= 1

	return result, nil
}

// Мера сложности: O(1) амортизированно из-за возможных аллокаций на расширение массива
func (st *Stack[T]) Push(itm T) {
	st.values = append(st.values, itm)
	st.count += 1
}

/* META
#5. Очереди
#5 Дополнительная задача
Добавьте функцию, которая обращает все элементы в очереди в обратном порядке.

Рефлексия к задаче:
	Временная сложность O(n), пространственная сложность O(n) из-за создания нового массива.
	Можно добиться O(1), если сделать два указателя идущих с разных концов друг на встречу другу, меняя элементы местами.
*/

func (q *Queue[T]) Reverse() {
	if q.Size() == 0 {
		return
	}

	reversed := make([]T, 0, len(q.elements))

	for i := len(q.elements) - 1; i >= 0; i-- {
		reversed = append(reversed, q.elements[i])
	}

	q.elements = reversed
}

/* META
#5. Очереди
#6 Дополнительная задача
Реализуйте круговую (циклическую буферную) очередь статическим массивом фиксированного размера. Добавьте ей метод проверки, полна ли она (при этом добавление новых элементов невозможно).
Обеспечьте эффективное управление указателями начала и конца очереди в рамках массива, чтобы избежать неоправданных сдвигов данных.

WIP

Рефлексия к задаче:
	WIP
*/
