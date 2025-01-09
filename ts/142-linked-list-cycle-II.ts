/*
Given the head of a linked list, return the node where the cycle begins.If there is no cycle, return null.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer.Internally, pos is used to denote the index of the node that tail's next pointer is connected to (0-indexed). It is -1 if there is no cycle. Note that pos is not passed as a parameter.

Do not modify the linked list.
*/
class ListNode {
  val: number
  next: ListNode | null
  constructor(val?: number, next?: ListNode | null) {
    this.val = (val === undefined ? 0 : val)
    this.next = (next === undefined ? null : next)
  }
}


function detectCycle(head: ListNode | null): ListNode | null {
  let slow = head
  let fast = head

  while (slow && fast) {
    slow = slow.next
    fast = (fast.next)?.next ?? null

    if (slow === fast) {
      slow = head

      while (slow && fast) {
        slow = slow.next
        fast = fast.next

        if (slow === fast) {
          return slow
        }
      }
    }
  }

  return null;
}
