// import Dependencies._
ThisBuild / scalaVersion := "2.12.3"
ThisBuild / organization := "jeronimo.garcia-loygorri"
ThisBuild / version      := "0.1.0-SNAPSHOT"

//Dependencies
val scalaTest = "org.scalatest" %% "scalatest" % "3.0.4"

// Root project
lazy val root = (project in file("."))
  .settings(
    name := "Codewars katas",
    libraryDependencies += scalaTest % "test"
  )

