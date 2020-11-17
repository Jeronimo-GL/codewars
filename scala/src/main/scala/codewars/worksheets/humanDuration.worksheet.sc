import scala.math._

val multipliers = List(
    (1.0, "second"),
    (60.0, "minute"),
    (60*60.0, "hour"),
    (60*60*24.0, "day"),
    (60*60*24*30.0, "month"),
    (60*60*24*30*12.0, "year")
)

val n =132030240


def procTime(time: Double, multipliers: Seq[(Double, String)]) : Seq[(Double, String)] = {
    multipliers match {
        case Nil => Nil
        case (a, s) :: tail => procTime(time - floor(time/a)*a, tail) :+ (floor(time / a), s)
    }
}

val numResult = procTime(n, multipliers.reverse)
    .filter(x => x._1>0)
    .map(x => s"${x._1.toInt} ${x._2}${if (x._1.toInt >1) "s" else ""}")
    .reverse

def concatenate(s: Seq[String]) : String = {
    s match {
        case a:: Nil => a
        case a::b::Nil => a + " and " + b
        case a::b::tail => a + ", " + concatenate(b::tail)
    }
}

concatenate(numResult)