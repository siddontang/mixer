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
const ADMIN = 57415
const SHOW = 57416
const DATABASES = 57417
const TABLES = 57418
const CREATE = 57419
const ALTER = 57420
const DROP = 57421
const RENAME = 57422
const TABLE = 57423
const INDEX = 57424
const VIEW = 57425
const TO = 57426
const IGNORE = 57427
const IF = 57428
const UNIQUE = 57429
const USING = 57430

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
	"ADMIN",
	"SHOW",
	"DATABASES",
	"TABLES",
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

const yyNprod = 214
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 600

var yyAct = []int{

	104, 310, 177, 377, 69, 179, 345, 131, 101, 264,
	302, 112, 255, 257, 91, 194, 102, 220, 71, 180,
	3, 90, 34, 35, 36, 37, 202, 153, 154, 83,
	386, 386, 189, 58, 386, 44, 148, 46, 308, 148,
	271, 47, 95, 76, 73, 148, 218, 78, 218, 49,
	80, 50, 72, 60, 84, 134, 107, 52, 53, 54,
	356, 111, 327, 329, 117, 96, 247, 355, 354, 77,
	79, 94, 108, 109, 110, 51, 74, 249, 130, 336,
	99, 388, 387, 256, 115, 385, 138, 335, 133, 307,
	297, 143, 146, 289, 150, 145, 295, 248, 152, 217,
	328, 127, 181, 98, 56, 57, 182, 113, 114, 92,
	256, 122, 300, 237, 118, 129, 153, 154, 59, 185,
	351, 188, 73, 153, 154, 73, 192, 70, 198, 197,
	72, 338, 303, 72, 267, 137, 199, 116, 66, 321,
	353, 176, 178, 196, 322, 216, 212, 146, 166, 167,
	168, 96, 226, 198, 352, 238, 227, 82, 230, 224,
	325, 235, 236, 225, 239, 240, 241, 242, 243, 244,
	245, 246, 213, 231, 278, 279, 280, 281, 282, 208,
	283, 284, 124, 195, 215, 319, 96, 96, 324, 323,
	320, 73, 73, 303, 124, 260, 228, 229, 206, 72,
	262, 209, 218, 268, 120, 251, 253, 123, 263, 259,
	126, 269, 85, 362, 147, 73, 340, 59, 372, 273,
	371, 274, 370, 72, 331, 139, 276, 272, 164, 165,
	166, 167, 168, 259, 224, 275, 291, 292, 266, 223,
	288, 144, 195, 186, 18, 19, 20, 21, 222, 290,
	278, 279, 280, 281, 282, 96, 283, 284, 148, 205,
	207, 204, 361, 34, 35, 36, 37, 299, 18, 184,
	183, 306, 22, 309, 296, 305, 214, 161, 162, 163,
	164, 165, 166, 167, 168, 124, 190, 191, 224, 224,
	89, 119, 287, 74, 317, 318, 334, 191, 223, 301,
	364, 365, 151, 337, 332, 330, 314, 222, 286, 73,
	313, 342, 383, 211, 343, 346, 210, 341, 59, 193,
	67, 347, 27, 28, 29, 135, 30, 32, 31, 132,
	384, 23, 24, 26, 25, 128, 357, 125, 81, 38,
	121, 358, 161, 162, 163, 164, 165, 166, 167, 168,
	359, 339, 86, 146, 18, 65, 368, 360, 366, 40,
	41, 42, 43, 294, 374, 346, 390, 200, 376, 375,
	55, 378, 378, 378, 73, 379, 380, 258, 381, 252,
	18, 107, 72, 136, 63, 61, 111, 391, 141, 117,
	350, 392, 367, 393, 369, 107, 94, 108, 109, 110,
	111, 142, 311, 117, 232, 99, 233, 234, 312, 115,
	74, 108, 109, 110, 265, 349, 316, 195, 88, 99,
	68, 389, 373, 115, 18, 39, 140, 87, 98, 17,
	16, 15, 113, 114, 92, 14, 13, 12, 201, 118,
	107, 45, 98, 270, 203, 111, 113, 114, 117, 48,
	18, 75, 261, 118, 382, 74, 108, 109, 110, 363,
	344, 348, 116, 315, 99, 250, 298, 187, 115, 254,
	111, 106, 103, 117, 105, 304, 116, 100, 155, 97,
	74, 108, 109, 110, 326, 221, 277, 98, 219, 144,
	93, 113, 114, 115, 285, 149, 111, 62, 118, 117,
	33, 64, 11, 10, 9, 8, 74, 108, 109, 110,
	156, 160, 158, 159, 7, 144, 113, 114, 6, 115,
	5, 116, 4, 118, 2, 1, 0, 0, 0, 172,
	173, 174, 175, 0, 169, 170, 171, 0, 0, 0,
	0, 0, 113, 114, 0, 0, 116, 0, 0, 118,
	0, 0, 0, 0, 0, 0, 157, 161, 162, 163,
	164, 165, 166, 167, 168, 0, 0, 0, 0, 0,
	333, 0, 116, 161, 162, 163, 164, 165, 166, 167,
	168, 293, 0, 0, 161, 162, 163, 164, 165, 166,
	167, 168, 161, 162, 163, 164, 165, 166, 167, 168,
}
var yyPact = []int{

	239, -1000, -1000, 214, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -61, -49, -21, -39, -1000, -1000, -1000,
	-1000, 14, 182, 419, 368, -1000, -1000, -1000, 366, -1000,
	326, 285, 411, 41, -58, -28, 182, -1000, -26, 182,
	-1000, 303, -72, 182, -72, 323, -1000, 409, 246, -1000,
	-1000, -1000, 36, -1000, 252, 285, 307, 35, 285, 141,
	302, -1000, 165, -1000, 25, 300, 48, 182, -1000, 294,
	-1000, -44, 290, 363, 71, 182, 285, 378, 471, 471,
	205, -1000, -1000, 283, 22, 58, 489, -1000, 420, 375,
	-1000, -1000, -1000, 471, 226, 225, -1000, 199, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 471, -1000,
	253, 258, 284, 407, 258, -1000, 471, 182, -1000, 347,
	-77, -1000, 166, -1000, 281, -1000, -1000, 278, -1000, 243,
	-1000, 420, 471, 524, 445, -5, 524, 204, 36, -1000,
	-1000, 182, 83, 420, 420, 471, 197, 383, 471, 471,
	88, 471, 471, 471, 471, 471, 471, 471, 471, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 489, -38, -7,
	-27, 489, -1000, 361, 36, -1000, 419, 4, 524, 349,
	258, 258, 232, -1000, 401, 420, -1000, 524, -1000, -1000,
	-1000, 70, 182, -1000, -59, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 349, 258, 58, 524, -1000, 471, 173,
	196, 273, 263, 17, -1000, -1000, -1000, -1000, -1000, -1000,
	524, -1000, 197, 471, 471, 524, 516, -1000, 338, 157,
	157, 157, 75, 75, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -8, 36, -14, 31, -1000, 420, 68, 197, 214,
	129, -15, -1000, 401, 387, 394, 58, 275, -1000, -1000,
	271, -1000, -1000, 141, 524, 405, 204, 204, -1000, -1000,
	131, 85, 135, 134, 106, 0, -1000, 270, 120, 269,
	-1000, 524, 505, 471, -1000, -1000, -17, -1000, -3, -1000,
	471, 51, -1000, 321, 163, -1000, -1000, -1000, 258, 387,
	-1000, 471, 471, -1000, -1000, 403, 376, 196, 56, -1000,
	100, -1000, 86, -1000, -1000, -1000, -1000, -29, -30, -37,
	-1000, -1000, -1000, 471, 524, -1000, -1000, 524, 471, 319,
	197, -1000, -1000, 209, 160, -1000, 274, -1000, 401, 420,
	471, 420, -1000, -1000, 178, 176, 174, 524, 524, 415,
	-1000, 471, 471, -1000, -1000, -1000, 387, 58, 149, 58,
	182, 182, 182, 258, 524, -1000, 296, -19, -1000, -22,
	-23, 141, -1000, 414, 345, -1000, 182, -1000, -1000, -1000,
	182, -1000, 182, -1000,
}
var yyPgo = []int{

	0, 525, 524, 19, 522, 520, 518, 514, 505, 504,
	503, 502, 339, 501, 500, 497, 21, 14, 495, 494,
	490, 488, 17, 486, 485, 138, 484, 3, 15, 42,
	479, 478, 13, 477, 2, 16, 5, 475, 474, 11,
	472, 8, 471, 469, 12, 467, 466, 463, 461, 9,
	460, 6, 459, 1, 454, 32, 452, 10, 4, 18,
	157, 451, 449, 444, 443, 441, 438, 0, 7, 437,
	436, 435, 431, 430, 429, 427, 426, 425,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 3, 3, 3,
	4, 4, 72, 72, 5, 6, 7, 7, 69, 70,
	71, 74, 73, 73, 8, 8, 8, 9, 9, 9,
	10, 11, 11, 11, 77, 12, 13, 13, 14, 14,
	14, 14, 14, 15, 15, 16, 16, 17, 17, 17,
	20, 20, 18, 18, 18, 21, 21, 22, 22, 22,
	22, 19, 19, 19, 23, 23, 23, 23, 23, 23,
	23, 23, 23, 24, 24, 24, 25, 25, 26, 26,
	26, 26, 27, 27, 28, 28, 76, 76, 76, 75,
	75, 29, 29, 29, 29, 29, 30, 30, 30, 30,
	30, 30, 30, 30, 30, 30, 31, 31, 31, 31,
	31, 31, 31, 32, 32, 37, 37, 35, 35, 39,
	36, 36, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 38,
	38, 40, 40, 40, 42, 45, 45, 43, 43, 44,
	46, 46, 41, 41, 33, 33, 33, 33, 47, 47,
	48, 48, 49, 49, 50, 50, 51, 52, 52, 52,
	53, 53, 53, 54, 54, 54, 55, 55, 56, 56,
	57, 57, 58, 58, 59, 60, 60, 61, 61, 62,
	62, 63, 63, 63, 63, 63, 64, 64, 65, 65,
	66, 66, 67, 68,
}
var yyR2 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 4, 12, 3,
	7, 7, 6, 6, 8, 7, 3, 4, 1, 1,
	1, 5, 2, 4, 5, 8, 4, 6, 7, 4,
	5, 4, 5, 5, 0, 2, 0, 2, 1, 2,
	1, 1, 1, 0, 1, 1, 3, 1, 2, 3,
	1, 1, 0, 1, 2, 1, 3, 3, 3, 3,
	5, 0, 1, 2, 1, 1, 2, 3, 2, 3,
	2, 2, 2, 1, 3, 1, 1, 3, 0, 5,
	5, 5, 1, 3, 0, 2, 0, 2, 2, 0,
	2, 1, 3, 3, 2, 3, 3, 3, 4, 3,
	4, 5, 6, 3, 4, 2, 1, 1, 1, 1,
	1, 1, 1, 2, 1, 1, 3, 3, 1, 3,
	1, 3, 1, 1, 1, 3, 3, 3, 3, 3,
	3, 3, 3, 2, 3, 4, 5, 4, 1, 1,
	1, 1, 1, 1, 5, 0, 1, 1, 2, 4,
	0, 2, 1, 3, 1, 1, 1, 1, 0, 3,
	0, 2, 0, 3, 1, 3, 2, 0, 1, 1,
	0, 2, 4, 0, 2, 4, 0, 3, 1, 3,
	0, 5, 1, 3, 3, 0, 2, 0, 3, 0,
	1, 1, 1, 1, 1, 1, 0, 1, 0, 1,
	0, 2, 1, 0,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -69, -70, -71, -72, -73, -74, 5, 6,
	7, 8, 33, 92, 93, 95, 94, 83, 84, 85,
	87, 89, 88, -14, 49, 50, 51, 52, -12, -77,
	-12, -12, -12, -12, 96, -65, 98, 102, -62, 98,
	100, 96, 96, 97, 98, -12, 90, 91, -67, 35,
	-3, 17, -15, 18, -13, 29, -25, 35, 9, -58,
	86, -59, -41, -67, 35, -61, 101, 97, -67, 96,
	-67, 35, -60, 101, -67, -60, 29, -75, 9, 44,
	-16, -17, 73, -20, 35, -29, -34, -30, 67, 44,
	-33, -41, -35, -40, -67, -38, -42, 20, 36, 37,
	38, 25, -39, 71, 72, 48, 101, 28, 78, 39,
	-25, 33, 76, -25, 53, 35, 45, 76, 35, 67,
	-67, -68, 35, -68, 99, 35, 20, 64, -67, -25,
	-76, 10, 23, -34, 44, -36, -34, 9, 53, -18,
	-67, 19, 76, 65, 66, -31, 21, 67, 23, 24,
	22, 68, 69, 70, 71, 72, 73, 74, 75, 45,
	46, 47, 40, 41, 42, 43, -29, -34, -29, -36,
	-3, -34, -34, 44, 44, -39, 44, -45, -34, -55,
	33, 44, -58, 35, -28, 10, -59, -34, -67, -68,
	20, -66, 103, -63, 95, 93, 32, 94, 13, 35,
	35, 35, -68, -55, 33, -29, -34, 104, 53, -21,
	-22, -24, 44, 35, -39, -17, -67, 73, -29, -29,
	-34, -35, 21, 23, 24, -34, -34, 25, 67, -34,
	-34, -34, -34, -34, -34, -34, -34, 104, 104, 104,
	104, -16, 18, -16, -43, -44, 79, -32, 28, -3,
	-58, -56, -41, -28, -49, 13, -29, 64, -67, -68,
	-64, 99, -32, -58, -34, -28, 53, -23, 54, 55,
	56, 57, 58, 60, 61, -19, 35, 19, -22, 76,
	-35, -34, -34, 65, 25, 104, -16, 104, -46, -44,
	81, -29, -57, 64, -37, -35, -57, 104, 53, -49,
	-53, 15, 14, 35, 35, -47, 11, -22, -22, 54,
	59, 54, 59, 54, 54, 54, -26, 62, 100, 63,
	35, 104, 35, 65, -34, 104, 82, -34, 80, 30,
	53, -41, -53, -34, -50, -51, -34, -68, -48, 12,
	14, 64, 54, 54, 97, 97, 97, -34, -34, 31,
	-35, 53, 53, -52, 26, 27, -49, -29, -36, -29,
	44, 44, 44, 7, -34, -51, -53, -27, -67, -27,
	-27, -58, -54, 16, 34, 104, 53, 104, 104, 7,
	21, -67, -67, -67,
}
var yyDef = []int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 44, 44,
	44, 44, 44, 208, 199, 0, 0, 28, 29, 30,
	44, 0, 0, 0, 48, 50, 51, 52, 53, 46,
	0, 0, 0, 0, 197, 0, 0, 209, 0, 0,
	200, 0, 195, 0, 195, 0, 32, 99, 0, 212,
	19, 49, 0, 54, 45, 0, 0, 86, 0, 26,
	0, 192, 0, 162, 212, 0, 0, 0, 213, 0,
	213, 0, 0, 0, 0, 0, 0, 96, 0, 0,
	17, 55, 57, 62, 212, 60, 61, 101, 0, 0,
	132, 133, 134, 0, 162, 0, 148, 0, 164, 165,
	166, 167, 128, 151, 152, 153, 149, 150, 155, 47,
	186, 0, 0, 94, 0, 27, 0, 0, 213, 0,
	210, 36, 0, 39, 0, 41, 196, 0, 213, 186,
	33, 0, 0, 100, 0, 0, 130, 0, 0, 58,
	63, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 116,
	117, 118, 119, 120, 121, 122, 104, 0, 0, 0,
	0, 130, 143, 0, 0, 115, 0, 0, 156, 0,
	0, 0, 94, 87, 172, 0, 193, 194, 163, 34,
	198, 0, 0, 213, 206, 201, 202, 203, 204, 205,
	40, 42, 43, 0, 0, 97, 98, 31, 0, 94,
	65, 71, 0, 83, 85, 56, 64, 59, 102, 103,
	106, 107, 0, 0, 0, 109, 0, 113, 0, 135,
	136, 137, 138, 139, 140, 141, 142, 105, 127, 129,
	144, 0, 0, 0, 160, 157, 0, 190, 0, 124,
	190, 0, 188, 172, 180, 0, 95, 0, 211, 37,
	0, 207, 22, 23, 131, 168, 0, 0, 74, 75,
	0, 0, 0, 0, 0, 88, 72, 0, 0, 0,
	108, 110, 0, 0, 114, 145, 0, 147, 0, 158,
	0, 0, 20, 0, 123, 125, 21, 187, 0, 180,
	25, 0, 0, 213, 38, 170, 0, 66, 69, 76,
	0, 78, 0, 80, 81, 82, 67, 0, 0, 0,
	73, 68, 84, 0, 111, 146, 154, 161, 0, 0,
	0, 189, 24, 181, 173, 174, 177, 35, 172, 0,
	0, 0, 77, 79, 0, 0, 0, 112, 159, 0,
	126, 0, 0, 176, 178, 179, 180, 171, 169, 70,
	0, 0, 0, 0, 182, 175, 183, 0, 92, 0,
	0, 191, 18, 0, 0, 89, 0, 90, 91, 184,
	0, 93, 0, 185,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 75, 68, 3,
	44, 104, 73, 71, 53, 72, 76, 74, 3, 3,
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
	97, 98, 99, 100, 101, 102, 103,
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
		//line sql.y:170
		{
			SetParseTree(yylex, yyS[yypt-0].statement)
		}
	case 2:
		//line sql.y:176
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
		yyVAL.statement = yyS[yypt-0].statement
	case 16:
		yyVAL.statement = yyS[yypt-0].statement
	case 17:
		//line sql.y:196
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyS[yypt-2].bytes2), Distinct: yyS[yypt-1].str, SelectExprs: yyS[yypt-0].selectExprs}
		}
	case 18:
		//line sql.y:200
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyS[yypt-10].bytes2), Distinct: yyS[yypt-9].str, SelectExprs: yyS[yypt-8].selectExprs, From: yyS[yypt-6].tableExprs, Where: NewWhere(AST_WHERE, yyS[yypt-5].boolExpr), GroupBy: GroupBy(yyS[yypt-4].valExprs), Having: NewWhere(AST_HAVING, yyS[yypt-3].boolExpr), OrderBy: yyS[yypt-2].orderBy, Limit: yyS[yypt-1].limit, Lock: yyS[yypt-0].str}
		}
	case 19:
		//line sql.y:204
		{
			yyVAL.selStmt = &Union{Type: yyS[yypt-1].str, Left: yyS[yypt-2].selStmt, Right: yyS[yypt-0].selStmt}
		}
	case 20:
		//line sql.y:211
		{
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: yyS[yypt-2].columns, Rows: yyS[yypt-1].insRows, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 21:
		//line sql.y:215
		{
			cols := make(Columns, 0, len(yyS[yypt-1].updateExprs))
			vals := make(ValTuple, 0, len(yyS[yypt-1].updateExprs))
			for _, col := range yyS[yypt-1].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 22:
		//line sql.y:227
		{
			yyVAL.statement = &Replace{Comments: Comments(yyS[yypt-4].bytes2), Table: yyS[yypt-2].tableName, Columns: yyS[yypt-1].columns, Rows: yyS[yypt-0].insRows}
		}
	case 23:
		//line sql.y:231
		{
			cols := make(Columns, 0, len(yyS[yypt-0].updateExprs))
			vals := make(ValTuple, 0, len(yyS[yypt-0].updateExprs))
			for _, col := range yyS[yypt-0].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Replace{Comments: Comments(yyS[yypt-4].bytes2), Table: yyS[yypt-2].tableName, Columns: cols, Rows: Values{vals}}
		}
	case 24:
		//line sql.y:244
		{
			yyVAL.statement = &Update{Comments: Comments(yyS[yypt-6].bytes2), Table: yyS[yypt-5].tableName, Exprs: yyS[yypt-3].updateExprs, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 25:
		//line sql.y:250
		{
			yyVAL.statement = &Delete{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 26:
		//line sql.y:256
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-1].bytes2), Exprs: yyS[yypt-0].updateExprs}
		}
	case 27:
		//line sql.y:260
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: StrVal(yyS[yypt-0].bytes)}}}
		}
	case 28:
		//line sql.y:266
		{
			yyVAL.statement = &Begin{}
		}
	case 29:
		//line sql.y:272
		{
			yyVAL.statement = &Commit{}
		}
	case 30:
		//line sql.y:278
		{
			yyVAL.statement = &Rollback{}
		}
	case 31:
		//line sql.y:284
		{
			yyVAL.statement = &Admin{Name: yyS[yypt-3].bytes, Values: yyS[yypt-1].valExprs}
		}
	case 32:
		//line sql.y:290
		{
			yyVAL.statement = &Show{Name: "databases"}
		}
	case 33:
		//line sql.y:294
		{
			yyVAL.statement = &Show{Name: "tables", From: yyS[yypt-1].valExpr, LikeOrWhere: yyS[yypt-0].expr}
		}
	case 34:
		//line sql.y:300
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 35:
		//line sql.y:304
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 36:
		//line sql.y:309
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 37:
		//line sql.y:315
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-2].bytes}
		}
	case 38:
		//line sql.y:319
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-3].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 39:
		//line sql.y:324
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 40:
		//line sql.y:330
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 41:
		//line sql.y:336
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-0].bytes}
		}
	case 42:
		//line sql.y:340
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 43:
		//line sql.y:345
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-1].bytes}
		}
	case 44:
		//line sql.y:350
		{
			SetAllowComments(yylex, true)
		}
	case 45:
		//line sql.y:354
		{
			yyVAL.bytes2 = yyS[yypt-0].bytes2
			SetAllowComments(yylex, false)
		}
	case 46:
		//line sql.y:360
		{
			yyVAL.bytes2 = nil
		}
	case 47:
		//line sql.y:364
		{
			yyVAL.bytes2 = append(yyS[yypt-1].bytes2, yyS[yypt-0].bytes)
		}
	case 48:
		//line sql.y:370
		{
			yyVAL.str = AST_UNION
		}
	case 49:
		//line sql.y:374
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 50:
		//line sql.y:378
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 51:
		//line sql.y:382
		{
			yyVAL.str = AST_EXCEPT
		}
	case 52:
		//line sql.y:386
		{
			yyVAL.str = AST_INTERSECT
		}
	case 53:
		//line sql.y:391
		{
			yyVAL.str = ""
		}
	case 54:
		//line sql.y:395
		{
			yyVAL.str = AST_DISTINCT
		}
	case 55:
		//line sql.y:401
		{
			yyVAL.selectExprs = SelectExprs{yyS[yypt-0].selectExpr}
		}
	case 56:
		//line sql.y:405
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyS[yypt-0].selectExpr)
		}
	case 57:
		//line sql.y:411
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 58:
		//line sql.y:415
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyS[yypt-1].expr, As: yyS[yypt-0].bytes}
		}
	case 59:
		//line sql.y:419
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyS[yypt-2].bytes}
		}
	case 60:
		//line sql.y:425
		{
			yyVAL.expr = yyS[yypt-0].boolExpr
		}
	case 61:
		//line sql.y:429
		{
			yyVAL.expr = yyS[yypt-0].valExpr
		}
	case 62:
		//line sql.y:434
		{
			yyVAL.bytes = nil
		}
	case 63:
		//line sql.y:438
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 64:
		//line sql.y:442
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 65:
		//line sql.y:448
		{
			yyVAL.tableExprs = TableExprs{yyS[yypt-0].tableExpr}
		}
	case 66:
		//line sql.y:452
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyS[yypt-0].tableExpr)
		}
	case 67:
		//line sql.y:458
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyS[yypt-2].smTableExpr, As: yyS[yypt-1].bytes, Hints: yyS[yypt-0].indexHints}
		}
	case 68:
		//line sql.y:462
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyS[yypt-1].tableExpr}
		}
	case 69:
		//line sql.y:466
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-2].tableExpr, Join: yyS[yypt-1].str, RightExpr: yyS[yypt-0].tableExpr}
		}
	case 70:
		//line sql.y:470
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-4].tableExpr, Join: yyS[yypt-3].str, RightExpr: yyS[yypt-2].tableExpr, On: yyS[yypt-0].boolExpr}
		}
	case 71:
		//line sql.y:475
		{
			yyVAL.bytes = nil
		}
	case 72:
		//line sql.y:479
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 73:
		//line sql.y:483
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 74:
		//line sql.y:489
		{
			yyVAL.str = AST_JOIN
		}
	case 75:
		//line sql.y:493
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 76:
		//line sql.y:497
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 77:
		//line sql.y:501
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 78:
		//line sql.y:505
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 79:
		//line sql.y:509
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 80:
		//line sql.y:513
		{
			yyVAL.str = AST_JOIN
		}
	case 81:
		//line sql.y:517
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 82:
		//line sql.y:521
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 83:
		//line sql.y:527
		{
			yyVAL.smTableExpr = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 84:
		//line sql.y:531
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 85:
		//line sql.y:535
		{
			yyVAL.smTableExpr = yyS[yypt-0].subquery
		}
	case 86:
		//line sql.y:541
		{
			yyVAL.tableName = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 87:
		//line sql.y:545
		{
			yyVAL.tableName = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 88:
		//line sql.y:550
		{
			yyVAL.indexHints = nil
		}
	case 89:
		//line sql.y:554
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyS[yypt-1].bytes2}
		}
	case 90:
		//line sql.y:558
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyS[yypt-1].bytes2}
		}
	case 91:
		//line sql.y:562
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyS[yypt-1].bytes2}
		}
	case 92:
		//line sql.y:568
		{
			yyVAL.bytes2 = [][]byte{yyS[yypt-0].bytes}
		}
	case 93:
		//line sql.y:572
		{
			yyVAL.bytes2 = append(yyS[yypt-2].bytes2, yyS[yypt-0].bytes)
		}
	case 94:
		//line sql.y:577
		{
			yyVAL.boolExpr = nil
		}
	case 95:
		//line sql.y:581
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 96:
		//line sql.y:586
		{
			yyVAL.expr = nil
		}
	case 97:
		//line sql.y:590
		{
			yyVAL.expr = yyS[yypt-0].boolExpr
		}
	case 98:
		//line sql.y:594
		{
			yyVAL.expr = yyS[yypt-0].valExpr
		}
	case 99:
		//line sql.y:599
		{
			yyVAL.valExpr = nil
		}
	case 100:
		//line sql.y:603
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 101:
		yyVAL.boolExpr = yyS[yypt-0].boolExpr
	case 102:
		//line sql.y:610
		{
			yyVAL.boolExpr = &AndExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 103:
		//line sql.y:614
		{
			yyVAL.boolExpr = &OrExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 104:
		//line sql.y:618
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyS[yypt-0].boolExpr}
		}
	case 105:
		//line sql.y:622
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyS[yypt-1].boolExpr}
		}
	case 106:
		//line sql.y:628
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: yyS[yypt-1].str, Right: yyS[yypt-0].valExpr}
		}
	case 107:
		//line sql.y:632
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_IN, Right: yyS[yypt-0].tuple}
		}
	case 108:
		//line sql.y:636
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_IN, Right: yyS[yypt-0].tuple}
		}
	case 109:
		//line sql.y:640
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 110:
		//line sql.y:644
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 111:
		//line sql.y:648
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-4].valExpr, Operator: AST_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 112:
		//line sql.y:652
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-5].valExpr, Operator: AST_NOT_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 113:
		//line sql.y:656
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyS[yypt-2].valExpr}
		}
	case 114:
		//line sql.y:660
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyS[yypt-3].valExpr}
		}
	case 115:
		//line sql.y:664
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyS[yypt-0].subquery}
		}
	case 116:
		//line sql.y:670
		{
			yyVAL.str = AST_EQ
		}
	case 117:
		//line sql.y:674
		{
			yyVAL.str = AST_LT
		}
	case 118:
		//line sql.y:678
		{
			yyVAL.str = AST_GT
		}
	case 119:
		//line sql.y:682
		{
			yyVAL.str = AST_LE
		}
	case 120:
		//line sql.y:686
		{
			yyVAL.str = AST_GE
		}
	case 121:
		//line sql.y:690
		{
			yyVAL.str = AST_NE
		}
	case 122:
		//line sql.y:694
		{
			yyVAL.str = AST_NSE
		}
	case 123:
		//line sql.y:700
		{
			yyVAL.insRows = yyS[yypt-0].values
		}
	case 124:
		//line sql.y:704
		{
			yyVAL.insRows = yyS[yypt-0].selStmt
		}
	case 125:
		//line sql.y:710
		{
			yyVAL.values = Values{yyS[yypt-0].tuple}
		}
	case 126:
		//line sql.y:714
		{
			yyVAL.values = append(yyS[yypt-2].values, yyS[yypt-0].tuple)
		}
	case 127:
		//line sql.y:720
		{
			yyVAL.tuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 128:
		//line sql.y:724
		{
			yyVAL.tuple = yyS[yypt-0].subquery
		}
	case 129:
		//line sql.y:730
		{
			yyVAL.subquery = &Subquery{yyS[yypt-1].selStmt}
		}
	case 130:
		//line sql.y:736
		{
			yyVAL.valExprs = ValExprs{yyS[yypt-0].valExpr}
		}
	case 131:
		//line sql.y:740
		{
			yyVAL.valExprs = append(yyS[yypt-2].valExprs, yyS[yypt-0].valExpr)
		}
	case 132:
		//line sql.y:746
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 133:
		//line sql.y:750
		{
			yyVAL.valExpr = yyS[yypt-0].colName
		}
	case 134:
		//line sql.y:754
		{
			yyVAL.valExpr = yyS[yypt-0].tuple
		}
	case 135:
		//line sql.y:758
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITAND, Right: yyS[yypt-0].valExpr}
		}
	case 136:
		//line sql.y:762
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITOR, Right: yyS[yypt-0].valExpr}
		}
	case 137:
		//line sql.y:766
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITXOR, Right: yyS[yypt-0].valExpr}
		}
	case 138:
		//line sql.y:770
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_PLUS, Right: yyS[yypt-0].valExpr}
		}
	case 139:
		//line sql.y:774
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MINUS, Right: yyS[yypt-0].valExpr}
		}
	case 140:
		//line sql.y:778
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MULT, Right: yyS[yypt-0].valExpr}
		}
	case 141:
		//line sql.y:782
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_DIV, Right: yyS[yypt-0].valExpr}
		}
	case 142:
		//line sql.y:786
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MOD, Right: yyS[yypt-0].valExpr}
		}
	case 143:
		//line sql.y:790
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
	case 144:
		//line sql.y:805
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-2].bytes}
		}
	case 145:
		//line sql.y:809
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 146:
		//line sql.y:813
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-4].bytes, Distinct: true, Exprs: yyS[yypt-1].selectExprs}
		}
	case 147:
		//line sql.y:817
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 148:
		//line sql.y:821
		{
			yyVAL.valExpr = yyS[yypt-0].caseExpr
		}
	case 149:
		//line sql.y:827
		{
			yyVAL.bytes = IF_BYTES
		}
	case 150:
		//line sql.y:831
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 151:
		//line sql.y:837
		{
			yyVAL.byt = AST_UPLUS
		}
	case 152:
		//line sql.y:841
		{
			yyVAL.byt = AST_UMINUS
		}
	case 153:
		//line sql.y:845
		{
			yyVAL.byt = AST_TILDA
		}
	case 154:
		//line sql.y:851
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyS[yypt-3].valExpr, Whens: yyS[yypt-2].whens, Else: yyS[yypt-1].valExpr}
		}
	case 155:
		//line sql.y:856
		{
			yyVAL.valExpr = nil
		}
	case 156:
		//line sql.y:860
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 157:
		//line sql.y:866
		{
			yyVAL.whens = []*When{yyS[yypt-0].when}
		}
	case 158:
		//line sql.y:870
		{
			yyVAL.whens = append(yyS[yypt-1].whens, yyS[yypt-0].when)
		}
	case 159:
		//line sql.y:876
		{
			yyVAL.when = &When{Cond: yyS[yypt-2].boolExpr, Val: yyS[yypt-0].valExpr}
		}
	case 160:
		//line sql.y:881
		{
			yyVAL.valExpr = nil
		}
	case 161:
		//line sql.y:885
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 162:
		//line sql.y:891
		{
			yyVAL.colName = &ColName{Name: yyS[yypt-0].bytes}
		}
	case 163:
		//line sql.y:895
		{
			yyVAL.colName = &ColName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 164:
		//line sql.y:901
		{
			yyVAL.valExpr = StrVal(yyS[yypt-0].bytes)
		}
	case 165:
		//line sql.y:905
		{
			yyVAL.valExpr = NumVal(yyS[yypt-0].bytes)
		}
	case 166:
		//line sql.y:909
		{
			yyVAL.valExpr = ValArg(yyS[yypt-0].bytes)
		}
	case 167:
		//line sql.y:913
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 168:
		//line sql.y:918
		{
			yyVAL.valExprs = nil
		}
	case 169:
		//line sql.y:922
		{
			yyVAL.valExprs = yyS[yypt-0].valExprs
		}
	case 170:
		//line sql.y:927
		{
			yyVAL.boolExpr = nil
		}
	case 171:
		//line sql.y:931
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 172:
		//line sql.y:936
		{
			yyVAL.orderBy = nil
		}
	case 173:
		//line sql.y:940
		{
			yyVAL.orderBy = yyS[yypt-0].orderBy
		}
	case 174:
		//line sql.y:946
		{
			yyVAL.orderBy = OrderBy{yyS[yypt-0].order}
		}
	case 175:
		//line sql.y:950
		{
			yyVAL.orderBy = append(yyS[yypt-2].orderBy, yyS[yypt-0].order)
		}
	case 176:
		//line sql.y:956
		{
			yyVAL.order = &Order{Expr: yyS[yypt-1].valExpr, Direction: yyS[yypt-0].str}
		}
	case 177:
		//line sql.y:961
		{
			yyVAL.str = AST_ASC
		}
	case 178:
		//line sql.y:965
		{
			yyVAL.str = AST_ASC
		}
	case 179:
		//line sql.y:969
		{
			yyVAL.str = AST_DESC
		}
	case 180:
		//line sql.y:974
		{
			yyVAL.limit = nil
		}
	case 181:
		//line sql.y:978
		{
			yyVAL.limit = &Limit{Rowcount: yyS[yypt-0].valExpr}
		}
	case 182:
		//line sql.y:982
		{
			yyVAL.limit = &Limit{Offset: yyS[yypt-2].valExpr, Rowcount: yyS[yypt-0].valExpr}
		}
	case 183:
		//line sql.y:987
		{
			yyVAL.str = ""
		}
	case 184:
		//line sql.y:991
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 185:
		//line sql.y:995
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
	case 186:
		//line sql.y:1008
		{
			yyVAL.columns = nil
		}
	case 187:
		//line sql.y:1012
		{
			yyVAL.columns = yyS[yypt-1].columns
		}
	case 188:
		//line sql.y:1018
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyS[yypt-0].colName}}
		}
	case 189:
		//line sql.y:1022
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyS[yypt-0].colName})
		}
	case 190:
		//line sql.y:1027
		{
			yyVAL.updateExprs = nil
		}
	case 191:
		//line sql.y:1031
		{
			yyVAL.updateExprs = yyS[yypt-0].updateExprs
		}
	case 192:
		//line sql.y:1037
		{
			yyVAL.updateExprs = UpdateExprs{yyS[yypt-0].updateExpr}
		}
	case 193:
		//line sql.y:1041
		{
			yyVAL.updateExprs = append(yyS[yypt-2].updateExprs, yyS[yypt-0].updateExpr)
		}
	case 194:
		//line sql.y:1047
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyS[yypt-2].colName, Expr: yyS[yypt-0].valExpr}
		}
	case 195:
		//line sql.y:1052
		{
			yyVAL.empty = struct{}{}
		}
	case 196:
		//line sql.y:1054
		{
			yyVAL.empty = struct{}{}
		}
	case 197:
		//line sql.y:1057
		{
			yyVAL.empty = struct{}{}
		}
	case 198:
		//line sql.y:1059
		{
			yyVAL.empty = struct{}{}
		}
	case 199:
		//line sql.y:1062
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		//line sql.y:1064
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		//line sql.y:1068
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		//line sql.y:1070
		{
			yyVAL.empty = struct{}{}
		}
	case 203:
		//line sql.y:1072
		{
			yyVAL.empty = struct{}{}
		}
	case 204:
		//line sql.y:1074
		{
			yyVAL.empty = struct{}{}
		}
	case 205:
		//line sql.y:1076
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		//line sql.y:1079
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		//line sql.y:1081
		{
			yyVAL.empty = struct{}{}
		}
	case 208:
		//line sql.y:1084
		{
			yyVAL.empty = struct{}{}
		}
	case 209:
		//line sql.y:1086
		{
			yyVAL.empty = struct{}{}
		}
	case 210:
		//line sql.y:1089
		{
			yyVAL.empty = struct{}{}
		}
	case 211:
		//line sql.y:1091
		{
			yyVAL.empty = struct{}{}
		}
	case 212:
		//line sql.y:1095
		{
			yyVAL.bytes = bytes.ToLower(yyS[yypt-0].bytes)
		}
	case 213:
		//line sql.y:1100
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
