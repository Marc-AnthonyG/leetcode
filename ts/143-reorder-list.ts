/*
You are given the head of a singly linked-list. The list can be represented as:

L0 → L1 → … → Ln - 1 → Ln
Reorder the list to be on the following form:

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
You may not modify the values in the list's nodes. Only nodes themselves may be changed.
*/

interface ListNode {
  val: number
  next: ListNode | null
}

/**
 Do not return anything, modify head in-place instead.
 */
function reorderList(head: ListNode | null): void {
  let reverseList: ListNode | null = null
  let current = head

  while (current) {
    const nextInReverseList = current.next

    if (nextInReverseList) {
      current.next = nextInReverseList.next
      current = nextInReverseList.next

      nextInReverseList.next = reverseList
      reverseList = nextInReverseList
    } else {
      current == null
    }
  }

  current = head
  while (current && reverseList) {
    const nextInCurrent = current.next
    const nextInReverseList = reverseList.next

    reverseList.next = nextInCurrent
    current.next = reverseList

    current = nextInCurrent
    reverseList = nextInReverseList
  }
};

//[1,2,3,4,5]
