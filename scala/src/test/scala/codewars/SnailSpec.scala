package codewars

import org.scalatest._
import Snail.snail

class SnailSpec extends FunSuite {
  test("Must work for empty box") {
    assert(snail(List(List())) == List())
  }

  test("Must work for 1x1 box") {
    assert(snail(List(List(1))) == List(1))
  }

  test("Must work for 2x2 box") {
    assert(snail(List(List(1,2), List(3,4))) == List(1,2,4,3))
  }

  test("Must work for a 3x3 box") {
    val box= List(List(1,2,3), List(4,5,6), List(7,8,9))
    assert(snail(box)== List(1,2,3,6,9,8,7,4,5))
  }

  test("Must work for a 4x4 box") {
    val box= List(
      List(1,2,3,4),
      List(5,6,7,8),
      List(9,10,11,12),
      List(13,14,15,16))
    assert(snail(box) == List(1,2,3,4,8,12,16,15,14,13,9,5,6,7,11,10))
  }
}
