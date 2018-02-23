package codewars

import org.scalatest._
import codewars.CreatePhoneNumbers._

class PhoneNumberSpec extends FunSpec with Matchers{
  describe("The Phone number creator"){
    it ("Should create proper phone numbers"){
      createPhoneNumber(Seq(1, 2, 3, 4, 5, 6, 7, 8, 9, 0)) should be  ("(123) 456-7890")
    }
  }
}
