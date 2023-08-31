

local:
	rm -rf tmp \
	&& mkdir tmp \
	&& cp -r ../pkgs tmp/pkgs \
	&& docker compose --env-file ./.env up \
	&& rm -rf tmp