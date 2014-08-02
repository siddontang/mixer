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
const REPLACE = 57414
const CREATE = 57415
const ALTER = 57416
const DROP = 57417
const RENAME = 57418
const TABLE = 57419
const INDEX = 57420
const VIEW = 57421
const TO = 57422
const IGNORE = 57423
const IF = 57424
const UNIQUE = 57425
const USING = 57426

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
	"REPLACE",
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

const yyNprod = 203
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 564

var yyAct = []int{

	72, 292, 169, 359, 62, 130, 327, 70, 246, 138,
	219, 239, 284, 200, 101, 205, 69, 100, 179, 174,
	368, 110, 63, 80, 131, 3, 78, 96, 105, 85,
	30, 31, 32, 33, 153, 154, 66, 75, 76, 77,
	65, 368, 88, 67, 90, 109, 368, 93, 59, 83,
	148, 97, 261, 290, 53, 148, 64, 141, 106, 267,
	268, 269, 270, 271, 148, 272, 273, 370, 108, 238,
	339, 338, 81, 82, 127, 49, 50, 51, 41, 86,
	43, 198, 132, 46, 44, 47, 201, 337, 369, 134,
	137, 89, 40, 367, 296, 92, 48, 295, 84, 145,
	289, 278, 251, 140, 150, 317, 201, 112, 255, 197,
	115, 249, 171, 152, 65, 313, 315, 65, 177, 183,
	182, 118, 184, 185, 186, 187, 188, 189, 190, 191,
	64, 106, 106, 64, 172, 153, 154, 168, 170, 181,
	114, 16, 17, 18, 19, 202, 193, 195, 146, 314,
	298, 106, 225, 183, 136, 215, 196, 91, 229, 153,
	154, 234, 235, 224, 230, 334, 216, 116, 285, 20,
	257, 223, 124, 125, 126, 236, 65, 65, 285, 144,
	242, 307, 227, 228, 117, 305, 308, 336, 211, 335,
	306, 311, 64, 244, 310, 226, 245, 106, 309, 241,
	252, 122, 123, 124, 125, 126, 258, 209, 116, 248,
	212, 197, 250, 254, 180, 344, 259, 237, 65, 25,
	26, 27, 263, 28, 21, 22, 24, 23, 262, 147,
	256, 180, 277, 322, 64, 280, 281, 264, 222, 279,
	95, 241, 217, 354, 111, 223, 353, 221, 287, 30,
	31, 32, 33, 176, 291, 288, 16, 265, 297, 78,
	352, 79, 85, 173, 208, 210, 207, 129, 175, 66,
	75, 76, 77, 148, 116, 128, 303, 304, 79, 176,
	113, 91, 83, 276, 151, 320, 222, 66, 318, 223,
	223, 65, 98, 324, 316, 221, 325, 328, 343, 275,
	91, 329, 300, 365, 299, 81, 82, 323, 214, 330,
	213, 178, 86, 119, 120, 121, 122, 123, 124, 125,
	126, 366, 340, 60, 142, 139, 135, 94, 341, 321,
	342, 84, 99, 58, 283, 231, 132, 232, 233, 350,
	348, 372, 16, 203, 143, 56, 356, 328, 54, 293,
	358, 357, 333, 360, 360, 360, 65, 361, 362, 332,
	363, 349, 294, 351, 194, 240, 110, 247, 302, 373,
	61, 78, 64, 374, 85, 375, 180, 371, 355, 16,
	110, 104, 75, 76, 77, 78, 35, 15, 85, 14,
	109, 34, 13, 12, 83, 104, 75, 76, 77, 267,
	268, 269, 270, 271, 109, 272, 273, 204, 83, 36,
	37, 38, 39, 108, 42, 260, 206, 81, 82, 102,
	52, 16, 45, 87, 86, 243, 364, 108, 345, 326,
	331, 81, 82, 102, 301, 16, 110, 253, 86, 133,
	199, 78, 74, 84, 85, 71, 192, 73, 286, 68,
	155, 66, 75, 76, 77, 78, 107, 84, 85, 312,
	109, 220, 266, 218, 83, 66, 75, 76, 77, 103,
	274, 149, 55, 29, 79, 346, 347, 57, 83, 11,
	10, 9, 8, 108, 7, 6, 5, 81, 82, 4,
	156, 160, 158, 159, 86, 2, 1, 0, 0, 0,
	0, 81, 82, 0, 0, 0, 0, 0, 86, 164,
	165, 166, 167, 84, 161, 162, 163, 119, 120, 121,
	122, 123, 124, 125, 126, 319, 0, 84, 119, 120,
	121, 122, 123, 124, 125, 126, 157, 119, 120, 121,
	122, 123, 124, 125, 126, 282, 0, 0, 119, 120,
	121, 122, 123, 124, 125, 126, 119, 120, 121, 122,
	123, 124, 125, 126,
}
var yyPact = []int{

	136, -1000, -1000, 200, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	6, -14, -11, 4, -17, -1000, -1000, -1000, -1000, 374,
	331, -1000, -1000, -1000, 327, -1000, 304, 288, 361, 252,
	234, -55, -2, 246, -1000, 3, 246, -1000, 292, -70,
	246, -70, 303, -1000, -1000, 360, -1000, 205, 288, 247,
	64, 288, 155, -1000, 139, -1000, 45, 488, -1000, -1000,
	-1000, 234, 231, 223, -1000, -1000, -1000, -1000, -1000, 430,
	-1000, -1000, -1000, -1000, -1000, -1000, 234, 291, 87, 246,
	-1000, -1000, 290, -1000, -38, 289, 324, 115, 246, 288,
	220, -1000, -1000, 265, 37, 94, 469, -1000, 1, 416,
	219, -1000, 235, 252, 276, 366, 252, 234, 246, 234,
	234, 234, 234, 234, 234, 234, 234, -1000, 346, 360,
	56, -19, 488, 7, 488, -1000, 323, -84, -1000, 175,
	-1000, 275, -1000, -1000, 273, -1000, 209, 203, 360, -1000,
	-1000, 246, 122, 1, 1, 234, 217, 314, 234, 234,
	150, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 469,
	-31, 469, -1000, 374, 337, 252, 252, 221, -1000, 354,
	1, -1000, 488, -1000, 130, 130, 130, 99, 99, -1000,
	-1000, -1000, -1000, 11, 360, 2, -1000, 234, -1000, 27,
	-1000, 1, -1000, -1000, 106, 246, -1000, -43, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 337, 252, 204, 345,
	264, 251, 25, -1000, -1000, -1000, -1000, -1000, -1000, 488,
	-1000, 217, 234, 234, 488, 480, -1000, 309, -1000, 104,
	217, 200, 114, 0, -1000, 354, 334, 348, 94, -1000,
	-3, -1000, 488, 12, -1000, 234, 70, 269, -1000, -1000,
	267, -1000, -1000, 155, 357, 203, 203, -1000, -1000, 131,
	127, 144, 140, 137, 53, -1000, 259, 5, 253, -1000,
	488, 460, 234, -1000, -1000, 299, 180, -1000, -1000, -1000,
	252, 334, -1000, 234, 234, -1000, -1000, 488, 234, -1000,
	-1000, 347, 338, 345, 101, -1000, 135, -1000, 133, -1000,
	-1000, -1000, -1000, -6, -22, -23, -1000, -1000, -1000, 234,
	488, 297, 217, -1000, -1000, 245, 162, -1000, 449, 488,
	-1000, 354, 1, 234, 1, -1000, -1000, 216, 202, 199,
	488, 371, -1000, 234, 234, -1000, -1000, -1000, 334, 94,
	158, 94, 246, 246, 246, 252, 488, -1000, 287, -7,
	-1000, -12, -33, 155, -1000, 370, 320, -1000, 246, -1000,
	-1000, -1000, 246, -1000, 246, -1000,
}
var yyPgo = []int{

	0, 496, 495, 24, 489, 486, 485, 484, 482, 481,
	480, 479, 391, 477, 473, 472, 17, 14, 471, 470,
	469, 463, 10, 462, 461, 48, 459, 3, 18, 28,
	456, 450, 11, 449, 2, 7, 5, 448, 447, 23,
	445, 16, 442, 440, 13, 439, 437, 434, 430, 8,
	429, 6, 428, 1, 426, 19, 425, 12, 4, 22,
	240, 423, 422, 416, 415, 414, 407, 0, 9, 393,
	392, 389, 387, 386,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 3, 3, 4, 4, 72,
	72, 5, 6, 7, 7, 69, 70, 71, 8, 8,
	8, 9, 9, 9, 10, 11, 11, 11, 73, 12,
	13, 13, 14, 14, 14, 14, 14, 15, 15, 16,
	16, 17, 17, 17, 20, 20, 18, 18, 18, 21,
	21, 22, 22, 22, 22, 19, 19, 19, 23, 23,
	23, 23, 23, 23, 23, 23, 23, 24, 24, 24,
	25, 25, 26, 26, 26, 26, 27, 27, 28, 28,
	29, 29, 29, 29, 29, 30, 30, 30, 30, 30,
	30, 30, 30, 30, 30, 31, 31, 31, 31, 31,
	31, 31, 32, 32, 37, 37, 35, 35, 39, 36,
	36, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 38, 38,
	40, 40, 40, 42, 45, 45, 43, 43, 44, 46,
	46, 41, 41, 33, 33, 33, 33, 47, 47, 48,
	48, 49, 49, 50, 50, 51, 52, 52, 52, 53,
	53, 53, 54, 54, 54, 55, 55, 56, 56, 57,
	57, 58, 58, 59, 60, 60, 61, 61, 62, 62,
	63, 63, 63, 63, 63, 64, 64, 65, 65, 66,
	66, 67, 68,
}
var yyR2 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 12, 3, 7, 7, 6,
	6, 8, 7, 3, 3, 1, 1, 1, 5, 8,
	4, 6, 7, 4, 5, 4, 5, 5, 0, 2,
	0, 2, 1, 2, 1, 1, 1, 0, 1, 1,
	3, 1, 2, 3, 1, 1, 0, 1, 2, 1,
	3, 3, 3, 3, 5, 0, 1, 2, 1, 1,
	2, 3, 2, 3, 2, 2, 2, 1, 3, 1,
	1, 3, 0, 5, 5, 5, 1, 3, 0, 2,
	1, 3, 3, 2, 3, 3, 3, 4, 3, 4,
	5, 6, 3, 4, 2, 1, 1, 1, 1, 1,
	1, 1, 2, 1, 1, 3, 3, 1, 3, 1,
	3, 1, 1, 1, 3, 3, 3, 3, 3, 3,
	3, 3, 2, 3, 4, 5, 4, 1, 1, 1,
	1, 1, 1, 5, 0, 1, 1, 2, 4, 0,
	2, 1, 3, 1, 1, 1, 1, 0, 3, 0,
	2, 0, 3, 1, 3, 2, 0, 1, 1, 0,
	2, 4, 0, 2, 4, 0, 3, 1, 3, 0,
	5, 1, 3, 3, 0, 2, 0, 3, 0, 1,
	1, 1, 1, 1, 1, 0, 1, 0, 1, 0,
	2, 1, 0,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -69, -70, -71, -72, 5, 6, 7, 8,
	33, 88, 89, 91, 90, 83, 84, 85, 87, -14,
	49, 50, 51, 52, -12, -73, -12, -12, -12, -12,
	86, 92, -65, 94, 98, -62, 94, 96, 92, 92,
	93, 94, -12, -3, 17, -15, 18, -13, 29, -25,
	35, 9, -58, -59, -41, -67, 35, -34, -33, -41,
	-35, -40, -67, -38, -42, 36, 37, 38, 25, 44,
	-39, 71, 72, 48, 97, 28, 78, -61, 97, 93,
	-67, 35, 92, -67, 35, -60, 97, -67, -60, 29,
	-16, -17, 73, -20, 35, -29, -34, -30, 67, 44,
	20, 39, -25, 33, 76, -25, 53, 45, 76, 68,
	69, 70, 71, 72, 73, 74, 75, -34, 44, 44,
	-36, -3, -34, -45, -34, 35, 67, -67, -68, 35,
	-68, 95, 35, 20, 64, -67, -25, 9, 53, -18,
	-67, 19, 76, 65, 66, -31, 21, 67, 23, 24,
	22, 45, 46, 47, 40, 41, 42, 43, -29, -34,
	-29, -34, -39, 44, -55, 33, 44, -58, 35, -28,
	10, -59, -34, -67, -34, -34, -34, -34, -34, -34,
	-34, -34, 100, -16, 18, -16, 100, 53, 100, -43,
	-44, 79, -68, 20, -66, 99, -63, 91, 89, 32,
	90, 13, 35, 35, 35, -68, -55, 33, -21, -22,
	-24, 44, 35, -39, -17, -67, 73, -29, -29, -34,
	-35, 21, 23, 24, -34, -34, 25, 67, 100, -32,
	28, -3, -58, -56, -41, -28, -49, 13, -29, 100,
	-16, 100, -34, -46, -44, 81, -29, 64, -67, -68,
	-64, 95, -32, -58, -28, 53, -23, 54, 55, 56,
	57, 58, 60, 61, -19, 35, 19, -22, 76, -35,
	-34, -34, 65, 25, -57, 64, -37, -35, -57, 100,
	53, -49, -53, 15, 14, 100, 82, -34, 80, 35,
	35, -47, 11, -22, -22, 54, 59, 54, 59, 54,
	54, 54, -26, 62, 96, 63, 35, 100, 35, 65,
	-34, 30, 53, -41, -53, -34, -50, -51, -34, -34,
	-68, -48, 12, 14, 64, 54, 54, 93, 93, 93,
	-34, 31, -35, 53, 53, -52, 26, 27, -49, -29,
	-36, -29, 44, 44, 44, 7, -34, -51, -53, -27,
	-67, -27, -27, -58, -54, 16, 34, 100, 53, 100,
	100, 7, 21, -67, -67, -67,
}
var yyDef = []int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 38, 38, 38, 38,
	38, 197, 188, 0, 0, 25, 26, 27, 38, 0,
	42, 44, 45, 46, 47, 40, 0, 0, 0, 0,
	0, 186, 0, 0, 198, 0, 0, 189, 0, 184,
	0, 184, 0, 16, 43, 0, 48, 39, 0, 0,
	80, 0, 23, 181, 0, 151, 201, 24, 121, 122,
	123, 0, 151, 0, 137, 153, 154, 155, 156, 0,
	117, 140, 141, 142, 138, 139, 144, 0, 0, 0,
	202, 201, 0, 202, 0, 0, 0, 0, 0, 0,
	0, 49, 51, 56, 201, 54, 55, 90, 0, 0,
	0, 41, 175, 0, 0, 88, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 132, 0, 0,
	0, 0, 119, 0, 145, 202, 0, 199, 30, 0,
	33, 0, 35, 185, 0, 202, 175, 0, 0, 52,
	57, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 105, 106, 107, 108, 109, 110, 111, 93, 0,
	0, 119, 104, 0, 0, 0, 0, 88, 81, 161,
	0, 182, 183, 152, 124, 125, 126, 127, 128, 129,
	130, 131, 133, 0, 0, 0, 116, 0, 118, 149,
	146, 0, 28, 187, 0, 0, 202, 195, 190, 191,
	192, 193, 194, 34, 36, 37, 0, 0, 88, 59,
	65, 0, 77, 79, 50, 58, 53, 91, 92, 95,
	96, 0, 0, 0, 98, 0, 102, 0, 94, 179,
	0, 113, 179, 0, 177, 161, 169, 0, 89, 134,
	0, 136, 120, 0, 147, 0, 0, 0, 200, 31,
	0, 196, 19, 20, 157, 0, 0, 68, 69, 0,
	0, 0, 0, 0, 82, 66, 0, 0, 0, 97,
	99, 0, 0, 103, 17, 0, 112, 114, 18, 176,
	0, 169, 22, 0, 0, 135, 143, 150, 0, 202,
	32, 159, 0, 60, 63, 70, 0, 72, 0, 74,
	75, 76, 61, 0, 0, 0, 67, 62, 78, 0,
	100, 0, 0, 178, 21, 170, 162, 163, 166, 148,
	29, 161, 0, 0, 0, 71, 73, 0, 0, 0,
	101, 0, 115, 0, 0, 165, 167, 168, 169, 160,
	158, 64, 0, 0, 0, 0, 171, 164, 172, 0,
	86, 0, 0, 180, 15, 0, 0, 83, 0, 84,
	85, 173, 0, 87, 0, 174,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 75, 68, 3,
	44, 100, 73, 71, 53, 72, 76, 74, 3, 3,
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
	97, 98, 99,
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
		//line sql.y:158
		{
			SetParseTree(yylex, yyS[yypt-0].statement)
		}
	case 2:
		//line sql.y:164
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
		yyVAL.statement = yyS[yypt-0].statement
	case 15:
		//line sql.y:182
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyS[yypt-10].bytes2), Distinct: yyS[yypt-9].str, SelectExprs: yyS[yypt-8].selectExprs, From: yyS[yypt-6].tableExprs, Where: NewWhere(AST_WHERE, yyS[yypt-5].boolExpr), GroupBy: GroupBy(yyS[yypt-4].valExprs), Having: NewWhere(AST_HAVING, yyS[yypt-3].boolExpr), OrderBy: yyS[yypt-2].orderBy, Limit: yyS[yypt-1].limit, Lock: yyS[yypt-0].str}
		}
	case 16:
		//line sql.y:186
		{
			yyVAL.selStmt = &Union{Type: yyS[yypt-1].str, Left: yyS[yypt-2].selStmt, Right: yyS[yypt-0].selStmt}
		}
	case 17:
		//line sql.y:193
		{
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: yyS[yypt-2].columns, Rows: yyS[yypt-1].insRows, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 18:
		//line sql.y:197
		{
			cols := make(Columns, 0, len(yyS[yypt-1].updateExprs))
			vals := make(ValTuple, 0, len(yyS[yypt-1].updateExprs))
			for _, col := range yyS[yypt-1].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 19:
		//line sql.y:209
		{
			yyVAL.statement = &Replace{Comments: Comments(yyS[yypt-4].bytes2), Table: yyS[yypt-2].tableName, Columns: yyS[yypt-1].columns, Rows: yyS[yypt-0].insRows}
		}
	case 20:
		//line sql.y:213
		{
			cols := make(Columns, 0, len(yyS[yypt-0].updateExprs))
			vals := make(ValTuple, 0, len(yyS[yypt-0].updateExprs))
			for _, col := range yyS[yypt-0].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Replace{Comments: Comments(yyS[yypt-4].bytes2), Table: yyS[yypt-2].tableName, Columns: cols, Rows: Values{vals}}
		}
	case 21:
		//line sql.y:226
		{
			yyVAL.statement = &Update{Comments: Comments(yyS[yypt-6].bytes2), Table: yyS[yypt-5].tableName, Exprs: yyS[yypt-3].updateExprs, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 22:
		//line sql.y:232
		{
			yyVAL.statement = &Delete{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 23:
		//line sql.y:238
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-1].bytes2), Exprs: yyS[yypt-0].updateExprs}
		}
	case 24:
		//line sql.y:242
		{
			yyVAL.statement = &Set{Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: yyS[yypt-0].valExpr}}}
		}
	case 25:
		//line sql.y:248
		{
			yyVAL.statement = &Begin{}
		}
	case 26:
		//line sql.y:254
		{
			yyVAL.statement = &Commit{}
		}
	case 27:
		//line sql.y:260
		{
			yyVAL.statement = &Rollback{}
		}
	case 28:
		//line sql.y:266
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 29:
		//line sql.y:270
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 30:
		//line sql.y:275
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 31:
		//line sql.y:281
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-2].bytes}
		}
	case 32:
		//line sql.y:285
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-3].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 33:
		//line sql.y:290
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 34:
		//line sql.y:296
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 35:
		//line sql.y:302
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-0].bytes}
		}
	case 36:
		//line sql.y:306
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 37:
		//line sql.y:311
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-1].bytes}
		}
	case 38:
		//line sql.y:316
		{
			SetAllowComments(yylex, true)
		}
	case 39:
		//line sql.y:320
		{
			yyVAL.bytes2 = yyS[yypt-0].bytes2
			SetAllowComments(yylex, false)
		}
	case 40:
		//line sql.y:326
		{
			yyVAL.bytes2 = nil
		}
	case 41:
		//line sql.y:330
		{
			yyVAL.bytes2 = append(yyS[yypt-1].bytes2, yyS[yypt-0].bytes)
		}
	case 42:
		//line sql.y:336
		{
			yyVAL.str = AST_UNION
		}
	case 43:
		//line sql.y:340
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 44:
		//line sql.y:344
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 45:
		//line sql.y:348
		{
			yyVAL.str = AST_EXCEPT
		}
	case 46:
		//line sql.y:352
		{
			yyVAL.str = AST_INTERSECT
		}
	case 47:
		//line sql.y:357
		{
			yyVAL.str = ""
		}
	case 48:
		//line sql.y:361
		{
			yyVAL.str = AST_DISTINCT
		}
	case 49:
		//line sql.y:367
		{
			yyVAL.selectExprs = SelectExprs{yyS[yypt-0].selectExpr}
		}
	case 50:
		//line sql.y:371
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyS[yypt-0].selectExpr)
		}
	case 51:
		//line sql.y:377
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 52:
		//line sql.y:381
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyS[yypt-1].expr, As: yyS[yypt-0].bytes}
		}
	case 53:
		//line sql.y:385
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyS[yypt-2].bytes}
		}
	case 54:
		//line sql.y:391
		{
			yyVAL.expr = yyS[yypt-0].boolExpr
		}
	case 55:
		//line sql.y:395
		{
			yyVAL.expr = yyS[yypt-0].valExpr
		}
	case 56:
		//line sql.y:400
		{
			yyVAL.bytes = nil
		}
	case 57:
		//line sql.y:404
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 58:
		//line sql.y:408
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 59:
		//line sql.y:414
		{
			yyVAL.tableExprs = TableExprs{yyS[yypt-0].tableExpr}
		}
	case 60:
		//line sql.y:418
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyS[yypt-0].tableExpr)
		}
	case 61:
		//line sql.y:424
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyS[yypt-2].smTableExpr, As: yyS[yypt-1].bytes, Hints: yyS[yypt-0].indexHints}
		}
	case 62:
		//line sql.y:428
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyS[yypt-1].tableExpr}
		}
	case 63:
		//line sql.y:432
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-2].tableExpr, Join: yyS[yypt-1].str, RightExpr: yyS[yypt-0].tableExpr}
		}
	case 64:
		//line sql.y:436
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-4].tableExpr, Join: yyS[yypt-3].str, RightExpr: yyS[yypt-2].tableExpr, On: yyS[yypt-0].boolExpr}
		}
	case 65:
		//line sql.y:441
		{
			yyVAL.bytes = nil
		}
	case 66:
		//line sql.y:445
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 67:
		//line sql.y:449
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 68:
		//line sql.y:455
		{
			yyVAL.str = AST_JOIN
		}
	case 69:
		//line sql.y:459
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 70:
		//line sql.y:463
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 71:
		//line sql.y:467
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 72:
		//line sql.y:471
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 73:
		//line sql.y:475
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 74:
		//line sql.y:479
		{
			yyVAL.str = AST_JOIN
		}
	case 75:
		//line sql.y:483
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 76:
		//line sql.y:487
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 77:
		//line sql.y:493
		{
			yyVAL.smTableExpr = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 78:
		//line sql.y:497
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 79:
		//line sql.y:501
		{
			yyVAL.smTableExpr = yyS[yypt-0].subquery
		}
	case 80:
		//line sql.y:507
		{
			yyVAL.tableName = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 81:
		//line sql.y:511
		{
			yyVAL.tableName = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 82:
		//line sql.y:516
		{
			yyVAL.indexHints = nil
		}
	case 83:
		//line sql.y:520
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyS[yypt-1].bytes2}
		}
	case 84:
		//line sql.y:524
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyS[yypt-1].bytes2}
		}
	case 85:
		//line sql.y:528
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyS[yypt-1].bytes2}
		}
	case 86:
		//line sql.y:534
		{
			yyVAL.bytes2 = [][]byte{yyS[yypt-0].bytes}
		}
	case 87:
		//line sql.y:538
		{
			yyVAL.bytes2 = append(yyS[yypt-2].bytes2, yyS[yypt-0].bytes)
		}
	case 88:
		//line sql.y:543
		{
			yyVAL.boolExpr = nil
		}
	case 89:
		//line sql.y:547
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 90:
		yyVAL.boolExpr = yyS[yypt-0].boolExpr
	case 91:
		//line sql.y:554
		{
			yyVAL.boolExpr = &AndExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 92:
		//line sql.y:558
		{
			yyVAL.boolExpr = &OrExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 93:
		//line sql.y:562
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyS[yypt-0].boolExpr}
		}
	case 94:
		//line sql.y:566
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyS[yypt-1].boolExpr}
		}
	case 95:
		//line sql.y:572
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: yyS[yypt-1].str, Right: yyS[yypt-0].valExpr}
		}
	case 96:
		//line sql.y:576
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_IN, Right: yyS[yypt-0].tuple}
		}
	case 97:
		//line sql.y:580
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_IN, Right: yyS[yypt-0].tuple}
		}
	case 98:
		//line sql.y:584
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 99:
		//line sql.y:588
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 100:
		//line sql.y:592
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-4].valExpr, Operator: AST_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 101:
		//line sql.y:596
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-5].valExpr, Operator: AST_NOT_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 102:
		//line sql.y:600
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyS[yypt-2].valExpr}
		}
	case 103:
		//line sql.y:604
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyS[yypt-3].valExpr}
		}
	case 104:
		//line sql.y:608
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyS[yypt-0].subquery}
		}
	case 105:
		//line sql.y:614
		{
			yyVAL.str = AST_EQ
		}
	case 106:
		//line sql.y:618
		{
			yyVAL.str = AST_LT
		}
	case 107:
		//line sql.y:622
		{
			yyVAL.str = AST_GT
		}
	case 108:
		//line sql.y:626
		{
			yyVAL.str = AST_LE
		}
	case 109:
		//line sql.y:630
		{
			yyVAL.str = AST_GE
		}
	case 110:
		//line sql.y:634
		{
			yyVAL.str = AST_NE
		}
	case 111:
		//line sql.y:638
		{
			yyVAL.str = AST_NSE
		}
	case 112:
		//line sql.y:644
		{
			yyVAL.insRows = yyS[yypt-0].values
		}
	case 113:
		//line sql.y:648
		{
			yyVAL.insRows = yyS[yypt-0].selStmt
		}
	case 114:
		//line sql.y:654
		{
			yyVAL.values = Values{yyS[yypt-0].tuple}
		}
	case 115:
		//line sql.y:658
		{
			yyVAL.values = append(yyS[yypt-2].values, yyS[yypt-0].tuple)
		}
	case 116:
		//line sql.y:664
		{
			yyVAL.tuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 117:
		//line sql.y:668
		{
			yyVAL.tuple = yyS[yypt-0].subquery
		}
	case 118:
		//line sql.y:674
		{
			yyVAL.subquery = &Subquery{yyS[yypt-1].selStmt}
		}
	case 119:
		//line sql.y:680
		{
			yyVAL.valExprs = ValExprs{yyS[yypt-0].valExpr}
		}
	case 120:
		//line sql.y:684
		{
			yyVAL.valExprs = append(yyS[yypt-2].valExprs, yyS[yypt-0].valExpr)
		}
	case 121:
		//line sql.y:690
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 122:
		//line sql.y:694
		{
			yyVAL.valExpr = yyS[yypt-0].colName
		}
	case 123:
		//line sql.y:698
		{
			yyVAL.valExpr = yyS[yypt-0].tuple
		}
	case 124:
		//line sql.y:702
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITAND, Right: yyS[yypt-0].valExpr}
		}
	case 125:
		//line sql.y:706
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITOR, Right: yyS[yypt-0].valExpr}
		}
	case 126:
		//line sql.y:710
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITXOR, Right: yyS[yypt-0].valExpr}
		}
	case 127:
		//line sql.y:714
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_PLUS, Right: yyS[yypt-0].valExpr}
		}
	case 128:
		//line sql.y:718
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MINUS, Right: yyS[yypt-0].valExpr}
		}
	case 129:
		//line sql.y:722
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MULT, Right: yyS[yypt-0].valExpr}
		}
	case 130:
		//line sql.y:726
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_DIV, Right: yyS[yypt-0].valExpr}
		}
	case 131:
		//line sql.y:730
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MOD, Right: yyS[yypt-0].valExpr}
		}
	case 132:
		//line sql.y:734
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
	case 133:
		//line sql.y:749
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-2].bytes}
		}
	case 134:
		//line sql.y:753
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 135:
		//line sql.y:757
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-4].bytes, Distinct: true, Exprs: yyS[yypt-1].selectExprs}
		}
	case 136:
		//line sql.y:761
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 137:
		//line sql.y:765
		{
			yyVAL.valExpr = yyS[yypt-0].caseExpr
		}
	case 138:
		//line sql.y:771
		{
			yyVAL.bytes = IF_BYTES
		}
	case 139:
		//line sql.y:775
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 140:
		//line sql.y:781
		{
			yyVAL.byt = AST_UPLUS
		}
	case 141:
		//line sql.y:785
		{
			yyVAL.byt = AST_UMINUS
		}
	case 142:
		//line sql.y:789
		{
			yyVAL.byt = AST_TILDA
		}
	case 143:
		//line sql.y:795
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyS[yypt-3].valExpr, Whens: yyS[yypt-2].whens, Else: yyS[yypt-1].valExpr}
		}
	case 144:
		//line sql.y:800
		{
			yyVAL.valExpr = nil
		}
	case 145:
		//line sql.y:804
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 146:
		//line sql.y:810
		{
			yyVAL.whens = []*When{yyS[yypt-0].when}
		}
	case 147:
		//line sql.y:814
		{
			yyVAL.whens = append(yyS[yypt-1].whens, yyS[yypt-0].when)
		}
	case 148:
		//line sql.y:820
		{
			yyVAL.when = &When{Cond: yyS[yypt-2].boolExpr, Val: yyS[yypt-0].valExpr}
		}
	case 149:
		//line sql.y:825
		{
			yyVAL.valExpr = nil
		}
	case 150:
		//line sql.y:829
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 151:
		//line sql.y:835
		{
			yyVAL.colName = &ColName{Name: yyS[yypt-0].bytes}
		}
	case 152:
		//line sql.y:839
		{
			yyVAL.colName = &ColName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 153:
		//line sql.y:845
		{
			yyVAL.valExpr = StrVal(yyS[yypt-0].bytes)
		}
	case 154:
		//line sql.y:849
		{
			yyVAL.valExpr = NumVal(yyS[yypt-0].bytes)
		}
	case 155:
		//line sql.y:853
		{
			yyVAL.valExpr = ValArg(yyS[yypt-0].bytes)
		}
	case 156:
		//line sql.y:857
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 157:
		//line sql.y:862
		{
			yyVAL.valExprs = nil
		}
	case 158:
		//line sql.y:866
		{
			yyVAL.valExprs = yyS[yypt-0].valExprs
		}
	case 159:
		//line sql.y:871
		{
			yyVAL.boolExpr = nil
		}
	case 160:
		//line sql.y:875
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 161:
		//line sql.y:880
		{
			yyVAL.orderBy = nil
		}
	case 162:
		//line sql.y:884
		{
			yyVAL.orderBy = yyS[yypt-0].orderBy
		}
	case 163:
		//line sql.y:890
		{
			yyVAL.orderBy = OrderBy{yyS[yypt-0].order}
		}
	case 164:
		//line sql.y:894
		{
			yyVAL.orderBy = append(yyS[yypt-2].orderBy, yyS[yypt-0].order)
		}
	case 165:
		//line sql.y:900
		{
			yyVAL.order = &Order{Expr: yyS[yypt-1].valExpr, Direction: yyS[yypt-0].str}
		}
	case 166:
		//line sql.y:905
		{
			yyVAL.str = AST_ASC
		}
	case 167:
		//line sql.y:909
		{
			yyVAL.str = AST_ASC
		}
	case 168:
		//line sql.y:913
		{
			yyVAL.str = AST_DESC
		}
	case 169:
		//line sql.y:918
		{
			yyVAL.limit = nil
		}
	case 170:
		//line sql.y:922
		{
			yyVAL.limit = &Limit{Rowcount: yyS[yypt-0].valExpr}
		}
	case 171:
		//line sql.y:926
		{
			yyVAL.limit = &Limit{Offset: yyS[yypt-2].valExpr, Rowcount: yyS[yypt-0].valExpr}
		}
	case 172:
		//line sql.y:931
		{
			yyVAL.str = ""
		}
	case 173:
		//line sql.y:935
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 174:
		//line sql.y:939
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
	case 175:
		//line sql.y:952
		{
			yyVAL.columns = nil
		}
	case 176:
		//line sql.y:956
		{
			yyVAL.columns = yyS[yypt-1].columns
		}
	case 177:
		//line sql.y:962
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyS[yypt-0].colName}}
		}
	case 178:
		//line sql.y:966
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyS[yypt-0].colName})
		}
	case 179:
		//line sql.y:971
		{
			yyVAL.updateExprs = nil
		}
	case 180:
		//line sql.y:975
		{
			yyVAL.updateExprs = yyS[yypt-0].updateExprs
		}
	case 181:
		//line sql.y:981
		{
			yyVAL.updateExprs = UpdateExprs{yyS[yypt-0].updateExpr}
		}
	case 182:
		//line sql.y:985
		{
			yyVAL.updateExprs = append(yyS[yypt-2].updateExprs, yyS[yypt-0].updateExpr)
		}
	case 183:
		//line sql.y:991
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyS[yypt-2].colName, Expr: yyS[yypt-0].valExpr}
		}
	case 184:
		//line sql.y:996
		{
			yyVAL.empty = struct{}{}
		}
	case 185:
		//line sql.y:998
		{
			yyVAL.empty = struct{}{}
		}
	case 186:
		//line sql.y:1001
		{
			yyVAL.empty = struct{}{}
		}
	case 187:
		//line sql.y:1003
		{
			yyVAL.empty = struct{}{}
		}
	case 188:
		//line sql.y:1006
		{
			yyVAL.empty = struct{}{}
		}
	case 189:
		//line sql.y:1008
		{
			yyVAL.empty = struct{}{}
		}
	case 190:
		//line sql.y:1012
		{
			yyVAL.empty = struct{}{}
		}
	case 191:
		//line sql.y:1014
		{
			yyVAL.empty = struct{}{}
		}
	case 192:
		//line sql.y:1016
		{
			yyVAL.empty = struct{}{}
		}
	case 193:
		//line sql.y:1018
		{
			yyVAL.empty = struct{}{}
		}
	case 194:
		//line sql.y:1020
		{
			yyVAL.empty = struct{}{}
		}
	case 195:
		//line sql.y:1023
		{
			yyVAL.empty = struct{}{}
		}
	case 196:
		//line sql.y:1025
		{
			yyVAL.empty = struct{}{}
		}
	case 197:
		//line sql.y:1028
		{
			yyVAL.empty = struct{}{}
		}
	case 198:
		//line sql.y:1030
		{
			yyVAL.empty = struct{}{}
		}
	case 199:
		//line sql.y:1033
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		//line sql.y:1035
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		//line sql.y:1039
		{
			yyVAL.bytes = bytes.ToLower(yyS[yypt-0].bytes)
		}
	case 202:
		//line sql.y:1044
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
