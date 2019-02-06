#!	/usr/bin/python

################################################################################
#	author.py  -  Feb-06-2019 by aldebap
#
#	Author entity class
################################################################################

import  json

class Author:

    def __init__( self, _id, _name ):
        self.id = _id
        self.name = _name

    @classmethod
    def serialize( cls, _ref ):
        attributes = {
            'id': _ref.id
            , 'name': _ref.name
        }

        return attributes

#    def listAll( self ):

def author_listAll():
    authorList = []

    authorList.append( Author( "1234", "Asimov, Isaac" ) )
    authorList.append( Author( "2345", "Kim, Gene" ) )
    authorList.append( Author( "2345", "Humble, Jez" ) )
    authorList.append( Author( "2345", "Debois, Patrick" ) )
    authorList.append( Author( "2345", "Willis, John" ) )
    authorList.append( Author( "3456", "Olofsson, Peter" ) )

    serializedAuthorList = []
    for author in authorList:
        serializedAuthorList.append( Author.serialize( author ) )

    return json.dumps( { 'authors': serializedAuthorList } )
