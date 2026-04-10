package main

import (
	"fmt"
	"sort"
)

/* META
#9. Ассоциативный массив
#1 Регулярная задача
Выясните, как в языке программирования, которым вы пользуетесь, реализован тип данных Словарь.

	Словарь в языке Go организован как "мапа": map, например: map[string]int{}

Рефлексия к задаче:
	Часто используемая СД в случаях, когда нужно быстро получить элемент, не проходясь по всему списку,
	или организовать хранение различающихся элементов, исключая дубликаты.
*/

/* META
#9. Ассоциативный массив
#5 Дополнительная задача
Реализуйте словарь с использованием упорядоченного списка по ключу для оптимизации производительности поиска.
Оцените временную сложность операций вставки, удаления и поиска в таком словаре.

	Временная сложность операций:
		- вставка O(n) — за счет поиска O(log(n)) + сдвиг списка O(n)
		- удаления O(n) — за счет поиска O(log(n)) + сдвиг списка O(n)
		- поиска O(log(n)) — бинарный поиск

Рефлексия к задаче:
	Реализация на упорядочном списке выгодна, когда операций чтения/поиска значительно больше, чем операций записи, 
	
	Для частых вставок и удалений лучше подходит хеш-таблица (амортизированное O(1)).
*/

type KeyValue[T any] struct {
	key   string // сортировка по этому полю
	value T
}

type NativeDictionaryWithOrderedList[T any] struct {
	orderedList []KeyValue[T]
}

func (ndo *NativeDictionaryWithOrderedList[T]) IsKey(key string) bool {
	// бинарный поиск O(log(n))
	idx := sort.Search(len(ndo.orderedList), func(i int) bool {
		return ndo.orderedList[i].key >= key
	})

	return ndo.orderedList[idx].key == key
}

func (ndo *NativeDictionaryWithOrderedList[T]) Get(key string) (T, error) {
	// бинарный поиск O(log(n))
	idx := sort.Search(len(ndo.orderedList), func(i int) bool {
		return ndo.orderedList[i].key >= key
	})

	if idx < len(ndo.orderedList) && ndo.orderedList[idx].key == key {
		return ndo.orderedList[idx].value, nil
	}

	var zero T
	return zero, fmt.Errorf("key not found")
}

func (ndo *NativeDictionaryWithOrderedList[T]) Put(key string, value T) {
	idx := sort.Search(len(ndo.orderedList), func(i int) bool {
		return ndo.orderedList[i].key >= key
	})

	// при существовании ключа — обновляем значение
	if idx < len(ndo.orderedList) && ndo.orderedList[idx].key == key {
		ndo.orderedList[idx].value = value
		return
	}

	// новая вставка
	ndo.orderedList = append(ndo.orderedList, KeyValue[T]{})

	copy(ndo.orderedList[idx+1:], ndo.orderedList[idx:])

	ndo.orderedList[idx] = KeyValue[T]{key: key, value: value}
}

/* META
#9. Ассоциативный массив
#6 Дополнительная задача
Создайте словарь, в котором ключи представлены битовыми строками фиксированной длины.
Реализуйте методы добавления, удаления и поиска элементов, используя битовые операции для ускорения работы.

	...

Рефлексия к задаче:
	...
*/
