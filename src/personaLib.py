#!	/usr/bin/python

################################################################################
#	personaLib.py  -  Feb-01-2019 by aldebap
#
#	The "personaLib" web server
################################################################################

import os
import sys

if __name__ == '__main__':

    os.environ.setdefault( 'DJANGO_SETTINGS_MODULE', 'settings' )

    try:
        from django.core.management import execute_from_command_line
    except ImportError as exc:
        raise ImportError( "Couldn't import Django. Check if it's installed and available on your PYTHONPATH variable" ) from exc

    execute_from_command_line( sys.argv )
