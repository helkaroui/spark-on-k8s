#!/usr/bin/env bash

if [ "$1" = "webserver" ]
then
	exec airflow webserver
	#exec airflow scheduler
fi

if [ "$1" = "scheduler" ]
then
	exec airflow scheduler
fi