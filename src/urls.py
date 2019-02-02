#!	/usr/bin/python

################################################################################
#	urls.py  -  Feb-01-2019 by aldebap
#
#	URLs for the "personaLib" web server
################################################################################

from django.urls import path

import  views

urlpatterns = [
    path( 'index/', views.index ),
    path( 'personaLib/book', views.bookList ),
]
