package codewars

import scala.math._

object Rectangles {
    
    def rectangleRotation(h:Int, w: Int): Int = {
        def isEven(number: Int):Boolean  = {
            number % 2 == 0
        }

        val wMax = floor(w / sqrt(2)).toInt
        val hMax = floor(h / sqrt(2)).toInt
        val dif = if (isEven(hMax + wMax)) 0 else 1
        (wMax * hMax) + (wMax + 1) * (hMax + 1) - dif
  }
}
