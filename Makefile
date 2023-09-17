batch:
	echo "Starting Skaffold using profile: batch"
	skaffold dev -n default -p batch --status-check=false

dev:
	echo "Running Skaffold in Hot reloading mode"
	skaffold dev -n default -p dev --status-check=false
