package codewars

import org.scalatest.flatspec.AnyFlatSpec
import org.scalatest.matchers.should.Matchers

class IncDecSpec extends AnyFlatSpec with Matchers{
    "totalIncDec(0)" should "return 1" in {
    TotalIncDec.totalIncDec(0) should be (1)
  }
  "totalIncDec(1)" should "return 10" in {
    TotalIncDec.totalIncDec(1) should be (10)
  }
  "totalIncDec(2)" should "return 100" in {
    TotalIncDec.totalIncDec(2) should be (100)
  }
  "totalIncDec(3)" should "return 475" in {
    TotalIncDec.totalIncDec(3) should be (475)
  }
  "totalIncDec(4)" should "return 1675" in {
    TotalIncDec.totalIncDec(4) should be (1675)
  }
  
}
