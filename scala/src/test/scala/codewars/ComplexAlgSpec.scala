package codewars

import org.scalatest.funspec._
import codewars.CustomCodewarsMatchers._
import org.scalatest.matchers.should.Matchers


class ComplexAlgSpec extends AnyFunSpec with Matchers {
  describe("The complex algorthim function"){

    it("should pass de example tests"){
      ComplexAlgorithm.log(Array[Double](20, 0)) should complexEqual (Array[Double](2.995732273553991, 0))
      ComplexAlgorithm.log(Array[Double](-1, 0)) should complexEqual (Array[Double](0, 3.141592653589793))
      ComplexAlgorithm.log(Array[Double](1, 1)) should complexEqual (Array[Double](0.346573590279973, 0.785398163397448))
    }

    it("should pass the fixed tests"){
      ComplexAlgorithm.log(Array[Double](-24, 7)) should complexEqual (Array[Double](3.218875824868201, 2.857798544381465))
      ComplexAlgorithm.log(Array[Double](0, 1)) should complexEqual (Array[Double](0, 1.570796326794897))
      ComplexAlgorithm.log(Array[Double](1, 0)) should complexEqual (Array[Double](0, 0))
      ComplexAlgorithm.log(Array[Double](8834d / 1997d, 93387d / 42243d)) should complexEqual (Array[Double](1.598433436873693, 0.463447163733142))
      ComplexAlgorithm.log(Array[Double](28347, 73623)) should complexEqual (Array[Double](11.275831160977204, 1.203262110987354))
    }

    it("Should throw an exception if called on the poles"){
      intercept[ArithmeticException] {
        ComplexAlgorithm.log(Array[Double](0, 0))
      }
    }
  }
}
