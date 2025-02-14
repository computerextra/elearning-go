# run templ generation in watch mode to detect all .templ files and
# re-create _templ.txt files on change, then send reload event to browser.
# default ulr: http://localhost:7331
live/templ:
	templ generate --watch --proxy="http://localhost:3000" --open-browser=false

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

# run tailwindcss to generate style.css bundle in watch mode
live/tailwind:
	pnpm tailwindcss -i ./static/css/input.css -o ./static/css/style.css --minify --watch=forever

# watch for any js or css change in the assets/ folder, then reload the browser via templ proxy
live/sync_assets:
	go run github.com/cosmtrek/air@v1.51.0 \
	--build.cmd "templ generate --notify-proxy" \
	--build.bin true \
	--build.delay "100" \
	--build.exclude_dir "" \
	--build.exclude_dir "node_modules" \
	--build.include_dir "static" \
	--build.include_ext "js,css"

# start all 4 watch processes in parallel
dev:
	make -j5 db/generate live/tailwind live/templ live/server live/sync_assets

build:
	go generate
	go build