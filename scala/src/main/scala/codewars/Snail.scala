package codewars

object Snail extends App{
  
  def snail(xs: List[List[Int]]): List[Int] = {
    xs match{
      case List(List()) => List()
      case List(List(x)) => List(x)
      case _ => 
        var pos = (0,0)
        var dir = (0,1)
        var visited : List[(Int, Int)] = List((0,0))
        var route: List[Int] = List(xs(pos._1)(pos._2))
        do{
          val newPos=nextPos(pos, dir)
          visited = visited :+ newPos
          pos=newPos
          route = route :+xs(newPos._1)(newPos._2)
          dir = nextDir(xs, pos, dir, visited)
        }while(dir._1 != 0 || dir._2 != 0)
        route
    }
  }

  // Decides if it is possible to move somewhere
  def nextDir(box: List[List[Int]], pos: (Int, Int),
    dir: (Int, Int), visited: List[(Int, Int)]): (Int, Int) ={
    var newDir = dir
    for (i <- 0 to 3){
      if (isValid(nextPos(pos, newDir), box.length, visited)) return newDir
      else  newDir=turnLeft(newDir)
    }
    (0,0)
  }


  // Changes the direction clockwise
  def turnLeft(dir: (Int, Int)): (Int, Int) = {
    dir match {
      case ( 0, 1) => ( 1, 0) // from E to S
      case ( 1, 0) => ( 0,-1) // from S to W
      case ( 0,-1) => (-1, 0) // from W to N
      case (-1, 0) => ( 0, 1) // from N to E
      case _ => (0,0)
    }
  }


  // Calculates the next position of the snail regardles of the size of the box
  def nextPos(pos :(Int, Int), dir:(Int, Int)): (Int, Int) = (pos._1 + dir._1, pos._2 + dir._2)

  // Calculates if the new position is valid in the box including the visited positions
  def isValid(pos: (Int, Int), dim: Int, visited: List[(Int, Int)]):Boolean = {
    pos._1 >=0 && pos._1 < dim &&
      pos._2 >=0 && pos._2<dim &&
      !visited.contains(pos)
  }

  val box=List(List(1,2,3,4), List(12,13,14,5), List(11,16,15,6), List(10,9,8,7))
  println(snail(box))
}
