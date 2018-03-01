package codewars

import org.scalatest._

class KprimesSpec extends FlatSpec with Matchers{
  "Kprimes" must "be 1 for 1" in {
    KPrimes.countKprimes(1, 0, 1) should equal ("")
  }

  it must "work for 2, 0 ,100" in {
    KPrimes.countKprimes(2, 0, 100) should equal ("4, 6, 9, 10, 14, 15, 21, 22, 25, 26, 33, 34, 35, 38, 39, 46, 49, 51, 55, 57, 58, 62, 65, 69, 74, 77, 82, 85, 86, 87, 91, 93, 94, 95")
  }

  it must "return nothing for 12, 100000, 100100)" in {
    KPrimes.countKprimes(12, 100000, 100100) should equal ("")
  }

  it must "work with 1, 2, 30" in {
    KPrimes.countKprimes(1, 2, 30) should equal ("2, 3, 5, 7, 11, 13, 17, 19, 23, 29")
  }

  it must "work with 8,  10000000, 10000200" in {
    KPrimes.countKprimes(8, 10000000, 10000200) should equal ("10000096, 10000152, 10000165, 10000200")
  }
}
