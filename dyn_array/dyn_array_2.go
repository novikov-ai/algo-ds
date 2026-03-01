package main

/* META
#3. Динамические массивы
#6 Доп. задача
Реализуйте динамический массив на основе банковского метода.

Рефлексия к задаче #6:
	Метод Insert (по сложности схож с обычной реализацией)
	Сложность пространственная O(n)
	Сложность временная: O(n)

	Но аммортизированно мы приближаемся к O(1) по памяти и времени.
*/

const operationCost = 3

type DynArrayBank[T any] struct {
	count         int
	capacity      int
	array         []T
	operationsSum int
}

func (da *DynArrayBank[T]) Init() {
	da.count = 0
	da.operationsSum = 0
	da.MakeArray(16)
}

func (da *DynArrayBank[T]) MakeArray(sz int) {
	arr := make([]T, sz)

	copy(arr, da.array[:da.count])

	da.capacity = sz
	da.array = arr
}

func (da *DynArrayBank[T]) IsCostEnoughForAllocate() bool {
	return da.operationsSum >= da.count
}

func (da *DynArrayBank[T]) Insert(itm T, index int) error {
	if index < 0 || index > da.count {
		return fmt.Errorf("bad index '%d'", index)
	}

	da.operationsSum += operationCost

	if da.IsCostEnoughForAllocate() {
		da.MakeArray(da.capacity * 2)
		da.operationsSum = 0
	}

	right := da.array[index:da.count]
	shiftedRight := da.array[index+1 : da.count+1]

	copy(shiftedRight, right)

	da.array[index] = itm

	da.count++

	return nil
}

/* META
#3. Динамические массивы
#7 Доп. задача
Реализуйте многомерный динамический массив: произвольное количество измерений,
при этом каждое измерение может внутри масштабироваться по потребности.
В конструкторе задаётся число измерений и размер по каждому из них.
Обращаться к такому массиву надо как к обычному многомерному, например: myArr[1,2,3].

Сложность пространственная [WIP]
Сложность временная: [WIP]

Рефлексия к задаче #7:
	[WIP]
*/
