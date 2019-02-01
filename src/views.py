#!	/usr/bin/python

################################################################################
#	views.py  -  Feb-01-2019 by aldebap
#
#	Views for the "personaLib" web server
################################################################################

from django.http import HttpResponse
from django.shortcuts import render

#from logo.parser import logoParser

def index( _request ):
    return render( _request, 'personaLib.html', {} )

def personaLib( _request ):
    if "POST" == _request.method:
#        return HttpResponse( logoParser( _request.POST[ "script" ] ) )
        return HttpResponse( "alert( 'RestFul API --> JSon' );" )

    return HttpResponse( "alert( 'RestFul API --> JSon' );" )
