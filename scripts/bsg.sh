#! /bin/sh
##
## bsg.sh --
##
##	Script for the ball-sim-go program collection.
##
##	Preparation;
##
##	Install the following packages:
##	Ubuntu
##		sudo apt-get install xorg-dev libgl1-mesa-dev
##
## Commands;
##

prg=$(basename $0)
dir=$(readlink -f $(dirname $0))
top=$(readlink -f $dir/..)
me=$dir/$prg
bins="bounce collision"

die() {
	echo "$(date +"%T") ERROR: $*" >&2
	exit 1
}
log() {
	echo "$(date +"%T") $prg: $*" >&2
}
help() {
	grep '^##' $0 | cut -c3-
	exit 0
}
test -n "$1" || help
echo "$1" | grep -qi "^help\|-h" && help

##	build [--clean]
##		Compiles ball-sim-go programs
##
cmd_build() {

	if test "$__clean" = "yes";then
		for b in $bins;do
			rm -f $top/$b
		done
	fi

	for b in $bins;do
		go build -o $top/$b $top/cmd/$b/main.go || die "build failed $b"
		log "build passed $top/$b"
	done

}

##	test
##		Unit test the ball-sim-go programs
##
cmd_test() {

	go test $top/pkg/... || die "test failed"
	log "test passed"

}

##	format
##		Lint and format check
##
cmd_format() {

	golint -set_exit_status $top/... || die "golint failed"
	log "golint passed"
	fmt=$(gofmt -l $top)
	test -z $fmt || die "gofmt failed $fmt"
	log "gofmt passed"

}

##	smoketest
##
##
cmd_smoketest() {

	cmd_build
	cmd_test
	cmd_format

}

# Get the command
cmd=$1
shift
grep -q "^cmd_$cmd()" $0 || die "Invalid command [$cmd]"

while echo "$1" | grep -q '^--'; do
	if echo $1 | grep -q =; then
		o=$(echo "$1" | cut -d= -f1 | sed -e 's,-,_,g')
		v=$(echo "$1" | cut -d= -f2-)
		eval "$o=\"$v\""
	else
		o=$(echo "$1" | sed -e 's,-,_,g')
		eval "$o=yes"
	fi
	shift
done
unset o v
long_opts=`set | grep '^__' | cut -d= -f1`

# Execute command
trap "die Interrupted" INT TERM
cmd_$cmd "$@"
status=$?
rm -rf $tmp
exit $status
