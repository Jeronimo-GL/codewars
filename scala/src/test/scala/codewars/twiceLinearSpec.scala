package codewars

import org.scalatest._


class TwiceLinearSpec extends FlatSpec with Matchers {
  "Twicelinear" should "work" in {
    DblLinear.dblLinear(50) should be (175)
    DblLinear.dblLinear(100) should be (447)
    DblLinear.dblLinear(500) should be (3355)
    DblLinear.dblLinear(1000) should be (8488)
  }
}
