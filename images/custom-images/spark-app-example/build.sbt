val scala3Version = "3.3.1"

lazy val root = project
  .in(file("."))
  .settings(
    name := "spark-app-example",
    version := "0.1.0-SNAPSHOT",

    scalaVersion := scala3Version,

    assembly / assemblyJarName := "spark-app-example.jar",

    libraryDependencies ++= Seq(
      "org.apache.spark" %% "spark-sql" % "3.5.0" % Provided cross CrossVersion.for3Use2_13,
      "org.apache.spark" %% "spark-streaming" % "3.5.0" % Provided cross CrossVersion.for3Use2_13,

      "org.scalatest" % "scalatest_3" % "3.2.15" % Test excludeAll (
          ExclusionRule(organization = "org.scala-lang.modules")
        ),
      "org.apache.spark" %% "spark-core" % "3.5.0" % Test cross CrossVersion.for3Use2_13,
      "org.apache.spark" %% "spark-sql" % "3.5.0" % Test cross CrossVersion.for3Use2_13
    )
  )
