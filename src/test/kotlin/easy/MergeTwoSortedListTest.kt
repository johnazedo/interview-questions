package easy

import datastructures.ListNode
import org.junit.jupiter.params.ParameterizedTest
import org.junit.jupiter.params.provider.Arguments
import org.junit.jupiter.params.provider.MethodSource
import java.util.stream.Stream
import kotlin.test.assertEquals

class MergeTwoSortedListTest {

    @ParameterizedTest(name = "given two linked lists {0} and {1}, then it should return a merged sorted linked list {2}")
    @MethodSource("arguments")
    fun execute(
        list1: ListNode?,
        list2: ListNode?,
        expected: ListNode?
    ) {
        var expectedNode = expected
        var actual = MergeTwoSortedList().mergeTwoLists(list1, list2)
        while(expectedNode != null && actual != null){
            if(expectedNode.`val` != actual.`val`) {
                throw Throwable("Not equals numbers ${expectedNode.`val`} != ${actual.`val`}")
            }
            expectedNode = expectedNode.next
            actual = actual.next
        }

        assert(true)
    }

    private companion object {
        @JvmStatic
        fun arguments(): Stream<Arguments> = Stream.of(
            Arguments.of(
                ListNode(1).putNext(ListNode(2).putNext(ListNode(4))),
                ListNode(1).putNext(ListNode(3).putNext(ListNode(4))),
                ListNode(1).putNext(
                    ListNode(1).putNext(
                        ListNode(2).putNext(
                            ListNode(3).putNext(
                                ListNode(4).putNext(
                                    ListNode(4)
                                )
                            )
                        )
                    )
                )
            )
        )
    }
}