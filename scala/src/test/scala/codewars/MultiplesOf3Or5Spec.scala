package codewars

import org.scalatest._

class MultiplesOf3Or5Spec extends FunSpec with Matchers{

  import MultiplesOf3Or5._

  describe("Multiples of 3 or 5") {
    it("solution of 10 ") {
      solution(10) should be (23)
    }

    it("solution of 20") {
      solution(20) should be (78)
    }

    it("should return 0 for solution(0)"){
      solution(0) should be (0)
    }

    it("should return 0 for solution(1)"){
      solution(1) should be (0)
    }

    it("it shouldf work for big numbers"){
      assert(solution(1000000)> 0)
    }
  }
}
