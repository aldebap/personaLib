#!	/usr/bin/ksh

#	color constants
RED='\033[0;31m'
GREEN='\033[0;32m'
LIGHTGRAY='\033[0;37m'
NOCOLOR='\033[0m'

#	set environment
export  PROJECT_CONFIG_FILE='build.cfg'
export  TARGETDIR='bin'

#	TODO: create an option to set the target dir

#   check if go toolchain is available
if [ -z "${GOROOT}" ]
then
    exit "error: GOROOT environment variable must be set"
fi

OUTPUT=$( which go 2> /dev/null )
if [ $? -ne 0 ]
then
    exit "error: 'go' compiler is required to run this script"
fi

#   check if jq is available
OUTPUT=$( which jq 2> /dev/null )
if [ $? -ne 0 ]
then
    exit "error: 'jq' is required to run this script"
fi

#	check if config file exists
if [ ! -f ${PROJECT_CONFIG_FILE} ]
then
    exit "error: build configration file not found: ${PROJECT_CONFIG_FILE}"
fi

#	load configuration from config file
export	PROJECT_CONFIG="$( cat ${PROJECT_CONFIG_FILE} )"
export  PROJECT_NAME="$( echo ${PROJECT_CONFIG} | jq '.project' | tr -d '"' )"
export  PROJECT_DESCRIPTION="$( echo ${PROJECT_CONFIG} | jq '.description' | tr -d '"' )"
export  PROJECT_SOURCEDIR="$( echo ${PROJECT_CONFIG} | jq '.source_directory' | tr -d '"' )"
export  PROJECT_SOURCE="$( echo ${PROJECT_CONFIG} | jq '.source | .[]' | tr -d '"' )"
export  PROJECT_PACKAGES="$( echo ${PROJECT_CONFIG} | jq '.packages | .[]' | tr -d '"' )"
export	PROJECT_TARGET="$( echo ${PROJECT_CONFIG} | jq '.target' | tr -d '"' )"
export  PROJECT_DEPENDENCIES="$( echo ${PROJECT_CONFIG} | jq '.dependencies | .[]' | tr -d '"' )"

#	TODO: validate the values from config file

#	TODO: set default values for optional config variables

#	if necessary, create target directory
if [ ! -d "${TARGETDIR}" ]
then
	mkdir "${TARGETDIR}"
fi

#	if necessary, set default target
export	TARGET_LIST="$*"

if [ -z "${TARGET_LIST}" ]
then
	export	TARGET_LIST='compile'
fi

#	TODO: create a better CLI option parser
if [ "${TARGET_LIST}" == 'all' ]
then
	export	TARGET_LIST='dependencies compile test package verify'
fi

#	help message
if [ "${TARGET_LIST}" == 'help' -o "${TARGET_LIST}" == '--help' ]
then
	cat <<HELP_MESSAGE
$( basename ${0} ): specify one or more targets:
	clean: remove all required packages and the target file
	dependencies: install all dependencies
	compile: compile and link source code
	test: execute unit tests
	package: package the target in a container
	verify: execute integration tests
	run: run the target
HELP_MESSAGE

	exit 0
fi

#	perform every specified target
for TARGET in ${TARGET_LIST}
do
	case ${TARGET} in 

	#	remove all required packages and the target file
		clean )
		echo -e "[build] ${GREEN}removing required packages and the target file${NOCOLOR}"

		#rm -rf pkg/*
		#	TODO: is it possible to remove a package installed in GOPATH dir ?
		echo rm -rf "${PROJECT_SOURCEDIR}/go.mod" "${PROJECT_SOURCEDIR}/go.sum"

		for PACKAGE in "${PROJECT_PACKAGES}"
		do
			echo cd "${PACKAGE}"
			echo rm -rf "${PROJECT_SOURCEDIR}/go.mod" "${PROJECT_SOURCEDIR}/go.sum"
		done

		echo rm -f "${TARGETDIR}/${PROJECT_TARGET}"
		;;

	#	install all dependencies
		dependencies )
		echo '+++ downloading and installing all dependencies'

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
		;;

	#	compile and link source code
		compile )
		echo -e "[build] ${GREEN}compiling and linking source code${NOCOLOR}"

		rm -f "${TARGETDIR}/${PROJECT_TARGET}"

		BUILD_DIR="$( pwd )"
		if [ ! -z "${PROJECT_SOURCEDIR}" -a "${PROJECT_SOURCEDIR}" != '.' ]
		then
			cd "${PROJECT_SOURCEDIR}"
		fi
		go build -o "${BUILD_DIR}/${TARGETDIR}/${PROJECT_TARGET}" ${PROJECT_SOURCE}
		cd "${BUILD_DIR}"
		;;

	#	execute unit tests
		test )
		echo -e "[build] ${GREEN}running unit tests${NOCOLOR}"

		export	TEST_RESULT=0

		BUILD_DIR="$( pwd )"
		if [ ! -z "${PROJECT_SOURCEDIR}" -a "${PROJECT_SOURCEDIR}" != '.' ]
		then
			cd "${PROJECT_SOURCEDIR}"
		fi

		#	TODO: the verbose for the unit tests should be a CLI option
		go test -v .
		if [ $? -ne 0 ]
		then
			export	TEST_RESULT=-1
		fi

		for PACKAGE in "${PROJECT_PACKAGES}"
		do
			echo cd "${PACKAGE}"
			cd "${PACKAGE}"
			go test -v .
			if [ $? -ne 0 ]
			then
				export	TEST_RESULT=-1
			fi
			cd ..
		done

		if [ ${TEST_RESULT} -ne 0 ]
		then
			exit ${TEST_RESULT}
		fi

		cd "${BUILD_DIR}"
		;;

	#	package the target in a Docker image
		package )
		echo '+++ package the target in a Docker image'
		#	TODO: adjust the target
		docker build --tag shopping-cart .
		;;

	#	execute integration tests
		verify )
		echo '+++ execute integration tests'

		#	TODO: adjust the target
		docker-compose up -d
		newman run 'test/Integrated Tests.postman_collection.json' --environment 'test/Localhost.postman_environment.json'
		docker-compose stop
		;;

	#	run the target
		run )
		echo '+++ run the target'

		#	TODO: adjust the target
		docker-compose up
		;;

	#	unknown target
		* )
		echo "[error] unknown target ${TARGET}"
		;;
	esac
done
