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

const yyNprod = 204
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 559

var yyAct = []int{

	94, 292, 160, 359, 61, 85, 327, 162, 91, 121,
	246, 200, 284, 237, 239, 102, 63, 186, 76, 178,
	92, 368, 80, 81, 163, 3, 173, 259, 260, 261,
	262, 263, 16, 264, 265, 30, 31, 32, 33, 68,
	65, 136, 137, 70, 368, 40, 73, 42, 64, 253,
	77, 43, 101, 368, 52, 107, 131, 86, 124, 290,
	131, 131, 66, 98, 99, 100, 72, 229, 370, 47,
	120, 166, 45, 313, 46, 105, 227, 338, 337, 128,
	48, 49, 50, 123, 133, 336, 230, 309, 311, 69,
	217, 369, 164, 318, 159, 161, 165, 238, 103, 104,
	367, 71, 238, 317, 282, 108, 289, 279, 277, 270,
	66, 172, 65, 169, 228, 65, 176, 135, 182, 181,
	64, 310, 117, 64, 106, 136, 137, 112, 183, 136,
	137, 180, 218, 333, 86, 206, 182, 119, 196, 207,
	320, 210, 208, 209, 215, 216, 204, 219, 220, 221,
	222, 223, 224, 225, 226, 205, 197, 149, 150, 151,
	211, 62, 147, 148, 149, 150, 151, 114, 285, 231,
	86, 86, 249, 58, 127, 65, 65, 335, 285, 242,
	334, 346, 347, 64, 244, 248, 192, 250, 303, 301,
	233, 235, 307, 304, 302, 306, 245, 251, 241, 65,
	305, 114, 229, 255, 344, 190, 179, 64, 193, 179,
	322, 116, 254, 354, 269, 130, 272, 273, 204, 256,
	75, 16, 241, 144, 145, 146, 147, 148, 149, 150,
	151, 110, 276, 271, 113, 101, 353, 86, 107, 30,
	31, 32, 33, 109, 283, 66, 98, 99, 100, 257,
	281, 203, 114, 129, 166, 288, 291, 278, 105, 131,
	202, 287, 189, 191, 188, 198, 203, 174, 352, 299,
	300, 78, 166, 204, 204, 202, 175, 316, 175, 170,
	168, 103, 104, 167, 268, 319, 134, 71, 108, 66,
	365, 65, 314, 324, 312, 296, 325, 328, 295, 323,
	267, 195, 71, 194, 177, 329, 59, 106, 366, 321,
	125, 122, 234, 118, 97, 115, 74, 111, 339, 101,
	341, 16, 107, 340, 79, 57, 275, 372, 184, 84,
	98, 99, 100, 126, 53, 231, 55, 349, 89, 351,
	350, 348, 105, 342, 240, 293, 356, 328, 332, 294,
	358, 357, 247, 360, 360, 360, 65, 361, 362, 331,
	363, 88, 179, 34, 64, 103, 104, 82, 298, 373,
	16, 97, 108, 374, 371, 375, 101, 60, 355, 107,
	16, 36, 37, 38, 39, 97, 84, 98, 99, 100,
	101, 106, 51, 107, 232, 89, 35, 15, 14, 105,
	66, 98, 99, 100, 259, 260, 261, 262, 263, 89,
	264, 265, 212, 105, 213, 214, 13, 12, 88, 185,
	41, 252, 103, 104, 82, 16, 17, 18, 19, 108,
	187, 44, 88, 67, 243, 364, 103, 104, 345, 97,
	326, 330, 297, 108, 101, 280, 171, 107, 106, 236,
	96, 93, 95, 20, 66, 98, 99, 100, 286, 90,
	138, 87, 106, 89, 308, 201, 258, 105, 199, 83,
	266, 132, 54, 29, 139, 143, 141, 142, 144, 145,
	146, 147, 148, 149, 150, 151, 88, 56, 11, 10,
	103, 104, 9, 155, 156, 157, 158, 108, 152, 153,
	154, 8, 7, 25, 26, 27, 6, 28, 21, 22,
	24, 23, 5, 4, 343, 2, 106, 1, 0, 0,
	140, 144, 145, 146, 147, 148, 149, 150, 151, 144,
	145, 146, 147, 148, 149, 150, 151, 315, 0, 0,
	144, 145, 146, 147, 148, 149, 150, 151, 274, 0,
	0, 144, 145, 146, 147, 148, 149, 150, 151,
}
var yyPact = []int{

	420, -1000, -1000, 190, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -47, -22, -23, -12, -1000, -1000, -1000, -1000, 375,
	317, -1000, -1000, -1000, 318, -1000, 296, 271, 368, 75,
	-58, -4, 252, -1000, -26, 252, -1000, 281, -79, 252,
	-79, 295, -1000, -1000, 351, -1000, 204, 271, 284, 51,
	271, 148, 280, -1000, 166, -1000, 46, 278, 70, 252,
	-1000, -1000, 276, -1000, -37, 275, 313, 110, 252, 271,
	206, -1000, -1000, 267, 41, 64, 453, -1000, 419, 365,
	-1000, -1000, -1000, 210, 239, 236, -1000, 235, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 210, -1000,
	234, 254, 269, 352, 254, -1000, 210, 252, -1000, 308,
	-82, -1000, 173, -1000, 268, -1000, -1000, 266, -1000, 232,
	231, 351, -1000, -1000, 252, 66, 419, 419, 210, 228,
	391, 210, 210, 65, 210, 210, 210, 210, 210, 210,
	210, 210, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	453, -24, 14, -14, 453, -1000, 27, 294, 351, -1000,
	375, 18, 410, 316, 254, 254, 199, -1000, 339, 419,
	-1000, 410, -1000, -1000, -1000, 108, 252, -1000, -46, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 316, 254, 196,
	350, 265, 216, 33, -1000, -1000, -1000, -1000, -1000, -1000,
	410, -1000, 228, 210, 210, 410, 483, -1000, 301, 91,
	91, 91, 84, 84, -1000, -1000, -1000, -1000, -1000, 210,
	-1000, 410, -1000, 8, 351, 7, 23, -1000, 419, 104,
	228, 190, 114, 6, -1000, 339, 330, 335, 64, 263,
	-1000, -1000, 260, -1000, -1000, 148, 357, 231, 231, -1000,
	-1000, 135, 134, 146, 141, 138, 25, -1000, 259, -27,
	257, -1000, 410, 472, 210, -1000, 410, -1000, 3, -1000,
	11, -1000, 210, 60, -1000, 279, 157, -1000, -1000, -1000,
	254, 330, -1000, 210, 210, -1000, -1000, 347, 334, 350,
	69, -1000, 126, -1000, 123, -1000, -1000, -1000, -1000, -8,
	-15, -16, -1000, -1000, -1000, 210, 410, -1000, -1000, 410,
	210, 289, 228, -1000, -1000, 461, 151, -1000, 155, -1000,
	339, 419, 210, 419, -1000, -1000, 224, 192, 169, 410,
	410, 371, -1000, 210, 210, -1000, -1000, -1000, 330, 64,
	149, 64, 252, 252, 252, 254, 410, -1000, 274, 0,
	-1000, -9, -32, 148, -1000, 367, 306, -1000, 252, -1000,
	-1000, -1000, 252, -1000, 252, -1000,
}
var yyPgo = []int{

	0, 517, 515, 24, 513, 512, 506, 502, 501, 492,
	489, 488, 363, 487, 473, 472, 22, 23, 471, 470,
	469, 468, 11, 466, 465, 173, 464, 3, 19, 5,
	461, 460, 14, 459, 2, 20, 7, 458, 452, 15,
	451, 8, 450, 449, 13, 446, 445, 442, 441, 10,
	440, 6, 438, 1, 435, 26, 434, 12, 4, 16,
	220, 433, 431, 430, 421, 420, 419, 0, 9, 417,
	416, 398, 397, 396,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 3, 3, 3, 4, 4,
	72, 72, 5, 6, 7, 7, 69, 70, 71, 8,
	8, 8, 9, 9, 9, 10, 11, 11, 11, 73,
	12, 13, 13, 14, 14, 14, 14, 14, 15, 15,
	16, 16, 17, 17, 17, 20, 20, 18, 18, 18,
	21, 21, 22, 22, 22, 22, 19, 19, 19, 23,
	23, 23, 23, 23, 23, 23, 23, 23, 24, 24,
	24, 25, 25, 26, 26, 26, 26, 27, 27, 28,
	28, 29, 29, 29, 29, 29, 30, 30, 30, 30,
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
	1, 1, 1, 1, 1, 4, 12, 3, 7, 7,
	6, 6, 8, 7, 3, 4, 1, 1, 1, 5,
	8, 4, 6, 7, 4, 5, 4, 5, 5, 0,
	2, 0, 2, 1, 2, 1, 1, 1, 0, 1,
	1, 3, 1, 2, 3, 1, 1, 0, 1, 2,
	1, 3, 3, 3, 3, 5, 0, 1, 2, 1,
	1, 2, 3, 2, 3, 2, 2, 2, 1, 3,
	1, 1, 3, 0, 5, 5, 5, 1, 3, 0,
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
	-10, -11, -69, -70, -71, -72, 5, 6, 7, 8,
	33, 88, 89, 91, 90, 83, 84, 85, 87, -14,
	49, 50, 51, 52, -12, -73, -12, -12, -12, -12,
	92, -65, 94, 98, -62, 94, 96, 92, 92, 93,
	94, -12, -3, 17, -15, 18, -13, 29, -25, 35,
	9, -58, 86, -59, -41, -67, 35, -61, 97, 93,
	-67, 35, 92, -67, 35, -60, 97, -67, -60, 29,
	-16, -17, 73, -20, 35, -29, -34, -30, 67, 44,
	-33, -41, -35, -40, -67, -38, -42, 20, 36, 37,
	38, 25, -39, 71, 72, 48, 97, 28, 78, 39,
	-25, 33, 76, -25, 53, 35, 45, 76, 35, 67,
	-67, -68, 35, -68, 95, 35, 20, 64, -67, -25,
	9, 53, -18, -67, 19, 76, 65, 66, -31, 21,
	67, 23, 24, 22, 68, 69, 70, 71, 72, 73,
	74, 75, 45, 46, 47, 40, 41, 42, 43, -29,
	-34, -29, -36, -3, -34, -34, 44, 44, 44, -39,
	44, -45, -34, -55, 33, 44, -58, 35, -28, 10,
	-59, -34, -67, -68, 20, -66, 99, -63, 91, 89,
	32, 90, 13, 35, 35, 35, -68, -55, 33, -21,
	-22, -24, 44, 35, -39, -17, -67, 73, -29, -29,
	-34, -35, 21, 23, 24, -34, -34, 25, 67, -34,
	-34, -34, -34, -34, -34, -34, -34, 100, 100, 53,
	100, -34, 100, -16, 18, -16, -43, -44, 79, -32,
	28, -3, -58, -56, -41, -28, -49, 13, -29, 64,
	-67, -68, -64, 95, -32, -58, -28, 53, -23, 54,
	55, 56, 57, 58, 60, 61, -19, 35, 19, -22,
	76, -35, -34, -34, 65, 25, -34, 100, -16, 100,
	-46, -44, 81, -29, -57, 64, -37, -35, -57, 100,
	53, -49, -53, 15, 14, 35, 35, -47, 11, -22,
	-22, 54, 59, 54, 59, 54, 54, 54, -26, 62,
	96, 63, 35, 100, 35, 65, -34, 100, 82, -34,
	80, 30, 53, -41, -53, -34, -50, -51, -34, -68,
	-48, 12, 14, 64, 54, 54, 93, 93, 93, -34,
	-34, 31, -35, 53, 53, -52, 26, 27, -49, -29,
	-36, -29, 44, 44, 44, 7, -34, -51, -53, -27,
	-67, -27, -27, -58, -54, 16, 34, 100, 53, 100,
	100, 7, 21, -67, -67, -67,
}
var yyDef = []int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 39, 39, 39, 39,
	39, 198, 189, 0, 0, 26, 27, 28, 39, 0,
	43, 45, 46, 47, 48, 41, 0, 0, 0, 0,
	187, 0, 0, 199, 0, 0, 190, 0, 185, 0,
	185, 0, 17, 44, 0, 49, 40, 0, 0, 81,
	0, 24, 0, 182, 0, 152, 202, 0, 0, 0,
	203, 202, 0, 203, 0, 0, 0, 0, 0, 0,
	15, 50, 52, 57, 202, 55, 56, 91, 0, 0,
	122, 123, 124, 0, 152, 0, 138, 0, 154, 155,
	156, 157, 118, 141, 142, 143, 139, 140, 145, 42,
	176, 0, 0, 89, 0, 25, 0, 0, 203, 0,
	200, 31, 0, 34, 0, 36, 186, 0, 203, 176,
	0, 0, 53, 58, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 106, 107, 108, 109, 110, 111, 112, 94,
	0, 0, 0, 0, 120, 133, 0, 0, 0, 105,
	0, 0, 146, 0, 0, 0, 89, 82, 162, 0,
	183, 184, 153, 29, 188, 0, 0, 203, 196, 191,
	192, 193, 194, 195, 35, 37, 38, 0, 0, 89,
	60, 66, 0, 78, 80, 51, 59, 54, 92, 93,
	96, 97, 0, 0, 0, 99, 0, 103, 0, 125,
	126, 127, 128, 129, 130, 131, 132, 95, 117, 0,
	119, 120, 134, 0, 0, 0, 150, 147, 0, 180,
	0, 114, 180, 0, 178, 162, 170, 0, 90, 0,
	201, 32, 0, 197, 20, 21, 158, 0, 0, 69,
	70, 0, 0, 0, 0, 0, 83, 67, 0, 0,
	0, 98, 100, 0, 0, 104, 121, 135, 0, 137,
	0, 148, 0, 0, 18, 0, 113, 115, 19, 177,
	0, 170, 23, 0, 0, 203, 33, 160, 0, 61,
	64, 71, 0, 73, 0, 75, 76, 77, 62, 0,
	0, 0, 68, 63, 79, 0, 101, 136, 144, 151,
	0, 0, 0, 179, 22, 171, 163, 164, 167, 30,
	162, 0, 0, 0, 72, 74, 0, 0, 0, 102,
	149, 0, 116, 0, 0, 166, 168, 169, 170, 161,
	159, 65, 0, 0, 0, 0, 172, 165, 173, 0,
	87, 0, 0, 181, 16, 0, 0, 84, 0, 85,
	86, 174, 0, 88, 0, 175,
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
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyS[yypt-2].bytes2), Distinct: yyS[yypt-1].str, SelectExprs: yyS[yypt-0].selectExprs}
		}
	case 16:
		//line sql.y:186
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyS[yypt-10].bytes2), Distinct: yyS[yypt-9].str, SelectExprs: yyS[yypt-8].selectExprs, From: yyS[yypt-6].tableExprs, Where: NewWhere(AST_WHERE, yyS[yypt-5].boolExpr), GroupBy: GroupBy(yyS[yypt-4].valExprs), Having: NewWhere(AST_HAVING, yyS[yypt-3].boolExpr), OrderBy: yyS[yypt-2].orderBy, Limit: yyS[yypt-1].limit, Lock: yyS[yypt-0].str}
		}
	case 17:
		//line sql.y:190
		{
			yyVAL.selStmt = &Union{Type: yyS[yypt-1].str, Left: yyS[yypt-2].selStmt, Right: yyS[yypt-0].selStmt}
		}
	case 18:
		//line sql.y:197
		{
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: yyS[yypt-2].columns, Rows: yyS[yypt-1].insRows, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 19:
		//line sql.y:201
		{
			cols := make(Columns, 0, len(yyS[yypt-1].updateExprs))
			vals := make(ValTuple, 0, len(yyS[yypt-1].updateExprs))
			for _, col := range yyS[yypt-1].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 20:
		//line sql.y:213
		{
			yyVAL.statement = &Replace{Comments: Comments(yyS[yypt-4].bytes2), Table: yyS[yypt-2].tableName, Columns: yyS[yypt-1].columns, Rows: yyS[yypt-0].insRows}
		}
	case 21:
		//line sql.y:217
		{
			cols := make(Columns, 0, len(yyS[yypt-0].updateExprs))
			vals := make(ValTuple, 0, len(yyS[yypt-0].updateExprs))
			for _, col := range yyS[yypt-0].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Replace{Comments: Comments(yyS[yypt-4].bytes2), Table: yyS[yypt-2].tableName, Columns: cols, Rows: Values{vals}}
		}
	case 22:
		//line sql.y:230
		{
			yyVAL.statement = &Update{Comments: Comments(yyS[yypt-6].bytes2), Table: yyS[yypt-5].tableName, Exprs: yyS[yypt-3].updateExprs, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 23:
		//line sql.y:236
		{
			yyVAL.statement = &Delete{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 24:
		//line sql.y:242
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-1].bytes2), Exprs: yyS[yypt-0].updateExprs}
		}
	case 25:
		//line sql.y:246
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: StrVal(yyS[yypt-0].bytes)}}}
		}
	case 26:
		//line sql.y:252
		{
			yyVAL.statement = &Begin{}
		}
	case 27:
		//line sql.y:258
		{
			yyVAL.statement = &Commit{}
		}
	case 28:
		//line sql.y:264
		{
			yyVAL.statement = &Rollback{}
		}
	case 29:
		//line sql.y:270
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 30:
		//line sql.y:274
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 31:
		//line sql.y:279
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 32:
		//line sql.y:285
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-2].bytes}
		}
	case 33:
		//line sql.y:289
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-3].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 34:
		//line sql.y:294
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 35:
		//line sql.y:300
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 36:
		//line sql.y:306
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-0].bytes}
		}
	case 37:
		//line sql.y:310
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 38:
		//line sql.y:315
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-1].bytes}
		}
	case 39:
		//line sql.y:320
		{
			SetAllowComments(yylex, true)
		}
	case 40:
		//line sql.y:324
		{
			yyVAL.bytes2 = yyS[yypt-0].bytes2
			SetAllowComments(yylex, false)
		}
	case 41:
		//line sql.y:330
		{
			yyVAL.bytes2 = nil
		}
	case 42:
		//line sql.y:334
		{
			yyVAL.bytes2 = append(yyS[yypt-1].bytes2, yyS[yypt-0].bytes)
		}
	case 43:
		//line sql.y:340
		{
			yyVAL.str = AST_UNION
		}
	case 44:
		//line sql.y:344
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 45:
		//line sql.y:348
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 46:
		//line sql.y:352
		{
			yyVAL.str = AST_EXCEPT
		}
	case 47:
		//line sql.y:356
		{
			yyVAL.str = AST_INTERSECT
		}
	case 48:
		//line sql.y:361
		{
			yyVAL.str = ""
		}
	case 49:
		//line sql.y:365
		{
			yyVAL.str = AST_DISTINCT
		}
	case 50:
		//line sql.y:371
		{
			yyVAL.selectExprs = SelectExprs{yyS[yypt-0].selectExpr}
		}
	case 51:
		//line sql.y:375
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyS[yypt-0].selectExpr)
		}
	case 52:
		//line sql.y:381
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 53:
		//line sql.y:385
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyS[yypt-1].expr, As: yyS[yypt-0].bytes}
		}
	case 54:
		//line sql.y:389
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyS[yypt-2].bytes}
		}
	case 55:
		//line sql.y:395
		{
			yyVAL.expr = yyS[yypt-0].boolExpr
		}
	case 56:
		//line sql.y:399
		{
			yyVAL.expr = yyS[yypt-0].valExpr
		}
	case 57:
		//line sql.y:404
		{
			yyVAL.bytes = nil
		}
	case 58:
		//line sql.y:408
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 59:
		//line sql.y:412
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 60:
		//line sql.y:418
		{
			yyVAL.tableExprs = TableExprs{yyS[yypt-0].tableExpr}
		}
	case 61:
		//line sql.y:422
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyS[yypt-0].tableExpr)
		}
	case 62:
		//line sql.y:428
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyS[yypt-2].smTableExpr, As: yyS[yypt-1].bytes, Hints: yyS[yypt-0].indexHints}
		}
	case 63:
		//line sql.y:432
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyS[yypt-1].tableExpr}
		}
	case 64:
		//line sql.y:436
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-2].tableExpr, Join: yyS[yypt-1].str, RightExpr: yyS[yypt-0].tableExpr}
		}
	case 65:
		//line sql.y:440
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-4].tableExpr, Join: yyS[yypt-3].str, RightExpr: yyS[yypt-2].tableExpr, On: yyS[yypt-0].boolExpr}
		}
	case 66:
		//line sql.y:445
		{
			yyVAL.bytes = nil
		}
	case 67:
		//line sql.y:449
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 68:
		//line sql.y:453
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 69:
		//line sql.y:459
		{
			yyVAL.str = AST_JOIN
		}
	case 70:
		//line sql.y:463
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 71:
		//line sql.y:467
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 72:
		//line sql.y:471
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 73:
		//line sql.y:475
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 74:
		//line sql.y:479
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 75:
		//line sql.y:483
		{
			yyVAL.str = AST_JOIN
		}
	case 76:
		//line sql.y:487
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 77:
		//line sql.y:491
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 78:
		//line sql.y:497
		{
			yyVAL.smTableExpr = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 79:
		//line sql.y:501
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 80:
		//line sql.y:505
		{
			yyVAL.smTableExpr = yyS[yypt-0].subquery
		}
	case 81:
		//line sql.y:511
		{
			yyVAL.tableName = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 82:
		//line sql.y:515
		{
			yyVAL.tableName = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 83:
		//line sql.y:520
		{
			yyVAL.indexHints = nil
		}
	case 84:
		//line sql.y:524
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyS[yypt-1].bytes2}
		}
	case 85:
		//line sql.y:528
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyS[yypt-1].bytes2}
		}
	case 86:
		//line sql.y:532
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyS[yypt-1].bytes2}
		}
	case 87:
		//line sql.y:538
		{
			yyVAL.bytes2 = [][]byte{yyS[yypt-0].bytes}
		}
	case 88:
		//line sql.y:542
		{
			yyVAL.bytes2 = append(yyS[yypt-2].bytes2, yyS[yypt-0].bytes)
		}
	case 89:
		//line sql.y:547
		{
			yyVAL.boolExpr = nil
		}
	case 90:
		//line sql.y:551
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 91:
		yyVAL.boolExpr = yyS[yypt-0].boolExpr
	case 92:
		//line sql.y:558
		{
			yyVAL.boolExpr = &AndExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 93:
		//line sql.y:562
		{
			yyVAL.boolExpr = &OrExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 94:
		//line sql.y:566
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyS[yypt-0].boolExpr}
		}
	case 95:
		//line sql.y:570
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyS[yypt-1].boolExpr}
		}
	case 96:
		//line sql.y:576
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: yyS[yypt-1].str, Right: yyS[yypt-0].valExpr}
		}
	case 97:
		//line sql.y:580
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_IN, Right: yyS[yypt-0].tuple}
		}
	case 98:
		//line sql.y:584
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_IN, Right: yyS[yypt-0].tuple}
		}
	case 99:
		//line sql.y:588
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 100:
		//line sql.y:592
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 101:
		//line sql.y:596
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-4].valExpr, Operator: AST_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 102:
		//line sql.y:600
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-5].valExpr, Operator: AST_NOT_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 103:
		//line sql.y:604
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyS[yypt-2].valExpr}
		}
	case 104:
		//line sql.y:608
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyS[yypt-3].valExpr}
		}
	case 105:
		//line sql.y:612
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyS[yypt-0].subquery}
		}
	case 106:
		//line sql.y:618
		{
			yyVAL.str = AST_EQ
		}
	case 107:
		//line sql.y:622
		{
			yyVAL.str = AST_LT
		}
	case 108:
		//line sql.y:626
		{
			yyVAL.str = AST_GT
		}
	case 109:
		//line sql.y:630
		{
			yyVAL.str = AST_LE
		}
	case 110:
		//line sql.y:634
		{
			yyVAL.str = AST_GE
		}
	case 111:
		//line sql.y:638
		{
			yyVAL.str = AST_NE
		}
	case 112:
		//line sql.y:642
		{
			yyVAL.str = AST_NSE
		}
	case 113:
		//line sql.y:648
		{
			yyVAL.insRows = yyS[yypt-0].values
		}
	case 114:
		//line sql.y:652
		{
			yyVAL.insRows = yyS[yypt-0].selStmt
		}
	case 115:
		//line sql.y:658
		{
			yyVAL.values = Values{yyS[yypt-0].tuple}
		}
	case 116:
		//line sql.y:662
		{
			yyVAL.values = append(yyS[yypt-2].values, yyS[yypt-0].tuple)
		}
	case 117:
		//line sql.y:668
		{
			yyVAL.tuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 118:
		//line sql.y:672
		{
			yyVAL.tuple = yyS[yypt-0].subquery
		}
	case 119:
		//line sql.y:678
		{
			yyVAL.subquery = &Subquery{yyS[yypt-1].selStmt}
		}
	case 120:
		//line sql.y:684
		{
			yyVAL.valExprs = ValExprs{yyS[yypt-0].valExpr}
		}
	case 121:
		//line sql.y:688
		{
			yyVAL.valExprs = append(yyS[yypt-2].valExprs, yyS[yypt-0].valExpr)
		}
	case 122:
		//line sql.y:694
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 123:
		//line sql.y:698
		{
			yyVAL.valExpr = yyS[yypt-0].colName
		}
	case 124:
		//line sql.y:702
		{
			yyVAL.valExpr = yyS[yypt-0].tuple
		}
	case 125:
		//line sql.y:706
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITAND, Right: yyS[yypt-0].valExpr}
		}
	case 126:
		//line sql.y:710
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITOR, Right: yyS[yypt-0].valExpr}
		}
	case 127:
		//line sql.y:714
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITXOR, Right: yyS[yypt-0].valExpr}
		}
	case 128:
		//line sql.y:718
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_PLUS, Right: yyS[yypt-0].valExpr}
		}
	case 129:
		//line sql.y:722
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MINUS, Right: yyS[yypt-0].valExpr}
		}
	case 130:
		//line sql.y:726
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MULT, Right: yyS[yypt-0].valExpr}
		}
	case 131:
		//line sql.y:730
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_DIV, Right: yyS[yypt-0].valExpr}
		}
	case 132:
		//line sql.y:734
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MOD, Right: yyS[yypt-0].valExpr}
		}
	case 133:
		//line sql.y:738
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
	case 134:
		//line sql.y:753
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-2].bytes}
		}
	case 135:
		//line sql.y:757
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 136:
		//line sql.y:761
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-4].bytes, Distinct: true, Exprs: yyS[yypt-1].selectExprs}
		}
	case 137:
		//line sql.y:765
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 138:
		//line sql.y:769
		{
			yyVAL.valExpr = yyS[yypt-0].caseExpr
		}
	case 139:
		//line sql.y:775
		{
			yyVAL.bytes = IF_BYTES
		}
	case 140:
		//line sql.y:779
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 141:
		//line sql.y:785
		{
			yyVAL.byt = AST_UPLUS
		}
	case 142:
		//line sql.y:789
		{
			yyVAL.byt = AST_UMINUS
		}
	case 143:
		//line sql.y:793
		{
			yyVAL.byt = AST_TILDA
		}
	case 144:
		//line sql.y:799
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyS[yypt-3].valExpr, Whens: yyS[yypt-2].whens, Else: yyS[yypt-1].valExpr}
		}
	case 145:
		//line sql.y:804
		{
			yyVAL.valExpr = nil
		}
	case 146:
		//line sql.y:808
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 147:
		//line sql.y:814
		{
			yyVAL.whens = []*When{yyS[yypt-0].when}
		}
	case 148:
		//line sql.y:818
		{
			yyVAL.whens = append(yyS[yypt-1].whens, yyS[yypt-0].when)
		}
	case 149:
		//line sql.y:824
		{
			yyVAL.when = &When{Cond: yyS[yypt-2].boolExpr, Val: yyS[yypt-0].valExpr}
		}
	case 150:
		//line sql.y:829
		{
			yyVAL.valExpr = nil
		}
	case 151:
		//line sql.y:833
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 152:
		//line sql.y:839
		{
			yyVAL.colName = &ColName{Name: yyS[yypt-0].bytes}
		}
	case 153:
		//line sql.y:843
		{
			yyVAL.colName = &ColName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 154:
		//line sql.y:849
		{
			yyVAL.valExpr = StrVal(yyS[yypt-0].bytes)
		}
	case 155:
		//line sql.y:853
		{
			yyVAL.valExpr = NumVal(yyS[yypt-0].bytes)
		}
	case 156:
		//line sql.y:857
		{
			yyVAL.valExpr = ValArg(yyS[yypt-0].bytes)
		}
	case 157:
		//line sql.y:861
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 158:
		//line sql.y:866
		{
			yyVAL.valExprs = nil
		}
	case 159:
		//line sql.y:870
		{
			yyVAL.valExprs = yyS[yypt-0].valExprs
		}
	case 160:
		//line sql.y:875
		{
			yyVAL.boolExpr = nil
		}
	case 161:
		//line sql.y:879
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 162:
		//line sql.y:884
		{
			yyVAL.orderBy = nil
		}
	case 163:
		//line sql.y:888
		{
			yyVAL.orderBy = yyS[yypt-0].orderBy
		}
	case 164:
		//line sql.y:894
		{
			yyVAL.orderBy = OrderBy{yyS[yypt-0].order}
		}
	case 165:
		//line sql.y:898
		{
			yyVAL.orderBy = append(yyS[yypt-2].orderBy, yyS[yypt-0].order)
		}
	case 166:
		//line sql.y:904
		{
			yyVAL.order = &Order{Expr: yyS[yypt-1].valExpr, Direction: yyS[yypt-0].str}
		}
	case 167:
		//line sql.y:909
		{
			yyVAL.str = AST_ASC
		}
	case 168:
		//line sql.y:913
		{
			yyVAL.str = AST_ASC
		}
	case 169:
		//line sql.y:917
		{
			yyVAL.str = AST_DESC
		}
	case 170:
		//line sql.y:922
		{
			yyVAL.limit = nil
		}
	case 171:
		//line sql.y:926
		{
			yyVAL.limit = &Limit{Rowcount: yyS[yypt-0].valExpr}
		}
	case 172:
		//line sql.y:930
		{
			yyVAL.limit = &Limit{Offset: yyS[yypt-2].valExpr, Rowcount: yyS[yypt-0].valExpr}
		}
	case 173:
		//line sql.y:935
		{
			yyVAL.str = ""
		}
	case 174:
		//line sql.y:939
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 175:
		//line sql.y:943
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
	case 176:
		//line sql.y:956
		{
			yyVAL.columns = nil
		}
	case 177:
		//line sql.y:960
		{
			yyVAL.columns = yyS[yypt-1].columns
		}
	case 178:
		//line sql.y:966
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyS[yypt-0].colName}}
		}
	case 179:
		//line sql.y:970
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyS[yypt-0].colName})
		}
	case 180:
		//line sql.y:975
		{
			yyVAL.updateExprs = nil
		}
	case 181:
		//line sql.y:979
		{
			yyVAL.updateExprs = yyS[yypt-0].updateExprs
		}
	case 182:
		//line sql.y:985
		{
			yyVAL.updateExprs = UpdateExprs{yyS[yypt-0].updateExpr}
		}
	case 183:
		//line sql.y:989
		{
			yyVAL.updateExprs = append(yyS[yypt-2].updateExprs, yyS[yypt-0].updateExpr)
		}
	case 184:
		//line sql.y:995
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyS[yypt-2].colName, Expr: yyS[yypt-0].valExpr}
		}
	case 185:
		//line sql.y:1000
		{
			yyVAL.empty = struct{}{}
		}
	case 186:
		//line sql.y:1002
		{
			yyVAL.empty = struct{}{}
		}
	case 187:
		//line sql.y:1005
		{
			yyVAL.empty = struct{}{}
		}
	case 188:
		//line sql.y:1007
		{
			yyVAL.empty = struct{}{}
		}
	case 189:
		//line sql.y:1010
		{
			yyVAL.empty = struct{}{}
		}
	case 190:
		//line sql.y:1012
		{
			yyVAL.empty = struct{}{}
		}
	case 191:
		//line sql.y:1016
		{
			yyVAL.empty = struct{}{}
		}
	case 192:
		//line sql.y:1018
		{
			yyVAL.empty = struct{}{}
		}
	case 193:
		//line sql.y:1020
		{
			yyVAL.empty = struct{}{}
		}
	case 194:
		//line sql.y:1022
		{
			yyVAL.empty = struct{}{}
		}
	case 195:
		//line sql.y:1024
		{
			yyVAL.empty = struct{}{}
		}
	case 196:
		//line sql.y:1027
		{
			yyVAL.empty = struct{}{}
		}
	case 197:
		//line sql.y:1029
		{
			yyVAL.empty = struct{}{}
		}
	case 198:
		//line sql.y:1032
		{
			yyVAL.empty = struct{}{}
		}
	case 199:
		//line sql.y:1034
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		//line sql.y:1037
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		//line sql.y:1039
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		//line sql.y:1043
		{
			yyVAL.bytes = bytes.ToLower(yyS[yypt-0].bytes)
		}
	case 203:
		//line sql.y:1048
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
