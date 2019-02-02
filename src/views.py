#!	/usr/bin/python

################################################################################
#	views.py  -  Feb-01-2019 by aldebap
#
#	Views for the "personaLib" web server
################################################################################

from django.http import HttpResponse
from django.shortcuts import render

from personaLib.book import listAll

def index( _request ):
    return render( _request, 'personaLib.html', {} )

def bookList( _request ):
    if "GET" == _request.method:
#        return HttpResponse( logoParser( _request.POST[ "script" ] ) )
        return HttpResponse( listAll() )
#        return HttpResponse( "{ \"title\": \"A medida do universo\", \"author\": \"Asimov, Isaac\" }" )

    return HttpResponse( "alert( 'Unsuported method' );" )
