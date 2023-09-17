package token

type Token_type int

const (
	LEFT_BRACKET Token_type = iota
	RIGHT_BRACKET
	LEFT_CURLY
	RIGHT_CURLY

	MULTiPLY
	SUBRACT
	DIVIDE
	MODULUS
	ADDITION

	AND
	OR

	GREATER_THAN
	GREATER_THAN_OR_EQUAL
	LESS_THAN
	LESS_THAN_OR_EQUAL
	EQUAL
	DOUBLE_EQUAL
	NOT_EQUAL
	NEGATION

	NEWLINE
	TAB
	//

	//Data types
	STRING
	FLOAT
	INTEGER
	TRUE
	FALSE

	//RESERVED KEY WORDS
	IDENTIFIER
	CONST    //done
	VAR      //done
	LET      //done
	BE       //done
	FOR      //done
	WHILE    //done
	BREAK    //done
	CONTINUE //done
	RANGE    //done
	IF       //done
	ELSE     //done
	IF_ELSE
	CASE   //done
	SWITCH //done
	EOL

	STRING_R  //done
	FLOAT_R   //done
	INTEGER_R //done

	EOF
)

type Token struct {
	Type Token_type
	Char string
}
