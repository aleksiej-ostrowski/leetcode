""" 21.py """

# ------------------------------ #
#                                #
#  version 0.0.1                 #
#                                #
#  Aleksiej Ostrowski, 2022      #
#                                #
#  https://aleksiej.com          #
#                                #
# ------------------------------ #  

from typing import Optional

import copy

# Definition for singly-linked list.

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:

        flag1 = list1 == None 
        flag2 = list2 == None

        if flag1 and flag2:
            return None

        node = ListNode()
        head = node

        while True:

            p1 = 0
            p2 = 0

            if not flag1:  
                p1 = list1.val

            if not flag2:
                p2 = list2.val  

            current = 0
            if not flag1 and not flag2:
                current = min(p1, p2)
                if p1 < p2:
                    list1 = list1.next
                else:
                    list2 = list2.next

            if not flag1 and flag2:
                current = p1
                list1 = list1.next

            if flag1 and not flag2:
                current = p2
                list2 = list2.next

            node.val = current

            flag1 = list1 == None 
            flag2 = list2 == None

            if flag1 and flag2:
                break

            new_node = ListNode()
            node.next = new_node
            node = new_node

        return head  


def list_to_ListNode(lst):
    head = ListNode().next
    for x in lst[::-1]:
        head = ListNode(x, head)
    return head    
        
def print_ListNode(head):
    while head != None:
        print(head.val)
        head = head.next

if __name__ == '__main__':

    a1 = [1,2,4]
    a2 = [1,3,4]

    l1 = list_to_ListNode(a1)
    l2 = list_to_ListNode(a2)

    print_ListNode(l1)
    print("-" * 20)
    print_ListNode(l2)
    print("-" * 20)

    res = Solution().mergeTwoLists(list1 = l1, list2 = l2)

    print_ListNode(res)

    print('ok')

