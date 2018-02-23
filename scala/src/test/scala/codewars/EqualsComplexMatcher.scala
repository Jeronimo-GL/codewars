package codewars

import org.scalatest._
import matchers._

trait CustomCodewarsMatchers {
  class EqualsComplexMatcher(expected: Array[Double]) extends Matcher[Array[Double]]{

    def isPartOK(expected: Double, actual: Double): Boolean = {
      (Math.abs( expected) <=1 && Math.abs(expected-actual)<1e-12) ||
      (Math.abs( expected) > 1 && Math.abs(expected-actual)<1e-10)
    }

    def apply(left: Array[Double]) = {
      if (isPartOK(expected(0), left(0)) && isPartOK(expected(1), left(1))) {
         MatchResult(true,
          "",
          "Numbers are equal enough"
        )
      } else if (isPartOK(expected(0), left(0))){
         MatchResult(false,
          f"Imaginary part is not equal enough. Expected: ${expected(1)}%.12f found: ${left(1)}%.12f",
          "Numbers are equal enough"
        )
      } else {
        MatchResult(false,
          f"Real part is not equal enough. Expected: ${expected(0)}%.12f found: ${left(0)}%.12f",
          "Numbers are equal enough"
        )
      }


    }
  }

  def complexEqual(expected: Array[Double]) = new EqualsComplexMatcher(expected)
}

object CustomCodewarsMatchers extends CustomCodewarsMatchers
