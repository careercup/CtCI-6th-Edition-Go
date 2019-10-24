package chapter2

func AddTwoLists(l1, l2 *DoublyLinkedList) int {
	multiplyFactor := 1
	res := 0
	n1, n2 := l1.head, l2.head
	for n1 != nil || n2 != nil {
		if n1 == nil {
			res += multiplyFactor * n2.value
			n2 = n2.next
		} else if n2 == nil {
			res += multiplyFactor * n1.value
			n1 = n1.next
		} else {
			res += multiplyFactor * (n1.value + n2.value)
			n1 = n1.next
			n2 = n2.next
		}
		multiplyFactor *= 10
	}
	return res
}

func AddTwoListsB(l1, l2 *DoublyLinkedList) *DoublyLinkedList {
	result := GetLinkedList()
	carryNumber := 0
	sum := 0
	n1, n2 := l1.head, l2.head
	for n1 != nil || n2 != nil {
		sum = carryNumber
		if n1 != nil {
			sum += n1.value
			n1 = n1.next
		}

		if n2 != nil {
			sum += n2.value
			n2 = n2.next
		}

		carryNumber = sum / 10
		result.Insert(sum % 10)

	}
	if carryNumber != 0 {
		result.Insert(carryNumber)
	}
	return result
}
