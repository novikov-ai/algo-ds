package main

import (
	"fmt"
	// "os"
)

type Stack[T any] struct {
	values []T
	count  int
}

func (st *Stack[T]) Size() int {
	return st.count
}

/* META
#4. Стек
#2 Регулярная задача
Переделайте реализацию стека так, чтобы она работала не с хвостом списка как с верхушкой стека, а с его головой.
Рефлексия к задаче:
	Сложность O(1) сохранена, мы просто изменили взятие элемента по индексу.
*/

func (st *Stack[T]) Peek() (T, error) {
	var result T

	if st.count == 0 {
		return result, fmt.Errorf("stack is empty")
	}

	result = st.values[0]

	return result, nil
}

// Мера сложности: O(1)
func (st *Stack[T]) Pop() (T, error) {
	var result T

	if st.count == 0 {
		return result, fmt.Errorf("stack is empty")
	}

	result = st.values[0]

	st.values = st.values[1:st.count]
	st.count -= 1

	return result, nil
}

// Мера сложности: O(1) амортизированно из-за возможных аллокаций на расширение массива
func (st *Stack[T]) Push(itm T) {
	st.values = append(st.values, itm)
	st.count += 1
}

/* META
#4. Стек
#3 Регулярная задача
Не запуская программу, скажите, как отработает такой цикл?
~~~
while (stack.size() > 0)
    stack.pop()
    stack.pop()
~~~

Рефлексия к задаче:
	Если в стеке перед запуском size — четное, то будет N/2 итераций, пока не опустошится стек.
	Если стек пустой, то мы даже не зайдем.
	Если в стеке перед запуском size — нечетное, то мы сделаем N/2-1 итераций перед тем как вернуть ошибку.
*/

/* META
#4. Стек
#4 Дополнительная задача
Напишите функцию, которая получает на вход строку и проверяет балансированность скобок.

Рефлексия к задаче:
	Сложность пространственная и временная: O(n)
*/

func IsBalanced(value string) bool {
	stack := Stack[rune]{}

	for _, v := range value {
		switch v {
		case rune('('):
			stack.Push(v)
		case rune(')'):
			_, err := stack.Pop()
			if err != nil {
				return false
			}
		}
	}

	return stack.Size() == 0
}

/* META
#4. Стек
#5 Дополнительная задача
Расширьте фукнцию из предыдущего примера, если скобки могут быть трех типов: (), {}, [].

Рефлексия к задаче:
	Сложность пространственная и временная: O(n)
*/

func IsBalancedV2(value string) bool {
	stack := Stack[rune]{}

	for _, v := range value {
		switch v {
		case rune('('), rune('{'), rune('['):
			stack.Push(v)
		case rune(')'), rune('}'), rune(']'):
			_, err := stack.Pop()
			if err != nil {
				return false
			}
		}
	}

	return stack.Size() == 0
}

/* META
#4. Стек
#6. Дополнительная задача
Добавьте в стек функцию, возвращающую текущий минимальный элемент в нём за O(1)

Рефлексия к задаче:
	Нам необходимо завести еще один стек, который будет наполняться каждый раз при вставке таким образом,
	что у нас будет стек минимальных значений.

	Очищаться такой стек будет при удалении соответствующих элементов из основного стека.
*/

func (st Stack[T]) IsMinimal() T {
	var result T

	value, err := st.valuesMinimum.Peek()
	if err != nil {
		result = value
	}

	return result
}

/* META
#4. Стек
#7. Дополнительная задача
Добавьте в стек функцию, которая возвращает среднее значение всех элементов в стеке за O(1).

Рефлексия к задаче:
	Для организации среднего значения нам нужно хранить сумму по всем значениям, которое будет обновляться
	при вставке или удалении.
*/

func (st Stack[T]) IsAverage() T {
	return st.sum / st.Size()
}

/* META
#4. Стек
#8. Дополнительная задача
Постфиксная запись выражения

Рефлексия к задаче:
	WIP
*/