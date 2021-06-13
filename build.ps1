function buildServer {
	docker build -t twostep -f server-dockerfile .
}

function buildClient {
	docker build -t twostep-test -f client-dockerfile .
}

function build {
	buildServer
	buildClient
}

function test {
	docker compose up
}