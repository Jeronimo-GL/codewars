import scala.math._

val w = 2 
val h = 4
def isEven(number: Int) = number % 2 == 0

val wMax = floor(w / sqrt(2)).toInt 
val hMax = floor(h / sqrt(2)).toInt 

if (!(isEven(hMax) && isEven(wMax))) {
  0
} else {
  (wMax * hMax) + (wMax +1 ) * (hMax + 1) 
}