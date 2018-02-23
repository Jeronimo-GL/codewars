package codewars

object ComplexAlgorithm extends App{

  def log(complex: Array[Double]): Array[Double] = {
    if (complex(0) != 0 || complex(1) != 0){
      return Array[Double](0,0)
    } else{
      throw new ArithmeticException("") // do it!
    }
  }
}
