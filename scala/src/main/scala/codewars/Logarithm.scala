package codewars

object ComplexAlgorithm extends App{

  def log(complex: Array[Double]): Array[Double] = {
    if (complex(0) != 0 || complex(1) != 0){
      val x = complex(0)
      val y = complex(1)
      return Array[Double](Math.log(Math.sqrt(x*x + y*y)), Math.atan2(y,x))

    } else{
      throw new ArithmeticException("")
    }
  }
}
