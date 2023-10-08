package dev.sharek.examples

import dev.sharek.examples.SparkPi.SparkPiLogic
import org.apache.spark.sql.SparkSession
import org.scalatest.funsuite.AnyFunSuite
import org.scalatest.matchers.must.Matchers._

class IntegrationTest extends AnyFunSuite {

  val spark = SparkSession
    .builder
    .appName("Spark Pi Test")
    .master("local[*]")
    .getOrCreate()

  test("SparkPi will run until completion without errors") {
    noException should be thrownBy ({
      SparkPiLogic(spark, Array("10"))
    })
  }

}
