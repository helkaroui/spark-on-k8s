val scala2Version = "2.13.2"

lazy val root = project
  .in(file("."))
  .settings(
    name := "spark-app-example",
    version := "0.1.0-SNAPSHOT",

    scalaVersion := scala2Version,

    assembly / assemblyJarName := "spark-app-example.jar",

    libraryDependencies ++= Seq(
      "org.apache.spark" %% "spark-sql" % "3.5.0" % Provided,
      "org.apache.spark" %% "spark-streaming" % "3.5.0" % Provided,

      "org.scalatest" %% "scalatest" % "3.2.15" % Test,
      "org.apache.spark" %% "spark-core" % "3.5.0" % Test,
      "org.apache.spark" %% "spark-sql" % "3.5.0" % Test
    )
  )
