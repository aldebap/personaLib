#!	/usr/bin/python

################################################################################
#	configuration.py  -  Feb-02-2019 by aldeba
#
#	PersonaLib configuration class
################################################################################

import	json

#	Configuration attributes class

class Configuration:

	def __init__( self, _dbSource, _dbUser, _dbPassword ):
		self.dbSource = _dbSource
		self.dbUser = _dbUser
		self.dbPassword = _dbPassword

	def __str__( self ):
		return 'dbSource: {0._dbSource!s} : {0._dbUser!s}@{0.dbPassword!s}'.format( self )

	@classmethod
	def serialize( cls, _ref ):
		attributes = {
			'dbSource': _ref.dbSource
			, 'dbUser': _ref.dbUser
			, 'dbPassword': _ref.dbPassword
		}

		return json.dumps( attributes )

	@classmethod
	def unserialize( cls, _stream ):
		attributes = json.loads( _stream )

		dbSource = attributes[ 'dbSource' ]
		dbUser = attributes[ 'dbUser' ]
		dbPassword = attributes[ 'dbPassword' ]

		configurationAux = Configuration( dbSource, dbUser, dbPassword )

		return configurationAux

#	ConfigurationFile  class

class ConfigurationFile:

	def __init__( self, _fileName ):
		self.fileName = _fileName

	def load( self ):
		with open( self.fileName, 'r' ) as fileHandler:
			attributes = json.load( fileHandler )

		return Configuration.unserialize( attributes )

	def save( self, configurationRef ):
		with open( self.fileName, 'w' ) as fileHandler:
			json.dump( Configuration.serialize( configurationRef ), fileHandler )
