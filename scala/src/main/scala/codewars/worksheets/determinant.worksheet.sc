val m1 = Array(Array(1))

val m2 = Array(Array(1,3), Array(2,5))

val m3 = Array(Array(2,5,3), Array(1,-2,-1), Array(1,3,4))

def subMatrix(pos:Int, a: Array[Array[Int]]):Array[Array[Int]] = {
    (a.take(pos-1) :++ a.takeRight( a.length - pos)).map(_.drop(1))
}


subMatrix(2, m3)
subMatrix(1, m2) 

for {
    x <- 1 to m3.length
    s = subMatrix(x, m3)
} yield(s)

(1 to m3.length).map(x => subMatrix(x, m3))

import scala.math._
pow(-1,3 )