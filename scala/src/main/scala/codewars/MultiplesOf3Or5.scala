package codewars
//import scala.collection.immutable.List

object MultiplesOf3Or5 {
  def solution(number: Int): Long = {
    if (number <= 1) return 0
    else {
      val gen = for {
        i <- 1L to (number.longValue()-1)
        if (i % 3 == 0 || i % 5 == 0)
      } yield i
      return gen.reduce(_ + _)
    }
  }
}
