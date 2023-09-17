package dev.sharek.examples

import org.apache.spark.sql.SparkSession
import org.apache.spark.sql.streaming.Trigger

def StreamingExample(args: Array[String]): Unit = {

  val spark = SparkSession
    .builder
    .appName("Streaming Example")
    .getOrCreate()

  StreamingExampleLogic(spark, args)
}

class StreamingExampleLogic(spark: SparkSession, args: Array[String]) {
  val df = spark.readStream
    .format("rate")
    .option("rowsPerSecond", 1)
    .option("numPartitions", 1)
    .option("rampUpTime", 1)
    .load()

  val rateRawData = df.selectExpr("CAST(timestamp AS STRING)", "CAST(value AS string)")

  val processingTimeStream = rateRawData.writeStream
    .format("console")
    .queryName("Micro Batch")
    .trigger(Trigger.ProcessingTime("1 seconds"))
    .start()

  spark.streams.awaitAnyTermination()
}