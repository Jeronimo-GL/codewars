// import Dependencies._

lazy val root = (project in file(".")).
  settings(
    inThisBuild(List(
      organization := "jeronimo.garcia-loygorri",
      scalaVersion := "2.12.3",
      version      := "0.1.0-SNAPSHOT"
    )),
    name := "Codewars katas",
    libraryDependencies += "org.scalatest" %% "scalatest" % "3.0.4" % "test"
  )
