import fileinput, collections

l = [ ( i, int( v ) ) for i, v in enumerate( fileinput.input() ) ]
d = collections.deque( l )
for i, v in l:
    d.rotate( -d.index( ( i, v ) ) )
    d.popleft()
    d.rotate( -v % len( d ) )
    d.appendleft( ( i, v ) )
d = collections.deque( [ v for k, v in d ] )
d.rotate( -d.index( 0 ) )
print( d[ 1000 ], d[ 2000 ], d[ 3000 ] )
