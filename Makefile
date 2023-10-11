dev:
	echo "Starting Skaffold using profile: dev"
	skaffold dev -n default -p dev --status-check=false --tail

prod:
	echo "Starting Skaffold using profile: prod"
	skaffold dev -n default -p prod --status-check=false --tail
