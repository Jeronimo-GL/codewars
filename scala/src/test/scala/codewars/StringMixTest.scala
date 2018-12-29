package codewars

import org.scalatest._
import org.scalatest.Assertions._

import StringMixTest._

class StringMixTest extends FlatSpec { 
  it should "pass basic tests" in {
    testing("Are they here", "yes, they are here", "2:eeeee/2:yy/=:hh/=:rr")
    testing("uuuuuu", "uuuuuu", "=:uuuuuu")
    testing("looping is fun but dangerous", "less dangerous than coding", 
            "1:ooo/1:uuu/2:sss/=:nnn/1:ii/2:aa/2:dd/2:ee/=:gg")
    testing(" In many languages", " there's a pair of functions", 
            "1:aaa/1:nnn/1:gg/2:ee/2:ff/2:ii/2:oo/2:rr/2:ss/2:tt")
    testing("Lords of the Fallen", "gamekult", "1:ee/1:ll/1:oo")
    testing("codewars", "codewars", "")

  }
}

object StringMixTest {

  private def testing(s1: String, s2: String, expect: String): Unit = {
    println("Testing:\n" + s1 + "\n" + s2)
    val actual: String = StringMix.mix(s1, s2)
    println("Actual: " + actual)
    println("Expect: " + expect)
    println("*")
    assertResult(expect){actual}
  }
}
