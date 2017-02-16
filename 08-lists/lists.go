/*
* if you are familiar with Arrays in JavaScript or Lists in any other languages
* then you might know that a List contains your data in a ordered sequence
* you can dynamically add and remove item*
* List is a linear data strucutre with a starting point and ending point
* we will work with slices here since
* they are flexible and extensible in comparsion to Arrays
 */

package list

// List - this will be our structure
type List struct {
	memory []interface{}
	length int
}

// Push - we need a method to add a value to the end of our List
func (list *List) Push(value interface{}) {
	// append will increase the capacity and realocate the memory for us
	list.memory = append(list.memory, value)

	// increase the length of our List
	list.length = list.length + 1
}

// Pop - we need a method to pop the last item from the List
func (list *List) Pop() interface{} {
	// if the memory is empty then return
	if list.length == 0 {
		return nil
	}

	// decrease length and shrink memory by one, also return popped value
	lastIndex := list.length - 1
	lastIndexValue := list.memory[lastIndex]
	list.memory = list.memory[:lastIndex]
	list.length = lastIndex

	return lastIndexValue

}

// Get - Get will give us a specific value based on its index
func (list *List) Get(index int) interface{} {
	return list.memory[index]
}