CURRENT_DIR=$(shell pwd)

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

exp:
	export DBURL='postgres://postgres:pass@localhost:5432/go11?sslmode=disable'

mig-up:
	migrate -path migrations -database 'postgres://postgres:pass@localhost:5432/go11?sslmode=disable' -verbose up

mig-down:
	migrate -path migrations -database ${DBURL} -verbose down


mig-create-users:
	migrate create -ext sql -dir migrations -seq users

mig-create-user_profiles:
	migrate create -ext sql -dir migrations -seq user_profiles

mig-insert:
	migrate create -ext sql -dir db/migrations -seq insert_table


# mig-delete:
#   rm -r db/migrations