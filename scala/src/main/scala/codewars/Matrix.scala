package codewars

import scala.math._

object Matrix {
    def subMatrix(pos:Int, a: Array[Array[Int]]):Array[Array[Int]] = {
        (a.take(pos-1) :++ a.takeRight( a.length - pos)).map(_.drop(1))
    }

    def determinant(matrix: Array[Array[Int]]): Int = {
        if (matrix.length == 1) matrix(0)(0) else
        (1 to matrix.length).map(x => -1*pow(-1.0, x).toInt*matrix(x-1)(0) * determinant(subMatrix(x, matrix))).foldLeft(0)(_+_)
    }
}
