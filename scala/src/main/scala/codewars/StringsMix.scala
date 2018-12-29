package codewars

object StringMix {

  def mix(s1: String, s2: String): String = {
    val counted=(countChars(s1, "1"):::countChars(s2, "2")).groupBy(_._1)
    val diff=counted.values.map(compStrs(_)).toList.sortWith(sortFunction)
    diff.map(toString) mkString "/"
  }
  
  def toString(str : (Int, Char, String)):String = {
    var retStr= str._3 + ":"
    for (i <- 1 to str._1){
      retStr+=str._2
    }
    retStr
  }
  
  def sortFunction(e1:(Int, Char, String), e2: (Int, Char, String)): Boolean  = {
    if  (e1._1 > e2._1) true
    else if (e1._1 == e2._1){
        toString(e1) < toString(e2)
    }
    else false
  }
  
  def compStrs(c1: List[(Char, Int, String)]): (Int, Char, String) = {
    c1 match {
      case e::Nil => (e._2, e._1, e._3)
      case a::b::Nil =>  if (a._2 == b._2) (a._2, a._1, "=")
        else if (a._2> b._2) (a._2, a._1, a._3)
        else (b._2, b._1, b._3)
      case Nil=> (0, ' ', "")
        
    }
  }

  def countChars(s: String, id:String):List[(Char, Int, String)] = {
    val valid="abcdefghijklmnopqrtsuvwxyz".toList
    s.toList.filter(valid.contains(_)).
      map((_, 1)).groupBy(_._1).
      map(x=>(x._1, x._2.length, id)).
      toList.sortWith(_._2>_._2).filter(_._2>1)
  }
  
  def main(args: Array[String]): Unit = {
    val s1 = "my&friend&Paul has heavy hats! &"
    val s2 = "my friend John has many many friends &"
    print(mix(s1, s2))
  }
}
