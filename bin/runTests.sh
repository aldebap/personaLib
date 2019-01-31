#!  /usr/bin/ksh

export  CURRENT_DIRECTORY="$( pwd )"
export  BASE_DIR=$( dirname "${CURRENT_DIRECTORY}/$0" )
export  BASE_DIR=$( dirname "${BASE_DIR}" )
export  TEST_DIRECTORY="${BASE_DIR}/test"
export  PYTHONPATH="${BASE_DIR}/src"

cd "${TEST_DIRECTORY}"
python3 -m unittest discover
cd "${CURRENT_DIRECTORY}"
