package codewars

package codewars
import org.scalatest._
import org.scalatest.funspec.AnyFunSpec
import org.scalatest.matchers.should.Matchers

class HumanDurationSpec extends AnyFunSpec with Matchers{
  describe("Human readable duration"){
      it ("Must pass the basic cases"){
        val tests = List(
            (0, "now"),
            (1, "1 second"),
            (62, "1 minute and 2 seconds"),
            (120, "2 minutes"),
            (3600, "1 hour"),
            (3662, "1 hour, 1 minute and 2 seconds"),
            (86400, "1 day"),
            (172800, "2 days"),
            (2678400, "31 days"),
            (61, "1 minute and 1 second"),
            (60, "1 minute")
        )

        tests.foreach {
            case (input, expected) =>
            HumanDuration.formatDuration(input) should be (expected)
        }
      }
  }
}
