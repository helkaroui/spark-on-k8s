spark-submit:
	echo "Starting Skaffold using profile: spark-submit"
	skaffold dev -n default -p spark-submit --status-check=false --tail

airflow:
	echo "Starting Skaffold using profile: airflow"
	skaffold dev -n default -p airflow --status-check=false --tail
