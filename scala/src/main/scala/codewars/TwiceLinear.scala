package codewars

//import scala.collection.mutable._

object DblLinear extends App{

  println(time {DblLinear.dblLinear(100)}) 
  println(time {DblLinear.dblLinear(1000)})
  println(time {DblLinear.dblLinear(10000)})

  def time[R](block: => R): R = {
    val t0 = System.nanoTime()
    val result = block    // call-by-name
    val t1 = System.nanoTime()
    print(f" Elapsed time: ${(t1 - t0)/1000000.0}%f ms ")
    result
  }

  def isIn(num: Int, comps:Array[Int], max: Int): Boolean = {
    var i=0
    while (i<=max){
      if (comps(i)==num) return true
      i+=1
    }
    return false
  }

  def dblLinear(n: Int): Int = {
    // your code 
    val twos = new Array[Int](n+1)
    val threes = new Array[Int](n+1)
    twos(0)=3
    threes(0)=4

    var next = 1
    var lastTwo = 0
    var lastThree = 0
    var nextTwo = 0
    var nextThree = 0

    for (i <- 1 to n){
      if (twos(nextTwo) <= threes(nextThree)){
        next=twos(nextTwo)
        nextTwo += 1
      } else {
        next= threes(nextThree)
        nextThree += 1
      }
//      if (!threes.slice(0, lastThree).contains(next*2+1)){
      if (!isIn(next*2+1, threes, lastThree)){  
        lastTwo +=1
        twos(lastTwo) =(next*2 +1)
      }
//      if (!twos.slice(0, lastTwo).contains(next*3+1)){
      if (!isIn(next*3+1, twos, lastTwo)){  
        lastThree += 1
        threes(lastThree)= (next*3 +1)

      }
    }
    return next
  }
}
