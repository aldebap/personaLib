#!	/usr/bin/python

################################################################################
#	book.py  -  Feb-02-2019 by aldebap
#
#	Book entity class
################################################################################

import  json

class Book:

    def __init__( self, _id, _title, _author ):
        self.id = _id
        self.title = _title
        self.author = _author

    @classmethod
    def serialize( cls, _ref ):
        attributes = {
            'id': _ref.id
            , 'title': _ref.title
            , 'author': _ref.author
        }

        return attributes

#    def listAll( self ):

def listAll():
    bookList = []

    bookList.append( Book( "1234", "A medida do universo", [ "Asimov, Isaac" ] ) )
    bookList.append( Book( "2345", "Manual de DevOps", [ "Kim, Gene", "Humble, Jez", "Debois, Patrick", "Willis, John" ] ) )
    bookList.append( Book( "3456", "Probability, Statistics, and Stochastic Processes", [ "Olofsson, Peter" ] ) )

    serializedBookList = []
    for book in bookList:
        serializedBookList.append( Book.serialize( book ) )

    return json.dumps( { 'books': serializedBookList } )
