airflow:
	echo "Starting Skaffold using profile: airflow"
	skaffold dev -n default -p airflow --status-check=false --tail

batch:
	echo "Starting Skaffold using profile: batch"
	skaffold dev -n default -p batch --status-check=false --tail

dev:
	echo "Running Skaffold in Hot reloading mode"
	skaffold dev -n default -p dev --status-check=false --tail
