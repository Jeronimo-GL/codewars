package worksheets


object sudokus {
	val blocks = for(r <- 0 to 2;
	    c <- 0 to 2)
	    yield(r,c)
	    
	blocks.foreach((r,c) => println(d"row:$r col:$c"))
}