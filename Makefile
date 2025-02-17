# run air to detect any go file changes to re-build and re-run the server
live/server:
	go run github.com/cosmtrek/air@v1.51.0 
# \
# --build.full_bin "go build -o ./tmp/bin/main.exe" --build.bin "./tmp/bin/main.exe" --build.delay "100" \
# --build.exclude_dir "node_modules" \
# --build.include_ext "go" \
# --build.stop_on_error false \
# --misc.clean_on_exit true

# generate db
db/generate:
	go run github.com/steebchen/prisma-client-go generate

db/push:
	go run github.com/steebchen/prisma-client-go db push
	go run github.com/steebchen/prisma-client-go generate



frontend/build:
	cd frontend && pnpm build

frontend/dev:
	cd frontend && pnpm dev

# start all 4 watch processes in parallel
dev:
	make -j2 db/generate live/server

build:
	make frontend/build
	make db/generate
	go generate
	go build