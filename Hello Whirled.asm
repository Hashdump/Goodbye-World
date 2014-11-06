*****************************************************************
* Caleb Begly							*
* November 5, 2014						*
* Hello World Whirled						*
* This assembly was designed for the Motorola HCS12 uProcessor	*
* but it should work on any processor that supports the CISC 	*
* instruction set.						*
* Adding or removing any single instruction will destroy it	*
* You have been warned.
*****************************************************************

	org	$800
ideas	ds.w	6
	dc.w	$0420	; 16 x The answer to the universe	; 
code	dc.b	"B",$15	; Negative ACK

	org	$A00
	lds	#$E00
	ldy	#ideas	
	jsr	beta
	asla  		; Stay observant
	std	0,Y
	std	0,Y
	swi

try	std	5,+Y	; The most beautiful line in the whole program

catch   psha
	lda	2,SP
	staa	1,Y+
	pula
	exg	D,X
	ldd	2,-X	; Get info from the most beautiful line
	std	2,Y+
	ldd	code	; Because we can
	lsrd		; Not whirled without sideways motion
	lsrd
	lsrd
	lsrd		; The little chips are dizzy now
	stab	1,Y+	; I attempt to store data on hydrocarbon nanorings
	ldx	#ideas	; The time has come for all good coders to burn their bad code. Too bad it's not on paper.
	jsr	$ff5e	; I cry a little when I do this.
	ldd	#done+6
	pshd
	rts
	
	
whyyes	ldx	0,SP
	inx
	ldd	2,X+
	staa	1,Y+	; Catch
	ldd	0,X
	staa	1,Y+
	jsr	alpha
	sts	$2045
	rts

alpha	ldx	0, SP
	ldd	0,X
	eora	#$10	; Chuck Norris tried this and FAILED
	std	2,Y+
	asrb		; When in doubt...
	sts	0,X
	jsr	badcode
	rts

badcode	ldx	0,SP
	ldd	6,-X 
	std	2,Y+
	ldd	#done	; Either this or the other one is a pointer. Doesn't matter which one.
	pshd	
	ldd	#catch  ; The other one.
	pshd	
	rts		; My ECE251 Professor is screaming by now

shift	lsra
	lsra
	lsra
	lsra

beta	ldx	0,SP	; Get address of the tower
	ldaa	0,X	; Identify all units
	staa	1,Y+	; Try
	ldd	#beta   
	stab	1,Y+	
	leas	-2, SP  
done	movw	#whyyes 0,SP ; All good things come to an end. Or do they?
	rts
	swi
	