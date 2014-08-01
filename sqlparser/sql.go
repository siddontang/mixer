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
const NAMES = 57413
const CREATE = 57414
const ALTER = 57415
const DROP = 57416
const RENAME = 57417
const TABLE = 57418
const INDEX = 57419
const VIEW = 57420
const TO = 57421
const IGNORE = 57422
const IF = 57423
const UNIQUE = 57424
const USING = 57425

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
	"NAMES",
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

const yyNprod = 200
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 540

var yyAct = []int{

	90, 283, 156, 350, 58, 158, 318, 118, 239, 87,
	275, 98, 230, 59, 77, 194, 182, 88, 159, 3,
	173, 76, 132, 133, 73, 250, 251, 252, 253, 254,
	15, 255, 256, 28, 29, 30, 31, 65, 62, 359,
	359, 67, 81, 359, 70, 127, 49, 60, 74, 281,
	97, 329, 127, 103, 82, 127, 221, 43, 223, 44,
	63, 94, 95, 96, 112, 328, 38, 117, 40, 113,
	304, 246, 41, 101, 121, 327, 125, 66, 120, 69,
	129, 68, 45, 224, 116, 361, 360, 188, 160, 358,
	309, 308, 161, 300, 302, 280, 99, 100, 270, 231,
	63, 268, 261, 104, 222, 164, 186, 167, 62, 189,
	131, 62, 171, 114, 176, 178, 177, 60, 108, 201,
	60, 102, 324, 179, 175, 276, 301, 155, 157, 326,
	82, 200, 178, 192, 46, 47, 48, 204, 198, 334,
	209, 210, 199, 213, 214, 215, 216, 217, 218, 219,
	220, 61, 242, 205, 140, 141, 142, 143, 144, 145,
	146, 147, 185, 187, 184, 82, 82, 132, 133, 124,
	62, 62, 325, 211, 235, 202, 203, 132, 133, 60,
	237, 298, 311, 243, 226, 228, 97, 234, 231, 103,
	273, 244, 238, 145, 146, 147, 63, 94, 95, 96,
	143, 144, 145, 146, 147, 113, 297, 296, 198, 101,
	263, 264, 260, 110, 247, 212, 110, 241, 227, 174,
	93, 223, 55, 174, 262, 97, 267, 276, 103, 294,
	82, 126, 99, 100, 295, 80, 94, 95, 96, 104,
	292, 335, 272, 313, 85, 293, 279, 282, 101, 269,
	111, 278, 28, 29, 30, 31, 15, 102, 337, 338,
	198, 198, 248, 345, 290, 291, 110, 84, 307, 344,
	343, 99, 100, 78, 274, 127, 310, 106, 104, 72,
	109, 197, 62, 113, 315, 165, 197, 316, 319, 163,
	196, 314, 162, 105, 320, 196, 102, 169, 68, 225,
	140, 141, 142, 143, 144, 145, 146, 147, 170, 330,
	63, 259, 306, 305, 331, 140, 141, 142, 143, 144,
	145, 146, 147, 130, 303, 356, 177, 258, 75, 341,
	339, 333, 287, 286, 191, 190, 172, 347, 319, 68,
	122, 349, 348, 357, 351, 351, 351, 62, 352, 353,
	119, 354, 115, 93, 56, 71, 60, 107, 97, 15,
	364, 103, 332, 15, 365, 340, 366, 342, 80, 94,
	95, 96, 312, 54, 93, 266, 32, 85, 363, 97,
	180, 101, 103, 123, 52, 206, 233, 207, 208, 63,
	94, 95, 96, 34, 35, 36, 37, 50, 85, 284,
	84, 323, 101, 285, 99, 100, 78, 240, 322, 289,
	174, 104, 57, 362, 15, 16, 17, 18, 346, 93,
	15, 84, 33, 14, 97, 99, 100, 103, 13, 102,
	12, 181, 104, 39, 63, 94, 95, 96, 245, 183,
	42, 64, 19, 85, 236, 168, 355, 101, 336, 317,
	102, 321, 288, 271, 135, 139, 137, 138, 140, 141,
	142, 143, 144, 145, 146, 147, 84, 166, 229, 92,
	99, 100, 89, 151, 152, 153, 154, 104, 148, 149,
	150, 265, 91, 277, 140, 141, 142, 143, 144, 145,
	146, 147, 24, 25, 26, 102, 20, 21, 23, 22,
	136, 140, 141, 142, 143, 144, 145, 146, 147, 250,
	251, 252, 253, 254, 86, 255, 256, 232, 134, 83,
	299, 195, 249, 193, 79, 257, 128, 51, 27, 53,
	11, 10, 9, 8, 7, 6, 5, 4, 2, 1,
}
var yyPact = []int{

	409, -1000, -1000, 203, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-25, -36, -9, 43, -1000, -1000, -1000, 415, 380, -1000,
	-1000, -1000, 366, -1000, 344, 319, 403, 65, -59, -15,
	263, -1000, -12, 263, -1000, 320, -72, 263, -72, -1000,
	-1000, 333, -1000, 254, 319, 324, 42, 319, 160, -1000,
	205, 161, -1000, 37, 317, 17, 263, -1000, -1000, 315,
	-1000, -20, 305, 363, 105, 263, 222, -1000, -1000, 304,
	34, 112, 433, -1000, 399, 354, -1000, -1000, -1000, 161,
	248, 245, -1000, 241, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 161, -1000, 264, 65, 301, 400,
	65, 161, 390, 25, 263, -1000, 360, -82, -1000, 74,
	-1000, 300, -1000, -1000, 299, -1000, 246, 333, -1000, -1000,
	263, 46, 399, 399, 161, 239, 364, 161, 161, 148,
	161, 161, 161, 161, 161, 161, 161, 161, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 433, -43, 5, -16,
	433, -1000, 200, 333, -1000, 415, 20, 390, 358, 65,
	275, 213, -1000, 394, 399, -1000, 390, 390, -1000, -1000,
	-1000, 88, 263, -1000, -23, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 209, 455, 292, 251, 26, -1000, -1000,
	-1000, -1000, -1000, -1000, 390, -1000, 239, 161, 161, 390,
	416, -1000, 350, 129, 129, 129, 120, 120, -1000, -1000,
	-1000, -1000, -1000, 161, -1000, -1000, 2, 333, -1, 109,
	-1000, 399, 61, 239, 203, 163, -4, -1000, 394, 384,
	389, 112, 298, -1000, -1000, 297, -1000, 398, 246, 246,
	-1000, -1000, 186, 175, 153, 152, 127, 31, -1000, 289,
	-29, 278, -1000, 390, 247, 161, -1000, 390, -1000, -8,
	-1000, 8, -1000, 161, 102, -1000, 342, 190, -1000, -1000,
	-1000, 275, 384, -1000, 161, 161, -1000, -1000, 396, 387,
	455, 58, -1000, 118, -1000, 75, -1000, -1000, -1000, -1000,
	-17, -27, -41, -1000, -1000, -1000, 161, 390, -1000, -1000,
	390, 161, 331, 239, -1000, -1000, 86, 188, -1000, 232,
	-1000, 394, 399, 161, 399, -1000, -1000, 226, 225, 219,
	390, 390, 411, -1000, 161, 161, -1000, -1000, -1000, 384,
	112, 168, 112, 263, 263, 263, 65, 390, -1000, 309,
	-10, -1000, -13, -14, 160, -1000, 406, 357, -1000, 263,
	-1000, -1000, -1000, 263, -1000, 263, -1000,
}
var yyPgo = []int{

	0, 539, 538, 18, 537, 536, 535, 534, 533, 532,
	531, 530, 376, 529, 528, 527, 21, 14, 526, 525,
	524, 523, 15, 522, 521, 222, 520, 3, 20, 42,
	519, 518, 517, 514, 2, 17, 5, 483, 482, 11,
	472, 9, 469, 468, 12, 467, 453, 452, 451, 8,
	449, 6, 448, 1, 446, 445, 444, 10, 4, 13,
	279, 441, 440, 439, 438, 433, 431, 0, 7, 430,
	428, 423, 422,
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
	59, 60, 60, 61, 61, 62, 62, 63, 63, 63,
	63, 63, 64, 64, 65, 65, 66, 66, 67, 68,
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
	2, 0, 2, 0, 3, 0, 1, 1, 1, 1,
	1, 1, 0, 1, 0, 1, 0, 2, 1, 0,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -69, -70, -71, 5, 6, 7, 8, 33,
	87, 88, 90, 89, 83, 84, 85, -14, 49, 50,
	51, 52, -12, -72, -12, -12, -12, -12, 91, -65,
	93, 97, -62, 93, 95, 91, 91, 92, 93, -3,
	17, -15, 18, -13, 29, -25, 35, 9, -58, -59,
	-41, 86, -67, 35, -61, 96, 92, -67, 35, 91,
	-67, 35, -60, 96, -67, -60, -16, -17, 73, -20,
	35, -29, -34, -30, 67, 44, -33, -41, -35, -40,
	-67, -38, -42, 20, 36, 37, 38, 25, -39, 71,
	72, 48, 96, 28, 78, 39, -25, 33, 76, -25,
	53, 45, -34, 44, 76, 35, 67, -67, -68, 35,
	-68, 94, 35, 20, 64, -67, 9, 53, -18, -67,
	19, 76, 65, 66, -31, 21, 67, 23, 24, 22,
	68, 69, 70, 71, 72, 73, 74, 75, 45, 46,
	47, 40, 41, 42, 43, -29, -34, -29, -36, -3,
	-34, -34, 44, 44, -39, 44, -45, -34, -55, 33,
	44, -58, 35, -28, 10, -59, -34, -34, -67, -68,
	20, -66, 98, -63, 90, 88, 32, 89, 13, 35,
	35, 35, -68, -21, -22, -24, 44, 35, -39, -17,
	-67, 73, -29, -29, -34, -35, 21, 23, 24, -34,
	-34, 25, 67, -34, -34, -34, -34, -34, -34, -34,
	-34, 99, 99, 53, 99, 99, -16, 18, -16, -43,
	-44, 79, -32, 28, -3, -58, -56, -41, -28, -49,
	13, -29, 64, -67, -68, -64, 94, -28, 53, -23,
	54, 55, 56, 57, 58, 60, 61, -19, 35, 19,
	-22, 76, -35, -34, -34, 65, 25, -34, 99, -16,
	99, -46, -44, 81, -29, -57, 64, -37, -35, -57,
	99, 53, -49, -53, 15, 14, 35, 35, -47, 11,
	-22, -22, 54, 59, 54, 59, 54, 54, 54, -26,
	62, 95, 63, 35, 99, 35, 65, -34, 99, 82,
	-34, 80, 30, 53, -41, -53, -34, -50, -51, -34,
	-68, -48, 12, 14, 64, 54, 54, 92, 92, 92,
	-34, -34, 31, -35, 53, 53, -52, 26, 27, -49,
	-29, -36, -29, 44, 44, 44, 7, -34, -51, -53,
	-27, -67, -27, -27, -58, -54, 16, 34, 99, 53,
	99, 99, 7, 21, -67, -67, -67,
}
var yyDef = []int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 34, 34, 34, 34, 34,
	194, 185, 0, 0, 21, 22, 23, 0, 38, 40,
	41, 42, 43, 36, 0, 0, 0, 0, 183, 0,
	0, 195, 0, 0, 186, 0, 181, 0, 181, 15,
	39, 0, 44, 35, 0, 0, 76, 0, 20, 177,
	0, 0, 147, 198, 0, 0, 0, 199, 198, 0,
	199, 0, 0, 0, 0, 0, 0, 45, 47, 52,
	198, 50, 51, 86, 0, 0, 117, 118, 119, 0,
	147, 0, 133, 0, 149, 150, 151, 152, 113, 136,
	137, 138, 134, 135, 140, 37, 171, 0, 0, 84,
	0, 0, 180, 0, 0, 199, 0, 196, 26, 0,
	29, 0, 31, 182, 0, 199, 0, 0, 48, 53,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 101, 102,
	103, 104, 105, 106, 107, 89, 0, 0, 0, 0,
	115, 128, 0, 0, 100, 0, 0, 141, 0, 0,
	0, 84, 77, 157, 0, 178, 179, 115, 148, 24,
	184, 0, 0, 199, 192, 187, 188, 189, 190, 191,
	30, 32, 33, 84, 55, 61, 0, 73, 75, 46,
	54, 49, 87, 88, 91, 92, 0, 0, 0, 94,
	0, 98, 0, 120, 121, 122, 123, 124, 125, 126,
	127, 90, 112, 0, 114, 129, 0, 0, 0, 145,
	142, 0, 175, 0, 109, 175, 0, 173, 157, 165,
	0, 85, 0, 197, 27, 0, 193, 153, 0, 0,
	64, 65, 0, 0, 0, 0, 0, 78, 62, 0,
	0, 0, 93, 95, 0, 0, 99, 116, 130, 0,
	132, 0, 143, 0, 0, 16, 0, 108, 110, 17,
	172, 0, 165, 19, 0, 0, 199, 28, 155, 0,
	56, 59, 66, 0, 68, 0, 70, 71, 72, 57,
	0, 0, 0, 63, 58, 74, 0, 96, 131, 139,
	146, 0, 0, 0, 174, 18, 166, 158, 159, 162,
	25, 157, 0, 0, 0, 67, 69, 0, 0, 0,
	97, 144, 0, 111, 0, 0, 161, 163, 164, 165,
	156, 154, 60, 0, 0, 0, 0, 167, 160, 168,
	0, 82, 0, 0, 176, 14, 0, 0, 79, 0,
	80, 81, 169, 0, 83, 0, 170,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 75, 68, 3,
	44, 99, 73, 71, 53, 72, 76, 74, 3, 3,
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
	97, 98,
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
		//line sql.y:154
		{
			SetParseTree(yylex, yyS[yypt-0].statement)
		}
	case 2:
		//line sql.y:160
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
		//line sql.y:177
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyS[yypt-10].bytes2), Distinct: yyS[yypt-9].str, SelectExprs: yyS[yypt-8].selectExprs, From: yyS[yypt-6].tableExprs, Where: NewWhere(AST_WHERE, yyS[yypt-5].boolExpr), GroupBy: GroupBy(yyS[yypt-4].valExprs), Having: NewWhere(AST_HAVING, yyS[yypt-3].boolExpr), OrderBy: yyS[yypt-2].orderBy, Limit: yyS[yypt-1].limit, Lock: yyS[yypt-0].str}
		}
	case 15:
		//line sql.y:181
		{
			yyVAL.selStmt = &Union{Type: yyS[yypt-1].str, Left: yyS[yypt-2].selStmt, Right: yyS[yypt-0].selStmt}
		}
	case 16:
		//line sql.y:188
		{
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: yyS[yypt-2].columns, Rows: yyS[yypt-1].insRows, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 17:
		//line sql.y:192
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
		//line sql.y:204
		{
			yyVAL.statement = &Update{Comments: Comments(yyS[yypt-6].bytes2), Table: yyS[yypt-5].tableName, Exprs: yyS[yypt-3].updateExprs, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 19:
		//line sql.y:210
		{
			yyVAL.statement = &Delete{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 20:
		//line sql.y:216
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-1].bytes2), Exprs: yyS[yypt-0].updateExprs}
		}
	case 21:
		//line sql.y:222
		{
			yyVAL.statement = &Begin{}
		}
	case 22:
		//line sql.y:228
		{
			yyVAL.statement = &Commit{}
		}
	case 23:
		//line sql.y:234
		{
			yyVAL.statement = &Rollback{}
		}
	case 24:
		//line sql.y:240
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 25:
		//line sql.y:244
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 26:
		//line sql.y:249
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 27:
		//line sql.y:255
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-2].bytes}
		}
	case 28:
		//line sql.y:259
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-3].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 29:
		//line sql.y:264
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 30:
		//line sql.y:270
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 31:
		//line sql.y:276
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-0].bytes}
		}
	case 32:
		//line sql.y:280
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 33:
		//line sql.y:285
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-1].bytes}
		}
	case 34:
		//line sql.y:290
		{
			SetAllowComments(yylex, true)
		}
	case 35:
		//line sql.y:294
		{
			yyVAL.bytes2 = yyS[yypt-0].bytes2
			SetAllowComments(yylex, false)
		}
	case 36:
		//line sql.y:300
		{
			yyVAL.bytes2 = nil
		}
	case 37:
		//line sql.y:304
		{
			yyVAL.bytes2 = append(yyS[yypt-1].bytes2, yyS[yypt-0].bytes)
		}
	case 38:
		//line sql.y:310
		{
			yyVAL.str = AST_UNION
		}
	case 39:
		//line sql.y:314
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 40:
		//line sql.y:318
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 41:
		//line sql.y:322
		{
			yyVAL.str = AST_EXCEPT
		}
	case 42:
		//line sql.y:326
		{
			yyVAL.str = AST_INTERSECT
		}
	case 43:
		//line sql.y:331
		{
			yyVAL.str = ""
		}
	case 44:
		//line sql.y:335
		{
			yyVAL.str = AST_DISTINCT
		}
	case 45:
		//line sql.y:341
		{
			yyVAL.selectExprs = SelectExprs{yyS[yypt-0].selectExpr}
		}
	case 46:
		//line sql.y:345
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyS[yypt-0].selectExpr)
		}
	case 47:
		//line sql.y:351
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 48:
		//line sql.y:355
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyS[yypt-1].expr, As: yyS[yypt-0].bytes}
		}
	case 49:
		//line sql.y:359
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyS[yypt-2].bytes}
		}
	case 50:
		//line sql.y:365
		{
			yyVAL.expr = yyS[yypt-0].boolExpr
		}
	case 51:
		//line sql.y:369
		{
			yyVAL.expr = yyS[yypt-0].valExpr
		}
	case 52:
		//line sql.y:374
		{
			yyVAL.bytes = nil
		}
	case 53:
		//line sql.y:378
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 54:
		//line sql.y:382
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 55:
		//line sql.y:388
		{
			yyVAL.tableExprs = TableExprs{yyS[yypt-0].tableExpr}
		}
	case 56:
		//line sql.y:392
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyS[yypt-0].tableExpr)
		}
	case 57:
		//line sql.y:398
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyS[yypt-2].smTableExpr, As: yyS[yypt-1].bytes, Hints: yyS[yypt-0].indexHints}
		}
	case 58:
		//line sql.y:402
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyS[yypt-1].tableExpr}
		}
	case 59:
		//line sql.y:406
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-2].tableExpr, Join: yyS[yypt-1].str, RightExpr: yyS[yypt-0].tableExpr}
		}
	case 60:
		//line sql.y:410
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-4].tableExpr, Join: yyS[yypt-3].str, RightExpr: yyS[yypt-2].tableExpr, On: yyS[yypt-0].boolExpr}
		}
	case 61:
		//line sql.y:415
		{
			yyVAL.bytes = nil
		}
	case 62:
		//line sql.y:419
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 63:
		//line sql.y:423
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 64:
		//line sql.y:429
		{
			yyVAL.str = AST_JOIN
		}
	case 65:
		//line sql.y:433
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 66:
		//line sql.y:437
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 67:
		//line sql.y:441
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 68:
		//line sql.y:445
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 69:
		//line sql.y:449
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 70:
		//line sql.y:453
		{
			yyVAL.str = AST_JOIN
		}
	case 71:
		//line sql.y:457
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 72:
		//line sql.y:461
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 73:
		//line sql.y:467
		{
			yyVAL.smTableExpr = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 74:
		//line sql.y:471
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 75:
		//line sql.y:475
		{
			yyVAL.smTableExpr = yyS[yypt-0].subquery
		}
	case 76:
		//line sql.y:481
		{
			yyVAL.tableName = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 77:
		//line sql.y:485
		{
			yyVAL.tableName = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 78:
		//line sql.y:490
		{
			yyVAL.indexHints = nil
		}
	case 79:
		//line sql.y:494
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyS[yypt-1].bytes2}
		}
	case 80:
		//line sql.y:498
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyS[yypt-1].bytes2}
		}
	case 81:
		//line sql.y:502
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyS[yypt-1].bytes2}
		}
	case 82:
		//line sql.y:508
		{
			yyVAL.bytes2 = [][]byte{yyS[yypt-0].bytes}
		}
	case 83:
		//line sql.y:512
		{
			yyVAL.bytes2 = append(yyS[yypt-2].bytes2, yyS[yypt-0].bytes)
		}
	case 84:
		//line sql.y:517
		{
			yyVAL.boolExpr = nil
		}
	case 85:
		//line sql.y:521
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 86:
		yyVAL.boolExpr = yyS[yypt-0].boolExpr
	case 87:
		//line sql.y:528
		{
			yyVAL.boolExpr = &AndExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 88:
		//line sql.y:532
		{
			yyVAL.boolExpr = &OrExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 89:
		//line sql.y:536
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyS[yypt-0].boolExpr}
		}
	case 90:
		//line sql.y:540
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyS[yypt-1].boolExpr}
		}
	case 91:
		//line sql.y:546
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: yyS[yypt-1].str, Right: yyS[yypt-0].valExpr}
		}
	case 92:
		//line sql.y:550
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_IN, Right: yyS[yypt-0].tuple}
		}
	case 93:
		//line sql.y:554
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_IN, Right: yyS[yypt-0].tuple}
		}
	case 94:
		//line sql.y:558
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 95:
		//line sql.y:562
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 96:
		//line sql.y:566
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-4].valExpr, Operator: AST_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 97:
		//line sql.y:570
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-5].valExpr, Operator: AST_NOT_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 98:
		//line sql.y:574
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyS[yypt-2].valExpr}
		}
	case 99:
		//line sql.y:578
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyS[yypt-3].valExpr}
		}
	case 100:
		//line sql.y:582
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyS[yypt-0].subquery}
		}
	case 101:
		//line sql.y:588
		{
			yyVAL.str = AST_EQ
		}
	case 102:
		//line sql.y:592
		{
			yyVAL.str = AST_LT
		}
	case 103:
		//line sql.y:596
		{
			yyVAL.str = AST_GT
		}
	case 104:
		//line sql.y:600
		{
			yyVAL.str = AST_LE
		}
	case 105:
		//line sql.y:604
		{
			yyVAL.str = AST_GE
		}
	case 106:
		//line sql.y:608
		{
			yyVAL.str = AST_NE
		}
	case 107:
		//line sql.y:612
		{
			yyVAL.str = AST_NSE
		}
	case 108:
		//line sql.y:618
		{
			yyVAL.insRows = yyS[yypt-0].values
		}
	case 109:
		//line sql.y:622
		{
			yyVAL.insRows = yyS[yypt-0].selStmt
		}
	case 110:
		//line sql.y:628
		{
			yyVAL.values = Values{yyS[yypt-0].tuple}
		}
	case 111:
		//line sql.y:632
		{
			yyVAL.values = append(yyS[yypt-2].values, yyS[yypt-0].tuple)
		}
	case 112:
		//line sql.y:638
		{
			yyVAL.tuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 113:
		//line sql.y:642
		{
			yyVAL.tuple = yyS[yypt-0].subquery
		}
	case 114:
		//line sql.y:648
		{
			yyVAL.subquery = &Subquery{yyS[yypt-1].selStmt}
		}
	case 115:
		//line sql.y:654
		{
			yyVAL.valExprs = ValExprs{yyS[yypt-0].valExpr}
		}
	case 116:
		//line sql.y:658
		{
			yyVAL.valExprs = append(yyS[yypt-2].valExprs, yyS[yypt-0].valExpr)
		}
	case 117:
		//line sql.y:664
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 118:
		//line sql.y:668
		{
			yyVAL.valExpr = yyS[yypt-0].colName
		}
	case 119:
		//line sql.y:672
		{
			yyVAL.valExpr = yyS[yypt-0].tuple
		}
	case 120:
		//line sql.y:676
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITAND, Right: yyS[yypt-0].valExpr}
		}
	case 121:
		//line sql.y:680
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITOR, Right: yyS[yypt-0].valExpr}
		}
	case 122:
		//line sql.y:684
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITXOR, Right: yyS[yypt-0].valExpr}
		}
	case 123:
		//line sql.y:688
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_PLUS, Right: yyS[yypt-0].valExpr}
		}
	case 124:
		//line sql.y:692
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MINUS, Right: yyS[yypt-0].valExpr}
		}
	case 125:
		//line sql.y:696
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MULT, Right: yyS[yypt-0].valExpr}
		}
	case 126:
		//line sql.y:700
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_DIV, Right: yyS[yypt-0].valExpr}
		}
	case 127:
		//line sql.y:704
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MOD, Right: yyS[yypt-0].valExpr}
		}
	case 128:
		//line sql.y:708
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
		//line sql.y:723
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-2].bytes}
		}
	case 130:
		//line sql.y:727
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 131:
		//line sql.y:731
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-4].bytes, Distinct: true, Exprs: yyS[yypt-1].selectExprs}
		}
	case 132:
		//line sql.y:735
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 133:
		//line sql.y:739
		{
			yyVAL.valExpr = yyS[yypt-0].caseExpr
		}
	case 134:
		//line sql.y:745
		{
			yyVAL.bytes = IF_BYTES
		}
	case 135:
		//line sql.y:749
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 136:
		//line sql.y:755
		{
			yyVAL.byt = AST_UPLUS
		}
	case 137:
		//line sql.y:759
		{
			yyVAL.byt = AST_UMINUS
		}
	case 138:
		//line sql.y:763
		{
			yyVAL.byt = AST_TILDA
		}
	case 139:
		//line sql.y:769
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyS[yypt-3].valExpr, Whens: yyS[yypt-2].whens, Else: yyS[yypt-1].valExpr}
		}
	case 140:
		//line sql.y:774
		{
			yyVAL.valExpr = nil
		}
	case 141:
		//line sql.y:778
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 142:
		//line sql.y:784
		{
			yyVAL.whens = []*When{yyS[yypt-0].when}
		}
	case 143:
		//line sql.y:788
		{
			yyVAL.whens = append(yyS[yypt-1].whens, yyS[yypt-0].when)
		}
	case 144:
		//line sql.y:794
		{
			yyVAL.when = &When{Cond: yyS[yypt-2].boolExpr, Val: yyS[yypt-0].valExpr}
		}
	case 145:
		//line sql.y:799
		{
			yyVAL.valExpr = nil
		}
	case 146:
		//line sql.y:803
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 147:
		//line sql.y:809
		{
			yyVAL.colName = &ColName{Name: yyS[yypt-0].bytes}
		}
	case 148:
		//line sql.y:813
		{
			yyVAL.colName = &ColName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 149:
		//line sql.y:819
		{
			yyVAL.valExpr = StrVal(yyS[yypt-0].bytes)
		}
	case 150:
		//line sql.y:823
		{
			yyVAL.valExpr = NumVal(yyS[yypt-0].bytes)
		}
	case 151:
		//line sql.y:827
		{
			yyVAL.valExpr = ValArg(yyS[yypt-0].bytes)
		}
	case 152:
		//line sql.y:831
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 153:
		//line sql.y:836
		{
			yyVAL.valExprs = nil
		}
	case 154:
		//line sql.y:840
		{
			yyVAL.valExprs = yyS[yypt-0].valExprs
		}
	case 155:
		//line sql.y:845
		{
			yyVAL.boolExpr = nil
		}
	case 156:
		//line sql.y:849
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 157:
		//line sql.y:854
		{
			yyVAL.orderBy = nil
		}
	case 158:
		//line sql.y:858
		{
			yyVAL.orderBy = yyS[yypt-0].orderBy
		}
	case 159:
		//line sql.y:864
		{
			yyVAL.orderBy = OrderBy{yyS[yypt-0].order}
		}
	case 160:
		//line sql.y:868
		{
			yyVAL.orderBy = append(yyS[yypt-2].orderBy, yyS[yypt-0].order)
		}
	case 161:
		//line sql.y:874
		{
			yyVAL.order = &Order{Expr: yyS[yypt-1].valExpr, Direction: yyS[yypt-0].str}
		}
	case 162:
		//line sql.y:879
		{
			yyVAL.str = AST_ASC
		}
	case 163:
		//line sql.y:883
		{
			yyVAL.str = AST_ASC
		}
	case 164:
		//line sql.y:887
		{
			yyVAL.str = AST_DESC
		}
	case 165:
		//line sql.y:892
		{
			yyVAL.limit = nil
		}
	case 166:
		//line sql.y:896
		{
			yyVAL.limit = &Limit{Rowcount: yyS[yypt-0].valExpr}
		}
	case 167:
		//line sql.y:900
		{
			yyVAL.limit = &Limit{Offset: yyS[yypt-2].valExpr, Rowcount: yyS[yypt-0].valExpr}
		}
	case 168:
		//line sql.y:905
		{
			yyVAL.str = ""
		}
	case 169:
		//line sql.y:909
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 170:
		//line sql.y:913
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
		//line sql.y:926
		{
			yyVAL.columns = nil
		}
	case 172:
		//line sql.y:930
		{
			yyVAL.columns = yyS[yypt-1].columns
		}
	case 173:
		//line sql.y:936
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyS[yypt-0].colName}}
		}
	case 174:
		//line sql.y:940
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyS[yypt-0].colName})
		}
	case 175:
		//line sql.y:945
		{
			yyVAL.updateExprs = nil
		}
	case 176:
		//line sql.y:949
		{
			yyVAL.updateExprs = yyS[yypt-0].updateExprs
		}
	case 177:
		//line sql.y:955
		{
			yyVAL.updateExprs = UpdateExprs{yyS[yypt-0].updateExpr}
		}
	case 178:
		//line sql.y:959
		{
			yyVAL.updateExprs = append(yyS[yypt-2].updateExprs, yyS[yypt-0].updateExpr)
		}
	case 179:
		//line sql.y:965
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyS[yypt-2].colName, Expr: yyS[yypt-0].valExpr}
		}
	case 180:
		//line sql.y:969
		{
			yyVAL.updateExpr = &UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: yyS[yypt-0].valExpr}
		}
	case 181:
		//line sql.y:974
		{
			yyVAL.empty = struct{}{}
		}
	case 182:
		//line sql.y:976
		{
			yyVAL.empty = struct{}{}
		}
	case 183:
		//line sql.y:979
		{
			yyVAL.empty = struct{}{}
		}
	case 184:
		//line sql.y:981
		{
			yyVAL.empty = struct{}{}
		}
	case 185:
		//line sql.y:984
		{
			yyVAL.empty = struct{}{}
		}
	case 186:
		//line sql.y:986
		{
			yyVAL.empty = struct{}{}
		}
	case 187:
		//line sql.y:990
		{
			yyVAL.empty = struct{}{}
		}
	case 188:
		//line sql.y:992
		{
			yyVAL.empty = struct{}{}
		}
	case 189:
		//line sql.y:994
		{
			yyVAL.empty = struct{}{}
		}
	case 190:
		//line sql.y:996
		{
			yyVAL.empty = struct{}{}
		}
	case 191:
		//line sql.y:998
		{
			yyVAL.empty = struct{}{}
		}
	case 192:
		//line sql.y:1001
		{
			yyVAL.empty = struct{}{}
		}
	case 193:
		//line sql.y:1003
		{
			yyVAL.empty = struct{}{}
		}
	case 194:
		//line sql.y:1006
		{
			yyVAL.empty = struct{}{}
		}
	case 195:
		//line sql.y:1008
		{
			yyVAL.empty = struct{}{}
		}
	case 196:
		//line sql.y:1011
		{
			yyVAL.empty = struct{}{}
		}
	case 197:
		//line sql.y:1013
		{
			yyVAL.empty = struct{}{}
		}
	case 198:
		//line sql.y:1017
		{
			yyVAL.bytes = bytes.ToLower(yyS[yypt-0].bytes)
		}
	case 199:
		//line sql.y:1022
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
