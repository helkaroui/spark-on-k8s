package dev.sharek.examples

import org.apache.spark.sql.SparkSession

import scala.math.random

class SparkPi(args: Array[String]) {

  val spark = SparkSession
    .builder
    .appName("Spark Pi")
    .getOrCreate()

  SparkPiLogic(spark, args)

  spark.stop()
}


class SparkPiLogic(spark: SparkSession, args: Array[String]) {

  val slices = if (args.length > 0) args(0).toInt else 2
  val numSecondsToSleep = if (args.length > 1) args(1).toInt else 10


  val n = math.min(100000L * slices, Int.MaxValue).toInt // avoid overflow
  val count = spark.sparkContext.parallelize(1 until n, slices).map { i =>
    val x = random * 2 - 1
    val y = random * 2 - 1
    if (x * x + y * y <= 1) 1 else 0
  }.reduce(_ + _)

  println(s"Pi is roughly ${4.0 * count / (n - 1)}")


  for (i <- 1 until numSecondsToSleep) {
    println(s"Alive for $i out of $numSecondsToSleep seconds")
    Thread.sleep(1000)
  }

}