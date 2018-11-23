package codewars

object Sudoku {

  def isValid(board: Array[Array[Int]]): Boolean = {
    val blocks = for(r <- 0 to 2;
	                  c <- 0 to 2)
	          yield(r,c)
	  blocks.forall(b => isBlockOK(getBlock(board, b._1, b._2))) &&
    (0 to 8).forall(row => isBlockOK(board(row))) &&
    (0 to 8).forall(col => isBlockOK(getColumn(board, col)))
  }
  
  def isBlockOK(block: Array[Int]):Boolean={
  	block.distinct.length==9 &&
  	block.distinct.forall((1 to 9).contains(_))
  }

	def getColumn(board: Array[Array[Int]], col: Int): Array[Int] = {
 		val r = for (row <- 0 to 8) yield board(row)(col)
 		r.toArray
  }
	
	def getBlock(board: Array[Array[Int]], row: Int, col: Int): Array[Int] ={
		val r=for ( i <- row*3 to row*3+2;
					     j <- col*3 to col*3+2)
				yield board(i)(j)
		
		r.toArray
	}
  
  
  def main(args: Array[String]): Unit = {
    println("Prueba de sudoku") 
    
    val blocks = for(r <- 0 to 2;
	                  c <- 0 to 2)
	          yield(r,c)
	    
	  blocks.foreach(p => println(p._1))
  }
}