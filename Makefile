coverage:
	cd algorithms && source setup.sh && go get -t -v github.com/stretchr/testify/assert
	cd algorithms && source setup.sh && go get -t -v github.com/jesperancinha/string.distance/algorithms
	cd algorithms && source setup.sh && go test -coverprofile=../coverage.out
