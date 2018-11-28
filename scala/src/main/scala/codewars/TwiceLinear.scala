package codewars

//import scala.collection.mutable._

object DblLinear extends App{

  println(time {DblLinear.dblLinear(10)})
//  println(time {DblLinear.dblLinear(1000)})
//  println(time {DblLinear.dblLinear(10000)})

  def time[R](block: => R): R = {
    val t0 = System.nanoTime()
    val result = block    // call-by-name
    val t1 = System.nanoTime()
    println("Elapsed time: " + (t1 - t0)/1000 + " ms")
    result
  }

  def dblLinear(n: Int): Int = {
    // your code
    val twos = new Array[Int](n)
    val threes = new Array[Int](n)
    twos(0)=3
    threes(0)=4

    var next = 1
    var lastTwo = 0
    var lastThree = 0

    for (i <- 2 to n+1){
      if (twos(lastTwo) <= threes(lastThree)){
        next=twos(lastTwo)
        lastTwo += 1
      } else {
        next= threes(lastThree)
        lastThree +=1
      }

      if (!threes.contains(next*2+1)){
        twos(lastTwo) =(next*2 +1)
        println("[2] "+ lastTwo + " " + twos(lastTwo))
      }

      if (!twos.contains(next*3+1)){
        threes(lastThree)= (next*3 +1)
        println("[3] "+ lastThree + " " + threes(lastThree))
      }
    }
    // for (i <- 0 to n-1){
    //   println(twos(i) + " " + threes(i))
    // }
    return next
  }
}
