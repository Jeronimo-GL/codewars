package codewars

object CreatePhoneNumbers {  // No need of app if we only use the tests
  def createPhoneNumber(numbers: Seq[Int]):String ={
    "(" + numbers.slice(0,3).mkString + ") " + numbers.slice(3, 6).mkString + "-" + numbers.slice(6, 10).mkString
  }
}
