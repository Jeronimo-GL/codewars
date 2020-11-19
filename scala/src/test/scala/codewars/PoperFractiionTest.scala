package codewars

import org.scalatest.flatspec.AnyFlatSpec
import org.scalatest.matchers.should.Matchers

class PoperFractiionTest extends AnyFlatSpec with Matchers {
  "properFractions(1)" should "return 0" in {
    ProperFractions.properFractions(1) should be (0)
  }
   "properFractions(2)" should "return 1" in {
    ProperFractions.properFractions(2) should be (1)
  }
   "properFractions(5)" should "return 4" in {
    ProperFractions.properFractions(5) should be (4)
  }
   "properFractions(15)" should "return 8" in {
    ProperFractions.properFractions(15) should be (8)
  }
   "properFractions(25)" should "return 20" in {
    ProperFractions.properFractions(25) should be (20)
  } 
}  
