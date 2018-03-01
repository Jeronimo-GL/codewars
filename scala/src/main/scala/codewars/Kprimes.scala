package codewars

object KPrimes extends App{

  def countFactors(number: Long, count: Long = 0): Long ={
    for(n <- 2L to number if (number % n == 0)) {
      return countFactors(number / n, count + 1)
    }
    count
  }

  def countKprimes(k:  Int, start: Long, nd: Long): String = {
    (start to nd).filter(countFactors(_) == k).mkString(", ")
  }
  
  def puzzle(s: Int): Int = {
    return 0
  }

  def primes(number: Int, list: List[Int] = List()): List[Int] = {
    for(n <- 2 to number if (number % n == 0)) {
      return primes(number / n, list :+ n)
    }
    list
  }

  val s =147
  val l7= countKprimes(7, 0, 143).split(",").map(_.trim.toInt).toList
  for(e <- l7)println(countKprimes(3, 0, s - e))

} 
