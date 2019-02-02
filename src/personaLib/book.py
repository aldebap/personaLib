#!	/usr/bin/python

################################################################################
#	book.py  -  Feb-02-2019 by aldebap
#
#	Book entity class
################################################################################

import  json

class Book:

    def __init__( self, _title, _author ):
        self.title = _title
        self.author = _author

    @classmethod
    def serialize( cls, _ref ):
        attributes = {
            'title': _ref.title
            , 'author': _ref.author
        }

        return json.dumps( attributes )

#    def listAll( self ):

def listAll():
    bookList = []

    bookList.append( Book( "A medida do universo", [ "Asimov, Isaac" ] ) )
    bookList.append( Book( "Manual de DevOps", [ "Kim, Gene", "Humble, Jez", "Debois, Patrick", "Willis, John" ] ) )
    bookList.append( Book( "Probability, Statistics, and Stochastic Processes", [ "Olofsson, Peter" ] ) )

    Book.serialize( bookList[ 0 ] )
