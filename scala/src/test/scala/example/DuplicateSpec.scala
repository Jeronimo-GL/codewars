package codewars

import org.scalatest._

class DulplicateSpec extends FlatSpec with Matchers {

  " din" should  " return (((" in {
    DuplicateEncoder.duplicateEncode("din") shouldEqual "((("
  }

  "recede" should  " return ()()()" in {
    DuplicateEncoder.duplicateEncode("recede") shouldEqual """()()()"""
  }

  " Success" should  " return )())())" in {
    DuplicateEncoder.duplicateEncode("Success") shouldEqual """)())())"""
  }

  """ (( @", """ should  """ return ))((""" in {
    DuplicateEncoder.duplicateEncode("""(( @", """) shouldEqual """))(("""
  }

}
