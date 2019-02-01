#!	/usr/bin/python

################################################################################
#	wsgi.py  -  Feb-01-2019 by aldebap
#
#	WSGI Application for the "personaLib" web server
################################################################################

import os

from django.core.wsgi import get_wsgi_application

os.environ.setdefault( 'DJANGO_SETTINGS_MODULE', 'settings' )

application = get_wsgi_application()
