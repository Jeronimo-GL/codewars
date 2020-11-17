package codewars

package codewars
import org.scalatest._
import org.scalatest.funspec.AnyFunSpec

class RectangleRotationSpec extends AnyFunSpec{
    describe ("The rectange"){
        it ("Basic cases"){
            assert(Rectangles.rectangleRotation(6, 4) === 23)
            assert(Rectangles.rectangleRotation(30, 2) === 65)
            assert(Rectangles.rectangleRotation(8, 6) === 49)
            assert(Rectangles.rectangleRotation(16, 20) === 333)
        }

        it ("More cases"){
            assert(Rectangles.rectangleRotation(4, 4) === 13)
            assert(Rectangles.rectangleRotation(2, 2) === 5)
            assert(Rectangles.rectangleRotation(2, 4) === 7)
        }
    }
}
