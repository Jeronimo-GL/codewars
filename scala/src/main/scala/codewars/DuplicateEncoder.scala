package codewars

object DuplicateEncoder extends App{
  val word=""" (( @", """
  val cs=word.groupBy(identity).mapValues(v => v.size)
  print(word.map((a:Char) =>if  (cs(a) == 1) """(""" else """)""" ))

  def duplicateEncode(word: String):String = {
    val lWord=word.toLowerCase
    val counts = lWord.groupBy(identity).mapValues(v => v.size)
    return lWord.map((a:Char) =>if  (counts(a) ==1) """(""" else """)""" ).mkString
  }
}
