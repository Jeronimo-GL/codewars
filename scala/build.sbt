// import Dependencies._
ThisBuild / scalaVersion := "2.13.3"
ThisBuild / organization := "jeronimo.garcia-loygorri"
ThisBuild / version      := "0.1.0-SNAPSHOT"

// Root project
lazy val root = (project in file("."))
  .settings(
    name := "Codewars katas",
    libraryDependencies ++= Seq(
      "org.scalactic" %% "scalactic" % "3.2.2",
      "org.scalatest" %% "scalatest" % "3.2.2" % "test"
    )
  )

