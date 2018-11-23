package codewars


object Magnets {
  def doubles(maxk: Int, maxn: Int): Double = {
    val pairs = for{
      k <- 1 to maxk;
      n <- 1 to maxn
    } yield v(k,n)

    pairs.toList.foldLeft(0d)(_+_)
  }

  def v(k: Int, n: Int): Double = {
    1 / (k*Math.pow((n+1).toDouble, (2*k).toDouble))
  }

  def main(args: Array[String]) = {
    println(doubles(1,3))
  }
}
