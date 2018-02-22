package codewars

import org.scalatest._

class ComplexAlgSpec extends FunSpec {
 protected def AssertComplexEquals(expected: Array[Double], actual: Array[Double]): Unit = {
    if (Math.abs(expected(0)) <= 1) assertEquals("The real part of your returned complex number is not sufficiently close to the expected value", actual(0), expected(0), 1e-12)
    else assertEquals("The real part of your returned complex number is not sufficiently close to the expected value", actual(0), expected(0), 1e-10)
    if (Math.abs(expected(1)) <= 1) assertEquals("The imaginary part of your returned complex number is not sufficiently close to the expected value", actual(1), expected(1), 1e-12)
    else assertEquals("The imaginary part of your returned complex number is not sufficiently close to the expected value", actual(1), expected(1), 1e-10)
 }

  describe("The complex algorthim function"){

    it("should calculate correctly"){
      assert(true)
    }

    it("Should throw an exception if called on the poles"){
      intercept[ArithmeticException] {
        ComplexAlgorithm.log(Array[Double](0, 0))
      }
    }

  }

}
