//line sql.y:6
package sqlparser

import __yyfmt__ "fmt"

//line sql.y:6
import "bytes"

func SetParseTree(yylex interface{}, stmt Statement) {
	yylex.(*Tokenizer).ParseTree = stmt
}

func SetAllowComments(yylex interface{}, allow bool) {
	yylex.(*Tokenizer).AllowComments = allow
}

func ForceEOF(yylex interface{}) {
	yylex.(*Tokenizer).ForceEOF = true
}

var (
	SHARE        = []byte("share")
	MODE         = []byte("mode")
	IF_BYTES     = []byte("if")
	VALUES_BYTES = []byte("values")
)

//line sql.y:31
type yySymType struct {
	yys         int
	empty       struct{}
	statement   Statement
	selStmt     SelectStatement
	byt         byte
	bytes       []byte
	bytes2      [][]byte
	str         string
	selectExprs SelectExprs
	selectExpr  SelectExpr
	columns     Columns
	colName     *ColName
	tableExprs  TableExprs
	tableExpr   TableExpr
	smTableExpr SimpleTableExpr
	tableName   *TableName
	indexHints  *IndexHints
	expr        Expr
	boolExpr    BoolExpr
	valExpr     ValExpr
	tuple       Tuple
	valExprs    ValExprs
	values      Values
	subquery    *Subquery
	caseExpr    *CaseExpr
	whens       []*When
	when        *When
	orderBy     OrderBy
	order       *Order
	limit       *Limit
	insRows     InsertRows
	updateExprs UpdateExprs
	updateExpr  *UpdateExpr
}

const LEX_ERROR = 57346
const SELECT = 57347
const INSERT = 57348
const UPDATE = 57349
const DELETE = 57350
const FROM = 57351
const WHERE = 57352
const GROUP = 57353
const HAVING = 57354
const ORDER = 57355
const BY = 57356
const LIMIT = 57357
const FOR = 57358
const ALL = 57359
const DISTINCT = 57360
const AS = 57361
const EXISTS = 57362
const IN = 57363
const IS = 57364
const LIKE = 57365
const BETWEEN = 57366
const NULL = 57367
const ASC = 57368
const DESC = 57369
const VALUES = 57370
const INTO = 57371
const DUPLICATE = 57372
const KEY = 57373
const DEFAULT = 57374
const SET = 57375
const LOCK = 57376
const ID = 57377
const STRING = 57378
const NUMBER = 57379
const VALUE_ARG = 57380
const COMMENT = 57381
const LE = 57382
const GE = 57383
const NE = 57384
const NULL_SAFE_EQUAL = 57385
const UNION = 57386
const MINUS = 57387
const EXCEPT = 57388
const INTERSECT = 57389
const JOIN = 57390
const STRAIGHT_JOIN = 57391
const LEFT = 57392
const RIGHT = 57393
const INNER = 57394
const OUTER = 57395
const CROSS = 57396
const NATURAL = 57397
const USE = 57398
const FORCE = 57399
const ON = 57400
const AND = 57401
const OR = 57402
const NOT = 57403
const UNARY = 57404
const CASE = 57405
const WHEN = 57406
const THEN = 57407
const ELSE = 57408
const END = 57409
const BEGIN = 57410
const COMMIT = 57411
const ROLLBACK = 57412
const CREATE = 57413
const ALTER = 57414
const DROP = 57415
const RENAME = 57416
const TABLE = 57417
const INDEX = 57418
const VIEW = 57419
const TO = 57420
const IGNORE = 57421
const IF = 57422
const UNIQUE = 57423
const USING = 57424

var yyToknames = []string{
	"LEX_ERROR",
	"SELECT",
	"INSERT",
	"UPDATE",
	"DELETE",
	"FROM",
	"WHERE",
	"GROUP",
	"HAVING",
	"ORDER",
	"BY",
	"LIMIT",
	"FOR",
	"ALL",
	"DISTINCT",
	"AS",
	"EXISTS",
	"IN",
	"IS",
	"LIKE",
	"BETWEEN",
	"NULL",
	"ASC",
	"DESC",
	"VALUES",
	"INTO",
	"DUPLICATE",
	"KEY",
	"DEFAULT",
	"SET",
	"LOCK",
	"ID",
	"STRING",
	"NUMBER",
	"VALUE_ARG",
	"COMMENT",
	"LE",
	"GE",
	"NE",
	"NULL_SAFE_EQUAL",
	" (",
	" =",
	" <",
	" >",
	" ~",
	"UNION",
	"MINUS",
	"EXCEPT",
	"INTERSECT",
	" ,",
	"JOIN",
	"STRAIGHT_JOIN",
	"LEFT",
	"RIGHT",
	"INNER",
	"OUTER",
	"CROSS",
	"NATURAL",
	"USE",
	"FORCE",
	"ON",
	"AND",
	"OR",
	"NOT",
	" &",
	" |",
	" ^",
	" +",
	" -",
	" *",
	" /",
	" %",
	" .",
	"UNARY",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"END",
	"BEGIN",
	"COMMIT",
	"ROLLBACK",
	"CREATE",
	"ALTER",
	"DROP",
	"RENAME",
	"TABLE",
	"INDEX",
	"VIEW",
	"TO",
	"IGNORE",
	"IF",
	"UNIQUE",
	"USING",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 199
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 556

var yyAct = []int{

	89, 281, 153, 348, 86, 80, 316, 155, 58, 97,
	237, 191, 273, 228, 76, 171, 59, 179, 72, 357,
	87, 64, 75, 156, 3, 244, 248, 249, 250, 251,
	252, 357, 253, 254, 118, 115, 357, 124, 61, 129,
	130, 66, 60, 225, 69, 92, 327, 326, 73, 279,
	96, 49, 325, 102, 81, 28, 29, 30, 31, 65,
	79, 93, 94, 95, 359, 68, 114, 298, 300, 84,
	302, 124, 218, 100, 124, 122, 358, 220, 38, 126,
	40, 356, 306, 43, 41, 44, 307, 157, 45, 152,
	154, 158, 83, 229, 278, 271, 98, 99, 77, 299,
	322, 229, 162, 103, 221, 117, 165, 61, 259, 128,
	61, 60, 175, 174, 60, 169, 268, 111, 107, 266,
	101, 67, 219, 223, 208, 113, 173, 81, 197, 175,
	46, 47, 48, 195, 201, 199, 200, 206, 207, 196,
	210, 211, 212, 213, 214, 215, 216, 217, 176, 109,
	129, 130, 274, 202, 142, 143, 144, 240, 189, 198,
	129, 130, 222, 81, 81, 309, 209, 121, 61, 61,
	109, 220, 60, 235, 333, 55, 233, 292, 239, 185,
	241, 274, 293, 224, 226, 236, 28, 29, 30, 31,
	232, 140, 141, 142, 143, 144, 324, 290, 183, 323,
	296, 186, 291, 195, 295, 258, 245, 261, 262, 294,
	110, 248, 249, 250, 251, 252, 242, 253, 254, 172,
	172, 92, 311, 265, 260, 123, 96, 71, 81, 102,
	105, 15, 343, 108, 194, 272, 62, 93, 94, 95,
	167, 270, 342, 193, 341, 84, 277, 280, 267, 100,
	159, 168, 276, 182, 184, 181, 195, 195, 288, 289,
	67, 194, 246, 109, 163, 161, 305, 160, 83, 124,
	193, 104, 98, 99, 308, 62, 74, 303, 301, 103,
	61, 257, 313, 332, 312, 314, 317, 137, 138, 139,
	140, 141, 142, 143, 144, 127, 101, 256, 137, 138,
	139, 140, 141, 142, 143, 144, 285, 328, 284, 354,
	188, 67, 329, 187, 170, 119, 116, 112, 56, 70,
	318, 106, 330, 54, 222, 310, 338, 355, 340, 339,
	337, 15, 331, 264, 361, 345, 317, 177, 120, 347,
	346, 52, 349, 349, 349, 61, 350, 351, 282, 60,
	15, 92, 50, 352, 231, 321, 96, 283, 362, 102,
	238, 320, 363, 287, 364, 92, 79, 93, 94, 95,
	96, 172, 57, 102, 203, 84, 204, 205, 360, 100,
	62, 93, 94, 95, 344, 15, 33, 14, 15, 84,
	13, 12, 178, 100, 39, 243, 180, 42, 83, 63,
	234, 166, 98, 99, 77, 353, 334, 315, 96, 103,
	319, 102, 83, 286, 269, 164, 98, 99, 62, 93,
	94, 95, 96, 103, 227, 102, 101, 159, 91, 32,
	88, 100, 62, 93, 94, 95, 15, 16, 17, 18,
	101, 159, 335, 336, 90, 100, 34, 35, 36, 37,
	275, 85, 230, 131, 98, 99, 82, 297, 192, 247,
	190, 103, 78, 255, 19, 125, 51, 27, 98, 99,
	53, 11, 10, 9, 8, 103, 7, 6, 101, 132,
	136, 134, 135, 5, 137, 138, 139, 140, 141, 142,
	143, 144, 101, 4, 2, 1, 0, 0, 148, 149,
	150, 151, 0, 145, 146, 147, 0, 0, 0, 0,
	0, 0, 0, 0, 24, 25, 26, 20, 21, 23,
	22, 0, 0, 0, 0, 133, 137, 138, 139, 140,
	141, 142, 143, 144, 304, 0, 0, 137, 138, 139,
	140, 141, 142, 143, 144, 263, 0, 0, 137, 138,
	139, 140, 141, 142, 143, 144,
}
var yyPact = []int{

	431, -1000, -1000, 137, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-12, -9, -2, 40, -1000, -1000, -1000, 380, 335, -1000,
	-1000, -1000, 323, -1000, 294, 283, 363, 240, -74, -32,
	225, -1000, -25, 225, -1000, 284, -77, 225, -77, -1000,
	-1000, 331, -1000, 232, 283, 288, 42, 283, 96, -1000,
	165, -1000, 41, 282, 58, 225, -1000, -1000, 281, -1000,
	-59, 280, 318, 103, 225, 216, -1000, -1000, 276, 33,
	95, 458, -1000, 201, 345, -1000, -1000, -1000, 397, 223,
	221, -1000, 220, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 397, -1000, 207, 240, 279, 361, 240,
	397, 225, -1000, 317, -80, -1000, 166, -1000, 278, -1000,
	-1000, 275, -1000, 199, 331, -1000, -1000, 225, 86, 201,
	201, 397, 206, 353, 397, 397, 99, 397, 397, 397,
	397, 397, 397, 397, 397, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 458, -26, 24, 6, 458, -1000, 383,
	25, 331, -1000, 380, 22, 219, 326, 240, 240, 210,
	-1000, 347, 201, -1000, 219, -1000, -1000, -1000, 93, 225,
	-1000, -68, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	209, 157, 262, 226, 32, -1000, -1000, -1000, -1000, -1000,
	-1000, 219, -1000, 206, 397, 397, 219, 480, -1000, 308,
	120, 120, 120, 81, 81, -1000, -1000, -1000, -1000, -1000,
	397, -1000, 219, -1000, 21, 331, 18, 14, -1000, 201,
	88, 206, 137, 117, -4, -1000, 347, 333, 343, 95,
	273, -1000, -1000, 271, -1000, 352, 199, 199, -1000, -1000,
	143, 123, 155, 150, 146, 5, -1000, 243, -28, 242,
	-1000, 219, 469, 397, -1000, 219, -1000, -16, -1000, 4,
	-1000, 397, 85, -1000, 295, 169, -1000, -1000, -1000, 240,
	333, -1000, 397, 397, -1000, -1000, 349, 341, 157, 36,
	-1000, 145, -1000, 142, -1000, -1000, -1000, -1000, -39, -44,
	-45, -1000, -1000, -1000, 397, 219, -1000, -1000, 219, 397,
	291, 206, -1000, -1000, 230, 121, -1000, 416, -1000, 347,
	201, 397, 201, -1000, -1000, 200, 198, 188, 219, 219,
	377, -1000, 397, 397, -1000, -1000, -1000, 333, 95, 118,
	95, 225, 225, 225, 240, 219, -1000, 293, -17, -1000,
	-22, -34, 96, -1000, 371, 313, -1000, 225, -1000, -1000,
	-1000, 225, -1000, 225, -1000,
}
var yyPgo = []int{

	0, 495, 494, 23, 493, 483, 477, 476, 474, 473,
	472, 471, 429, 470, 467, 466, 22, 14, 465, 463,
	462, 460, 11, 459, 458, 175, 457, 3, 15, 5,
	456, 453, 452, 451, 2, 20, 7, 450, 444, 9,
	430, 4, 428, 424, 13, 415, 414, 413, 410, 10,
	407, 6, 406, 1, 405, 401, 400, 12, 8, 16,
	227, 399, 397, 396, 395, 394, 392, 0, 35, 391,
	390, 387, 386,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 3, 3, 4, 4, 5, 6,
	7, 69, 70, 71, 8, 8, 8, 9, 9, 9,
	10, 11, 11, 11, 72, 12, 13, 13, 14, 14,
	14, 14, 14, 15, 15, 16, 16, 17, 17, 17,
	20, 20, 18, 18, 18, 21, 21, 22, 22, 22,
	22, 19, 19, 19, 23, 23, 23, 23, 23, 23,
	23, 23, 23, 24, 24, 24, 25, 25, 26, 26,
	26, 26, 27, 27, 28, 28, 29, 29, 29, 29,
	29, 30, 30, 30, 30, 30, 30, 30, 30, 30,
	30, 31, 31, 31, 31, 31, 31, 31, 32, 32,
	37, 37, 35, 35, 39, 36, 36, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 38, 38, 40, 40, 40, 42,
	45, 45, 43, 43, 44, 46, 46, 41, 41, 33,
	33, 33, 33, 47, 47, 48, 48, 49, 49, 50,
	50, 51, 52, 52, 52, 53, 53, 53, 54, 54,
	54, 55, 55, 56, 56, 57, 57, 58, 58, 59,
	60, 60, 61, 61, 62, 62, 63, 63, 63, 63,
	63, 64, 64, 65, 65, 66, 66, 67, 68,
}
var yyR2 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 12, 3, 7, 7, 8, 7,
	3, 1, 1, 1, 5, 8, 4, 6, 7, 4,
	5, 4, 5, 5, 0, 2, 0, 2, 1, 2,
	1, 1, 1, 0, 1, 1, 3, 1, 2, 3,
	1, 1, 0, 1, 2, 1, 3, 3, 3, 3,
	5, 0, 1, 2, 1, 1, 2, 3, 2, 3,
	2, 2, 2, 1, 3, 1, 1, 3, 0, 5,
	5, 5, 1, 3, 0, 2, 1, 3, 3, 2,
	3, 3, 3, 4, 3, 4, 5, 6, 3, 4,
	2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 3, 3, 1, 3, 1, 3, 1, 1, 1,
	3, 3, 3, 3, 3, 3, 3, 3, 2, 3,
	4, 5, 4, 1, 1, 1, 1, 1, 1, 5,
	0, 1, 1, 2, 4, 0, 2, 1, 3, 1,
	1, 1, 1, 0, 3, 0, 2, 0, 3, 1,
	3, 2, 0, 1, 1, 0, 2, 4, 0, 2,
	4, 0, 3, 1, 3, 0, 5, 1, 3, 3,
	0, 2, 0, 3, 0, 1, 1, 1, 1, 1,
	1, 0, 1, 0, 1, 0, 2, 1, 0,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -69, -70, -71, 5, 6, 7, 8, 33,
	86, 87, 89, 88, 83, 84, 85, -14, 49, 50,
	51, 52, -12, -72, -12, -12, -12, -12, 90, -65,
	92, 96, -62, 92, 94, 90, 90, 91, 92, -3,
	17, -15, 18, -13, 29, -25, 35, 9, -58, -59,
	-41, -67, 35, -61, 95, 91, -67, 35, 90, -67,
	35, -60, 95, -67, -60, -16, -17, 73, -20, 35,
	-29, -34, -30, 67, 44, -33, -41, -35, -40, -67,
	-38, -42, 20, 36, 37, 38, 25, -39, 71, 72,
	48, 95, 28, 78, 39, -25, 33, 76, -25, 53,
	45, 76, 35, 67, -67, -68, 35, -68, 93, 35,
	20, 64, -67, 9, 53, -18, -67, 19, 76, 65,
	66, -31, 21, 67, 23, 24, 22, 68, 69, 70,
	71, 72, 73, 74, 75, 45, 46, 47, 40, 41,
	42, 43, -29, -34, -29, -36, -3, -34, -34, 44,
	44, 44, -39, 44, -45, -34, -55, 33, 44, -58,
	35, -28, 10, -59, -34, -67, -68, 20, -66, 97,
	-63, 89, 87, 32, 88, 13, 35, 35, 35, -68,
	-21, -22, -24, 44, 35, -39, -17, -67, 73, -29,
	-29, -34, -35, 21, 23, 24, -34, -34, 25, 67,
	-34, -34, -34, -34, -34, -34, -34, -34, 98, 98,
	53, 98, -34, 98, -16, 18, -16, -43, -44, 79,
	-32, 28, -3, -58, -56, -41, -28, -49, 13, -29,
	64, -67, -68, -64, 93, -28, 53, -23, 54, 55,
	56, 57, 58, 60, 61, -19, 35, 19, -22, 76,
	-35, -34, -34, 65, 25, -34, 98, -16, 98, -46,
	-44, 81, -29, -57, 64, -37, -35, -57, 98, 53,
	-49, -53, 15, 14, 35, 35, -47, 11, -22, -22,
	54, 59, 54, 59, 54, 54, 54, -26, 62, 94,
	63, 35, 98, 35, 65, -34, 98, 82, -34, 80,
	30, 53, -41, -53, -34, -50, -51, -34, -68, -48,
	12, 14, 64, 54, 54, 91, 91, 91, -34, -34,
	31, -35, 53, 53, -52, 26, 27, -49, -29, -36,
	-29, 44, 44, 44, 7, -34, -51, -53, -27, -67,
	-27, -27, -58, -54, 16, 34, 98, 53, 98, 98,
	7, 21, -67, -67, -67,
}
var yyDef = []int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 34, 34, 34, 34, 34,
	193, 184, 0, 0, 21, 22, 23, 0, 38, 40,
	41, 42, 43, 36, 0, 0, 0, 0, 182, 0,
	0, 194, 0, 0, 185, 0, 180, 0, 180, 15,
	39, 0, 44, 35, 0, 0, 76, 0, 20, 177,
	0, 147, 197, 0, 0, 0, 198, 197, 0, 198,
	0, 0, 0, 0, 0, 0, 45, 47, 52, 197,
	50, 51, 86, 0, 0, 117, 118, 119, 0, 147,
	0, 133, 0, 149, 150, 151, 152, 113, 136, 137,
	138, 134, 135, 140, 37, 171, 0, 0, 84, 0,
	0, 0, 198, 0, 195, 26, 0, 29, 0, 31,
	181, 0, 198, 0, 0, 48, 53, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 101, 102, 103, 104, 105,
	106, 107, 89, 0, 0, 0, 0, 115, 128, 0,
	0, 0, 100, 0, 0, 141, 0, 0, 0, 84,
	77, 157, 0, 178, 179, 148, 24, 183, 0, 0,
	198, 191, 186, 187, 188, 189, 190, 30, 32, 33,
	84, 55, 61, 0, 73, 75, 46, 54, 49, 87,
	88, 91, 92, 0, 0, 0, 94, 0, 98, 0,
	120, 121, 122, 123, 124, 125, 126, 127, 90, 112,
	0, 114, 115, 129, 0, 0, 0, 145, 142, 0,
	175, 0, 109, 175, 0, 173, 157, 165, 0, 85,
	0, 196, 27, 0, 192, 153, 0, 0, 64, 65,
	0, 0, 0, 0, 0, 78, 62, 0, 0, 0,
	93, 95, 0, 0, 99, 116, 130, 0, 132, 0,
	143, 0, 0, 16, 0, 108, 110, 17, 172, 0,
	165, 19, 0, 0, 198, 28, 155, 0, 56, 59,
	66, 0, 68, 0, 70, 71, 72, 57, 0, 0,
	0, 63, 58, 74, 0, 96, 131, 139, 146, 0,
	0, 0, 174, 18, 166, 158, 159, 162, 25, 157,
	0, 0, 0, 67, 69, 0, 0, 0, 97, 144,
	0, 111, 0, 0, 161, 163, 164, 165, 156, 154,
	60, 0, 0, 0, 0, 167, 160, 168, 0, 82,
	0, 0, 176, 14, 0, 0, 79, 0, 80, 81,
	169, 0, 83, 0, 170,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 75, 68, 3,
	44, 98, 73, 71, 53, 72, 76, 74, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	46, 45, 47, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 70, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 69, 3, 48,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 49, 50, 51, 52, 54, 55, 56, 57,
	58, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	77, 78, 79, 80, 81, 82, 83, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line sql.y:151
		{
			SetParseTree(yylex, yyS[yypt-0].statement)
		}
	case 2:
		//line sql.y:157
		{
			yyVAL.statement = yyS[yypt-0].selStmt
		}
	case 3:
		yyVAL.statement = yyS[yypt-0].statement
	case 4:
		yyVAL.statement = yyS[yypt-0].statement
	case 5:
		yyVAL.statement = yyS[yypt-0].statement
	case 6:
		yyVAL.statement = yyS[yypt-0].statement
	case 7:
		yyVAL.statement = yyS[yypt-0].statement
	case 8:
		yyVAL.statement = yyS[yypt-0].statement
	case 9:
		yyVAL.statement = yyS[yypt-0].statement
	case 10:
		yyVAL.statement = yyS[yypt-0].statement
	case 11:
		yyVAL.statement = yyS[yypt-0].statement
	case 12:
		yyVAL.statement = yyS[yypt-0].statement
	case 13:
		yyVAL.statement = yyS[yypt-0].statement
	case 14:
		//line sql.y:174
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyS[yypt-10].bytes2), Distinct: yyS[yypt-9].str, SelectExprs: yyS[yypt-8].selectExprs, From: yyS[yypt-6].tableExprs, Where: NewWhere(AST_WHERE, yyS[yypt-5].boolExpr), GroupBy: GroupBy(yyS[yypt-4].valExprs), Having: NewWhere(AST_HAVING, yyS[yypt-3].boolExpr), OrderBy: yyS[yypt-2].orderBy, Limit: yyS[yypt-1].limit, Lock: yyS[yypt-0].str}
		}
	case 15:
		//line sql.y:178
		{
			yyVAL.selStmt = &Union{Type: yyS[yypt-1].str, Left: yyS[yypt-2].selStmt, Right: yyS[yypt-0].selStmt}
		}
	case 16:
		//line sql.y:185
		{
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: yyS[yypt-2].columns, Rows: yyS[yypt-1].insRows, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 17:
		//line sql.y:189
		{
			cols := make(Columns, 0, len(yyS[yypt-1].updateExprs))
			vals := make(ValTuple, 0, len(yyS[yypt-1].updateExprs))
			for _, col := range yyS[yypt-1].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 18:
		//line sql.y:201
		{
			yyVAL.statement = &Update{Comments: Comments(yyS[yypt-6].bytes2), Table: yyS[yypt-5].tableName, Exprs: yyS[yypt-3].updateExprs, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 19:
		//line sql.y:207
		{
			yyVAL.statement = &Delete{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 20:
		//line sql.y:213
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-1].bytes2), Exprs: yyS[yypt-0].updateExprs}
		}
	case 21:
		//line sql.y:219
		{
			yyVAL.statement = &Begin{}
		}
	case 22:
		//line sql.y:225
		{
			yyVAL.statement = &Commit{}
		}
	case 23:
		//line sql.y:231
		{
			yyVAL.statement = &Rollback{}
		}
	case 24:
		//line sql.y:237
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 25:
		//line sql.y:241
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 26:
		//line sql.y:246
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 27:
		//line sql.y:252
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-2].bytes}
		}
	case 28:
		//line sql.y:256
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-3].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 29:
		//line sql.y:261
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 30:
		//line sql.y:267
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 31:
		//line sql.y:273
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-0].bytes}
		}
	case 32:
		//line sql.y:277
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 33:
		//line sql.y:282
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-1].bytes}
		}
	case 34:
		//line sql.y:287
		{
			SetAllowComments(yylex, true)
		}
	case 35:
		//line sql.y:291
		{
			yyVAL.bytes2 = yyS[yypt-0].bytes2
			SetAllowComments(yylex, false)
		}
	case 36:
		//line sql.y:297
		{
			yyVAL.bytes2 = nil
		}
	case 37:
		//line sql.y:301
		{
			yyVAL.bytes2 = append(yyS[yypt-1].bytes2, yyS[yypt-0].bytes)
		}
	case 38:
		//line sql.y:307
		{
			yyVAL.str = AST_UNION
		}
	case 39:
		//line sql.y:311
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 40:
		//line sql.y:315
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 41:
		//line sql.y:319
		{
			yyVAL.str = AST_EXCEPT
		}
	case 42:
		//line sql.y:323
		{
			yyVAL.str = AST_INTERSECT
		}
	case 43:
		//line sql.y:328
		{
			yyVAL.str = ""
		}
	case 44:
		//line sql.y:332
		{
			yyVAL.str = AST_DISTINCT
		}
	case 45:
		//line sql.y:338
		{
			yyVAL.selectExprs = SelectExprs{yyS[yypt-0].selectExpr}
		}
	case 46:
		//line sql.y:342
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyS[yypt-0].selectExpr)
		}
	case 47:
		//line sql.y:348
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 48:
		//line sql.y:352
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyS[yypt-1].expr, As: yyS[yypt-0].bytes}
		}
	case 49:
		//line sql.y:356
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyS[yypt-2].bytes}
		}
	case 50:
		//line sql.y:362
		{
			yyVAL.expr = yyS[yypt-0].boolExpr
		}
	case 51:
		//line sql.y:366
		{
			yyVAL.expr = yyS[yypt-0].valExpr
		}
	case 52:
		//line sql.y:371
		{
			yyVAL.bytes = nil
		}
	case 53:
		//line sql.y:375
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 54:
		//line sql.y:379
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 55:
		//line sql.y:385
		{
			yyVAL.tableExprs = TableExprs{yyS[yypt-0].tableExpr}
		}
	case 56:
		//line sql.y:389
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyS[yypt-0].tableExpr)
		}
	case 57:
		//line sql.y:395
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyS[yypt-2].smTableExpr, As: yyS[yypt-1].bytes, Hints: yyS[yypt-0].indexHints}
		}
	case 58:
		//line sql.y:399
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyS[yypt-1].tableExpr}
		}
	case 59:
		//line sql.y:403
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-2].tableExpr, Join: yyS[yypt-1].str, RightExpr: yyS[yypt-0].tableExpr}
		}
	case 60:
		//line sql.y:407
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-4].tableExpr, Join: yyS[yypt-3].str, RightExpr: yyS[yypt-2].tableExpr, On: yyS[yypt-0].boolExpr}
		}
	case 61:
		//line sql.y:412
		{
			yyVAL.bytes = nil
		}
	case 62:
		//line sql.y:416
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 63:
		//line sql.y:420
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 64:
		//line sql.y:426
		{
			yyVAL.str = AST_JOIN
		}
	case 65:
		//line sql.y:430
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 66:
		//line sql.y:434
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 67:
		//line sql.y:438
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 68:
		//line sql.y:442
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 69:
		//line sql.y:446
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 70:
		//line sql.y:450
		{
			yyVAL.str = AST_JOIN
		}
	case 71:
		//line sql.y:454
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 72:
		//line sql.y:458
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 73:
		//line sql.y:464
		{
			yyVAL.smTableExpr = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 74:
		//line sql.y:468
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 75:
		//line sql.y:472
		{
			yyVAL.smTableExpr = yyS[yypt-0].subquery
		}
	case 76:
		//line sql.y:478
		{
			yyVAL.tableName = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 77:
		//line sql.y:482
		{
			yyVAL.tableName = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 78:
		//line sql.y:487
		{
			yyVAL.indexHints = nil
		}
	case 79:
		//line sql.y:491
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyS[yypt-1].bytes2}
		}
	case 80:
		//line sql.y:495
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyS[yypt-1].bytes2}
		}
	case 81:
		//line sql.y:499
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyS[yypt-1].bytes2}
		}
	case 82:
		//line sql.y:505
		{
			yyVAL.bytes2 = [][]byte{yyS[yypt-0].bytes}
		}
	case 83:
		//line sql.y:509
		{
			yyVAL.bytes2 = append(yyS[yypt-2].bytes2, yyS[yypt-0].bytes)
		}
	case 84:
		//line sql.y:514
		{
			yyVAL.boolExpr = nil
		}
	case 85:
		//line sql.y:518
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 86:
		yyVAL.boolExpr = yyS[yypt-0].boolExpr
	case 87:
		//line sql.y:525
		{
			yyVAL.boolExpr = &AndExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 88:
		//line sql.y:529
		{
			yyVAL.boolExpr = &OrExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 89:
		//line sql.y:533
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyS[yypt-0].boolExpr}
		}
	case 90:
		//line sql.y:537
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyS[yypt-1].boolExpr}
		}
	case 91:
		//line sql.y:543
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: yyS[yypt-1].str, Right: yyS[yypt-0].valExpr}
		}
	case 92:
		//line sql.y:547
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_IN, Right: yyS[yypt-0].tuple}
		}
	case 93:
		//line sql.y:551
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_IN, Right: yyS[yypt-0].tuple}
		}
	case 94:
		//line sql.y:555
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 95:
		//line sql.y:559
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 96:
		//line sql.y:563
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-4].valExpr, Operator: AST_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 97:
		//line sql.y:567
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-5].valExpr, Operator: AST_NOT_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 98:
		//line sql.y:571
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyS[yypt-2].valExpr}
		}
	case 99:
		//line sql.y:575
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyS[yypt-3].valExpr}
		}
	case 100:
		//line sql.y:579
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyS[yypt-0].subquery}
		}
	case 101:
		//line sql.y:585
		{
			yyVAL.str = AST_EQ
		}
	case 102:
		//line sql.y:589
		{
			yyVAL.str = AST_LT
		}
	case 103:
		//line sql.y:593
		{
			yyVAL.str = AST_GT
		}
	case 104:
		//line sql.y:597
		{
			yyVAL.str = AST_LE
		}
	case 105:
		//line sql.y:601
		{
			yyVAL.str = AST_GE
		}
	case 106:
		//line sql.y:605
		{
			yyVAL.str = AST_NE
		}
	case 107:
		//line sql.y:609
		{
			yyVAL.str = AST_NSE
		}
	case 108:
		//line sql.y:615
		{
			yyVAL.insRows = yyS[yypt-0].values
		}
	case 109:
		//line sql.y:619
		{
			yyVAL.insRows = yyS[yypt-0].selStmt
		}
	case 110:
		//line sql.y:625
		{
			yyVAL.values = Values{yyS[yypt-0].tuple}
		}
	case 111:
		//line sql.y:629
		{
			yyVAL.values = append(yyS[yypt-2].values, yyS[yypt-0].tuple)
		}
	case 112:
		//line sql.y:635
		{
			yyVAL.tuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 113:
		//line sql.y:639
		{
			yyVAL.tuple = yyS[yypt-0].subquery
		}
	case 114:
		//line sql.y:645
		{
			yyVAL.subquery = &Subquery{yyS[yypt-1].selStmt}
		}
	case 115:
		//line sql.y:651
		{
			yyVAL.valExprs = ValExprs{yyS[yypt-0].valExpr}
		}
	case 116:
		//line sql.y:655
		{
			yyVAL.valExprs = append(yyS[yypt-2].valExprs, yyS[yypt-0].valExpr)
		}
	case 117:
		//line sql.y:661
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 118:
		//line sql.y:665
		{
			yyVAL.valExpr = yyS[yypt-0].colName
		}
	case 119:
		//line sql.y:669
		{
			yyVAL.valExpr = yyS[yypt-0].tuple
		}
	case 120:
		//line sql.y:673
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITAND, Right: yyS[yypt-0].valExpr}
		}
	case 121:
		//line sql.y:677
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITOR, Right: yyS[yypt-0].valExpr}
		}
	case 122:
		//line sql.y:681
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITXOR, Right: yyS[yypt-0].valExpr}
		}
	case 123:
		//line sql.y:685
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_PLUS, Right: yyS[yypt-0].valExpr}
		}
	case 124:
		//line sql.y:689
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MINUS, Right: yyS[yypt-0].valExpr}
		}
	case 125:
		//line sql.y:693
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MULT, Right: yyS[yypt-0].valExpr}
		}
	case 126:
		//line sql.y:697
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_DIV, Right: yyS[yypt-0].valExpr}
		}
	case 127:
		//line sql.y:701
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MOD, Right: yyS[yypt-0].valExpr}
		}
	case 128:
		//line sql.y:705
		{
			if num, ok := yyS[yypt-0].valExpr.(NumVal); ok {
				switch yyS[yypt-1].byt {
				case '-':
					yyVAL.valExpr = append(NumVal("-"), num...)
				case '+':
					yyVAL.valExpr = num
				default:
					yyVAL.valExpr = &UnaryExpr{Operator: yyS[yypt-1].byt, Expr: yyS[yypt-0].valExpr}
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: yyS[yypt-1].byt, Expr: yyS[yypt-0].valExpr}
			}
		}
	case 129:
		//line sql.y:720
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-2].bytes}
		}
	case 130:
		//line sql.y:724
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 131:
		//line sql.y:728
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-4].bytes, Distinct: true, Exprs: yyS[yypt-1].selectExprs}
		}
	case 132:
		//line sql.y:732
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 133:
		//line sql.y:736
		{
			yyVAL.valExpr = yyS[yypt-0].caseExpr
		}
	case 134:
		//line sql.y:742
		{
			yyVAL.bytes = IF_BYTES
		}
	case 135:
		//line sql.y:746
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 136:
		//line sql.y:752
		{
			yyVAL.byt = AST_UPLUS
		}
	case 137:
		//line sql.y:756
		{
			yyVAL.byt = AST_UMINUS
		}
	case 138:
		//line sql.y:760
		{
			yyVAL.byt = AST_TILDA
		}
	case 139:
		//line sql.y:766
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyS[yypt-3].valExpr, Whens: yyS[yypt-2].whens, Else: yyS[yypt-1].valExpr}
		}
	case 140:
		//line sql.y:771
		{
			yyVAL.valExpr = nil
		}
	case 141:
		//line sql.y:775
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 142:
		//line sql.y:781
		{
			yyVAL.whens = []*When{yyS[yypt-0].when}
		}
	case 143:
		//line sql.y:785
		{
			yyVAL.whens = append(yyS[yypt-1].whens, yyS[yypt-0].when)
		}
	case 144:
		//line sql.y:791
		{
			yyVAL.when = &When{Cond: yyS[yypt-2].boolExpr, Val: yyS[yypt-0].valExpr}
		}
	case 145:
		//line sql.y:796
		{
			yyVAL.valExpr = nil
		}
	case 146:
		//line sql.y:800
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 147:
		//line sql.y:806
		{
			yyVAL.colName = &ColName{Name: yyS[yypt-0].bytes}
		}
	case 148:
		//line sql.y:810
		{
			yyVAL.colName = &ColName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 149:
		//line sql.y:816
		{
			yyVAL.valExpr = StrVal(yyS[yypt-0].bytes)
		}
	case 150:
		//line sql.y:820
		{
			yyVAL.valExpr = NumVal(yyS[yypt-0].bytes)
		}
	case 151:
		//line sql.y:824
		{
			yyVAL.valExpr = ValArg(yyS[yypt-0].bytes)
		}
	case 152:
		//line sql.y:828
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 153:
		//line sql.y:833
		{
			yyVAL.valExprs = nil
		}
	case 154:
		//line sql.y:837
		{
			yyVAL.valExprs = yyS[yypt-0].valExprs
		}
	case 155:
		//line sql.y:842
		{
			yyVAL.boolExpr = nil
		}
	case 156:
		//line sql.y:846
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 157:
		//line sql.y:851
		{
			yyVAL.orderBy = nil
		}
	case 158:
		//line sql.y:855
		{
			yyVAL.orderBy = yyS[yypt-0].orderBy
		}
	case 159:
		//line sql.y:861
		{
			yyVAL.orderBy = OrderBy{yyS[yypt-0].order}
		}
	case 160:
		//line sql.y:865
		{
			yyVAL.orderBy = append(yyS[yypt-2].orderBy, yyS[yypt-0].order)
		}
	case 161:
		//line sql.y:871
		{
			yyVAL.order = &Order{Expr: yyS[yypt-1].valExpr, Direction: yyS[yypt-0].str}
		}
	case 162:
		//line sql.y:876
		{
			yyVAL.str = AST_ASC
		}
	case 163:
		//line sql.y:880
		{
			yyVAL.str = AST_ASC
		}
	case 164:
		//line sql.y:884
		{
			yyVAL.str = AST_DESC
		}
	case 165:
		//line sql.y:889
		{
			yyVAL.limit = nil
		}
	case 166:
		//line sql.y:893
		{
			yyVAL.limit = &Limit{Rowcount: yyS[yypt-0].valExpr}
		}
	case 167:
		//line sql.y:897
		{
			yyVAL.limit = &Limit{Offset: yyS[yypt-2].valExpr, Rowcount: yyS[yypt-0].valExpr}
		}
	case 168:
		//line sql.y:902
		{
			yyVAL.str = ""
		}
	case 169:
		//line sql.y:906
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 170:
		//line sql.y:910
		{
			if !bytes.Equal(yyS[yypt-1].bytes, SHARE) {
				yylex.Error("expecting share")
				return 1
			}
			if !bytes.Equal(yyS[yypt-0].bytes, MODE) {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = AST_SHARE_MODE
		}
	case 171:
		//line sql.y:923
		{
			yyVAL.columns = nil
		}
	case 172:
		//line sql.y:927
		{
			yyVAL.columns = yyS[yypt-1].columns
		}
	case 173:
		//line sql.y:933
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyS[yypt-0].colName}}
		}
	case 174:
		//line sql.y:937
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyS[yypt-0].colName})
		}
	case 175:
		//line sql.y:942
		{
			yyVAL.updateExprs = nil
		}
	case 176:
		//line sql.y:946
		{
			yyVAL.updateExprs = yyS[yypt-0].updateExprs
		}
	case 177:
		//line sql.y:952
		{
			yyVAL.updateExprs = UpdateExprs{yyS[yypt-0].updateExpr}
		}
	case 178:
		//line sql.y:956
		{
			yyVAL.updateExprs = append(yyS[yypt-2].updateExprs, yyS[yypt-0].updateExpr)
		}
	case 179:
		//line sql.y:962
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyS[yypt-2].colName, Expr: yyS[yypt-0].valExpr}
		}
	case 180:
		//line sql.y:967
		{
			yyVAL.empty = struct{}{}
		}
	case 181:
		//line sql.y:969
		{
			yyVAL.empty = struct{}{}
		}
	case 182:
		//line sql.y:972
		{
			yyVAL.empty = struct{}{}
		}
	case 183:
		//line sql.y:974
		{
			yyVAL.empty = struct{}{}
		}
	case 184:
		//line sql.y:977
		{
			yyVAL.empty = struct{}{}
		}
	case 185:
		//line sql.y:979
		{
			yyVAL.empty = struct{}{}
		}
	case 186:
		//line sql.y:983
		{
			yyVAL.empty = struct{}{}
		}
	case 187:
		//line sql.y:985
		{
			yyVAL.empty = struct{}{}
		}
	case 188:
		//line sql.y:987
		{
			yyVAL.empty = struct{}{}
		}
	case 189:
		//line sql.y:989
		{
			yyVAL.empty = struct{}{}
		}
	case 190:
		//line sql.y:991
		{
			yyVAL.empty = struct{}{}
		}
	case 191:
		//line sql.y:994
		{
			yyVAL.empty = struct{}{}
		}
	case 192:
		//line sql.y:996
		{
			yyVAL.empty = struct{}{}
		}
	case 193:
		//line sql.y:999
		{
			yyVAL.empty = struct{}{}
		}
	case 194:
		//line sql.y:1001
		{
			yyVAL.empty = struct{}{}
		}
	case 195:
		//line sql.y:1004
		{
			yyVAL.empty = struct{}{}
		}
	case 196:
		//line sql.y:1006
		{
			yyVAL.empty = struct{}{}
		}
	case 197:
		//line sql.y:1010
		{
			yyVAL.bytes = bytes.ToLower(yyS[yypt-0].bytes)
		}
	case 198:
		//line sql.y:1015
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
