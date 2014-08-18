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
const CREATE = 57418
const ALTER = 57419
const DROP = 57420
const RENAME = 57421
const TABLE = 57422
const INDEX = 57423
const VIEW = 57424
const TO = 57425
const IGNORE = 57426
const IF = 57427
const UNIQUE = 57428
const USING = 57429

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

const yyNprod = 208
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 568

var yyAct = []int{

	101, 301, 170, 368, 336, 68, 172, 255, 98, 109,
	246, 211, 128, 293, 248, 187, 182, 99, 70, 195,
	87, 82, 377, 173, 3, 377, 88, 34, 35, 36,
	37, 146, 147, 57, 75, 269, 270, 271, 272, 273,
	92, 274, 275, 104, 72, 318, 320, 77, 108, 262,
	79, 114, 71, 131, 83, 377, 141, 59, 91, 105,
	106, 107, 299, 141, 93, 141, 209, 96, 78, 238,
	209, 112, 379, 347, 44, 378, 46, 127, 346, 51,
	47, 240, 319, 345, 322, 135, 49, 76, 50, 138,
	95, 143, 130, 137, 110, 111, 89, 56, 73, 174,
	327, 115, 280, 175, 247, 376, 326, 145, 124, 52,
	53, 54, 298, 288, 178, 286, 239, 119, 181, 72,
	208, 126, 72, 113, 185, 191, 190, 71, 58, 247,
	71, 291, 146, 147, 228, 342, 169, 171, 192, 294,
	189, 258, 138, 134, 93, 217, 191, 329, 205, 69,
	215, 221, 65, 206, 226, 227, 201, 230, 231, 232,
	233, 234, 235, 236, 237, 344, 218, 222, 216, 146,
	147, 18, 19, 20, 21, 199, 229, 343, 202, 93,
	93, 159, 160, 161, 72, 72, 316, 219, 220, 251,
	312, 121, 71, 253, 315, 313, 259, 242, 244, 22,
	310, 254, 294, 314, 188, 311, 250, 121, 72, 260,
	209, 353, 265, 264, 331, 214, 71, 117, 188, 123,
	120, 263, 363, 215, 213, 279, 266, 282, 283, 257,
	250, 34, 35, 36, 37, 198, 200, 197, 136, 140,
	362, 281, 81, 18, 361, 139, 93, 267, 179, 27,
	28, 29, 116, 30, 32, 31, 290, 23, 24, 26,
	25, 121, 300, 108, 287, 297, 114, 296, 157, 158,
	159, 160, 161, 73, 105, 106, 107, 215, 215, 308,
	309, 177, 139, 141, 176, 86, 112, 325, 292, 269,
	270, 271, 272, 273, 328, 274, 275, 84, 18, 58,
	72, 207, 333, 352, 183, 334, 337, 73, 332, 110,
	111, 323, 184, 278, 144, 184, 115, 338, 154, 155,
	156, 157, 158, 159, 160, 161, 321, 348, 214, 277,
	58, 305, 349, 374, 304, 204, 203, 213, 113, 186,
	66, 132, 129, 125, 138, 122, 80, 357, 359, 351,
	330, 375, 118, 350, 85, 365, 337, 18, 366, 367,
	64, 285, 369, 369, 369, 72, 370, 371, 381, 243,
	372, 104, 302, 71, 193, 133, 108, 18, 382, 114,
	249, 358, 383, 360, 384, 62, 91, 105, 106, 107,
	60, 223, 104, 224, 225, 96, 341, 108, 303, 112,
	114, 256, 340, 307, 188, 67, 380, 73, 105, 106,
	107, 364, 18, 39, 17, 16, 96, 15, 95, 14,
	112, 13, 110, 111, 89, 12, 194, 324, 45, 115,
	154, 155, 156, 157, 158, 159, 160, 161, 261, 95,
	104, 196, 48, 110, 111, 108, 74, 252, 114, 373,
	115, 113, 354, 335, 241, 73, 105, 106, 107, 355,
	356, 339, 306, 108, 96, 289, 114, 180, 112, 245,
	103, 100, 113, 73, 105, 106, 107, 102, 295, 97,
	148, 94, 139, 317, 212, 268, 112, 95, 210, 90,
	276, 110, 111, 149, 153, 151, 152, 142, 115, 61,
	33, 154, 155, 156, 157, 158, 159, 160, 161, 110,
	111, 63, 165, 166, 167, 168, 115, 162, 163, 164,
	113, 284, 11, 10, 154, 155, 156, 157, 158, 159,
	160, 161, 9, 8, 7, 6, 38, 5, 113, 150,
	154, 155, 156, 157, 158, 159, 160, 161, 154, 155,
	156, 157, 158, 159, 160, 161, 40, 41, 42, 43,
	4, 2, 1, 0, 0, 0, 0, 55,
}
var yyPact = []int{

	166, -1000, -1000, 182, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -21, -11, -16, 14, -1000, -1000, -1000,
	-1000, 7, 264, 407, 373, -1000, -1000, -1000, 367, -1000,
	331, 305, 396, 63, -66, -9, 264, -1000, -27, 264,
	-1000, 311, -79, 264, -79, 325, -1000, 241, -1000, -1000,
	-1000, 23, -1000, 213, 305, 319, 41, 305, 154, 310,
	-1000, 174, -1000, 32, 308, 54, 264, -1000, 307, -1000,
	-45, 306, 355, 79, 264, 305, 438, 230, -1000, -1000,
	295, 31, 104, 472, -1000, 420, 372, -1000, -1000, -1000,
	438, 240, 237, -1000, 204, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 438, -1000, 271, 272, 304,
	394, 272, -1000, 438, 264, -1000, 354, -83, -1000, 143,
	-1000, 301, -1000, -1000, 300, -1000, 268, 17, 480, 238,
	180, 23, -1000, -1000, 264, 93, 420, 420, 438, 201,
	370, 438, 438, 109, 438, 438, 438, 438, 438, 438,
	438, 438, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	472, -34, 13, -22, 472, -1000, 351, 23, -1000, 407,
	25, 480, 352, 272, 272, 208, -1000, 388, 420, -1000,
	480, -1000, -1000, -1000, 77, 264, -1000, -49, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 352, 272, -1000, 438,
	194, 235, 294, 293, 26, -1000, -1000, -1000, -1000, -1000,
	-1000, 480, -1000, 201, 438, 438, 480, 456, -1000, 336,
	197, 197, 197, 108, 108, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 12, 23, 10, 50, -1000, 420, 75, 201,
	182, 138, 9, -1000, 388, 357, 384, 104, 299, -1000,
	-1000, 296, -1000, -1000, 154, 480, 392, 180, 180, -1000,
	-1000, 146, 136, 149, 140, 132, -17, -1000, 291, -19,
	276, -1000, 480, 362, 438, -1000, -1000, 3, -1000, 18,
	-1000, 438, 67, -1000, 320, 161, -1000, -1000, -1000, 272,
	357, -1000, 438, 438, -1000, -1000, 390, 382, 235, 71,
	-1000, 123, -1000, 111, -1000, -1000, -1000, -1000, -13, -18,
	-23, -1000, -1000, -1000, 438, 480, -1000, -1000, 480, 438,
	322, 201, -1000, -1000, 250, 158, -1000, 433, -1000, 388,
	420, 438, 420, -1000, -1000, 200, 196, 178, 480, 480,
	404, -1000, 438, 438, -1000, -1000, -1000, 357, 104, 157,
	104, 264, 264, 264, 272, 480, -1000, 317, 2, -1000,
	-28, -31, 154, -1000, 399, 347, -1000, 264, -1000, -1000,
	-1000, 264, -1000, 264, -1000,
}
var yyPgo = []int{

	0, 562, 561, 23, 560, 537, 535, 534, 533, 532,
	523, 522, 536, 511, 500, 499, 20, 26, 497, 490,
	489, 488, 11, 485, 484, 152, 483, 3, 15, 40,
	481, 480, 14, 479, 2, 17, 6, 478, 477, 9,
	471, 8, 470, 469, 10, 467, 465, 462, 461, 7,
	453, 4, 452, 1, 449, 16, 447, 13, 5, 18,
	242, 446, 442, 441, 438, 428, 426, 0, 12, 425,
	421, 419, 417, 415, 414, 413,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 3, 3, 3,
	4, 4, 72, 72, 5, 6, 7, 7, 69, 70,
	71, 74, 73, 8, 8, 8, 9, 9, 9, 10,
	11, 11, 11, 75, 12, 13, 13, 14, 14, 14,
	14, 14, 15, 15, 16, 16, 17, 17, 17, 20,
	20, 18, 18, 18, 21, 21, 22, 22, 22, 22,
	19, 19, 19, 23, 23, 23, 23, 23, 23, 23,
	23, 23, 24, 24, 24, 25, 25, 26, 26, 26,
	26, 27, 27, 28, 28, 29, 29, 29, 29, 29,
	30, 30, 30, 30, 30, 30, 30, 30, 30, 30,
	31, 31, 31, 31, 31, 31, 31, 32, 32, 37,
	37, 35, 35, 39, 36, 36, 34, 34, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	34, 34, 34, 38, 38, 40, 40, 40, 42, 45,
	45, 43, 43, 44, 46, 46, 41, 41, 33, 33,
	33, 33, 47, 47, 48, 48, 49, 49, 50, 50,
	51, 52, 52, 52, 53, 53, 53, 54, 54, 54,
	55, 55, 56, 56, 57, 57, 58, 58, 59, 60,
	60, 61, 61, 62, 62, 63, 63, 63, 63, 63,
	64, 64, 65, 65, 66, 66, 67, 68,
}
var yyR2 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 4, 12, 3,
	7, 7, 6, 6, 8, 7, 3, 4, 1, 1,
	1, 5, 2, 5, 8, 4, 6, 7, 4, 5,
	4, 5, 5, 0, 2, 0, 2, 1, 2, 1,
	1, 1, 0, 1, 1, 3, 1, 2, 3, 1,
	1, 0, 1, 2, 1, 3, 3, 3, 3, 5,
	0, 1, 2, 1, 1, 2, 3, 2, 3, 2,
	2, 2, 1, 3, 1, 1, 3, 0, 5, 5,
	5, 1, 3, 0, 2, 1, 3, 3, 2, 3,
	3, 3, 4, 3, 4, 5, 6, 3, 4, 2,
	1, 1, 1, 1, 1, 1, 1, 2, 1, 1,
	3, 3, 1, 3, 1, 3, 1, 1, 1, 3,
	3, 3, 3, 3, 3, 3, 3, 2, 3, 4,
	5, 4, 1, 1, 1, 1, 1, 1, 5, 0,
	1, 1, 2, 4, 0, 2, 1, 3, 1, 1,
	1, 1, 0, 3, 0, 2, 0, 3, 1, 3,
	2, 0, 1, 1, 0, 2, 4, 0, 2, 4,
	0, 3, 1, 3, 0, 5, 1, 3, 3, 0,
	2, 0, 3, 0, 1, 1, 1, 1, 1, 1,
	0, 1, 0, 1, 0, 2, 1, 0,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -69, -70, -71, -72, -73, -74, 5, 6,
	7, 8, 33, 91, 92, 94, 93, 83, 84, 85,
	87, 89, 88, -14, 49, 50, 51, 52, -12, -75,
	-12, -12, -12, -12, 95, -65, 97, 101, -62, 97,
	99, 95, 95, 96, 97, -12, 90, -67, 35, -3,
	17, -15, 18, -13, 29, -25, 35, 9, -58, 86,
	-59, -41, -67, 35, -61, 100, 96, -67, 95, -67,
	35, -60, 100, -67, -60, 29, 44, -16, -17, 73,
	-20, 35, -29, -34, -30, 67, 44, -33, -41, -35,
	-40, -67, -38, -42, 20, 36, 37, 38, 25, -39,
	71, 72, 48, 100, 28, 78, 39, -25, 33, 76,
	-25, 53, 35, 45, 76, 35, 67, -67, -68, 35,
	-68, 98, 35, 20, 64, -67, -25, -36, -34, 44,
	9, 53, -18, -67, 19, 76, 65, 66, -31, 21,
	67, 23, 24, 22, 68, 69, 70, 71, 72, 73,
	74, 75, 45, 46, 47, 40, 41, 42, 43, -29,
	-34, -29, -36, -3, -34, -34, 44, 44, -39, 44,
	-45, -34, -55, 33, 44, -58, 35, -28, 10, -59,
	-34, -67, -68, 20, -66, 102, -63, 94, 92, 32,
	93, 13, 35, 35, 35, -68, -55, 33, 103, 53,
	-21, -22, -24, 44, 35, -39, -17, -67, 73, -29,
	-29, -34, -35, 21, 23, 24, -34, -34, 25, 67,
	-34, -34, -34, -34, -34, -34, -34, -34, 103, 103,
	103, 103, -16, 18, -16, -43, -44, 79, -32, 28,
	-3, -58, -56, -41, -28, -49, 13, -29, 64, -67,
	-68, -64, 98, -32, -58, -34, -28, 53, -23, 54,
	55, 56, 57, 58, 60, 61, -19, 35, 19, -22,
	76, -35, -34, -34, 65, 25, 103, -16, 103, -46,
	-44, 81, -29, -57, 64, -37, -35, -57, 103, 53,
	-49, -53, 15, 14, 35, 35, -47, 11, -22, -22,
	54, 59, 54, 59, 54, 54, 54, -26, 62, 99,
	63, 35, 103, 35, 65, -34, 103, 82, -34, 80,
	30, 53, -41, -53, -34, -50, -51, -34, -68, -48,
	12, 14, 64, 54, 54, 96, 96, 96, -34, -34,
	31, -35, 53, 53, -52, 26, 27, -49, -29, -36,
	-29, 44, 44, 44, 7, -34, -51, -53, -27, -67,
	-27, -27, -58, -54, 16, 34, 103, 53, 103, 103,
	7, 21, -67, -67, -67,
}
var yyDef = []int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 43, 43,
	43, 43, 43, 202, 193, 0, 0, 28, 29, 30,
	43, 0, 0, 0, 47, 49, 50, 51, 52, 45,
	0, 0, 0, 0, 191, 0, 0, 203, 0, 0,
	194, 0, 189, 0, 189, 0, 32, 0, 206, 19,
	48, 0, 53, 44, 0, 0, 85, 0, 26, 0,
	186, 0, 156, 206, 0, 0, 0, 207, 0, 207,
	0, 0, 0, 0, 0, 0, 0, 17, 54, 56,
	61, 206, 59, 60, 95, 0, 0, 126, 127, 128,
	0, 156, 0, 142, 0, 158, 159, 160, 161, 122,
	145, 146, 147, 143, 144, 149, 46, 180, 0, 0,
	93, 0, 27, 0, 0, 207, 0, 204, 35, 0,
	38, 0, 40, 190, 0, 207, 180, 0, 124, 0,
	0, 0, 57, 62, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 110, 111, 112, 113, 114, 115, 116, 98,
	0, 0, 0, 0, 124, 137, 0, 0, 109, 0,
	0, 150, 0, 0, 0, 93, 86, 166, 0, 187,
	188, 157, 33, 192, 0, 0, 207, 200, 195, 196,
	197, 198, 199, 39, 41, 42, 0, 0, 31, 0,
	93, 64, 70, 0, 82, 84, 55, 63, 58, 96,
	97, 100, 101, 0, 0, 0, 103, 0, 107, 0,
	129, 130, 131, 132, 133, 134, 135, 136, 99, 121,
	123, 138, 0, 0, 0, 154, 151, 0, 184, 0,
	118, 184, 0, 182, 166, 174, 0, 94, 0, 205,
	36, 0, 201, 22, 23, 125, 162, 0, 0, 73,
	74, 0, 0, 0, 0, 0, 87, 71, 0, 0,
	0, 102, 104, 0, 0, 108, 139, 0, 141, 0,
	152, 0, 0, 20, 0, 117, 119, 21, 181, 0,
	174, 25, 0, 0, 207, 37, 164, 0, 65, 68,
	75, 0, 77, 0, 79, 80, 81, 66, 0, 0,
	0, 72, 67, 83, 0, 105, 140, 148, 155, 0,
	0, 0, 183, 24, 175, 167, 168, 171, 34, 166,
	0, 0, 0, 76, 78, 0, 0, 0, 106, 153,
	0, 120, 0, 0, 170, 172, 173, 174, 165, 163,
	69, 0, 0, 0, 0, 176, 169, 177, 0, 91,
	0, 0, 185, 18, 0, 0, 88, 0, 89, 90,
	178, 0, 92, 0, 179,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 75, 68, 3,
	44, 103, 73, 71, 53, 72, 76, 74, 3, 3,
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
	97, 98, 99, 100, 101, 102,
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
		//line sql.y:167
		{
			SetParseTree(yylex, yyS[yypt-0].statement)
		}
	case 2:
		//line sql.y:173
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
		//line sql.y:193
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyS[yypt-2].bytes2), Distinct: yyS[yypt-1].str, SelectExprs: yyS[yypt-0].selectExprs}
		}
	case 18:
		//line sql.y:197
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyS[yypt-10].bytes2), Distinct: yyS[yypt-9].str, SelectExprs: yyS[yypt-8].selectExprs, From: yyS[yypt-6].tableExprs, Where: NewWhere(AST_WHERE, yyS[yypt-5].boolExpr), GroupBy: GroupBy(yyS[yypt-4].valExprs), Having: NewWhere(AST_HAVING, yyS[yypt-3].boolExpr), OrderBy: yyS[yypt-2].orderBy, Limit: yyS[yypt-1].limit, Lock: yyS[yypt-0].str}
		}
	case 19:
		//line sql.y:201
		{
			yyVAL.selStmt = &Union{Type: yyS[yypt-1].str, Left: yyS[yypt-2].selStmt, Right: yyS[yypt-0].selStmt}
		}
	case 20:
		//line sql.y:208
		{
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: yyS[yypt-2].columns, Rows: yyS[yypt-1].insRows, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 21:
		//line sql.y:212
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
		//line sql.y:224
		{
			yyVAL.statement = &Replace{Comments: Comments(yyS[yypt-4].bytes2), Table: yyS[yypt-2].tableName, Columns: yyS[yypt-1].columns, Rows: yyS[yypt-0].insRows}
		}
	case 23:
		//line sql.y:228
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
		//line sql.y:241
		{
			yyVAL.statement = &Update{Comments: Comments(yyS[yypt-6].bytes2), Table: yyS[yypt-5].tableName, Exprs: yyS[yypt-3].updateExprs, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 25:
		//line sql.y:247
		{
			yyVAL.statement = &Delete{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 26:
		//line sql.y:253
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-1].bytes2), Exprs: yyS[yypt-0].updateExprs}
		}
	case 27:
		//line sql.y:257
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: StrVal(yyS[yypt-0].bytes)}}}
		}
	case 28:
		//line sql.y:263
		{
			yyVAL.statement = &Begin{}
		}
	case 29:
		//line sql.y:269
		{
			yyVAL.statement = &Commit{}
		}
	case 30:
		//line sql.y:275
		{
			yyVAL.statement = &Rollback{}
		}
	case 31:
		//line sql.y:281
		{
			yyVAL.statement = &Admin{Name: yyS[yypt-3].bytes, Values: yyS[yypt-1].valExprs}
		}
	case 32:
		//line sql.y:287
		{
			yyVAL.statement = &Show{Name: "databases"}
		}
	case 33:
		//line sql.y:293
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 34:
		//line sql.y:297
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 35:
		//line sql.y:302
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 36:
		//line sql.y:308
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-2].bytes}
		}
	case 37:
		//line sql.y:312
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-3].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 38:
		//line sql.y:317
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 39:
		//line sql.y:323
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 40:
		//line sql.y:329
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-0].bytes}
		}
	case 41:
		//line sql.y:333
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 42:
		//line sql.y:338
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-1].bytes}
		}
	case 43:
		//line sql.y:343
		{
			SetAllowComments(yylex, true)
		}
	case 44:
		//line sql.y:347
		{
			yyVAL.bytes2 = yyS[yypt-0].bytes2
			SetAllowComments(yylex, false)
		}
	case 45:
		//line sql.y:353
		{
			yyVAL.bytes2 = nil
		}
	case 46:
		//line sql.y:357
		{
			yyVAL.bytes2 = append(yyS[yypt-1].bytes2, yyS[yypt-0].bytes)
		}
	case 47:
		//line sql.y:363
		{
			yyVAL.str = AST_UNION
		}
	case 48:
		//line sql.y:367
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 49:
		//line sql.y:371
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 50:
		//line sql.y:375
		{
			yyVAL.str = AST_EXCEPT
		}
	case 51:
		//line sql.y:379
		{
			yyVAL.str = AST_INTERSECT
		}
	case 52:
		//line sql.y:384
		{
			yyVAL.str = ""
		}
	case 53:
		//line sql.y:388
		{
			yyVAL.str = AST_DISTINCT
		}
	case 54:
		//line sql.y:394
		{
			yyVAL.selectExprs = SelectExprs{yyS[yypt-0].selectExpr}
		}
	case 55:
		//line sql.y:398
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyS[yypt-0].selectExpr)
		}
	case 56:
		//line sql.y:404
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 57:
		//line sql.y:408
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyS[yypt-1].expr, As: yyS[yypt-0].bytes}
		}
	case 58:
		//line sql.y:412
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyS[yypt-2].bytes}
		}
	case 59:
		//line sql.y:418
		{
			yyVAL.expr = yyS[yypt-0].boolExpr
		}
	case 60:
		//line sql.y:422
		{
			yyVAL.expr = yyS[yypt-0].valExpr
		}
	case 61:
		//line sql.y:427
		{
			yyVAL.bytes = nil
		}
	case 62:
		//line sql.y:431
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 63:
		//line sql.y:435
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 64:
		//line sql.y:441
		{
			yyVAL.tableExprs = TableExprs{yyS[yypt-0].tableExpr}
		}
	case 65:
		//line sql.y:445
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyS[yypt-0].tableExpr)
		}
	case 66:
		//line sql.y:451
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyS[yypt-2].smTableExpr, As: yyS[yypt-1].bytes, Hints: yyS[yypt-0].indexHints}
		}
	case 67:
		//line sql.y:455
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyS[yypt-1].tableExpr}
		}
	case 68:
		//line sql.y:459
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-2].tableExpr, Join: yyS[yypt-1].str, RightExpr: yyS[yypt-0].tableExpr}
		}
	case 69:
		//line sql.y:463
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-4].tableExpr, Join: yyS[yypt-3].str, RightExpr: yyS[yypt-2].tableExpr, On: yyS[yypt-0].boolExpr}
		}
	case 70:
		//line sql.y:468
		{
			yyVAL.bytes = nil
		}
	case 71:
		//line sql.y:472
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 72:
		//line sql.y:476
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 73:
		//line sql.y:482
		{
			yyVAL.str = AST_JOIN
		}
	case 74:
		//line sql.y:486
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 75:
		//line sql.y:490
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 76:
		//line sql.y:494
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 77:
		//line sql.y:498
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 78:
		//line sql.y:502
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 79:
		//line sql.y:506
		{
			yyVAL.str = AST_JOIN
		}
	case 80:
		//line sql.y:510
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 81:
		//line sql.y:514
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 82:
		//line sql.y:520
		{
			yyVAL.smTableExpr = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 83:
		//line sql.y:524
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 84:
		//line sql.y:528
		{
			yyVAL.smTableExpr = yyS[yypt-0].subquery
		}
	case 85:
		//line sql.y:534
		{
			yyVAL.tableName = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 86:
		//line sql.y:538
		{
			yyVAL.tableName = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 87:
		//line sql.y:543
		{
			yyVAL.indexHints = nil
		}
	case 88:
		//line sql.y:547
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyS[yypt-1].bytes2}
		}
	case 89:
		//line sql.y:551
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyS[yypt-1].bytes2}
		}
	case 90:
		//line sql.y:555
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyS[yypt-1].bytes2}
		}
	case 91:
		//line sql.y:561
		{
			yyVAL.bytes2 = [][]byte{yyS[yypt-0].bytes}
		}
	case 92:
		//line sql.y:565
		{
			yyVAL.bytes2 = append(yyS[yypt-2].bytes2, yyS[yypt-0].bytes)
		}
	case 93:
		//line sql.y:570
		{
			yyVAL.boolExpr = nil
		}
	case 94:
		//line sql.y:574
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 95:
		yyVAL.boolExpr = yyS[yypt-0].boolExpr
	case 96:
		//line sql.y:581
		{
			yyVAL.boolExpr = &AndExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 97:
		//line sql.y:585
		{
			yyVAL.boolExpr = &OrExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 98:
		//line sql.y:589
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyS[yypt-0].boolExpr}
		}
	case 99:
		//line sql.y:593
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyS[yypt-1].boolExpr}
		}
	case 100:
		//line sql.y:599
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: yyS[yypt-1].str, Right: yyS[yypt-0].valExpr}
		}
	case 101:
		//line sql.y:603
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_IN, Right: yyS[yypt-0].tuple}
		}
	case 102:
		//line sql.y:607
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_IN, Right: yyS[yypt-0].tuple}
		}
	case 103:
		//line sql.y:611
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 104:
		//line sql.y:615
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 105:
		//line sql.y:619
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-4].valExpr, Operator: AST_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 106:
		//line sql.y:623
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-5].valExpr, Operator: AST_NOT_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 107:
		//line sql.y:627
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyS[yypt-2].valExpr}
		}
	case 108:
		//line sql.y:631
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyS[yypt-3].valExpr}
		}
	case 109:
		//line sql.y:635
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyS[yypt-0].subquery}
		}
	case 110:
		//line sql.y:641
		{
			yyVAL.str = AST_EQ
		}
	case 111:
		//line sql.y:645
		{
			yyVAL.str = AST_LT
		}
	case 112:
		//line sql.y:649
		{
			yyVAL.str = AST_GT
		}
	case 113:
		//line sql.y:653
		{
			yyVAL.str = AST_LE
		}
	case 114:
		//line sql.y:657
		{
			yyVAL.str = AST_GE
		}
	case 115:
		//line sql.y:661
		{
			yyVAL.str = AST_NE
		}
	case 116:
		//line sql.y:665
		{
			yyVAL.str = AST_NSE
		}
	case 117:
		//line sql.y:671
		{
			yyVAL.insRows = yyS[yypt-0].values
		}
	case 118:
		//line sql.y:675
		{
			yyVAL.insRows = yyS[yypt-0].selStmt
		}
	case 119:
		//line sql.y:681
		{
			yyVAL.values = Values{yyS[yypt-0].tuple}
		}
	case 120:
		//line sql.y:685
		{
			yyVAL.values = append(yyS[yypt-2].values, yyS[yypt-0].tuple)
		}
	case 121:
		//line sql.y:691
		{
			yyVAL.tuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 122:
		//line sql.y:695
		{
			yyVAL.tuple = yyS[yypt-0].subquery
		}
	case 123:
		//line sql.y:701
		{
			yyVAL.subquery = &Subquery{yyS[yypt-1].selStmt}
		}
	case 124:
		//line sql.y:707
		{
			yyVAL.valExprs = ValExprs{yyS[yypt-0].valExpr}
		}
	case 125:
		//line sql.y:711
		{
			yyVAL.valExprs = append(yyS[yypt-2].valExprs, yyS[yypt-0].valExpr)
		}
	case 126:
		//line sql.y:717
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 127:
		//line sql.y:721
		{
			yyVAL.valExpr = yyS[yypt-0].colName
		}
	case 128:
		//line sql.y:725
		{
			yyVAL.valExpr = yyS[yypt-0].tuple
		}
	case 129:
		//line sql.y:729
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITAND, Right: yyS[yypt-0].valExpr}
		}
	case 130:
		//line sql.y:733
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITOR, Right: yyS[yypt-0].valExpr}
		}
	case 131:
		//line sql.y:737
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITXOR, Right: yyS[yypt-0].valExpr}
		}
	case 132:
		//line sql.y:741
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_PLUS, Right: yyS[yypt-0].valExpr}
		}
	case 133:
		//line sql.y:745
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MINUS, Right: yyS[yypt-0].valExpr}
		}
	case 134:
		//line sql.y:749
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MULT, Right: yyS[yypt-0].valExpr}
		}
	case 135:
		//line sql.y:753
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_DIV, Right: yyS[yypt-0].valExpr}
		}
	case 136:
		//line sql.y:757
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MOD, Right: yyS[yypt-0].valExpr}
		}
	case 137:
		//line sql.y:761
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
	case 138:
		//line sql.y:776
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-2].bytes}
		}
	case 139:
		//line sql.y:780
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 140:
		//line sql.y:784
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-4].bytes, Distinct: true, Exprs: yyS[yypt-1].selectExprs}
		}
	case 141:
		//line sql.y:788
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 142:
		//line sql.y:792
		{
			yyVAL.valExpr = yyS[yypt-0].caseExpr
		}
	case 143:
		//line sql.y:798
		{
			yyVAL.bytes = IF_BYTES
		}
	case 144:
		//line sql.y:802
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 145:
		//line sql.y:808
		{
			yyVAL.byt = AST_UPLUS
		}
	case 146:
		//line sql.y:812
		{
			yyVAL.byt = AST_UMINUS
		}
	case 147:
		//line sql.y:816
		{
			yyVAL.byt = AST_TILDA
		}
	case 148:
		//line sql.y:822
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyS[yypt-3].valExpr, Whens: yyS[yypt-2].whens, Else: yyS[yypt-1].valExpr}
		}
	case 149:
		//line sql.y:827
		{
			yyVAL.valExpr = nil
		}
	case 150:
		//line sql.y:831
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 151:
		//line sql.y:837
		{
			yyVAL.whens = []*When{yyS[yypt-0].when}
		}
	case 152:
		//line sql.y:841
		{
			yyVAL.whens = append(yyS[yypt-1].whens, yyS[yypt-0].when)
		}
	case 153:
		//line sql.y:847
		{
			yyVAL.when = &When{Cond: yyS[yypt-2].boolExpr, Val: yyS[yypt-0].valExpr}
		}
	case 154:
		//line sql.y:852
		{
			yyVAL.valExpr = nil
		}
	case 155:
		//line sql.y:856
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 156:
		//line sql.y:862
		{
			yyVAL.colName = &ColName{Name: yyS[yypt-0].bytes}
		}
	case 157:
		//line sql.y:866
		{
			yyVAL.colName = &ColName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 158:
		//line sql.y:872
		{
			yyVAL.valExpr = StrVal(yyS[yypt-0].bytes)
		}
	case 159:
		//line sql.y:876
		{
			yyVAL.valExpr = NumVal(yyS[yypt-0].bytes)
		}
	case 160:
		//line sql.y:880
		{
			yyVAL.valExpr = ValArg(yyS[yypt-0].bytes)
		}
	case 161:
		//line sql.y:884
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 162:
		//line sql.y:889
		{
			yyVAL.valExprs = nil
		}
	case 163:
		//line sql.y:893
		{
			yyVAL.valExprs = yyS[yypt-0].valExprs
		}
	case 164:
		//line sql.y:898
		{
			yyVAL.boolExpr = nil
		}
	case 165:
		//line sql.y:902
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 166:
		//line sql.y:907
		{
			yyVAL.orderBy = nil
		}
	case 167:
		//line sql.y:911
		{
			yyVAL.orderBy = yyS[yypt-0].orderBy
		}
	case 168:
		//line sql.y:917
		{
			yyVAL.orderBy = OrderBy{yyS[yypt-0].order}
		}
	case 169:
		//line sql.y:921
		{
			yyVAL.orderBy = append(yyS[yypt-2].orderBy, yyS[yypt-0].order)
		}
	case 170:
		//line sql.y:927
		{
			yyVAL.order = &Order{Expr: yyS[yypt-1].valExpr, Direction: yyS[yypt-0].str}
		}
	case 171:
		//line sql.y:932
		{
			yyVAL.str = AST_ASC
		}
	case 172:
		//line sql.y:936
		{
			yyVAL.str = AST_ASC
		}
	case 173:
		//line sql.y:940
		{
			yyVAL.str = AST_DESC
		}
	case 174:
		//line sql.y:945
		{
			yyVAL.limit = nil
		}
	case 175:
		//line sql.y:949
		{
			yyVAL.limit = &Limit{Rowcount: yyS[yypt-0].valExpr}
		}
	case 176:
		//line sql.y:953
		{
			yyVAL.limit = &Limit{Offset: yyS[yypt-2].valExpr, Rowcount: yyS[yypt-0].valExpr}
		}
	case 177:
		//line sql.y:958
		{
			yyVAL.str = ""
		}
	case 178:
		//line sql.y:962
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 179:
		//line sql.y:966
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
	case 180:
		//line sql.y:979
		{
			yyVAL.columns = nil
		}
	case 181:
		//line sql.y:983
		{
			yyVAL.columns = yyS[yypt-1].columns
		}
	case 182:
		//line sql.y:989
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyS[yypt-0].colName}}
		}
	case 183:
		//line sql.y:993
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyS[yypt-0].colName})
		}
	case 184:
		//line sql.y:998
		{
			yyVAL.updateExprs = nil
		}
	case 185:
		//line sql.y:1002
		{
			yyVAL.updateExprs = yyS[yypt-0].updateExprs
		}
	case 186:
		//line sql.y:1008
		{
			yyVAL.updateExprs = UpdateExprs{yyS[yypt-0].updateExpr}
		}
	case 187:
		//line sql.y:1012
		{
			yyVAL.updateExprs = append(yyS[yypt-2].updateExprs, yyS[yypt-0].updateExpr)
		}
	case 188:
		//line sql.y:1018
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyS[yypt-2].colName, Expr: yyS[yypt-0].valExpr}
		}
	case 189:
		//line sql.y:1023
		{
			yyVAL.empty = struct{}{}
		}
	case 190:
		//line sql.y:1025
		{
			yyVAL.empty = struct{}{}
		}
	case 191:
		//line sql.y:1028
		{
			yyVAL.empty = struct{}{}
		}
	case 192:
		//line sql.y:1030
		{
			yyVAL.empty = struct{}{}
		}
	case 193:
		//line sql.y:1033
		{
			yyVAL.empty = struct{}{}
		}
	case 194:
		//line sql.y:1035
		{
			yyVAL.empty = struct{}{}
		}
	case 195:
		//line sql.y:1039
		{
			yyVAL.empty = struct{}{}
		}
	case 196:
		//line sql.y:1041
		{
			yyVAL.empty = struct{}{}
		}
	case 197:
		//line sql.y:1043
		{
			yyVAL.empty = struct{}{}
		}
	case 198:
		//line sql.y:1045
		{
			yyVAL.empty = struct{}{}
		}
	case 199:
		//line sql.y:1047
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		//line sql.y:1050
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		//line sql.y:1052
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		//line sql.y:1055
		{
			yyVAL.empty = struct{}{}
		}
	case 203:
		//line sql.y:1057
		{
			yyVAL.empty = struct{}{}
		}
	case 204:
		//line sql.y:1060
		{
			yyVAL.empty = struct{}{}
		}
	case 205:
		//line sql.y:1062
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		//line sql.y:1066
		{
			yyVAL.bytes = bytes.ToLower(yyS[yypt-0].bytes)
		}
	case 207:
		//line sql.y:1071
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
