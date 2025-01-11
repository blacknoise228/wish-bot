wishql:
	sqlc generate -f ./core/wishbot/db/wish.yaml

shopql:
	sqlc generate -f ./core/shop/db/shop.yaml


.PHONY: wishql shopql