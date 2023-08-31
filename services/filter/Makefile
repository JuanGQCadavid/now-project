

local:
	rm -rf tmp \
	&& mkdir tmp \
	&& cp -r ../pkgs tmp/pkgs \
	&& docker compose --env-file ./.env up --detach --build \
	&& rm -rf tmp \
	&& docker-compose logs -f