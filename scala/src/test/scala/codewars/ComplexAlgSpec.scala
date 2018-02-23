package codewars

import org.scalatest._
import codewars.CustomCodewarsMatchers._


class ComplexAlgSpec extends FunSpec with Matchers {
  describe("The complex algorthim function"){

    it("Should work for normal numbers"){
      ComplexAlgorithm.log(Array[Double](20, 0)) should complexEqual (Array[Double](2.995732273553991, 0))
      ComplexAlgorithm.log(Array[Double](-1, 0)) should complexEqual (Array[Double](0, 3.141592653589793))
      ComplexAlgorithm.log(Array[Double](1, 1)) should complexEqual (Array[Double](0.346573590279973, 0.785398163397448))
    }

    it("Should throw an exception if called on the poles"){
      intercept[ArithmeticException] {
        ComplexAlgorithm.log(Array[Double](0, 0))
      }
    }

  }
}
