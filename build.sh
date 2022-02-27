#!	/usr/bin/ksh

###############################################################################
#	build a golang project described by a build config file
################################################################################

#	color constants
export	RED='\033[0;31m'
export	GREEN='\033[0;32m'
export	LIGHTGRAY='\033[0;37m'
export	NOCOLOR='\033[0m'

#	set environment

export	PROJECT_CONFIG_FILE='build.cfg'
export	TARGET_DIR='bin'
export	VERBOSE='false'
export	TARGET_LIST=''

#	function to check if the required tools are available
function checkRequiredTools {

	local	TMP_OUTPUT=''

	#	if go toolchain is available
	if [ -z "${GOROOT}" ]
	then
		"error: GOROOT environment variable must be set"
	fi

	TMP_OUTPUT=$( which go 2> /dev/null )
	if [ $? -ne 0 ]
	then
		exit "error: 'go' compiler is required to run this script"
	fi

	#	 if jq is available
	OUTPUT=$( which jq 2> /dev/null )
	if [ $? -ne 0 ]
	then
		exit "error: 'jq' is required to run this script"
	fi
}

#	function to load project configuration from config file
function loadProjectConfig {

	local	PROJECT_CONFIG="$( cat ${PROJECT_CONFIG_FILE} )"

	export	PROJECT_NAME="$( echo ${PROJECT_CONFIG} | jq '.project' | tr -d '"' )"
	export	PROJECT_DESCRIPTION="$( echo ${PROJECT_CONFIG} | jq '.description' | tr -d '"' )"
	export	PROJECT_SOURCE_DIR="$( echo ${PROJECT_CONFIG} | jq '.source_directory' | tr -d '"' )"
	export	PROJECT_SOURCE="$( echo ${PROJECT_CONFIG} | jq '.source | .[]' | tr -d '"' )"
	export	PROJECT_PACKAGES="$( echo ${PROJECT_CONFIG} | jq '.packages | .[]' | tr -d '"' )"
	export	PROJECT_TARGET="$( echo ${PROJECT_CONFIG} | jq '.target' | tr -d '"' )"
	export	PROJECT_DEPENDENCIES="$( echo ${PROJECT_CONFIG} | jq '.dependencies | .[]' | tr -d '"' )"

	#	TODO: validate the values from config file
	#	TODO: set default values for optional config variables
	#	TODO: change the config file to allow specify custom golang compilation options
	#	TODO: change the config file to allow specify the version of dependency packages
}

#	function to execute the "clean" target action
function cleanTarget {

	#	TODO: is it possible to remove a package installed in GOPATH dir ?
	#rm -rf pkg/*
	echo -e '[debug]' ${LIGHTGRAY} rm -rf "${PROJECT_SOURCE_DIR}/go.mod" "${PROJECT_SOURCE_DIR}/go.sum" ${NOCOLOR}

	for PACKAGE in ${PROJECT_PACKAGES}
	do
		echo -e cd "${PACKAGE}"
		echo -e rm -rf "${PROJECT_SOURCE_DIR}/go.mod" "${PROJECT_SOURCE_DIR}/go.sum"
	done

	echo -e rm -f "${TARGET_DIR}/${PROJECT_TARGET}"
}

#	function to execute the "dependencies" target action
function dependenciesTarget {

	#	TODO: adjust the target
	#	main packge dependencies
	cd src
	go mod init aldebap/${PROJECT_NAME}

	for MODULE in "${PROJECT_DEPENDENCIES}"
	do
		go get -u ${MODULE}
	done

	for PACKAGE in "${PROJECT_PACKAGES}"
	do
		echo "require ${PROJECT_NAME}/${PACKAGE} v0.0.0-unpublished" >> go.mod
		echo "replace ${PROJECT_NAME}/${PACKAGE} v0.0.0-unpublished => ../src/${PACKAGE}" >> go.mod
		#echo go mod edit -replace ${PROJECT_NAME}/${PACKAGE}=../src/${PACKAGE}
	done

	cd ..

	#	create project packages module files
	for PACKAGE in ${PROJECT_PACKAGES}
	do
		cd src/${PACKAGE}
		go mod init ${PROJECT_NAME}/${PACKAGE}
		cd ../..
	done
}

#	function to execute the "compile" target action
function compileTarget {

	local	BUILD_DIR="$( pwd )"

	rm -f "${TARGET_DIR}/${PROJECT_TARGET}"

	if [ ! -z "${PROJECT_SOURCE_DIR}" -a "${PROJECT_SOURCE_DIR}" != '.' ]
	then
		cd "${PROJECT_SOURCE_DIR}"
	fi
	go build -o "${BUILD_DIR}/${TARGET_DIR}/${PROJECT_TARGET}" ${PROJECT_SOURCE}
	cd "${BUILD_DIR}"
}

#	function to execute the "test" target action
function testTarget {

	local	BUILD_DIR="$( pwd )"
	local	TEST_RESULT=0
	local	GOTEST_FLAGS=''

	if [ ! -z "${PROJECT_SOURCE_DIR}" -a "${PROJECT_SOURCE_DIR}" != '.' ]
	then
		cd "${PROJECT_SOURCE_DIR}"
	fi

	if [ "${VERBOSE}" == 'true' ]
	then
		GOTEST_FLAGS='-v'
	fi

	go test "${GOTEST_FLAGS}" .
	if [ $? -ne 0 ]
	then
		export	TEST_RESULT=1
	fi

	for PACKAGE in ${PROJECT_PACKAGES}
	do
		echo cd "${PACKAGE}"
		cd "${PACKAGE}"
		go test "${GOTEST_FLAGS}" .
		if [ $? -ne 0 ]
		then
			export	TEST_RESULT=1
		fi
		cd ..
	done

	if [ ${TEST_RESULT} -ne 0 ]
	then
		echo -e "[build] ${RED}unit tests failed${NOCOLOR}"
		exit ${TEST_RESULT}
	fi

	cd "${BUILD_DIR}"
}

#	function to execute the "package" target action
function packageTarget {

	#	TODO: adjust the target
	docker build --tag shopping-cart .
}

#	function to execute the "verify" target action
function verifyTarget {

	#	TODO: adjust the target
	docker-compose up -d
	newman run 'test/Integrated Tests.postman_collection.json' --environment 'test/Localhost.postman_environment.json'
	docker-compose stop
}

#	function to execute the "run" target action
function runTarget {

	#	TODO: adjust the target
	docker-compose up
}

#	CLI arguments parsing

while [ true ]
do
	ARG=${1}

	if [ -z "${ARG}" ]
	then
		break
	fi

	case ${ARG} in

	#	help message option
		--help )
		cat <<HELP_MESSAGE
$( basename ${0} ): [options] targets
options:
	--config-file file-name: set the project config file name (default: ${PROJECT_CONFIG_FILE})
	--target-dir: set the target directory name (default: ${TARGET_DIR})
	--verbose: show detailed information during execution
	--help: show this help message

targets:
	all: the same as the specifying targets "dependencies compile test package verify"
	clean: remove all required packages and the target file
	dependencies: install all dependencies
	compile: compile and link source code
	test: execute unit tests
	package: package the target in a container
	verify: execute integration tests
	run: run the target
HELP_MESSAGE

		exit 0
		;;

	#	config file name option
		--config-file )
		PROJECT_CONFIG_FILE=${2}
		shift
		;;

	#	target directory name option
		--target-dir )
		TARGET_DIR=${2}
		shift
		;;

	#	verbose option
		--verbose )
		VERBOSE='true'
		shift
		;;

	#	if it's not an option, it's a target
		* )
		if [ -z "${TARGET_LIST}" ]
		then
			TARGET_LIST="${ARG}"
		else
			TARGET_LIST="${TARGET_LIST} ${ARG}"
		fi
		;;
	esac

	shift
done

#	check if config file exists
if [ -z "${PROJECT_CONFIG_FILE}" -o ! -f ${PROJECT_CONFIG_FILE} ]
then
	echo -e "[build] ${RED}error: build configration file not found: \"${PROJECT_CONFIG_FILE}\"${NOCOLOR}"
	exit 1
fi

#	if necessary, create target directory
if [ ! -d "${TARGET_DIR}" ]
then
	mkdir "${TARGET_DIR}"
fi

#	if necessary, set default target
if [ -z "${TARGET_LIST}" ]
then
	export	TARGET_LIST='compile'
fi

if [ "${TARGET_LIST}" == 'all' ]
then
	export	TARGET_LIST='dependencies compile test package verify'
fi

#	load project config files and perform every specified target
checkRequiredTools
loadProjectConfig

for TARGET in ${TARGET_LIST}
do
	case ${TARGET} in

		#	remove all required packages and the target file
		clean )
		echo -e "[build] ${TARGET}: ${GREEN}removing required packages and the target file${NOCOLOR}"

		cleanTarget
		;;

		#	install all dependencies
		dependencies )
		echo -e "[build] ${TARGET}: ${GREEN}downloading and installing all dependencies${NOCOLOR}"

		dependenciesTarget
		;;

		#	compile and link source code
		compile )
		echo -e "[build] ${TARGET}: ${GREEN}compiling and linking source code${NOCOLOR}"

		compileTarget
		;;

		#	execute unit tests
		test )
		echo -e "[build] ${TARGET}: ${GREEN}running unit tests${NOCOLOR}"

		testTarget
		;;

		#	package the target in a Docker image
		package )
		echo -e "[build] ${TARGET}: ${GREEN}package the target in a Docker image${NOCOLOR}"

		packageTarget
		;;

		#	execute integration tests
		verify )
		echo -e "[build] ${TARGET}: ${GREEN}execute integration tests${NOCOLOR}"

		verifyTarget
		;;

		#	run the target
		run )
		echo -e "[build] ${TARGET}: ${GREEN}running the project${NOCOLOR}"

		runTarget
		;;

		#	unknown target
		* )
		echo -e "[build] ${RED}error: unknown target ${TARGET}${NOCOLOR}"
		exit 1
		;;
	esac
done
