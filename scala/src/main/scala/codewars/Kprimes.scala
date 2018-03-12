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
    val l7 =ckToList(7, 0, s).map((x:Int) => (aux(s-x,ckToList(3, 0, s-x)).length ))
    return l7.reduce(_+_)
  }

  def ckToList(k: Int, from:Int, nd:Int):List[Int]={
    countKprimes(k, from, nd).split(",").map(_.trim.toInt).toList
  }


  def isPrime(n: Int): Boolean = (2 to math.sqrt(n).toInt) forall (x => n % x != 0)


  def aux(n:Int, ls:List[Int]):List[Int]={
    ls.map(n-_).filter((x:Int) => (isPrime(x) && x > 0))
  }


  def puzzle_Complete(s: Int): Int = {
    val l7 =ckToList(7, 0, s).map((x:Int) => (x, s-x,
      aux(s-x,ckToList(3, 0, s-x)).length
    ))
    l7.foreach(println)
    return 0
  }


/*  
  def primes(number: Int, list: List[Int] = List()): List[Int] = {
    for(n <- 2 to number if (number % n == 0)) {
      return primes(number / n, list :+ n)
    }
    list
  }
 */


  aux(15,List(8, 12))
  puzzle(200)
//  ctk1(72, List(8, 12, 18, 20, 27, 28, 30, 42, 44, 45, 50, 52, 63, 66, 68, 70))
} 
