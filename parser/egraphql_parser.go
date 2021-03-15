// Code generated from EGraphQL.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // EGraphQL

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 72, 728,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 4, 55,
	9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4, 60, 9,
	60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65, 9, 65,
	4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9, 70, 4,
	71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75, 4, 76,
	9, 76, 3, 2, 6, 2, 154, 10, 2, 13, 2, 14, 2, 155, 3, 3, 3, 3, 3, 3, 5,
	3, 161, 10, 3, 3, 4, 3, 4, 5, 4, 165, 10, 4, 3, 5, 3, 5, 5, 5, 169, 10,
	5, 3, 5, 5, 5, 172, 10, 5, 3, 5, 5, 5, 175, 10, 5, 3, 5, 3, 5, 3, 5, 5,
	5, 180, 10, 5, 3, 6, 3, 6, 3, 7, 3, 7, 6, 7, 186, 10, 7, 13, 7, 14, 7,
	187, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 5, 8, 195, 10, 8, 3, 9, 5, 9, 198, 10,
	9, 3, 9, 3, 9, 5, 9, 202, 10, 9, 3, 9, 5, 9, 205, 10, 9, 3, 9, 5, 9, 208,
	10, 9, 3, 10, 3, 10, 6, 10, 212, 10, 10, 13, 10, 14, 10, 213, 3, 10, 3,
	10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13,
	5, 13, 228, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 234, 10, 14, 3,
	14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 5, 17, 245,
	10, 17, 3, 17, 5, 17, 248, 10, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 261, 10, 18, 3, 19, 3, 19,
	3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3,
	25, 3, 25, 3, 25, 3, 25, 6, 25, 279, 10, 25, 13, 25, 14, 25, 280, 3, 25,
	3, 25, 5, 25, 285, 10, 25, 3, 26, 3, 26, 7, 26, 289, 10, 26, 12, 26, 14,
	26, 292, 11, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28,
	3, 28, 3, 29, 3, 29, 6, 29, 305, 10, 29, 13, 29, 14, 29, 306, 3, 29, 3,
	29, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 315, 10, 30, 3, 31, 3, 31, 3, 31,
	3, 32, 3, 32, 5, 32, 322, 10, 32, 3, 32, 3, 32, 5, 32, 326, 10, 32, 5,
	32, 328, 10, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 6, 35,
	337, 10, 35, 13, 35, 14, 35, 338, 3, 36, 3, 36, 3, 36, 5, 36, 344, 10,
	36, 3, 37, 3, 37, 3, 37, 5, 37, 349, 10, 37, 3, 38, 3, 38, 5, 38, 353,
	10, 38, 3, 39, 3, 39, 5, 39, 357, 10, 39, 3, 39, 3, 39, 6, 39, 361, 10,
	39, 13, 39, 14, 39, 362, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41,
	3, 41, 3, 41, 5, 41, 374, 10, 41, 3, 41, 3, 41, 6, 41, 378, 10, 41, 13,
	41, 14, 41, 379, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 5, 41, 387, 10, 41,
	3, 42, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3,
	44, 3, 44, 5, 44, 401, 10, 44, 3, 45, 3, 45, 3, 45, 5, 45, 406, 10, 45,
	3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 5, 46, 414, 10, 46, 3, 47, 5,
	47, 417, 10, 47, 3, 47, 3, 47, 3, 47, 5, 47, 422, 10, 47, 3, 48, 3, 48,
	3, 48, 3, 48, 3, 48, 3, 49, 5, 49, 430, 10, 49, 3, 49, 3, 49, 3, 49, 5,
	49, 435, 10, 49, 3, 49, 5, 49, 438, 10, 49, 3, 49, 5, 49, 441, 10, 49,
	3, 50, 3, 50, 3, 50, 5, 50, 446, 10, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3,
	50, 7, 50, 453, 10, 50, 12, 50, 14, 50, 456, 11, 50, 3, 51, 3, 51, 6, 51,
	460, 10, 51, 13, 51, 14, 51, 461, 3, 51, 3, 51, 3, 52, 5, 52, 467, 10,
	52, 3, 52, 3, 52, 5, 52, 471, 10, 52, 3, 52, 3, 52, 3, 52, 5, 52, 476,
	10, 52, 3, 53, 3, 53, 6, 53, 480, 10, 53, 13, 53, 14, 53, 481, 3, 53, 3,
	53, 3, 54, 5, 54, 487, 10, 54, 3, 54, 3, 54, 3, 54, 3, 54, 5, 54, 493,
	10, 54, 3, 54, 5, 54, 496, 10, 54, 3, 55, 3, 55, 3, 55, 3, 55, 5, 55, 502,
	10, 55, 3, 55, 5, 55, 505, 10, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3,
	55, 5, 55, 513, 10, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55,
	5, 55, 522, 10, 55, 3, 56, 5, 56, 525, 10, 56, 3, 56, 3, 56, 3, 56, 5,
	56, 530, 10, 56, 3, 56, 5, 56, 533, 10, 56, 3, 57, 3, 57, 3, 57, 3, 57,
	5, 57, 539, 10, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 5,
	57, 548, 10, 57, 3, 58, 5, 58, 551, 10, 58, 3, 58, 3, 58, 3, 58, 5, 58,
	556, 10, 58, 3, 58, 5, 58, 559, 10, 58, 3, 59, 3, 59, 5, 59, 563, 10, 59,
	3, 59, 3, 59, 3, 59, 7, 59, 568, 10, 59, 12, 59, 14, 59, 571, 11, 59, 3,
	60, 3, 60, 3, 60, 3, 60, 5, 60, 577, 10, 60, 3, 60, 3, 60, 3, 60, 3, 60,
	3, 60, 3, 60, 3, 60, 5, 60, 586, 10, 60, 3, 61, 5, 61, 589, 10, 61, 3,
	61, 3, 61, 3, 61, 5, 61, 594, 10, 61, 3, 61, 5, 61, 597, 10, 61, 3, 62,
	3, 62, 6, 62, 601, 10, 62, 13, 62, 14, 62, 602, 3, 62, 3, 62, 3, 63, 5,
	63, 608, 10, 63, 3, 63, 3, 63, 5, 63, 612, 10, 63, 3, 64, 3, 64, 3, 64,
	3, 64, 5, 64, 618, 10, 64, 3, 64, 3, 64, 3, 64, 3, 64, 3, 64, 3, 64, 3,
	64, 5, 64, 627, 10, 64, 3, 65, 5, 65, 630, 10, 65, 3, 65, 3, 65, 3, 65,
	5, 65, 635, 10, 65, 3, 65, 5, 65, 638, 10, 65, 3, 66, 3, 66, 6, 66, 642,
	10, 66, 13, 66, 14, 66, 643, 3, 66, 3, 66, 3, 67, 3, 67, 3, 67, 3, 67,
	5, 67, 652, 10, 67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 5,
	67, 661, 10, 67, 3, 68, 5, 68, 664, 10, 68, 3, 68, 3, 68, 3, 68, 3, 68,
	5, 68, 670, 10, 68, 3, 68, 3, 68, 3, 68, 3, 69, 3, 69, 3, 69, 7, 69, 678,
	10, 69, 12, 69, 14, 69, 681, 11, 69, 3, 70, 3, 70, 5, 70, 685, 10, 70,
	3, 71, 3, 71, 3, 72, 3, 72, 3, 73, 3, 73, 3, 74, 3, 74, 3, 74, 3, 74, 7,
	74, 697, 10, 74, 12, 74, 14, 74, 700, 11, 74, 3, 74, 3, 74, 3, 75, 3, 75,
	5, 75, 706, 10, 75, 3, 75, 3, 75, 3, 76, 3, 76, 3, 76, 5, 76, 713, 10,
	76, 3, 76, 3, 76, 3, 76, 5, 76, 718, 10, 76, 7, 76, 720, 10, 76, 12, 76,
	14, 76, 723, 11, 76, 3, 76, 3, 76, 3, 76, 3, 76, 2, 3, 98, 77, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
	44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78,
	80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112,
	114, 116, 118, 120, 122, 124, 126, 128, 130, 132, 134, 136, 138, 140, 142,
	144, 146, 148, 150, 2, 7, 3, 2, 3, 5, 3, 2, 14, 15, 3, 2, 60, 61, 3, 2,
	36, 42, 3, 2, 43, 53, 2, 766, 2, 153, 3, 2, 2, 2, 4, 160, 3, 2, 2, 2, 6,
	164, 3, 2, 2, 2, 8, 179, 3, 2, 2, 2, 10, 181, 3, 2, 2, 2, 12, 183, 3, 2,
	2, 2, 14, 194, 3, 2, 2, 2, 16, 197, 3, 2, 2, 2, 18, 209, 3, 2, 2, 2, 20,
	217, 3, 2, 2, 2, 22, 221, 3, 2, 2, 2, 24, 224, 3, 2, 2, 2, 26, 229, 3,
	2, 2, 2, 28, 237, 3, 2, 2, 2, 30, 239, 3, 2, 2, 2, 32, 242, 3, 2, 2, 2,
	34, 260, 3, 2, 2, 2, 36, 262, 3, 2, 2, 2, 38, 264, 3, 2, 2, 2, 40, 266,
	3, 2, 2, 2, 42, 268, 3, 2, 2, 2, 44, 270, 3, 2, 2, 2, 46, 272, 3, 2, 2,
	2, 48, 284, 3, 2, 2, 2, 50, 286, 3, 2, 2, 2, 52, 295, 3, 2, 2, 2, 54, 299,
	3, 2, 2, 2, 56, 302, 3, 2, 2, 2, 58, 310, 3, 2, 2, 2, 60, 316, 3, 2, 2,
	2, 62, 327, 3, 2, 2, 2, 64, 329, 3, 2, 2, 2, 66, 331, 3, 2, 2, 2, 68, 336,
	3, 2, 2, 2, 70, 340, 3, 2, 2, 2, 72, 348, 3, 2, 2, 2, 74, 352, 3, 2, 2,
	2, 76, 354, 3, 2, 2, 2, 78, 366, 3, 2, 2, 2, 80, 386, 3, 2, 2, 2, 82, 388,
	3, 2, 2, 2, 84, 392, 3, 2, 2, 2, 86, 400, 3, 2, 2, 2, 88, 405, 3, 2, 2,
	2, 90, 413, 3, 2, 2, 2, 92, 416, 3, 2, 2, 2, 94, 423, 3, 2, 2, 2, 96, 429,
	3, 2, 2, 2, 98, 442, 3, 2, 2, 2, 100, 457, 3, 2, 2, 2, 102, 466, 3, 2,
	2, 2, 104, 477, 3, 2, 2, 2, 106, 486, 3, 2, 2, 2, 108, 521, 3, 2, 2, 2,
	110, 524, 3, 2, 2, 2, 112, 547, 3, 2, 2, 2, 114, 550, 3, 2, 2, 2, 116,
	560, 3, 2, 2, 2, 118, 585, 3, 2, 2, 2, 120, 588, 3, 2, 2, 2, 122, 598,
	3, 2, 2, 2, 124, 607, 3, 2, 2, 2, 126, 626, 3, 2, 2, 2, 128, 629, 3, 2,
	2, 2, 130, 639, 3, 2, 2, 2, 132, 660, 3, 2, 2, 2, 134, 663, 3, 2, 2, 2,
	136, 674, 3, 2, 2, 2, 138, 684, 3, 2, 2, 2, 140, 686, 3, 2, 2, 2, 142,
	688, 3, 2, 2, 2, 144, 690, 3, 2, 2, 2, 146, 692, 3, 2, 2, 2, 148, 703,
	3, 2, 2, 2, 150, 709, 3, 2, 2, 2, 152, 154, 5, 4, 3, 2, 153, 152, 3, 2,
	2, 2, 154, 155, 3, 2, 2, 2, 155, 153, 3, 2, 2, 2, 155, 156, 3, 2, 2, 2,
	156, 3, 3, 2, 2, 2, 157, 161, 5, 6, 4, 2, 158, 161, 5, 72, 37, 2, 159,
	161, 5, 74, 38, 2, 160, 157, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 160, 159,
	3, 2, 2, 2, 161, 5, 3, 2, 2, 2, 162, 165, 5, 8, 5, 2, 163, 165, 5, 26,
	14, 2, 164, 162, 3, 2, 2, 2, 164, 163, 3, 2, 2, 2, 165, 7, 3, 2, 2, 2,
	166, 168, 5, 10, 6, 2, 167, 169, 5, 144, 73, 2, 168, 167, 3, 2, 2, 2, 168,
	169, 3, 2, 2, 2, 169, 171, 3, 2, 2, 2, 170, 172, 5, 56, 29, 2, 171, 170,
	3, 2, 2, 2, 171, 172, 3, 2, 2, 2, 172, 174, 3, 2, 2, 2, 173, 175, 5, 68,
	35, 2, 174, 173, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2,
	176, 177, 5, 12, 7, 2, 177, 180, 3, 2, 2, 2, 178, 180, 5, 12, 7, 2, 179,
	166, 3, 2, 2, 2, 179, 178, 3, 2, 2, 2, 180, 9, 3, 2, 2, 2, 181, 182, 9,
	2, 2, 2, 182, 11, 3, 2, 2, 2, 183, 185, 7, 6, 2, 2, 184, 186, 5, 14, 8,
	2, 185, 184, 3, 2, 2, 2, 186, 187, 3, 2, 2, 2, 187, 185, 3, 2, 2, 2, 187,
	188, 3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 190, 7, 7, 2, 2, 190, 13, 3,
	2, 2, 2, 191, 195, 5, 16, 9, 2, 192, 195, 5, 24, 13, 2, 193, 195, 5, 32,
	17, 2, 194, 191, 3, 2, 2, 2, 194, 192, 3, 2, 2, 2, 194, 193, 3, 2, 2, 2,
	195, 15, 3, 2, 2, 2, 196, 198, 5, 22, 12, 2, 197, 196, 3, 2, 2, 2, 197,
	198, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 201, 5, 144, 73, 2, 200, 202,
	5, 18, 10, 2, 201, 200, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 204, 3,
	2, 2, 2, 203, 205, 5, 68, 35, 2, 204, 203, 3, 2, 2, 2, 204, 205, 3, 2,
	2, 2, 205, 207, 3, 2, 2, 2, 206, 208, 5, 12, 7, 2, 207, 206, 3, 2, 2, 2,
	207, 208, 3, 2, 2, 2, 208, 17, 3, 2, 2, 2, 209, 211, 7, 8, 2, 2, 210, 212,
	5, 20, 11, 2, 211, 210, 3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 211, 3,
	2, 2, 2, 213, 214, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 216, 7, 9, 2,
	2, 216, 19, 3, 2, 2, 2, 217, 218, 5, 144, 73, 2, 218, 219, 7, 10, 2, 2,
	219, 220, 5, 34, 18, 2, 220, 21, 3, 2, 2, 2, 221, 222, 5, 144, 73, 2, 222,
	223, 7, 10, 2, 2, 223, 23, 3, 2, 2, 2, 224, 225, 7, 11, 2, 2, 225, 227,
	5, 28, 15, 2, 226, 228, 5, 68, 35, 2, 227, 226, 3, 2, 2, 2, 227, 228, 3,
	2, 2, 2, 228, 25, 3, 2, 2, 2, 229, 230, 7, 12, 2, 2, 230, 231, 5, 28, 15,
	2, 231, 233, 5, 30, 16, 2, 232, 234, 5, 68, 35, 2, 233, 232, 3, 2, 2, 2,
	233, 234, 3, 2, 2, 2, 234, 235, 3, 2, 2, 2, 235, 236, 5, 12, 7, 2, 236,
	27, 3, 2, 2, 2, 237, 238, 5, 144, 73, 2, 238, 29, 3, 2, 2, 2, 239, 240,
	7, 13, 2, 2, 240, 241, 5, 64, 33, 2, 241, 31, 3, 2, 2, 2, 242, 244, 7,
	11, 2, 2, 243, 245, 5, 30, 16, 2, 244, 243, 3, 2, 2, 2, 244, 245, 3, 2,
	2, 2, 245, 247, 3, 2, 2, 2, 246, 248, 5, 68, 35, 2, 247, 246, 3, 2, 2,
	2, 247, 248, 3, 2, 2, 2, 248, 249, 3, 2, 2, 2, 249, 250, 5, 12, 7, 2, 250,
	33, 3, 2, 2, 2, 251, 261, 5, 54, 28, 2, 252, 261, 5, 36, 19, 2, 253, 261,
	5, 38, 20, 2, 254, 261, 5, 42, 22, 2, 255, 261, 5, 40, 21, 2, 256, 261,
	5, 44, 23, 2, 257, 261, 5, 46, 24, 2, 258, 261, 5, 48, 25, 2, 259, 261,
	5, 50, 26, 2, 260, 251, 3, 2, 2, 2, 260, 252, 3, 2, 2, 2, 260, 253, 3,
	2, 2, 2, 260, 254, 3, 2, 2, 2, 260, 255, 3, 2, 2, 2, 260, 256, 3, 2, 2,
	2, 260, 257, 3, 2, 2, 2, 260, 258, 3, 2, 2, 2, 260, 259, 3, 2, 2, 2, 261,
	35, 3, 2, 2, 2, 262, 263, 7, 64, 2, 2, 263, 37, 3, 2, 2, 2, 264, 265, 7,
	63, 2, 2, 265, 39, 3, 2, 2, 2, 266, 267, 9, 3, 2, 2, 267, 41, 3, 2, 2,
	2, 268, 269, 9, 4, 2, 2, 269, 43, 3, 2, 2, 2, 270, 271, 7, 16, 2, 2, 271,
	45, 3, 2, 2, 2, 272, 273, 5, 144, 73, 2, 273, 47, 3, 2, 2, 2, 274, 275,
	7, 17, 2, 2, 275, 285, 7, 18, 2, 2, 276, 278, 7, 17, 2, 2, 277, 279, 5,
	34, 18, 2, 278, 277, 3, 2, 2, 2, 279, 280, 3, 2, 2, 2, 280, 278, 3, 2,
	2, 2, 280, 281, 3, 2, 2, 2, 281, 282, 3, 2, 2, 2, 282, 283, 7, 18, 2, 2,
	283, 285, 3, 2, 2, 2, 284, 274, 3, 2, 2, 2, 284, 276, 3, 2, 2, 2, 285,
	49, 3, 2, 2, 2, 286, 290, 7, 6, 2, 2, 287, 289, 5, 52, 27, 2, 288, 287,
	3, 2, 2, 2, 289, 292, 3, 2, 2, 2, 290, 288, 3, 2, 2, 2, 290, 291, 3, 2,
	2, 2, 291, 293, 3, 2, 2, 2, 292, 290, 3, 2, 2, 2, 293, 294, 7, 7, 2, 2,
	294, 51, 3, 2, 2, 2, 295, 296, 5, 144, 73, 2, 296, 297, 7, 10, 2, 2, 297,
	298, 5, 34, 18, 2, 298, 53, 3, 2, 2, 2, 299, 300, 7, 19, 2, 2, 300, 301,
	5, 144, 73, 2, 301, 55, 3, 2, 2, 2, 302, 304, 7, 8, 2, 2, 303, 305, 5,
	58, 30, 2, 304, 303, 3, 2, 2, 2, 305, 306, 3, 2, 2, 2, 306, 304, 3, 2,
	2, 2, 306, 307, 3, 2, 2, 2, 307, 308, 3, 2, 2, 2, 308, 309, 7, 9, 2, 2,
	309, 57, 3, 2, 2, 2, 310, 311, 5, 54, 28, 2, 311, 312, 7, 10, 2, 2, 312,
	314, 5, 62, 32, 2, 313, 315, 5, 60, 31, 2, 314, 313, 3, 2, 2, 2, 314, 315,
	3, 2, 2, 2, 315, 59, 3, 2, 2, 2, 316, 317, 7, 20, 2, 2, 317, 318, 5, 34,
	18, 2, 318, 61, 3, 2, 2, 2, 319, 321, 5, 64, 33, 2, 320, 322, 7, 21, 2,
	2, 321, 320, 3, 2, 2, 2, 321, 322, 3, 2, 2, 2, 322, 328, 3, 2, 2, 2, 323,
	325, 5, 66, 34, 2, 324, 326, 7, 21, 2, 2, 325, 324, 3, 2, 2, 2, 325, 326,
	3, 2, 2, 2, 326, 328, 3, 2, 2, 2, 327, 319, 3, 2, 2, 2, 327, 323, 3, 2,
	2, 2, 328, 63, 3, 2, 2, 2, 329, 330, 5, 144, 73, 2, 330, 65, 3, 2, 2, 2,
	331, 332, 7, 17, 2, 2, 332, 333, 5, 62, 32, 2, 333, 334, 7, 18, 2, 2, 334,
	67, 3, 2, 2, 2, 335, 337, 5, 70, 36, 2, 336, 335, 3, 2, 2, 2, 337, 338,
	3, 2, 2, 2, 338, 336, 3, 2, 2, 2, 338, 339, 3, 2, 2, 2, 339, 69, 3, 2,
	2, 2, 340, 341, 7, 22, 2, 2, 341, 343, 5, 144, 73, 2, 342, 344, 5, 18,
	10, 2, 343, 342, 3, 2, 2, 2, 343, 344, 3, 2, 2, 2, 344, 71, 3, 2, 2, 2,
	345, 349, 5, 76, 39, 2, 346, 349, 5, 86, 44, 2, 347, 349, 5, 134, 68, 2,
	348, 345, 3, 2, 2, 2, 348, 346, 3, 2, 2, 2, 348, 347, 3, 2, 2, 2, 349,
	73, 3, 2, 2, 2, 350, 353, 5, 80, 41, 2, 351, 353, 5, 90, 46, 2, 352, 350,
	3, 2, 2, 2, 352, 351, 3, 2, 2, 2, 353, 75, 3, 2, 2, 2, 354, 356, 7, 23,
	2, 2, 355, 357, 5, 68, 35, 2, 356, 355, 3, 2, 2, 2, 356, 357, 3, 2, 2,
	2, 357, 358, 3, 2, 2, 2, 358, 360, 7, 6, 2, 2, 359, 361, 5, 78, 40, 2,
	360, 359, 3, 2, 2, 2, 361, 362, 3, 2, 2, 2, 362, 360, 3, 2, 2, 2, 362,
	363, 3, 2, 2, 2, 363, 364, 3, 2, 2, 2, 364, 365, 7, 7, 2, 2, 365, 77, 3,
	2, 2, 2, 366, 367, 5, 10, 6, 2, 367, 368, 7, 10, 2, 2, 368, 369, 5, 64,
	33, 2, 369, 79, 3, 2, 2, 2, 370, 371, 7, 24, 2, 2, 371, 373, 7, 23, 2,
	2, 372, 374, 5, 68, 35, 2, 373, 372, 3, 2, 2, 2, 373, 374, 3, 2, 2, 2,
	374, 375, 3, 2, 2, 2, 375, 377, 7, 6, 2, 2, 376, 378, 5, 82, 42, 2, 377,
	376, 3, 2, 2, 2, 378, 379, 3, 2, 2, 2, 379, 377, 3, 2, 2, 2, 379, 380,
	3, 2, 2, 2, 380, 381, 3, 2, 2, 2, 381, 382, 7, 7, 2, 2, 382, 387, 3, 2,
	2, 2, 383, 384, 7, 24, 2, 2, 384, 385, 7, 23, 2, 2, 385, 387, 5, 68, 35,
	2, 386, 370, 3, 2, 2, 2, 386, 383, 3, 2, 2, 2, 387, 81, 3, 2, 2, 2, 388,
	389, 5, 10, 6, 2, 389, 390, 7, 10, 2, 2, 390, 391, 5, 64, 33, 2, 391, 83,
	3, 2, 2, 2, 392, 393, 5, 42, 22, 2, 393, 85, 3, 2, 2, 2, 394, 401, 5, 92,
	47, 2, 395, 401, 5, 114, 58, 2, 396, 401, 5, 120, 61, 2, 397, 401, 5, 88,
	45, 2, 398, 401, 5, 148, 75, 2, 399, 401, 5, 150, 76, 2, 400, 394, 3, 2,
	2, 2, 400, 395, 3, 2, 2, 2, 400, 396, 3, 2, 2, 2, 400, 397, 3, 2, 2, 2,
	400, 398, 3, 2, 2, 2, 400, 399, 3, 2, 2, 2, 401, 87, 3, 2, 2, 2, 402, 406,
	5, 96, 49, 2, 403, 406, 5, 110, 56, 2, 404, 406, 5, 128, 65, 2, 405, 402,
	3, 2, 2, 2, 405, 403, 3, 2, 2, 2, 405, 404, 3, 2, 2, 2, 406, 89, 3, 2,
	2, 2, 407, 414, 5, 94, 48, 2, 408, 414, 5, 108, 55, 2, 409, 414, 5, 112,
	57, 2, 410, 414, 5, 118, 60, 2, 411, 414, 5, 126, 64, 2, 412, 414, 5, 132,
	67, 2, 413, 407, 3, 2, 2, 2, 413, 408, 3, 2, 2, 2, 413, 409, 3, 2, 2, 2,
	413, 410, 3, 2, 2, 2, 413, 411, 3, 2, 2, 2, 413, 412, 3, 2, 2, 2, 414,
	91, 3, 2, 2, 2, 415, 417, 5, 84, 43, 2, 416, 415, 3, 2, 2, 2, 416, 417,
	3, 2, 2, 2, 417, 418, 3, 2, 2, 2, 418, 419, 7, 25, 2, 2, 419, 421, 5, 144,
	73, 2, 420, 422, 5, 68, 35, 2, 421, 420, 3, 2, 2, 2, 421, 422, 3, 2, 2,
	2, 422, 93, 3, 2, 2, 2, 423, 424, 7, 26, 2, 2, 424, 425, 7, 25, 2, 2, 425,
	426, 5, 144, 73, 2, 426, 427, 5, 68, 35, 2, 427, 95, 3, 2, 2, 2, 428, 430,
	5, 84, 43, 2, 429, 428, 3, 2, 2, 2, 429, 430, 3, 2, 2, 2, 430, 431, 3,
	2, 2, 2, 431, 432, 7, 27, 2, 2, 432, 434, 5, 144, 73, 2, 433, 435, 5, 98,
	50, 2, 434, 433, 3, 2, 2, 2, 434, 435, 3, 2, 2, 2, 435, 437, 3, 2, 2, 2,
	436, 438, 5, 68, 35, 2, 437, 436, 3, 2, 2, 2, 437, 438, 3, 2, 2, 2, 438,
	440, 3, 2, 2, 2, 439, 441, 5, 100, 51, 2, 440, 439, 3, 2, 2, 2, 440, 441,
	3, 2, 2, 2, 441, 97, 3, 2, 2, 2, 442, 443, 8, 50, 1, 2, 443, 445, 7, 28,
	2, 2, 444, 446, 7, 29, 2, 2, 445, 444, 3, 2, 2, 2, 445, 446, 3, 2, 2, 2,
	446, 447, 3, 2, 2, 2, 447, 448, 5, 64, 33, 2, 448, 454, 3, 2, 2, 2, 449,
	450, 12, 3, 2, 2, 450, 451, 7, 29, 2, 2, 451, 453, 5, 64, 33, 2, 452, 449,
	3, 2, 2, 2, 453, 456, 3, 2, 2, 2, 454, 452, 3, 2, 2, 2, 454, 455, 3, 2,
	2, 2, 455, 99, 3, 2, 2, 2, 456, 454, 3, 2, 2, 2, 457, 459, 7, 6, 2, 2,
	458, 460, 5, 102, 52, 2, 459, 458, 3, 2, 2, 2, 460, 461, 3, 2, 2, 2, 461,
	459, 3, 2, 2, 2, 461, 462, 3, 2, 2, 2, 462, 463, 3, 2, 2, 2, 463, 464,
	7, 7, 2, 2, 464, 101, 3, 2, 2, 2, 465, 467, 5, 84, 43, 2, 466, 465, 3,
	2, 2, 2, 466, 467, 3, 2, 2, 2, 467, 468, 3, 2, 2, 2, 468, 470, 5, 144,
	73, 2, 469, 471, 5, 104, 53, 2, 470, 469, 3, 2, 2, 2, 470, 471, 3, 2, 2,
	2, 471, 472, 3, 2, 2, 2, 472, 473, 7, 10, 2, 2, 473, 475, 5, 62, 32, 2,
	474, 476, 5, 68, 35, 2, 475, 474, 3, 2, 2, 2, 475, 476, 3, 2, 2, 2, 476,
	103, 3, 2, 2, 2, 477, 479, 7, 8, 2, 2, 478, 480, 5, 106, 54, 2, 479, 478,
	3, 2, 2, 2, 480, 481, 3, 2, 2, 2, 481, 479, 3, 2, 2, 2, 481, 482, 3, 2,
	2, 2, 482, 483, 3, 2, 2, 2, 483, 484, 7, 9, 2, 2, 484, 105, 3, 2, 2, 2,
	485, 487, 5, 84, 43, 2, 486, 485, 3, 2, 2, 2, 486, 487, 3, 2, 2, 2, 487,
	488, 3, 2, 2, 2, 488, 489, 5, 144, 73, 2, 489, 490, 7, 10, 2, 2, 490, 492,
	5, 62, 32, 2, 491, 493, 5, 60, 31, 2, 492, 491, 3, 2, 2, 2, 492, 493, 3,
	2, 2, 2, 493, 495, 3, 2, 2, 2, 494, 496, 5, 68, 35, 2, 495, 494, 3, 2,
	2, 2, 495, 496, 3, 2, 2, 2, 496, 107, 3, 2, 2, 2, 497, 498, 7, 24, 2, 2,
	498, 499, 7, 27, 2, 2, 499, 501, 5, 144, 73, 2, 500, 502, 5, 98, 50, 2,
	501, 500, 3, 2, 2, 2, 501, 502, 3, 2, 2, 2, 502, 504, 3, 2, 2, 2, 503,
	505, 5, 68, 35, 2, 504, 503, 3, 2, 2, 2, 504, 505, 3, 2, 2, 2, 505, 506,
	3, 2, 2, 2, 506, 507, 5, 100, 51, 2, 507, 522, 3, 2, 2, 2, 508, 509, 7,
	24, 2, 2, 509, 510, 7, 27, 2, 2, 510, 512, 5, 144, 73, 2, 511, 513, 5,
	98, 50, 2, 512, 511, 3, 2, 2, 2, 512, 513, 3, 2, 2, 2, 513, 514, 3, 2,
	2, 2, 514, 515, 5, 68, 35, 2, 515, 522, 3, 2, 2, 2, 516, 517, 7, 24, 2,
	2, 517, 518, 7, 27, 2, 2, 518, 519, 5, 144, 73, 2, 519, 520, 5, 98, 50,
	2, 520, 522, 3, 2, 2, 2, 521, 497, 3, 2, 2, 2, 521, 508, 3, 2, 2, 2, 521,
	516, 3, 2, 2, 2, 522, 109, 3, 2, 2, 2, 523, 525, 5, 84, 43, 2, 524, 523,
	3, 2, 2, 2, 524, 525, 3, 2, 2, 2, 525, 526, 3, 2, 2, 2, 526, 527, 7, 30,
	2, 2, 527, 529, 5, 144, 73, 2, 528, 530, 5, 68, 35, 2, 529, 528, 3, 2,
	2, 2, 529, 530, 3, 2, 2, 2, 530, 532, 3, 2, 2, 2, 531, 533, 5, 100, 51,
	2, 532, 531, 3, 2, 2, 2, 532, 533, 3, 2, 2, 2, 533, 111, 3, 2, 2, 2, 534,
	535, 7, 24, 2, 2, 535, 536, 7, 30, 2, 2, 536, 538, 5, 144, 73, 2, 537,
	539, 5, 68, 35, 2, 538, 537, 3, 2, 2, 2, 538, 539, 3, 2, 2, 2, 539, 540,
	3, 2, 2, 2, 540, 541, 5, 100, 51, 2, 541, 548, 3, 2, 2, 2, 542, 543, 7,
	24, 2, 2, 543, 544, 7, 30, 2, 2, 544, 545, 5, 144, 73, 2, 545, 546, 5,
	68, 35, 2, 546, 548, 3, 2, 2, 2, 547, 534, 3, 2, 2, 2, 547, 542, 3, 2,
	2, 2, 548, 113, 3, 2, 2, 2, 549, 551, 5, 84, 43, 2, 550, 549, 3, 2, 2,
	2, 550, 551, 3, 2, 2, 2, 551, 552, 3, 2, 2, 2, 552, 553, 7, 31, 2, 2, 553,
	555, 5, 144, 73, 2, 554, 556, 5, 68, 35, 2, 555, 554, 3, 2, 2, 2, 555,
	556, 3, 2, 2, 2, 556, 558, 3, 2, 2, 2, 557, 559, 5, 116, 59, 2, 558, 557,
	3, 2, 2, 2, 558, 559, 3, 2, 2, 2, 559, 115, 3, 2, 2, 2, 560, 562, 7, 20,
	2, 2, 561, 563, 7, 32, 2, 2, 562, 561, 3, 2, 2, 2, 562, 563, 3, 2, 2, 2,
	563, 564, 3, 2, 2, 2, 564, 569, 5, 64, 33, 2, 565, 566, 7, 32, 2, 2, 566,
	568, 5, 64, 33, 2, 567, 565, 3, 2, 2, 2, 568, 571, 3, 2, 2, 2, 569, 567,
	3, 2, 2, 2, 569, 570, 3, 2, 2, 2, 570, 117, 3, 2, 2, 2, 571, 569, 3, 2,
	2, 2, 572, 573, 7, 24, 2, 2, 573, 574, 7, 31, 2, 2, 574, 576, 5, 144, 73,
	2, 575, 577, 5, 68, 35, 2, 576, 575, 3, 2, 2, 2, 576, 577, 3, 2, 2, 2,
	577, 578, 3, 2, 2, 2, 578, 579, 5, 116, 59, 2, 579, 586, 3, 2, 2, 2, 580,
	581, 7, 24, 2, 2, 581, 582, 7, 31, 2, 2, 582, 583, 5, 144, 73, 2, 583,
	584, 5, 68, 35, 2, 584, 586, 3, 2, 2, 2, 585, 572, 3, 2, 2, 2, 585, 580,
	3, 2, 2, 2, 586, 119, 3, 2, 2, 2, 587, 589, 5, 84, 43, 2, 588, 587, 3,
	2, 2, 2, 588, 589, 3, 2, 2, 2, 589, 590, 3, 2, 2, 2, 590, 591, 7, 33, 2,
	2, 591, 593, 5, 144, 73, 2, 592, 594, 5, 68, 35, 2, 593, 592, 3, 2, 2,
	2, 593, 594, 3, 2, 2, 2, 594, 596, 3, 2, 2, 2, 595, 597, 5, 122, 62, 2,
	596, 595, 3, 2, 2, 2, 596, 597, 3, 2, 2, 2, 597, 121, 3, 2, 2, 2, 598,
	600, 7, 6, 2, 2, 599, 601, 5, 124, 63, 2, 600, 599, 3, 2, 2, 2, 601, 602,
	3, 2, 2, 2, 602, 600, 3, 2, 2, 2, 602, 603, 3, 2, 2, 2, 603, 604, 3, 2,
	2, 2, 604, 605, 7, 7, 2, 2, 605, 123, 3, 2, 2, 2, 606, 608, 5, 84, 43,
	2, 607, 606, 3, 2, 2, 2, 607, 608, 3, 2, 2, 2, 608, 609, 3, 2, 2, 2, 609,
	611, 5, 46, 24, 2, 610, 612, 5, 68, 35, 2, 611, 610, 3, 2, 2, 2, 611, 612,
	3, 2, 2, 2, 612, 125, 3, 2, 2, 2, 613, 614, 7, 24, 2, 2, 614, 615, 7, 33,
	2, 2, 615, 617, 5, 144, 73, 2, 616, 618, 5, 68, 35, 2, 617, 616, 3, 2,
	2, 2, 617, 618, 3, 2, 2, 2, 618, 619, 3, 2, 2, 2, 619, 620, 5, 122, 62,
	2, 620, 627, 3, 2, 2, 2, 621, 622, 7, 24, 2, 2, 622, 623, 7, 33, 2, 2,
	623, 624, 5, 144, 73, 2, 624, 625, 5, 68, 35, 2, 625, 627, 3, 2, 2, 2,
	626, 613, 3, 2, 2, 2, 626, 621, 3, 2, 2, 2, 627, 127, 3, 2, 2, 2, 628,
	630, 5, 84, 43, 2, 629, 628, 3, 2, 2, 2, 629, 630, 3, 2, 2, 2, 630, 631,
	3, 2, 2, 2, 631, 632, 7, 34, 2, 2, 632, 634, 5, 144, 73, 2, 633, 635, 5,
	68, 35, 2, 634, 633, 3, 2, 2, 2, 634, 635, 3, 2, 2, 2, 635, 637, 3, 2,
	2, 2, 636, 638, 5, 130, 66, 2, 637, 636, 3, 2, 2, 2, 637, 638, 3, 2, 2,
	2, 638, 129, 3, 2, 2, 2, 639, 641, 7, 6, 2, 2, 640, 642, 5, 106, 54, 2,
	641, 640, 3, 2, 2, 2, 642, 643, 3, 2, 2, 2, 643, 641, 3, 2, 2, 2, 643,
	644, 3, 2, 2, 2, 644, 645, 3, 2, 2, 2, 645, 646, 7, 7, 2, 2, 646, 131,
	3, 2, 2, 2, 647, 648, 7, 24, 2, 2, 648, 649, 7, 34, 2, 2, 649, 651, 5,
	144, 73, 2, 650, 652, 5, 68, 35, 2, 651, 650, 3, 2, 2, 2, 651, 652, 3,
	2, 2, 2, 652, 653, 3, 2, 2, 2, 653, 654, 5, 130, 66, 2, 654, 661, 3, 2,
	2, 2, 655, 656, 7, 24, 2, 2, 656, 657, 7, 34, 2, 2, 657, 658, 5, 144, 73,
	2, 658, 659, 5, 68, 35, 2, 659, 661, 3, 2, 2, 2, 660, 647, 3, 2, 2, 2,
	660, 655, 3, 2, 2, 2, 661, 133, 3, 2, 2, 2, 662, 664, 5, 84, 43, 2, 663,
	662, 3, 2, 2, 2, 663, 664, 3, 2, 2, 2, 664, 665, 3, 2, 2, 2, 665, 666,
	7, 35, 2, 2, 666, 667, 7, 22, 2, 2, 667, 669, 5, 144, 73, 2, 668, 670,
	5, 104, 53, 2, 669, 668, 3, 2, 2, 2, 669, 670, 3, 2, 2, 2, 670, 671, 3,
	2, 2, 2, 671, 672, 7, 13, 2, 2, 672, 673, 5, 136, 69, 2, 673, 135, 3, 2,
	2, 2, 674, 679, 5, 138, 70, 2, 675, 676, 7, 32, 2, 2, 676, 678, 5, 138,
	70, 2, 677, 675, 3, 2, 2, 2, 678, 681, 3, 2, 2, 2, 679, 677, 3, 2, 2, 2,
	679, 680, 3, 2, 2, 2, 680, 137, 3, 2, 2, 2, 681, 679, 3, 2, 2, 2, 682,
	685, 5, 140, 71, 2, 683, 685, 5, 142, 72, 2, 684, 682, 3, 2, 2, 2, 684,
	683, 3, 2, 2, 2, 685, 139, 3, 2, 2, 2, 686, 687, 9, 5, 2, 2, 687, 141,
	3, 2, 2, 2, 688, 689, 9, 6, 2, 2, 689, 143, 3, 2, 2, 2, 690, 691, 7, 59,
	2, 2, 691, 145, 3, 2, 2, 2, 692, 693, 7, 54, 2, 2, 693, 698, 5, 144, 73,
	2, 694, 695, 7, 67, 2, 2, 695, 697, 5, 144, 73, 2, 696, 694, 3, 2, 2, 2,
	697, 700, 3, 2, 2, 2, 698, 696, 3, 2, 2, 2, 698, 699, 3, 2, 2, 2, 699,
	701, 3, 2, 2, 2, 700, 698, 3, 2, 2, 2, 701, 702, 7, 55, 2, 2, 702, 147,
	3, 2, 2, 2, 703, 705, 7, 56, 2, 2, 704, 706, 5, 146, 74, 2, 705, 704, 3,
	2, 2, 2, 705, 706, 3, 2, 2, 2, 706, 707, 3, 2, 2, 2, 707, 708, 5, 88, 45,
	2, 708, 149, 3, 2, 2, 2, 709, 710, 7, 57, 2, 2, 710, 712, 5, 144, 73, 2,
	711, 713, 5, 146, 74, 2, 712, 711, 3, 2, 2, 2, 712, 713, 3, 2, 2, 2, 713,
	721, 3, 2, 2, 2, 714, 715, 7, 67, 2, 2, 715, 717, 5, 144, 73, 2, 716, 718,
	5, 146, 74, 2, 717, 716, 3, 2, 2, 2, 717, 718, 3, 2, 2, 2, 718, 720, 3,
	2, 2, 2, 719, 714, 3, 2, 2, 2, 720, 723, 3, 2, 2, 2, 721, 719, 3, 2, 2,
	2, 721, 722, 3, 2, 2, 2, 722, 724, 3, 2, 2, 2, 723, 721, 3, 2, 2, 2, 724,
	725, 7, 58, 2, 2, 725, 726, 5, 88, 45, 2, 726, 151, 3, 2, 2, 2, 96, 155,
	160, 164, 168, 171, 174, 179, 187, 194, 197, 201, 204, 207, 213, 227, 233,
	244, 247, 260, 280, 284, 290, 306, 314, 321, 325, 327, 338, 343, 348, 352,
	356, 362, 373, 379, 386, 400, 405, 413, 416, 421, 429, 434, 437, 440, 445,
	454, 461, 466, 470, 475, 481, 486, 492, 495, 501, 504, 512, 521, 524, 529,
	532, 538, 547, 550, 555, 558, 562, 569, 576, 585, 588, 593, 596, 602, 607,
	611, 617, 626, 629, 634, 637, 643, 651, 660, 663, 669, 679, 684, 698, 705,
	712, 717, 721,
}
var literalNames = []string{
	"", "'query'", "'mutation'", "'subscription'", "'{'", "'}'", "'('", "')'",
	"':'", "'...'", "'fragment'", "'on'", "'true'", "'false'", "'null'", "'['",
	"']'", "'$'", "'='", "'!'", "'@'", "'schema'", "'extend'", "'scalar'",
	"'extends'", "'type'", "'implements'", "'&'", "'interface'", "'union'",
	"'|'", "'enum'", "'input'", "'directive'", "'QUERY'", "'MUTATION'", "'SUBSCRIPTION'",
	"'FIELD'", "'FRAGMENT_DEFINITION'", "'FRAGMENT_SPREAD'", "'INLINE_FRAGMENT'",
	"'SCHEMA'", "'SCALAR'", "'OBJECT'", "'FIELD_DEFINITION'", "'ARGUMENT_DEFINITION'",
	"'INTERFACE'", "'UNION'", "'ENUM'", "'ENUM_VALUE'", "'INPUT_OBJECT'", "'INPUT_FIELD_DEFINITION'",
	"'<'", "'>'", "'template'", "'apply'", "'to'", "", "", "", "", "", "",
	"", "", "','", "", "", "'\uEFBBBF'", "'\uFEFF'", "'\u0000FEFF'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "NAME", "STRING", "BLOCK_STRING", "ID", "FLOAT", "INT", "PUNCTUATOR",
	"WS", "COMMA", "LineComment", "UNICODE_BOM", "UTF8_BOM", "UTF16_BOM", "UTF32_BOM",
}

var ruleNames = []string{
	"document", "definition", "executableDefinition", "operationDefinition",
	"operationType", "selectionSet", "selection", "field", "arguments", "argument",
	"alias", "fragmentSpread", "fragmentDefinition", "fragmentName", "typeCondition",
	"inlineFragment", "value", "intValue", "floatValue", "booleanValue", "stringValue",
	"nullValue", "enumValue", "listValue", "objectValue", "objectField", "variable",
	"variableDefinitions", "variableDefinition", "defaultValue", "type_", "namedType",
	"listType", "directives", "directive", "typeSystemDefinition", "typeSystemExtension",
	"schemaDefinition", "rootOperationTypeDefinition", "schemaExtension", "operationTypeDefinition",
	"description", "typeDefinition", "templetableTypeDefinition", "typeExtension",
	"scalarTypeDefinition", "scalarTypeExtension", "objectTypeDefinition",
	"implementsInterfaces", "fieldsDefinition", "fieldDefinition", "argumentsDefinition",
	"inputValueDefinition", "objectTypeExtension", "interfaceTypeDefinition",
	"interfaceTypeExtension", "unionTypeDefinition", "unionMemberTypes", "unionTypeExtension",
	"enumTypeDefinition", "enumValuesDefinition", "enumValueDefinition", "enumTypeExtension",
	"inputObjectTypeDefinition", "inputFieldsDefinition", "inputObjectTypeExtension",
	"directiveDefinition", "directiveLocations", "directiveLocation", "executableDirectiveLocation",
	"typeSystemDirectiveLocation", "name", "templateParametersDefinition",
	"templateTypeDefinition", "templateImplementedTypeDefinition",
}

type EGraphQLParser struct {
	*antlr.BaseParser
}

// NewEGraphQLParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *EGraphQLParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewEGraphQLParser(input antlr.TokenStream) *EGraphQLParser {
	this := new(EGraphQLParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "EGraphQL.g4"

	return this
}

// EGraphQLParser tokens.
const (
	EGraphQLParserEOF          = antlr.TokenEOF
	EGraphQLParserT__0         = 1
	EGraphQLParserT__1         = 2
	EGraphQLParserT__2         = 3
	EGraphQLParserT__3         = 4
	EGraphQLParserT__4         = 5
	EGraphQLParserT__5         = 6
	EGraphQLParserT__6         = 7
	EGraphQLParserT__7         = 8
	EGraphQLParserT__8         = 9
	EGraphQLParserT__9         = 10
	EGraphQLParserT__10        = 11
	EGraphQLParserT__11        = 12
	EGraphQLParserT__12        = 13
	EGraphQLParserT__13        = 14
	EGraphQLParserT__14        = 15
	EGraphQLParserT__15        = 16
	EGraphQLParserT__16        = 17
	EGraphQLParserT__17        = 18
	EGraphQLParserT__18        = 19
	EGraphQLParserT__19        = 20
	EGraphQLParserT__20        = 21
	EGraphQLParserT__21        = 22
	EGraphQLParserT__22        = 23
	EGraphQLParserT__23        = 24
	EGraphQLParserT__24        = 25
	EGraphQLParserT__25        = 26
	EGraphQLParserT__26        = 27
	EGraphQLParserT__27        = 28
	EGraphQLParserT__28        = 29
	EGraphQLParserT__29        = 30
	EGraphQLParserT__30        = 31
	EGraphQLParserT__31        = 32
	EGraphQLParserT__32        = 33
	EGraphQLParserT__33        = 34
	EGraphQLParserT__34        = 35
	EGraphQLParserT__35        = 36
	EGraphQLParserT__36        = 37
	EGraphQLParserT__37        = 38
	EGraphQLParserT__38        = 39
	EGraphQLParserT__39        = 40
	EGraphQLParserT__40        = 41
	EGraphQLParserT__41        = 42
	EGraphQLParserT__42        = 43
	EGraphQLParserT__43        = 44
	EGraphQLParserT__44        = 45
	EGraphQLParserT__45        = 46
	EGraphQLParserT__46        = 47
	EGraphQLParserT__47        = 48
	EGraphQLParserT__48        = 49
	EGraphQLParserT__49        = 50
	EGraphQLParserT__50        = 51
	EGraphQLParserT__51        = 52
	EGraphQLParserT__52        = 53
	EGraphQLParserT__53        = 54
	EGraphQLParserT__54        = 55
	EGraphQLParserT__55        = 56
	EGraphQLParserNAME         = 57
	EGraphQLParserSTRING       = 58
	EGraphQLParserBLOCK_STRING = 59
	EGraphQLParserID           = 60
	EGraphQLParserFLOAT        = 61
	EGraphQLParserINT          = 62
	EGraphQLParserPUNCTUATOR   = 63
	EGraphQLParserWS           = 64
	EGraphQLParserCOMMA        = 65
	EGraphQLParserLineComment  = 66
	EGraphQLParserUNICODE_BOM  = 67
	EGraphQLParserUTF8_BOM     = 68
	EGraphQLParserUTF16_BOM    = 69
	EGraphQLParserUTF32_BOM    = 70
)

// EGraphQLParser rules.
const (
	EGraphQLParserRULE_document                          = 0
	EGraphQLParserRULE_definition                        = 1
	EGraphQLParserRULE_executableDefinition              = 2
	EGraphQLParserRULE_operationDefinition               = 3
	EGraphQLParserRULE_operationType                     = 4
	EGraphQLParserRULE_selectionSet                      = 5
	EGraphQLParserRULE_selection                         = 6
	EGraphQLParserRULE_field                             = 7
	EGraphQLParserRULE_arguments                         = 8
	EGraphQLParserRULE_argument                          = 9
	EGraphQLParserRULE_alias                             = 10
	EGraphQLParserRULE_fragmentSpread                    = 11
	EGraphQLParserRULE_fragmentDefinition                = 12
	EGraphQLParserRULE_fragmentName                      = 13
	EGraphQLParserRULE_typeCondition                     = 14
	EGraphQLParserRULE_inlineFragment                    = 15
	EGraphQLParserRULE_value                             = 16
	EGraphQLParserRULE_intValue                          = 17
	EGraphQLParserRULE_floatValue                        = 18
	EGraphQLParserRULE_booleanValue                      = 19
	EGraphQLParserRULE_stringValue                       = 20
	EGraphQLParserRULE_nullValue                         = 21
	EGraphQLParserRULE_enumValue                         = 22
	EGraphQLParserRULE_listValue                         = 23
	EGraphQLParserRULE_objectValue                       = 24
	EGraphQLParserRULE_objectField                       = 25
	EGraphQLParserRULE_variable                          = 26
	EGraphQLParserRULE_variableDefinitions               = 27
	EGraphQLParserRULE_variableDefinition                = 28
	EGraphQLParserRULE_defaultValue                      = 29
	EGraphQLParserRULE_type_                             = 30
	EGraphQLParserRULE_namedType                         = 31
	EGraphQLParserRULE_listType                          = 32
	EGraphQLParserRULE_directives                        = 33
	EGraphQLParserRULE_directive                         = 34
	EGraphQLParserRULE_typeSystemDefinition              = 35
	EGraphQLParserRULE_typeSystemExtension               = 36
	EGraphQLParserRULE_schemaDefinition                  = 37
	EGraphQLParserRULE_rootOperationTypeDefinition       = 38
	EGraphQLParserRULE_schemaExtension                   = 39
	EGraphQLParserRULE_operationTypeDefinition           = 40
	EGraphQLParserRULE_description                       = 41
	EGraphQLParserRULE_typeDefinition                    = 42
	EGraphQLParserRULE_templetableTypeDefinition         = 43
	EGraphQLParserRULE_typeExtension                     = 44
	EGraphQLParserRULE_scalarTypeDefinition              = 45
	EGraphQLParserRULE_scalarTypeExtension               = 46
	EGraphQLParserRULE_objectTypeDefinition              = 47
	EGraphQLParserRULE_implementsInterfaces              = 48
	EGraphQLParserRULE_fieldsDefinition                  = 49
	EGraphQLParserRULE_fieldDefinition                   = 50
	EGraphQLParserRULE_argumentsDefinition               = 51
	EGraphQLParserRULE_inputValueDefinition              = 52
	EGraphQLParserRULE_objectTypeExtension               = 53
	EGraphQLParserRULE_interfaceTypeDefinition           = 54
	EGraphQLParserRULE_interfaceTypeExtension            = 55
	EGraphQLParserRULE_unionTypeDefinition               = 56
	EGraphQLParserRULE_unionMemberTypes                  = 57
	EGraphQLParserRULE_unionTypeExtension                = 58
	EGraphQLParserRULE_enumTypeDefinition                = 59
	EGraphQLParserRULE_enumValuesDefinition              = 60
	EGraphQLParserRULE_enumValueDefinition               = 61
	EGraphQLParserRULE_enumTypeExtension                 = 62
	EGraphQLParserRULE_inputObjectTypeDefinition         = 63
	EGraphQLParserRULE_inputFieldsDefinition             = 64
	EGraphQLParserRULE_inputObjectTypeExtension          = 65
	EGraphQLParserRULE_directiveDefinition               = 66
	EGraphQLParserRULE_directiveLocations                = 67
	EGraphQLParserRULE_directiveLocation                 = 68
	EGraphQLParserRULE_executableDirectiveLocation       = 69
	EGraphQLParserRULE_typeSystemDirectiveLocation       = 70
	EGraphQLParserRULE_name                              = 71
	EGraphQLParserRULE_templateParametersDefinition      = 72
	EGraphQLParserRULE_templateTypeDefinition            = 73
	EGraphQLParserRULE_templateImplementedTypeDefinition = 74
)

// IDocumentContext is an interface to support dynamic dispatch.
type IDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDocumentContext differentiates from other interfaces.
	IsDocumentContext()
}

type DocumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentContext() *DocumentContext {
	var p = new(DocumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_document
	return p
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) AllDefinition() []IDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefinitionContext)(nil)).Elem())
	var tst = make([]IDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefinitionContext)
		}
	}

	return tst
}

func (s *DocumentContext) Definition(i int) IDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (p *EGraphQLParser) Document() (localctx IDocumentContext) {
	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EGraphQLParserRULE_document)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<EGraphQLParserT__0)|(1<<EGraphQLParserT__1)|(1<<EGraphQLParserT__2)|(1<<EGraphQLParserT__3)|(1<<EGraphQLParserT__9)|(1<<EGraphQLParserT__20)|(1<<EGraphQLParserT__21)|(1<<EGraphQLParserT__22)|(1<<EGraphQLParserT__23)|(1<<EGraphQLParserT__24)|(1<<EGraphQLParserT__27)|(1<<EGraphQLParserT__28)|(1<<EGraphQLParserT__30))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(EGraphQLParserT__31-32))|(1<<(EGraphQLParserT__32-32))|(1<<(EGraphQLParserT__53-32))|(1<<(EGraphQLParserT__54-32))|(1<<(EGraphQLParserSTRING-32))|(1<<(EGraphQLParserBLOCK_STRING-32)))) != 0) {
		{
			p.SetState(150)
			p.Definition()
		}

		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_definition
	return p
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) ExecutableDefinition() IExecutableDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExecutableDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExecutableDefinitionContext)
}

func (s *DefinitionContext) TypeSystemDefinition() ITypeSystemDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSystemDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSystemDefinitionContext)
}

func (s *DefinitionContext) TypeSystemExtension() ITypeSystemExtensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSystemExtensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSystemExtensionContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (p *EGraphQLParser) Definition() (localctx IDefinitionContext) {
	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EGraphQLParserRULE_definition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(158)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case EGraphQLParserT__0, EGraphQLParserT__1, EGraphQLParserT__2, EGraphQLParserT__3, EGraphQLParserT__9:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(155)
			p.ExecutableDefinition()
		}

	case EGraphQLParserT__20, EGraphQLParserT__22, EGraphQLParserT__24, EGraphQLParserT__27, EGraphQLParserT__28, EGraphQLParserT__30, EGraphQLParserT__31, EGraphQLParserT__32, EGraphQLParserT__53, EGraphQLParserT__54, EGraphQLParserSTRING, EGraphQLParserBLOCK_STRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(156)
			p.TypeSystemDefinition()
		}

	case EGraphQLParserT__21, EGraphQLParserT__23:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(157)
			p.TypeSystemExtension()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExecutableDefinitionContext is an interface to support dynamic dispatch.
type IExecutableDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExecutableDefinitionContext differentiates from other interfaces.
	IsExecutableDefinitionContext()
}

type ExecutableDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExecutableDefinitionContext() *ExecutableDefinitionContext {
	var p = new(ExecutableDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_executableDefinition
	return p
}

func (*ExecutableDefinitionContext) IsExecutableDefinitionContext() {}

func NewExecutableDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExecutableDefinitionContext {
	var p = new(ExecutableDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_executableDefinition

	return p
}

func (s *ExecutableDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExecutableDefinitionContext) OperationDefinition() IOperationDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperationDefinitionContext)
}

func (s *ExecutableDefinitionContext) FragmentDefinition() IFragmentDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFragmentDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFragmentDefinitionContext)
}

func (s *ExecutableDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExecutableDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExecutableDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterExecutableDefinition(s)
	}
}

func (s *ExecutableDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitExecutableDefinition(s)
	}
}

func (p *EGraphQLParser) ExecutableDefinition() (localctx IExecutableDefinitionContext) {
	localctx = NewExecutableDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, EGraphQLParserRULE_executableDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(162)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case EGraphQLParserT__0, EGraphQLParserT__1, EGraphQLParserT__2, EGraphQLParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(160)
			p.OperationDefinition()
		}

	case EGraphQLParserT__9:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(161)
			p.FragmentDefinition()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOperationDefinitionContext is an interface to support dynamic dispatch.
type IOperationDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperationDefinitionContext differentiates from other interfaces.
	IsOperationDefinitionContext()
}

type OperationDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationDefinitionContext() *OperationDefinitionContext {
	var p = new(OperationDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_operationDefinition
	return p
}

func (*OperationDefinitionContext) IsOperationDefinitionContext() {}

func NewOperationDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationDefinitionContext {
	var p = new(OperationDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_operationDefinition

	return p
}

func (s *OperationDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationDefinitionContext) OperationType() IOperationTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperationTypeContext)
}

func (s *OperationDefinitionContext) SelectionSet() ISelectionSetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionSetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *OperationDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *OperationDefinitionContext) VariableDefinitions() IVariableDefinitionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDefinitionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDefinitionsContext)
}

func (s *OperationDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *OperationDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterOperationDefinition(s)
	}
}

func (s *OperationDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitOperationDefinition(s)
	}
}

func (p *EGraphQLParser) OperationDefinition() (localctx IOperationDefinitionContext) {
	localctx = NewOperationDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, EGraphQLParserRULE_operationDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(177)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case EGraphQLParserT__0, EGraphQLParserT__1, EGraphQLParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(164)
			p.OperationType()
		}
		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EGraphQLParserNAME {
			{
				p.SetState(165)
				p.Name()
			}

		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EGraphQLParserT__5 {
			{
				p.SetState(168)
				p.VariableDefinitions()
			}

		}
		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EGraphQLParserT__19 {
			{
				p.SetState(171)
				p.Directives()
			}

		}
		{
			p.SetState(174)
			p.SelectionSet()
		}

	case EGraphQLParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(176)
			p.SelectionSet()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOperationTypeContext is an interface to support dynamic dispatch.
type IOperationTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperationTypeContext differentiates from other interfaces.
	IsOperationTypeContext()
}

type OperationTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationTypeContext() *OperationTypeContext {
	var p = new(OperationTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_operationType
	return p
}

func (*OperationTypeContext) IsOperationTypeContext() {}

func NewOperationTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationTypeContext {
	var p = new(OperationTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_operationType

	return p
}

func (s *OperationTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *OperationTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterOperationType(s)
	}
}

func (s *OperationTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitOperationType(s)
	}
}

func (p *EGraphQLParser) OperationType() (localctx IOperationTypeContext) {
	localctx = NewOperationTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, EGraphQLParserRULE_operationType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<EGraphQLParserT__0)|(1<<EGraphQLParserT__1)|(1<<EGraphQLParserT__2))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISelectionSetContext is an interface to support dynamic dispatch.
type ISelectionSetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectionSetContext differentiates from other interfaces.
	IsSelectionSetContext()
}

type SelectionSetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectionSetContext() *SelectionSetContext {
	var p = new(SelectionSetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_selectionSet
	return p
}

func (*SelectionSetContext) IsSelectionSetContext() {}

func NewSelectionSetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectionSetContext {
	var p = new(SelectionSetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_selectionSet

	return p
}

func (s *SelectionSetContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectionSetContext) AllSelection() []ISelectionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISelectionContext)(nil)).Elem())
	var tst = make([]ISelectionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISelectionContext)
		}
	}

	return tst
}

func (s *SelectionSetContext) Selection(i int) ISelectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISelectionContext)
}

func (s *SelectionSetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectionSetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectionSetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterSelectionSet(s)
	}
}

func (s *SelectionSetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitSelectionSet(s)
	}
}

func (p *EGraphQLParser) SelectionSet() (localctx ISelectionSetContext) {
	localctx = NewSelectionSetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, EGraphQLParserRULE_selectionSet)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.Match(EGraphQLParserT__3)
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == EGraphQLParserT__8 || _la == EGraphQLParserNAME {
		{
			p.SetState(182)
			p.Selection()
		}

		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(187)
		p.Match(EGraphQLParserT__4)
	}

	return localctx
}

// ISelectionContext is an interface to support dynamic dispatch.
type ISelectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectionContext differentiates from other interfaces.
	IsSelectionContext()
}

type SelectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectionContext() *SelectionContext {
	var p = new(SelectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_selection
	return p
}

func (*SelectionContext) IsSelectionContext() {}

func NewSelectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectionContext {
	var p = new(SelectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_selection

	return p
}

func (s *SelectionContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectionContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *SelectionContext) FragmentSpread() IFragmentSpreadContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFragmentSpreadContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFragmentSpreadContext)
}

func (s *SelectionContext) InlineFragment() IInlineFragmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInlineFragmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInlineFragmentContext)
}

func (s *SelectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterSelection(s)
	}
}

func (s *SelectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitSelection(s)
	}
}

func (p *EGraphQLParser) Selection() (localctx ISelectionContext) {
	localctx = NewSelectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, EGraphQLParserRULE_selection)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(189)
			p.Field()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(190)
			p.FragmentSpread()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(191)
			p.InlineFragment()
		}

	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *FieldContext) Alias() IAliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAliasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAliasContext)
}

func (s *FieldContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *FieldContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *FieldContext) SelectionSet() ISelectionSetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionSetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *EGraphQLParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, EGraphQLParserRULE_field)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(195)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(194)
			p.Alias()
		}

	}
	{
		p.SetState(197)
		p.Name()
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__5 {
		{
			p.SetState(198)
			p.Arguments()
		}

	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__19 {
		{
			p.SetState(201)
			p.Directives()
		}

	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__3 {
		{
			p.SetState(204)
			p.SelectionSet()
		}

	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) AllArgument() []IArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentContext)(nil)).Elem())
	var tst = make([]IArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentContext)
		}
	}

	return tst
}

func (s *ArgumentsContext) Argument(i int) IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *EGraphQLParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, EGraphQLParserRULE_arguments)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(207)
		p.Match(EGraphQLParserT__5)
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == EGraphQLParserNAME {
		{
			p.SetState(208)
			p.Argument()
		}

		p.SetState(211)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(213)
		p.Match(EGraphQLParserT__6)
	}

	return localctx
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_argument
	return p
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ArgumentContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (p *EGraphQLParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, EGraphQLParserRULE_argument)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		p.Name()
	}
	{
		p.SetState(216)
		p.Match(EGraphQLParserT__7)
	}
	{
		p.SetState(217)
		p.Value()
	}

	return localctx
}

// IAliasContext is an interface to support dynamic dispatch.
type IAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAliasContext differentiates from other interfaces.
	IsAliasContext()
}

type AliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasContext() *AliasContext {
	var p = new(AliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_alias
	return p
}

func (*AliasContext) IsAliasContext() {}

func NewAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliasContext {
	var p = new(AliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_alias

	return p
}

func (s *AliasContext) GetParser() antlr.Parser { return s.parser }

func (s *AliasContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *AliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterAlias(s)
	}
}

func (s *AliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitAlias(s)
	}
}

func (p *EGraphQLParser) Alias() (localctx IAliasContext) {
	localctx = NewAliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, EGraphQLParserRULE_alias)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(219)
		p.Name()
	}
	{
		p.SetState(220)
		p.Match(EGraphQLParserT__7)
	}

	return localctx
}

// IFragmentSpreadContext is an interface to support dynamic dispatch.
type IFragmentSpreadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFragmentSpreadContext differentiates from other interfaces.
	IsFragmentSpreadContext()
}

type FragmentSpreadContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFragmentSpreadContext() *FragmentSpreadContext {
	var p = new(FragmentSpreadContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_fragmentSpread
	return p
}

func (*FragmentSpreadContext) IsFragmentSpreadContext() {}

func NewFragmentSpreadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FragmentSpreadContext {
	var p = new(FragmentSpreadContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_fragmentSpread

	return p
}

func (s *FragmentSpreadContext) GetParser() antlr.Parser { return s.parser }

func (s *FragmentSpreadContext) FragmentName() IFragmentNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFragmentNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFragmentNameContext)
}

func (s *FragmentSpreadContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *FragmentSpreadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FragmentSpreadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FragmentSpreadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterFragmentSpread(s)
	}
}

func (s *FragmentSpreadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitFragmentSpread(s)
	}
}

func (p *EGraphQLParser) FragmentSpread() (localctx IFragmentSpreadContext) {
	localctx = NewFragmentSpreadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, EGraphQLParserRULE_fragmentSpread)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
		p.Match(EGraphQLParserT__8)
	}
	{
		p.SetState(223)
		p.FragmentName()
	}
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__19 {
		{
			p.SetState(224)
			p.Directives()
		}

	}

	return localctx
}

// IFragmentDefinitionContext is an interface to support dynamic dispatch.
type IFragmentDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFragmentDefinitionContext differentiates from other interfaces.
	IsFragmentDefinitionContext()
}

type FragmentDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFragmentDefinitionContext() *FragmentDefinitionContext {
	var p = new(FragmentDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_fragmentDefinition
	return p
}

func (*FragmentDefinitionContext) IsFragmentDefinitionContext() {}

func NewFragmentDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FragmentDefinitionContext {
	var p = new(FragmentDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_fragmentDefinition

	return p
}

func (s *FragmentDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FragmentDefinitionContext) FragmentName() IFragmentNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFragmentNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFragmentNameContext)
}

func (s *FragmentDefinitionContext) TypeCondition() ITypeConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeConditionContext)
}

func (s *FragmentDefinitionContext) SelectionSet() ISelectionSetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionSetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *FragmentDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *FragmentDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FragmentDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FragmentDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterFragmentDefinition(s)
	}
}

func (s *FragmentDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitFragmentDefinition(s)
	}
}

func (p *EGraphQLParser) FragmentDefinition() (localctx IFragmentDefinitionContext) {
	localctx = NewFragmentDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, EGraphQLParserRULE_fragmentDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		p.Match(EGraphQLParserT__9)
	}
	{
		p.SetState(228)
		p.FragmentName()
	}
	{
		p.SetState(229)
		p.TypeCondition()
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__19 {
		{
			p.SetState(230)
			p.Directives()
		}

	}
	{
		p.SetState(233)
		p.SelectionSet()
	}

	return localctx
}

// IFragmentNameContext is an interface to support dynamic dispatch.
type IFragmentNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFragmentNameContext differentiates from other interfaces.
	IsFragmentNameContext()
}

type FragmentNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFragmentNameContext() *FragmentNameContext {
	var p = new(FragmentNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_fragmentName
	return p
}

func (*FragmentNameContext) IsFragmentNameContext() {}

func NewFragmentNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FragmentNameContext {
	var p = new(FragmentNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_fragmentName

	return p
}

func (s *FragmentNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FragmentNameContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *FragmentNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FragmentNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FragmentNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterFragmentName(s)
	}
}

func (s *FragmentNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitFragmentName(s)
	}
}

func (p *EGraphQLParser) FragmentName() (localctx IFragmentNameContext) {
	localctx = NewFragmentNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, EGraphQLParserRULE_fragmentName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Name()
	}

	return localctx
}

// ITypeConditionContext is an interface to support dynamic dispatch.
type ITypeConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeConditionContext differentiates from other interfaces.
	IsTypeConditionContext()
}

type TypeConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeConditionContext() *TypeConditionContext {
	var p = new(TypeConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_typeCondition
	return p
}

func (*TypeConditionContext) IsTypeConditionContext() {}

func NewTypeConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeConditionContext {
	var p = new(TypeConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_typeCondition

	return p
}

func (s *TypeConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeConditionContext) NamedType() INamedTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamedTypeContext)
}

func (s *TypeConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterTypeCondition(s)
	}
}

func (s *TypeConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitTypeCondition(s)
	}
}

func (p *EGraphQLParser) TypeCondition() (localctx ITypeConditionContext) {
	localctx = NewTypeConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, EGraphQLParserRULE_typeCondition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)
		p.Match(EGraphQLParserT__10)
	}
	{
		p.SetState(238)
		p.NamedType()
	}

	return localctx
}

// IInlineFragmentContext is an interface to support dynamic dispatch.
type IInlineFragmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInlineFragmentContext differentiates from other interfaces.
	IsInlineFragmentContext()
}

type InlineFragmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineFragmentContext() *InlineFragmentContext {
	var p = new(InlineFragmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_inlineFragment
	return p
}

func (*InlineFragmentContext) IsInlineFragmentContext() {}

func NewInlineFragmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineFragmentContext {
	var p = new(InlineFragmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_inlineFragment

	return p
}

func (s *InlineFragmentContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineFragmentContext) SelectionSet() ISelectionSetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionSetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *InlineFragmentContext) TypeCondition() ITypeConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeConditionContext)
}

func (s *InlineFragmentContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InlineFragmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineFragmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InlineFragmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterInlineFragment(s)
	}
}

func (s *InlineFragmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitInlineFragment(s)
	}
}

func (p *EGraphQLParser) InlineFragment() (localctx IInlineFragmentContext) {
	localctx = NewInlineFragmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, EGraphQLParserRULE_inlineFragment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)
		p.Match(EGraphQLParserT__8)
	}
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__10 {
		{
			p.SetState(241)
			p.TypeCondition()
		}

	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__19 {
		{
			p.SetState(244)
			p.Directives()
		}

	}
	{
		p.SetState(247)
		p.SelectionSet()
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ValueContext) IntValue() IIntValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntValueContext)
}

func (s *ValueContext) FloatValue() IFloatValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloatValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloatValueContext)
}

func (s *ValueContext) StringValue() IStringValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
}

func (s *ValueContext) BooleanValue() IBooleanValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanValueContext)
}

func (s *ValueContext) NullValue() INullValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INullValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INullValueContext)
}

func (s *ValueContext) EnumValue() IEnumValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumValueContext)
}

func (s *ValueContext) ListValue() IListValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListValueContext)
}

func (s *ValueContext) ObjectValue() IObjectValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectValueContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *EGraphQLParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, EGraphQLParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(258)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case EGraphQLParserT__16:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(249)
			p.Variable()
		}

	case EGraphQLParserINT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(250)
			p.IntValue()
		}

	case EGraphQLParserFLOAT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(251)
			p.FloatValue()
		}

	case EGraphQLParserSTRING, EGraphQLParserBLOCK_STRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(252)
			p.StringValue()
		}

	case EGraphQLParserT__11, EGraphQLParserT__12:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(253)
			p.BooleanValue()
		}

	case EGraphQLParserT__13:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(254)
			p.NullValue()
		}

	case EGraphQLParserNAME:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(255)
			p.EnumValue()
		}

	case EGraphQLParserT__14:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(256)
			p.ListValue()
		}

	case EGraphQLParserT__3:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(257)
			p.ObjectValue()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIntValueContext is an interface to support dynamic dispatch.
type IIntValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntValueContext differentiates from other interfaces.
	IsIntValueContext()
}

type IntValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntValueContext() *IntValueContext {
	var p = new(IntValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_intValue
	return p
}

func (*IntValueContext) IsIntValueContext() {}

func NewIntValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntValueContext {
	var p = new(IntValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_intValue

	return p
}

func (s *IntValueContext) GetParser() antlr.Parser { return s.parser }

func (s *IntValueContext) INT() antlr.TerminalNode {
	return s.GetToken(EGraphQLParserINT, 0)
}

func (s *IntValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterIntValue(s)
	}
}

func (s *IntValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitIntValue(s)
	}
}

func (p *EGraphQLParser) IntValue() (localctx IIntValueContext) {
	localctx = NewIntValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, EGraphQLParserRULE_intValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(260)
		p.Match(EGraphQLParserINT)
	}

	return localctx
}

// IFloatValueContext is an interface to support dynamic dispatch.
type IFloatValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloatValueContext differentiates from other interfaces.
	IsFloatValueContext()
}

type FloatValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatValueContext() *FloatValueContext {
	var p = new(FloatValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_floatValue
	return p
}

func (*FloatValueContext) IsFloatValueContext() {}

func NewFloatValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatValueContext {
	var p = new(FloatValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_floatValue

	return p
}

func (s *FloatValueContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatValueContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(EGraphQLParserFLOAT, 0)
}

func (s *FloatValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterFloatValue(s)
	}
}

func (s *FloatValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitFloatValue(s)
	}
}

func (p *EGraphQLParser) FloatValue() (localctx IFloatValueContext) {
	localctx = NewFloatValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, EGraphQLParserRULE_floatValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(262)
		p.Match(EGraphQLParserFLOAT)
	}

	return localctx
}

// IBooleanValueContext is an interface to support dynamic dispatch.
type IBooleanValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanValueContext differentiates from other interfaces.
	IsBooleanValueContext()
}

type BooleanValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanValueContext() *BooleanValueContext {
	var p = new(BooleanValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_booleanValue
	return p
}

func (*BooleanValueContext) IsBooleanValueContext() {}

func NewBooleanValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanValueContext {
	var p = new(BooleanValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_booleanValue

	return p
}

func (s *BooleanValueContext) GetParser() antlr.Parser { return s.parser }
func (s *BooleanValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterBooleanValue(s)
	}
}

func (s *BooleanValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitBooleanValue(s)
	}
}

func (p *EGraphQLParser) BooleanValue() (localctx IBooleanValueContext) {
	localctx = NewBooleanValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, EGraphQLParserRULE_booleanValue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)
		_la = p.GetTokenStream().LA(1)

		if !(_la == EGraphQLParserT__11 || _la == EGraphQLParserT__12) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IStringValueContext is an interface to support dynamic dispatch.
type IStringValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringValueContext differentiates from other interfaces.
	IsStringValueContext()
}

type StringValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringValueContext() *StringValueContext {
	var p = new(StringValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_stringValue
	return p
}

func (*StringValueContext) IsStringValueContext() {}

func NewStringValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringValueContext {
	var p = new(StringValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_stringValue

	return p
}

func (s *StringValueContext) GetParser() antlr.Parser { return s.parser }

func (s *StringValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(EGraphQLParserSTRING, 0)
}

func (s *StringValueContext) BLOCK_STRING() antlr.TerminalNode {
	return s.GetToken(EGraphQLParserBLOCK_STRING, 0)
}

func (s *StringValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterStringValue(s)
	}
}

func (s *StringValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitStringValue(s)
	}
}

func (p *EGraphQLParser) StringValue() (localctx IStringValueContext) {
	localctx = NewStringValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, EGraphQLParserRULE_stringValue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		_la = p.GetTokenStream().LA(1)

		if !(_la == EGraphQLParserSTRING || _la == EGraphQLParserBLOCK_STRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INullValueContext is an interface to support dynamic dispatch.
type INullValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNullValueContext differentiates from other interfaces.
	IsNullValueContext()
}

type NullValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNullValueContext() *NullValueContext {
	var p = new(NullValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_nullValue
	return p
}

func (*NullValueContext) IsNullValueContext() {}

func NewNullValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NullValueContext {
	var p = new(NullValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_nullValue

	return p
}

func (s *NullValueContext) GetParser() antlr.Parser { return s.parser }
func (s *NullValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NullValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterNullValue(s)
	}
}

func (s *NullValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitNullValue(s)
	}
}

func (p *EGraphQLParser) NullValue() (localctx INullValueContext) {
	localctx = NewNullValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, EGraphQLParserRULE_nullValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.Match(EGraphQLParserT__13)
	}

	return localctx
}

// IEnumValueContext is an interface to support dynamic dispatch.
type IEnumValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumValueContext differentiates from other interfaces.
	IsEnumValueContext()
}

type EnumValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueContext() *EnumValueContext {
	var p = new(EnumValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_enumValue
	return p
}

func (*EnumValueContext) IsEnumValueContext() {}

func NewEnumValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueContext {
	var p = new(EnumValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_enumValue

	return p
}

func (s *EnumValueContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *EnumValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterEnumValue(s)
	}
}

func (s *EnumValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitEnumValue(s)
	}
}

func (p *EGraphQLParser) EnumValue() (localctx IEnumValueContext) {
	localctx = NewEnumValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, EGraphQLParserRULE_enumValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(270)
		p.Name()
	}

	return localctx
}

// IListValueContext is an interface to support dynamic dispatch.
type IListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListValueContext differentiates from other interfaces.
	IsListValueContext()
}

type ListValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListValueContext() *ListValueContext {
	var p = new(ListValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_listValue
	return p
}

func (*ListValueContext) IsListValueContext() {}

func NewListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListValueContext {
	var p = new(ListValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_listValue

	return p
}

func (s *ListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ListValueContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *ListValueContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ListValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterListValue(s)
	}
}

func (s *ListValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitListValue(s)
	}
}

func (p *EGraphQLParser) ListValue() (localctx IListValueContext) {
	localctx = NewListValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, EGraphQLParserRULE_listValue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(272)
			p.Match(EGraphQLParserT__14)
		}
		{
			p.SetState(273)
			p.Match(EGraphQLParserT__15)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(274)
			p.Match(EGraphQLParserT__14)
		}
		p.SetState(276)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<EGraphQLParserT__3)|(1<<EGraphQLParserT__11)|(1<<EGraphQLParserT__12)|(1<<EGraphQLParserT__13)|(1<<EGraphQLParserT__14)|(1<<EGraphQLParserT__16))) != 0) || (((_la-57)&-(0x1f+1)) == 0 && ((1<<uint((_la-57)))&((1<<(EGraphQLParserNAME-57))|(1<<(EGraphQLParserSTRING-57))|(1<<(EGraphQLParserBLOCK_STRING-57))|(1<<(EGraphQLParserFLOAT-57))|(1<<(EGraphQLParserINT-57)))) != 0) {
			{
				p.SetState(275)
				p.Value()
			}

			p.SetState(278)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(280)
			p.Match(EGraphQLParserT__15)
		}

	}

	return localctx
}

// IObjectValueContext is an interface to support dynamic dispatch.
type IObjectValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectValueContext differentiates from other interfaces.
	IsObjectValueContext()
}

type ObjectValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectValueContext() *ObjectValueContext {
	var p = new(ObjectValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_objectValue
	return p
}

func (*ObjectValueContext) IsObjectValueContext() {}

func NewObjectValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectValueContext {
	var p = new(ObjectValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_objectValue

	return p
}

func (s *ObjectValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectValueContext) AllObjectField() []IObjectFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IObjectFieldContext)(nil)).Elem())
	var tst = make([]IObjectFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IObjectFieldContext)
		}
	}

	return tst
}

func (s *ObjectValueContext) ObjectField(i int) IObjectFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IObjectFieldContext)
}

func (s *ObjectValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterObjectValue(s)
	}
}

func (s *ObjectValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitObjectValue(s)
	}
}

func (p *EGraphQLParser) ObjectValue() (localctx IObjectValueContext) {
	localctx = NewObjectValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, EGraphQLParserRULE_objectValue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(284)
		p.Match(EGraphQLParserT__3)
	}
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == EGraphQLParserNAME {
		{
			p.SetState(285)
			p.ObjectField()
		}

		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(291)
		p.Match(EGraphQLParserT__4)
	}

	return localctx
}

// IObjectFieldContext is an interface to support dynamic dispatch.
type IObjectFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectFieldContext differentiates from other interfaces.
	IsObjectFieldContext()
}

type ObjectFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectFieldContext() *ObjectFieldContext {
	var p = new(ObjectFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_objectField
	return p
}

func (*ObjectFieldContext) IsObjectFieldContext() {}

func NewObjectFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectFieldContext {
	var p = new(ObjectFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_objectField

	return p
}

func (s *ObjectFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectFieldContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ObjectFieldContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ObjectFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterObjectField(s)
	}
}

func (s *ObjectFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitObjectField(s)
	}
}

func (p *EGraphQLParser) ObjectField() (localctx IObjectFieldContext) {
	localctx = NewObjectFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, EGraphQLParserRULE_objectField)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		p.Name()
	}
	{
		p.SetState(294)
		p.Match(EGraphQLParserT__7)
	}
	{
		p.SetState(295)
		p.Value()
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *EGraphQLParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, EGraphQLParserRULE_variable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(297)
		p.Match(EGraphQLParserT__16)
	}
	{
		p.SetState(298)
		p.Name()
	}

	return localctx
}

// IVariableDefinitionsContext is an interface to support dynamic dispatch.
type IVariableDefinitionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDefinitionsContext differentiates from other interfaces.
	IsVariableDefinitionsContext()
}

type VariableDefinitionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDefinitionsContext() *VariableDefinitionsContext {
	var p = new(VariableDefinitionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_variableDefinitions
	return p
}

func (*VariableDefinitionsContext) IsVariableDefinitionsContext() {}

func NewVariableDefinitionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDefinitionsContext {
	var p = new(VariableDefinitionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_variableDefinitions

	return p
}

func (s *VariableDefinitionsContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDefinitionsContext) AllVariableDefinition() []IVariableDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableDefinitionContext)(nil)).Elem())
	var tst = make([]IVariableDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableDefinitionContext)
		}
	}

	return tst
}

func (s *VariableDefinitionsContext) VariableDefinition(i int) IVariableDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableDefinitionContext)
}

func (s *VariableDefinitionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDefinitionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDefinitionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterVariableDefinitions(s)
	}
}

func (s *VariableDefinitionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitVariableDefinitions(s)
	}
}

func (p *EGraphQLParser) VariableDefinitions() (localctx IVariableDefinitionsContext) {
	localctx = NewVariableDefinitionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, EGraphQLParserRULE_variableDefinitions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(300)
		p.Match(EGraphQLParserT__5)
	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == EGraphQLParserT__16 {
		{
			p.SetState(301)
			p.VariableDefinition()
		}

		p.SetState(304)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(306)
		p.Match(EGraphQLParserT__6)
	}

	return localctx
}

// IVariableDefinitionContext is an interface to support dynamic dispatch.
type IVariableDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDefinitionContext differentiates from other interfaces.
	IsVariableDefinitionContext()
}

type VariableDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDefinitionContext() *VariableDefinitionContext {
	var p = new(VariableDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_variableDefinition
	return p
}

func (*VariableDefinitionContext) IsVariableDefinitionContext() {}

func NewVariableDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDefinitionContext {
	var p = new(VariableDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_variableDefinition

	return p
}

func (s *VariableDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDefinitionContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *VariableDefinitionContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *VariableDefinitionContext) DefaultValue() IDefaultValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultValueContext)
}

func (s *VariableDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterVariableDefinition(s)
	}
}

func (s *VariableDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitVariableDefinition(s)
	}
}

func (p *EGraphQLParser) VariableDefinition() (localctx IVariableDefinitionContext) {
	localctx = NewVariableDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, EGraphQLParserRULE_variableDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(308)
		p.Variable()
	}
	{
		p.SetState(309)
		p.Match(EGraphQLParserT__7)
	}
	{
		p.SetState(310)
		p.Type_()
	}
	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__17 {
		{
			p.SetState(311)
			p.DefaultValue()
		}

	}

	return localctx
}

// IDefaultValueContext is an interface to support dynamic dispatch.
type IDefaultValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefaultValueContext differentiates from other interfaces.
	IsDefaultValueContext()
}

type DefaultValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultValueContext() *DefaultValueContext {
	var p = new(DefaultValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_defaultValue
	return p
}

func (*DefaultValueContext) IsDefaultValueContext() {}

func NewDefaultValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultValueContext {
	var p = new(DefaultValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_defaultValue

	return p
}

func (s *DefaultValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultValueContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *DefaultValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterDefaultValue(s)
	}
}

func (s *DefaultValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitDefaultValue(s)
	}
}

func (p *EGraphQLParser) DefaultValue() (localctx IDefaultValueContext) {
	localctx = NewDefaultValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, EGraphQLParserRULE_defaultValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)
		p.Match(EGraphQLParserT__17)
	}
	{
		p.SetState(315)
		p.Value()
	}

	return localctx
}

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_type_
	return p
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) NamedType() INamedTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamedTypeContext)
}

func (s *Type_Context) ListType() IListTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListTypeContext)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitType_(s)
	}
}

func (p *EGraphQLParser) Type_() (localctx IType_Context) {
	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, EGraphQLParserRULE_type_)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(325)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case EGraphQLParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(317)
			p.NamedType()
		}
		p.SetState(319)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EGraphQLParserT__18 {
			{
				p.SetState(318)
				p.Match(EGraphQLParserT__18)
			}

		}

	case EGraphQLParserT__14:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(321)
			p.ListType()
		}
		p.SetState(323)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EGraphQLParserT__18 {
			{
				p.SetState(322)
				p.Match(EGraphQLParserT__18)
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INamedTypeContext is an interface to support dynamic dispatch.
type INamedTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamedTypeContext differentiates from other interfaces.
	IsNamedTypeContext()
}

type NamedTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamedTypeContext() *NamedTypeContext {
	var p = new(NamedTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_namedType
	return p
}

func (*NamedTypeContext) IsNamedTypeContext() {}

func NewNamedTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedTypeContext {
	var p = new(NamedTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_namedType

	return p
}

func (s *NamedTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedTypeContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *NamedTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterNamedType(s)
	}
}

func (s *NamedTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitNamedType(s)
	}
}

func (p *EGraphQLParser) NamedType() (localctx INamedTypeContext) {
	localctx = NewNamedTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, EGraphQLParserRULE_namedType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(327)
		p.Name()
	}

	return localctx
}

// IListTypeContext is an interface to support dynamic dispatch.
type IListTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListTypeContext differentiates from other interfaces.
	IsListTypeContext()
}

type ListTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListTypeContext() *ListTypeContext {
	var p = new(ListTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_listType
	return p
}

func (*ListTypeContext) IsListTypeContext() {}

func NewListTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListTypeContext {
	var p = new(ListTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_listType

	return p
}

func (s *ListTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ListTypeContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ListTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterListType(s)
	}
}

func (s *ListTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitListType(s)
	}
}

func (p *EGraphQLParser) ListType() (localctx IListTypeContext) {
	localctx = NewListTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, EGraphQLParserRULE_listType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(329)
		p.Match(EGraphQLParserT__14)
	}
	{
		p.SetState(330)
		p.Type_()
	}
	{
		p.SetState(331)
		p.Match(EGraphQLParserT__15)
	}

	return localctx
}

// IDirectivesContext is an interface to support dynamic dispatch.
type IDirectivesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectivesContext differentiates from other interfaces.
	IsDirectivesContext()
}

type DirectivesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectivesContext() *DirectivesContext {
	var p = new(DirectivesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_directives
	return p
}

func (*DirectivesContext) IsDirectivesContext() {}

func NewDirectivesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectivesContext {
	var p = new(DirectivesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_directives

	return p
}

func (s *DirectivesContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectivesContext) AllDirective() []IDirectiveContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDirectiveContext)(nil)).Elem())
	var tst = make([]IDirectiveContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDirectiveContext)
		}
	}

	return tst
}

func (s *DirectivesContext) Directive(i int) IDirectiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDirectiveContext)
}

func (s *DirectivesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectivesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectivesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterDirectives(s)
	}
}

func (s *DirectivesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitDirectives(s)
	}
}

func (p *EGraphQLParser) Directives() (localctx IDirectivesContext) {
	localctx = NewDirectivesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, EGraphQLParserRULE_directives)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == EGraphQLParserT__19 {
		{
			p.SetState(333)
			p.Directive()
		}

		p.SetState(336)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDirectiveContext is an interface to support dynamic dispatch.
type IDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectiveContext differentiates from other interfaces.
	IsDirectiveContext()
}

type DirectiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveContext() *DirectiveContext {
	var p = new(DirectiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_directive
	return p
}

func (*DirectiveContext) IsDirectiveContext() {}

func NewDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveContext {
	var p = new(DirectiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_directive

	return p
}

func (s *DirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *DirectiveContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *DirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterDirective(s)
	}
}

func (s *DirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitDirective(s)
	}
}

func (p *EGraphQLParser) Directive() (localctx IDirectiveContext) {
	localctx = NewDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, EGraphQLParserRULE_directive)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(338)
		p.Match(EGraphQLParserT__19)
	}
	{
		p.SetState(339)
		p.Name()
	}
	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__5 {
		{
			p.SetState(340)
			p.Arguments()
		}

	}

	return localctx
}

// ITypeSystemDefinitionContext is an interface to support dynamic dispatch.
type ITypeSystemDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSystemDefinitionContext differentiates from other interfaces.
	IsTypeSystemDefinitionContext()
}

type TypeSystemDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSystemDefinitionContext() *TypeSystemDefinitionContext {
	var p = new(TypeSystemDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_typeSystemDefinition
	return p
}

func (*TypeSystemDefinitionContext) IsTypeSystemDefinitionContext() {}

func NewTypeSystemDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSystemDefinitionContext {
	var p = new(TypeSystemDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_typeSystemDefinition

	return p
}

func (s *TypeSystemDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSystemDefinitionContext) SchemaDefinition() ISchemaDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchemaDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISchemaDefinitionContext)
}

func (s *TypeSystemDefinitionContext) TypeDefinition() ITypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDefinitionContext)
}

func (s *TypeSystemDefinitionContext) DirectiveDefinition() IDirectiveDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectiveDefinitionContext)
}

func (s *TypeSystemDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSystemDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSystemDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterTypeSystemDefinition(s)
	}
}

func (s *TypeSystemDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitTypeSystemDefinition(s)
	}
}

func (p *EGraphQLParser) TypeSystemDefinition() (localctx ITypeSystemDefinitionContext) {
	localctx = NewTypeSystemDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, EGraphQLParserRULE_typeSystemDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(346)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(343)
			p.SchemaDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(344)
			p.TypeDefinition()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(345)
			p.DirectiveDefinition()
		}

	}

	return localctx
}

// ITypeSystemExtensionContext is an interface to support dynamic dispatch.
type ITypeSystemExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSystemExtensionContext differentiates from other interfaces.
	IsTypeSystemExtensionContext()
}

type TypeSystemExtensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSystemExtensionContext() *TypeSystemExtensionContext {
	var p = new(TypeSystemExtensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_typeSystemExtension
	return p
}

func (*TypeSystemExtensionContext) IsTypeSystemExtensionContext() {}

func NewTypeSystemExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSystemExtensionContext {
	var p = new(TypeSystemExtensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_typeSystemExtension

	return p
}

func (s *TypeSystemExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSystemExtensionContext) SchemaExtension() ISchemaExtensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchemaExtensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISchemaExtensionContext)
}

func (s *TypeSystemExtensionContext) TypeExtension() ITypeExtensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeExtensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeExtensionContext)
}

func (s *TypeSystemExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSystemExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSystemExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterTypeSystemExtension(s)
	}
}

func (s *TypeSystemExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitTypeSystemExtension(s)
	}
}

func (p *EGraphQLParser) TypeSystemExtension() (localctx ITypeSystemExtensionContext) {
	localctx = NewTypeSystemExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, EGraphQLParserRULE_typeSystemExtension)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(348)
			p.SchemaExtension()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(349)
			p.TypeExtension()
		}

	}

	return localctx
}

// ISchemaDefinitionContext is an interface to support dynamic dispatch.
type ISchemaDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSchemaDefinitionContext differentiates from other interfaces.
	IsSchemaDefinitionContext()
}

type SchemaDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchemaDefinitionContext() *SchemaDefinitionContext {
	var p = new(SchemaDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_schemaDefinition
	return p
}

func (*SchemaDefinitionContext) IsSchemaDefinitionContext() {}

func NewSchemaDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchemaDefinitionContext {
	var p = new(SchemaDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_schemaDefinition

	return p
}

func (s *SchemaDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *SchemaDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *SchemaDefinitionContext) AllRootOperationTypeDefinition() []IRootOperationTypeDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRootOperationTypeDefinitionContext)(nil)).Elem())
	var tst = make([]IRootOperationTypeDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRootOperationTypeDefinitionContext)
		}
	}

	return tst
}

func (s *SchemaDefinitionContext) RootOperationTypeDefinition(i int) IRootOperationTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRootOperationTypeDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRootOperationTypeDefinitionContext)
}

func (s *SchemaDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchemaDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SchemaDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterSchemaDefinition(s)
	}
}

func (s *SchemaDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitSchemaDefinition(s)
	}
}

func (p *EGraphQLParser) SchemaDefinition() (localctx ISchemaDefinitionContext) {
	localctx = NewSchemaDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, EGraphQLParserRULE_schemaDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(352)
		p.Match(EGraphQLParserT__20)
	}
	p.SetState(354)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__19 {
		{
			p.SetState(353)
			p.Directives()
		}

	}
	{
		p.SetState(356)
		p.Match(EGraphQLParserT__3)
	}
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<EGraphQLParserT__0)|(1<<EGraphQLParserT__1)|(1<<EGraphQLParserT__2))) != 0) {
		{
			p.SetState(357)
			p.RootOperationTypeDefinition()
		}

		p.SetState(360)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(362)
		p.Match(EGraphQLParserT__4)
	}

	return localctx
}

// IRootOperationTypeDefinitionContext is an interface to support dynamic dispatch.
type IRootOperationTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRootOperationTypeDefinitionContext differentiates from other interfaces.
	IsRootOperationTypeDefinitionContext()
}

type RootOperationTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootOperationTypeDefinitionContext() *RootOperationTypeDefinitionContext {
	var p = new(RootOperationTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_rootOperationTypeDefinition
	return p
}

func (*RootOperationTypeDefinitionContext) IsRootOperationTypeDefinitionContext() {}

func NewRootOperationTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootOperationTypeDefinitionContext {
	var p = new(RootOperationTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_rootOperationTypeDefinition

	return p
}

func (s *RootOperationTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *RootOperationTypeDefinitionContext) OperationType() IOperationTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperationTypeContext)
}

func (s *RootOperationTypeDefinitionContext) NamedType() INamedTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamedTypeContext)
}

func (s *RootOperationTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootOperationTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootOperationTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterRootOperationTypeDefinition(s)
	}
}

func (s *RootOperationTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitRootOperationTypeDefinition(s)
	}
}

func (p *EGraphQLParser) RootOperationTypeDefinition() (localctx IRootOperationTypeDefinitionContext) {
	localctx = NewRootOperationTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, EGraphQLParserRULE_rootOperationTypeDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(364)
		p.OperationType()
	}
	{
		p.SetState(365)
		p.Match(EGraphQLParserT__7)
	}
	{
		p.SetState(366)
		p.NamedType()
	}

	return localctx
}

// ISchemaExtensionContext is an interface to support dynamic dispatch.
type ISchemaExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSchemaExtensionContext differentiates from other interfaces.
	IsSchemaExtensionContext()
}

type SchemaExtensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchemaExtensionContext() *SchemaExtensionContext {
	var p = new(SchemaExtensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_schemaExtension
	return p
}

func (*SchemaExtensionContext) IsSchemaExtensionContext() {}

func NewSchemaExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchemaExtensionContext {
	var p = new(SchemaExtensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_schemaExtension

	return p
}

func (s *SchemaExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *SchemaExtensionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *SchemaExtensionContext) AllOperationTypeDefinition() []IOperationTypeDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOperationTypeDefinitionContext)(nil)).Elem())
	var tst = make([]IOperationTypeDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOperationTypeDefinitionContext)
		}
	}

	return tst
}

func (s *SchemaExtensionContext) OperationTypeDefinition(i int) IOperationTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationTypeDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOperationTypeDefinitionContext)
}

func (s *SchemaExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchemaExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SchemaExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterSchemaExtension(s)
	}
}

func (s *SchemaExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitSchemaExtension(s)
	}
}

func (p *EGraphQLParser) SchemaExtension() (localctx ISchemaExtensionContext) {
	localctx = NewSchemaExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, EGraphQLParserRULE_schemaExtension)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(368)
			p.Match(EGraphQLParserT__21)
		}
		{
			p.SetState(369)
			p.Match(EGraphQLParserT__20)
		}
		p.SetState(371)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EGraphQLParserT__19 {
			{
				p.SetState(370)
				p.Directives()
			}

		}
		{
			p.SetState(373)
			p.Match(EGraphQLParserT__3)
		}
		p.SetState(375)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<EGraphQLParserT__0)|(1<<EGraphQLParserT__1)|(1<<EGraphQLParserT__2))) != 0) {
			{
				p.SetState(374)
				p.OperationTypeDefinition()
			}

			p.SetState(377)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(379)
			p.Match(EGraphQLParserT__4)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(381)
			p.Match(EGraphQLParserT__21)
		}
		{
			p.SetState(382)
			p.Match(EGraphQLParserT__20)
		}
		{
			p.SetState(383)
			p.Directives()
		}

	}

	return localctx
}

// IOperationTypeDefinitionContext is an interface to support dynamic dispatch.
type IOperationTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperationTypeDefinitionContext differentiates from other interfaces.
	IsOperationTypeDefinitionContext()
}

type OperationTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationTypeDefinitionContext() *OperationTypeDefinitionContext {
	var p = new(OperationTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_operationTypeDefinition
	return p
}

func (*OperationTypeDefinitionContext) IsOperationTypeDefinitionContext() {}

func NewOperationTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationTypeDefinitionContext {
	var p = new(OperationTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_operationTypeDefinition

	return p
}

func (s *OperationTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationTypeDefinitionContext) OperationType() IOperationTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperationTypeContext)
}

func (s *OperationTypeDefinitionContext) NamedType() INamedTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamedTypeContext)
}

func (s *OperationTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterOperationTypeDefinition(s)
	}
}

func (s *OperationTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitOperationTypeDefinition(s)
	}
}

func (p *EGraphQLParser) OperationTypeDefinition() (localctx IOperationTypeDefinitionContext) {
	localctx = NewOperationTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, EGraphQLParserRULE_operationTypeDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(386)
		p.OperationType()
	}
	{
		p.SetState(387)
		p.Match(EGraphQLParserT__7)
	}
	{
		p.SetState(388)
		p.NamedType()
	}

	return localctx
}

// IDescriptionContext is an interface to support dynamic dispatch.
type IDescriptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDescriptionContext differentiates from other interfaces.
	IsDescriptionContext()
}

type DescriptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescriptionContext() *DescriptionContext {
	var p = new(DescriptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_description
	return p
}

func (*DescriptionContext) IsDescriptionContext() {}

func NewDescriptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DescriptionContext {
	var p = new(DescriptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_description

	return p
}

func (s *DescriptionContext) GetParser() antlr.Parser { return s.parser }

func (s *DescriptionContext) StringValue() IStringValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
}

func (s *DescriptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescriptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DescriptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterDescription(s)
	}
}

func (s *DescriptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitDescription(s)
	}
}

func (p *EGraphQLParser) Description() (localctx IDescriptionContext) {
	localctx = NewDescriptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, EGraphQLParserRULE_description)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(390)
		p.StringValue()
	}

	return localctx
}

// ITypeDefinitionContext is an interface to support dynamic dispatch.
type ITypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDefinitionContext differentiates from other interfaces.
	IsTypeDefinitionContext()
}

type TypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDefinitionContext() *TypeDefinitionContext {
	var p = new(TypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_typeDefinition
	return p
}

func (*TypeDefinitionContext) IsTypeDefinitionContext() {}

func NewTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDefinitionContext {
	var p = new(TypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_typeDefinition

	return p
}

func (s *TypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDefinitionContext) ScalarTypeDefinition() IScalarTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScalarTypeDefinitionContext)
}

func (s *TypeDefinitionContext) UnionTypeDefinition() IUnionTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionTypeDefinitionContext)
}

func (s *TypeDefinitionContext) EnumTypeDefinition() IEnumTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumTypeDefinitionContext)
}

func (s *TypeDefinitionContext) TempletableTypeDefinition() ITempletableTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITempletableTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITempletableTypeDefinitionContext)
}

func (s *TypeDefinitionContext) TemplateTypeDefinition() ITemplateTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemplateTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITemplateTypeDefinitionContext)
}

func (s *TypeDefinitionContext) TemplateImplementedTypeDefinition() ITemplateImplementedTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemplateImplementedTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITemplateImplementedTypeDefinitionContext)
}

func (s *TypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterTypeDefinition(s)
	}
}

func (s *TypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitTypeDefinition(s)
	}
}

func (p *EGraphQLParser) TypeDefinition() (localctx ITypeDefinitionContext) {
	localctx = NewTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, EGraphQLParserRULE_typeDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(398)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(392)
			p.ScalarTypeDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(393)
			p.UnionTypeDefinition()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(394)
			p.EnumTypeDefinition()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(395)
			p.TempletableTypeDefinition()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(396)
			p.TemplateTypeDefinition()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(397)
			p.TemplateImplementedTypeDefinition()
		}

	}

	return localctx
}

// ITempletableTypeDefinitionContext is an interface to support dynamic dispatch.
type ITempletableTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTempletableTypeDefinitionContext differentiates from other interfaces.
	IsTempletableTypeDefinitionContext()
}

type TempletableTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTempletableTypeDefinitionContext() *TempletableTypeDefinitionContext {
	var p = new(TempletableTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_templetableTypeDefinition
	return p
}

func (*TempletableTypeDefinitionContext) IsTempletableTypeDefinitionContext() {}

func NewTempletableTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TempletableTypeDefinitionContext {
	var p = new(TempletableTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_templetableTypeDefinition

	return p
}

func (s *TempletableTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *TempletableTypeDefinitionContext) ObjectTypeDefinition() IObjectTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectTypeDefinitionContext)
}

func (s *TempletableTypeDefinitionContext) InterfaceTypeDefinition() IInterfaceTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterfaceTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInterfaceTypeDefinitionContext)
}

func (s *TempletableTypeDefinitionContext) InputObjectTypeDefinition() IInputObjectTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputObjectTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInputObjectTypeDefinitionContext)
}

func (s *TempletableTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TempletableTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TempletableTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterTempletableTypeDefinition(s)
	}
}

func (s *TempletableTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitTempletableTypeDefinition(s)
	}
}

func (p *EGraphQLParser) TempletableTypeDefinition() (localctx ITempletableTypeDefinitionContext) {
	localctx = NewTempletableTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, EGraphQLParserRULE_templetableTypeDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(403)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(400)
			p.ObjectTypeDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(401)
			p.InterfaceTypeDefinition()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(402)
			p.InputObjectTypeDefinition()
		}

	}

	return localctx
}

// ITypeExtensionContext is an interface to support dynamic dispatch.
type ITypeExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeExtensionContext differentiates from other interfaces.
	IsTypeExtensionContext()
}

type TypeExtensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeExtensionContext() *TypeExtensionContext {
	var p = new(TypeExtensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_typeExtension
	return p
}

func (*TypeExtensionContext) IsTypeExtensionContext() {}

func NewTypeExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeExtensionContext {
	var p = new(TypeExtensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_typeExtension

	return p
}

func (s *TypeExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeExtensionContext) ScalarTypeExtension() IScalarTypeExtensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarTypeExtensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScalarTypeExtensionContext)
}

func (s *TypeExtensionContext) ObjectTypeExtension() IObjectTypeExtensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectTypeExtensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectTypeExtensionContext)
}

func (s *TypeExtensionContext) InterfaceTypeExtension() IInterfaceTypeExtensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterfaceTypeExtensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInterfaceTypeExtensionContext)
}

func (s *TypeExtensionContext) UnionTypeExtension() IUnionTypeExtensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionTypeExtensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionTypeExtensionContext)
}

func (s *TypeExtensionContext) EnumTypeExtension() IEnumTypeExtensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumTypeExtensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumTypeExtensionContext)
}

func (s *TypeExtensionContext) InputObjectTypeExtension() IInputObjectTypeExtensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputObjectTypeExtensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInputObjectTypeExtensionContext)
}

func (s *TypeExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterTypeExtension(s)
	}
}

func (s *TypeExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitTypeExtension(s)
	}
}

func (p *EGraphQLParser) TypeExtension() (localctx ITypeExtensionContext) {
	localctx = NewTypeExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, EGraphQLParserRULE_typeExtension)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(411)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(405)
			p.ScalarTypeExtension()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(406)
			p.ObjectTypeExtension()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(407)
			p.InterfaceTypeExtension()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(408)
			p.UnionTypeExtension()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(409)
			p.EnumTypeExtension()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(410)
			p.InputObjectTypeExtension()
		}

	}

	return localctx
}

// IScalarTypeDefinitionContext is an interface to support dynamic dispatch.
type IScalarTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScalarTypeDefinitionContext differentiates from other interfaces.
	IsScalarTypeDefinitionContext()
}

type ScalarTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScalarTypeDefinitionContext() *ScalarTypeDefinitionContext {
	var p = new(ScalarTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_scalarTypeDefinition
	return p
}

func (*ScalarTypeDefinitionContext) IsScalarTypeDefinitionContext() {}

func NewScalarTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScalarTypeDefinitionContext {
	var p = new(ScalarTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_scalarTypeDefinition

	return p
}

func (s *ScalarTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ScalarTypeDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ScalarTypeDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *ScalarTypeDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *ScalarTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScalarTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterScalarTypeDefinition(s)
	}
}

func (s *ScalarTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitScalarTypeDefinition(s)
	}
}

func (p *EGraphQLParser) ScalarTypeDefinition() (localctx IScalarTypeDefinitionContext) {
	localctx = NewScalarTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, EGraphQLParserRULE_scalarTypeDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(414)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserSTRING || _la == EGraphQLParserBLOCK_STRING {
		{
			p.SetState(413)
			p.Description()
		}

	}
	{
		p.SetState(416)
		p.Match(EGraphQLParserT__22)
	}
	{
		p.SetState(417)
		p.Name()
	}
	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__19 {
		{
			p.SetState(418)
			p.Directives()
		}

	}

	return localctx
}

// IScalarTypeExtensionContext is an interface to support dynamic dispatch.
type IScalarTypeExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScalarTypeExtensionContext differentiates from other interfaces.
	IsScalarTypeExtensionContext()
}

type ScalarTypeExtensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScalarTypeExtensionContext() *ScalarTypeExtensionContext {
	var p = new(ScalarTypeExtensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_scalarTypeExtension
	return p
}

func (*ScalarTypeExtensionContext) IsScalarTypeExtensionContext() {}

func NewScalarTypeExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScalarTypeExtensionContext {
	var p = new(ScalarTypeExtensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_scalarTypeExtension

	return p
}

func (s *ScalarTypeExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *ScalarTypeExtensionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ScalarTypeExtensionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *ScalarTypeExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarTypeExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScalarTypeExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterScalarTypeExtension(s)
	}
}

func (s *ScalarTypeExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitScalarTypeExtension(s)
	}
}

func (p *EGraphQLParser) ScalarTypeExtension() (localctx IScalarTypeExtensionContext) {
	localctx = NewScalarTypeExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, EGraphQLParserRULE_scalarTypeExtension)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(421)
		p.Match(EGraphQLParserT__23)
	}
	{
		p.SetState(422)
		p.Match(EGraphQLParserT__22)
	}
	{
		p.SetState(423)
		p.Name()
	}
	{
		p.SetState(424)
		p.Directives()
	}

	return localctx
}

// IObjectTypeDefinitionContext is an interface to support dynamic dispatch.
type IObjectTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectTypeDefinitionContext differentiates from other interfaces.
	IsObjectTypeDefinitionContext()
}

type ObjectTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectTypeDefinitionContext() *ObjectTypeDefinitionContext {
	var p = new(ObjectTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_objectTypeDefinition
	return p
}

func (*ObjectTypeDefinitionContext) IsObjectTypeDefinitionContext() {}

func NewObjectTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectTypeDefinitionContext {
	var p = new(ObjectTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_objectTypeDefinition

	return p
}

func (s *ObjectTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectTypeDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ObjectTypeDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *ObjectTypeDefinitionContext) ImplementsInterfaces() IImplementsInterfacesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImplementsInterfacesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImplementsInterfacesContext)
}

func (s *ObjectTypeDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *ObjectTypeDefinitionContext) FieldsDefinition() IFieldsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldsDefinitionContext)
}

func (s *ObjectTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterObjectTypeDefinition(s)
	}
}

func (s *ObjectTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitObjectTypeDefinition(s)
	}
}

func (p *EGraphQLParser) ObjectTypeDefinition() (localctx IObjectTypeDefinitionContext) {
	localctx = NewObjectTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, EGraphQLParserRULE_objectTypeDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(427)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserSTRING || _la == EGraphQLParserBLOCK_STRING {
		{
			p.SetState(426)
			p.Description()
		}

	}
	{
		p.SetState(429)
		p.Match(EGraphQLParserT__24)
	}
	{
		p.SetState(430)
		p.Name()
	}
	p.SetState(432)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__25 {
		{
			p.SetState(431)
			p.implementsInterfaces(0)
		}

	}
	p.SetState(435)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__19 {
		{
			p.SetState(434)
			p.Directives()
		}

	}
	p.SetState(438)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(437)
			p.FieldsDefinition()
		}

	}

	return localctx
}

// IImplementsInterfacesContext is an interface to support dynamic dispatch.
type IImplementsInterfacesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImplementsInterfacesContext differentiates from other interfaces.
	IsImplementsInterfacesContext()
}

type ImplementsInterfacesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImplementsInterfacesContext() *ImplementsInterfacesContext {
	var p = new(ImplementsInterfacesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_implementsInterfaces
	return p
}

func (*ImplementsInterfacesContext) IsImplementsInterfacesContext() {}

func NewImplementsInterfacesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImplementsInterfacesContext {
	var p = new(ImplementsInterfacesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_implementsInterfaces

	return p
}

func (s *ImplementsInterfacesContext) GetParser() antlr.Parser { return s.parser }

func (s *ImplementsInterfacesContext) NamedType() INamedTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamedTypeContext)
}

func (s *ImplementsInterfacesContext) ImplementsInterfaces() IImplementsInterfacesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImplementsInterfacesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImplementsInterfacesContext)
}

func (s *ImplementsInterfacesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImplementsInterfacesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImplementsInterfacesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterImplementsInterfaces(s)
	}
}

func (s *ImplementsInterfacesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitImplementsInterfaces(s)
	}
}

func (p *EGraphQLParser) ImplementsInterfaces() (localctx IImplementsInterfacesContext) {
	return p.implementsInterfaces(0)
}

func (p *EGraphQLParser) implementsInterfaces(_p int) (localctx IImplementsInterfacesContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewImplementsInterfacesContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IImplementsInterfacesContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 96
	p.EnterRecursionRule(localctx, 96, EGraphQLParserRULE_implementsInterfaces, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(441)
		p.Match(EGraphQLParserT__25)
	}
	p.SetState(443)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__26 {
		{
			p.SetState(442)
			p.Match(EGraphQLParserT__26)
		}

	}
	{
		p.SetState(445)
		p.NamedType()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(452)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewImplementsInterfacesContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, EGraphQLParserRULE_implementsInterfaces)
			p.SetState(447)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(448)
				p.Match(EGraphQLParserT__26)
			}
			{
				p.SetState(449)
				p.NamedType()
			}

		}
		p.SetState(454)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())
	}

	return localctx
}

// IFieldsDefinitionContext is an interface to support dynamic dispatch.
type IFieldsDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldsDefinitionContext differentiates from other interfaces.
	IsFieldsDefinitionContext()
}

type FieldsDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsDefinitionContext() *FieldsDefinitionContext {
	var p = new(FieldsDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_fieldsDefinition
	return p
}

func (*FieldsDefinitionContext) IsFieldsDefinitionContext() {}

func NewFieldsDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsDefinitionContext {
	var p = new(FieldsDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_fieldsDefinition

	return p
}

func (s *FieldsDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsDefinitionContext) AllFieldDefinition() []IFieldDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldDefinitionContext)(nil)).Elem())
	var tst = make([]IFieldDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldDefinitionContext)
		}
	}

	return tst
}

func (s *FieldsDefinitionContext) FieldDefinition(i int) IFieldDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldDefinitionContext)
}

func (s *FieldsDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterFieldsDefinition(s)
	}
}

func (s *FieldsDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitFieldsDefinition(s)
	}
}

func (p *EGraphQLParser) FieldsDefinition() (localctx IFieldsDefinitionContext) {
	localctx = NewFieldsDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, EGraphQLParserRULE_fieldsDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(455)
		p.Match(EGraphQLParserT__3)
	}
	p.SetState(457)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-57)&-(0x1f+1)) == 0 && ((1<<uint((_la-57)))&((1<<(EGraphQLParserNAME-57))|(1<<(EGraphQLParserSTRING-57))|(1<<(EGraphQLParserBLOCK_STRING-57)))) != 0) {
		{
			p.SetState(456)
			p.FieldDefinition()
		}

		p.SetState(459)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(461)
		p.Match(EGraphQLParserT__4)
	}

	return localctx
}

// IFieldDefinitionContext is an interface to support dynamic dispatch.
type IFieldDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldDefinitionContext differentiates from other interfaces.
	IsFieldDefinitionContext()
}

type FieldDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldDefinitionContext() *FieldDefinitionContext {
	var p = new(FieldDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_fieldDefinition
	return p
}

func (*FieldDefinitionContext) IsFieldDefinitionContext() {}

func NewFieldDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldDefinitionContext {
	var p = new(FieldDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_fieldDefinition

	return p
}

func (s *FieldDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *FieldDefinitionContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *FieldDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *FieldDefinitionContext) ArgumentsDefinition() IArgumentsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsDefinitionContext)
}

func (s *FieldDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *FieldDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterFieldDefinition(s)
	}
}

func (s *FieldDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitFieldDefinition(s)
	}
}

func (p *EGraphQLParser) FieldDefinition() (localctx IFieldDefinitionContext) {
	localctx = NewFieldDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, EGraphQLParserRULE_fieldDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(464)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserSTRING || _la == EGraphQLParserBLOCK_STRING {
		{
			p.SetState(463)
			p.Description()
		}

	}
	{
		p.SetState(466)
		p.Name()
	}
	p.SetState(468)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__5 {
		{
			p.SetState(467)
			p.ArgumentsDefinition()
		}

	}
	{
		p.SetState(470)
		p.Match(EGraphQLParserT__7)
	}
	{
		p.SetState(471)
		p.Type_()
	}
	p.SetState(473)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__19 {
		{
			p.SetState(472)
			p.Directives()
		}

	}

	return localctx
}

// IArgumentsDefinitionContext is an interface to support dynamic dispatch.
type IArgumentsDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsDefinitionContext differentiates from other interfaces.
	IsArgumentsDefinitionContext()
}

type ArgumentsDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsDefinitionContext() *ArgumentsDefinitionContext {
	var p = new(ArgumentsDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_argumentsDefinition
	return p
}

func (*ArgumentsDefinitionContext) IsArgumentsDefinitionContext() {}

func NewArgumentsDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsDefinitionContext {
	var p = new(ArgumentsDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_argumentsDefinition

	return p
}

func (s *ArgumentsDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsDefinitionContext) AllInputValueDefinition() []IInputValueDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInputValueDefinitionContext)(nil)).Elem())
	var tst = make([]IInputValueDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInputValueDefinitionContext)
		}
	}

	return tst
}

func (s *ArgumentsDefinitionContext) InputValueDefinition(i int) IInputValueDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputValueDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInputValueDefinitionContext)
}

func (s *ArgumentsDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterArgumentsDefinition(s)
	}
}

func (s *ArgumentsDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitArgumentsDefinition(s)
	}
}

func (p *EGraphQLParser) ArgumentsDefinition() (localctx IArgumentsDefinitionContext) {
	localctx = NewArgumentsDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, EGraphQLParserRULE_argumentsDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(475)
		p.Match(EGraphQLParserT__5)
	}
	p.SetState(477)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-57)&-(0x1f+1)) == 0 && ((1<<uint((_la-57)))&((1<<(EGraphQLParserNAME-57))|(1<<(EGraphQLParserSTRING-57))|(1<<(EGraphQLParserBLOCK_STRING-57)))) != 0) {
		{
			p.SetState(476)
			p.InputValueDefinition()
		}

		p.SetState(479)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(481)
		p.Match(EGraphQLParserT__6)
	}

	return localctx
}

// IInputValueDefinitionContext is an interface to support dynamic dispatch.
type IInputValueDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInputValueDefinitionContext differentiates from other interfaces.
	IsInputValueDefinitionContext()
}

type InputValueDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputValueDefinitionContext() *InputValueDefinitionContext {
	var p = new(InputValueDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_inputValueDefinition
	return p
}

func (*InputValueDefinitionContext) IsInputValueDefinitionContext() {}

func NewInputValueDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputValueDefinitionContext {
	var p = new(InputValueDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_inputValueDefinition

	return p
}

func (s *InputValueDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *InputValueDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *InputValueDefinitionContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *InputValueDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *InputValueDefinitionContext) DefaultValue() IDefaultValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultValueContext)
}

func (s *InputValueDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InputValueDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputValueDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputValueDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterInputValueDefinition(s)
	}
}

func (s *InputValueDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitInputValueDefinition(s)
	}
}

func (p *EGraphQLParser) InputValueDefinition() (localctx IInputValueDefinitionContext) {
	localctx = NewInputValueDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, EGraphQLParserRULE_inputValueDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(484)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserSTRING || _la == EGraphQLParserBLOCK_STRING {
		{
			p.SetState(483)
			p.Description()
		}

	}
	{
		p.SetState(486)
		p.Name()
	}
	{
		p.SetState(487)
		p.Match(EGraphQLParserT__7)
	}
	{
		p.SetState(488)
		p.Type_()
	}
	p.SetState(490)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__17 {
		{
			p.SetState(489)
			p.DefaultValue()
		}

	}
	p.SetState(493)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__19 {
		{
			p.SetState(492)
			p.Directives()
		}

	}

	return localctx
}

// IObjectTypeExtensionContext is an interface to support dynamic dispatch.
type IObjectTypeExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectTypeExtensionContext differentiates from other interfaces.
	IsObjectTypeExtensionContext()
}

type ObjectTypeExtensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectTypeExtensionContext() *ObjectTypeExtensionContext {
	var p = new(ObjectTypeExtensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_objectTypeExtension
	return p
}

func (*ObjectTypeExtensionContext) IsObjectTypeExtensionContext() {}

func NewObjectTypeExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectTypeExtensionContext {
	var p = new(ObjectTypeExtensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_objectTypeExtension

	return p
}

func (s *ObjectTypeExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectTypeExtensionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ObjectTypeExtensionContext) FieldsDefinition() IFieldsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldsDefinitionContext)
}

func (s *ObjectTypeExtensionContext) ImplementsInterfaces() IImplementsInterfacesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImplementsInterfacesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImplementsInterfacesContext)
}

func (s *ObjectTypeExtensionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *ObjectTypeExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectTypeExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectTypeExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterObjectTypeExtension(s)
	}
}

func (s *ObjectTypeExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitObjectTypeExtension(s)
	}
}

func (p *EGraphQLParser) ObjectTypeExtension() (localctx IObjectTypeExtensionContext) {
	localctx = NewObjectTypeExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, EGraphQLParserRULE_objectTypeExtension)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(519)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 58, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(495)
			p.Match(EGraphQLParserT__21)
		}
		{
			p.SetState(496)
			p.Match(EGraphQLParserT__24)
		}
		{
			p.SetState(497)
			p.Name()
		}
		p.SetState(499)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EGraphQLParserT__25 {
			{
				p.SetState(498)
				p.implementsInterfaces(0)
			}

		}
		p.SetState(502)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EGraphQLParserT__19 {
			{
				p.SetState(501)
				p.Directives()
			}

		}
		{
			p.SetState(504)
			p.FieldsDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(506)
			p.Match(EGraphQLParserT__21)
		}
		{
			p.SetState(507)
			p.Match(EGraphQLParserT__24)
		}
		{
			p.SetState(508)
			p.Name()
		}
		p.SetState(510)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EGraphQLParserT__25 {
			{
				p.SetState(509)
				p.implementsInterfaces(0)
			}

		}
		{
			p.SetState(512)
			p.Directives()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(514)
			p.Match(EGraphQLParserT__21)
		}
		{
			p.SetState(515)
			p.Match(EGraphQLParserT__24)
		}
		{
			p.SetState(516)
			p.Name()
		}
		{
			p.SetState(517)
			p.implementsInterfaces(0)
		}

	}

	return localctx
}

// IInterfaceTypeDefinitionContext is an interface to support dynamic dispatch.
type IInterfaceTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInterfaceTypeDefinitionContext differentiates from other interfaces.
	IsInterfaceTypeDefinitionContext()
}

type InterfaceTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterfaceTypeDefinitionContext() *InterfaceTypeDefinitionContext {
	var p = new(InterfaceTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_interfaceTypeDefinition
	return p
}

func (*InterfaceTypeDefinitionContext) IsInterfaceTypeDefinitionContext() {}

func NewInterfaceTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceTypeDefinitionContext {
	var p = new(InterfaceTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_interfaceTypeDefinition

	return p
}

func (s *InterfaceTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceTypeDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *InterfaceTypeDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *InterfaceTypeDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InterfaceTypeDefinitionContext) FieldsDefinition() IFieldsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldsDefinitionContext)
}

func (s *InterfaceTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterfaceTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterInterfaceTypeDefinition(s)
	}
}

func (s *InterfaceTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitInterfaceTypeDefinition(s)
	}
}

func (p *EGraphQLParser) InterfaceTypeDefinition() (localctx IInterfaceTypeDefinitionContext) {
	localctx = NewInterfaceTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, EGraphQLParserRULE_interfaceTypeDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(522)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserSTRING || _la == EGraphQLParserBLOCK_STRING {
		{
			p.SetState(521)
			p.Description()
		}

	}
	{
		p.SetState(524)
		p.Match(EGraphQLParserT__27)
	}
	{
		p.SetState(525)
		p.Name()
	}
	p.SetState(527)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__19 {
		{
			p.SetState(526)
			p.Directives()
		}

	}
	p.SetState(530)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 61, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(529)
			p.FieldsDefinition()
		}

	}

	return localctx
}

// IInterfaceTypeExtensionContext is an interface to support dynamic dispatch.
type IInterfaceTypeExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInterfaceTypeExtensionContext differentiates from other interfaces.
	IsInterfaceTypeExtensionContext()
}

type InterfaceTypeExtensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterfaceTypeExtensionContext() *InterfaceTypeExtensionContext {
	var p = new(InterfaceTypeExtensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_interfaceTypeExtension
	return p
}

func (*InterfaceTypeExtensionContext) IsInterfaceTypeExtensionContext() {}

func NewInterfaceTypeExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceTypeExtensionContext {
	var p = new(InterfaceTypeExtensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_interfaceTypeExtension

	return p
}

func (s *InterfaceTypeExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceTypeExtensionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *InterfaceTypeExtensionContext) FieldsDefinition() IFieldsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldsDefinitionContext)
}

func (s *InterfaceTypeExtensionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InterfaceTypeExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceTypeExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterfaceTypeExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterInterfaceTypeExtension(s)
	}
}

func (s *InterfaceTypeExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitInterfaceTypeExtension(s)
	}
}

func (p *EGraphQLParser) InterfaceTypeExtension() (localctx IInterfaceTypeExtensionContext) {
	localctx = NewInterfaceTypeExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, EGraphQLParserRULE_interfaceTypeExtension)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(545)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 63, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(532)
			p.Match(EGraphQLParserT__21)
		}
		{
			p.SetState(533)
			p.Match(EGraphQLParserT__27)
		}
		{
			p.SetState(534)
			p.Name()
		}
		p.SetState(536)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EGraphQLParserT__19 {
			{
				p.SetState(535)
				p.Directives()
			}

		}
		{
			p.SetState(538)
			p.FieldsDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(540)
			p.Match(EGraphQLParserT__21)
		}
		{
			p.SetState(541)
			p.Match(EGraphQLParserT__27)
		}
		{
			p.SetState(542)
			p.Name()
		}
		{
			p.SetState(543)
			p.Directives()
		}

	}

	return localctx
}

// IUnionTypeDefinitionContext is an interface to support dynamic dispatch.
type IUnionTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionTypeDefinitionContext differentiates from other interfaces.
	IsUnionTypeDefinitionContext()
}

type UnionTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionTypeDefinitionContext() *UnionTypeDefinitionContext {
	var p = new(UnionTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_unionTypeDefinition
	return p
}

func (*UnionTypeDefinitionContext) IsUnionTypeDefinitionContext() {}

func NewUnionTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionTypeDefinitionContext {
	var p = new(UnionTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_unionTypeDefinition

	return p
}

func (s *UnionTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionTypeDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *UnionTypeDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *UnionTypeDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *UnionTypeDefinitionContext) UnionMemberTypes() IUnionMemberTypesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionMemberTypesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionMemberTypesContext)
}

func (s *UnionTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterUnionTypeDefinition(s)
	}
}

func (s *UnionTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitUnionTypeDefinition(s)
	}
}

func (p *EGraphQLParser) UnionTypeDefinition() (localctx IUnionTypeDefinitionContext) {
	localctx = NewUnionTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, EGraphQLParserRULE_unionTypeDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(548)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserSTRING || _la == EGraphQLParserBLOCK_STRING {
		{
			p.SetState(547)
			p.Description()
		}

	}
	{
		p.SetState(550)
		p.Match(EGraphQLParserT__28)
	}
	{
		p.SetState(551)
		p.Name()
	}
	p.SetState(553)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__19 {
		{
			p.SetState(552)
			p.Directives()
		}

	}
	p.SetState(556)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__17 {
		{
			p.SetState(555)
			p.UnionMemberTypes()
		}

	}

	return localctx
}

// IUnionMemberTypesContext is an interface to support dynamic dispatch.
type IUnionMemberTypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionMemberTypesContext differentiates from other interfaces.
	IsUnionMemberTypesContext()
}

type UnionMemberTypesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionMemberTypesContext() *UnionMemberTypesContext {
	var p = new(UnionMemberTypesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_unionMemberTypes
	return p
}

func (*UnionMemberTypesContext) IsUnionMemberTypesContext() {}

func NewUnionMemberTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionMemberTypesContext {
	var p = new(UnionMemberTypesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_unionMemberTypes

	return p
}

func (s *UnionMemberTypesContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionMemberTypesContext) AllNamedType() []INamedTypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INamedTypeContext)(nil)).Elem())
	var tst = make([]INamedTypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INamedTypeContext)
		}
	}

	return tst
}

func (s *UnionMemberTypesContext) NamedType(i int) INamedTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedTypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INamedTypeContext)
}

func (s *UnionMemberTypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionMemberTypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionMemberTypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterUnionMemberTypes(s)
	}
}

func (s *UnionMemberTypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitUnionMemberTypes(s)
	}
}

func (p *EGraphQLParser) UnionMemberTypes() (localctx IUnionMemberTypesContext) {
	localctx = NewUnionMemberTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, EGraphQLParserRULE_unionMemberTypes)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(558)
		p.Match(EGraphQLParserT__17)
	}
	p.SetState(560)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__29 {
		{
			p.SetState(559)
			p.Match(EGraphQLParserT__29)
		}

	}
	{
		p.SetState(562)
		p.NamedType()
	}
	p.SetState(567)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == EGraphQLParserT__29 {
		{
			p.SetState(563)
			p.Match(EGraphQLParserT__29)
		}
		{
			p.SetState(564)
			p.NamedType()
		}

		p.SetState(569)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IUnionTypeExtensionContext is an interface to support dynamic dispatch.
type IUnionTypeExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionTypeExtensionContext differentiates from other interfaces.
	IsUnionTypeExtensionContext()
}

type UnionTypeExtensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionTypeExtensionContext() *UnionTypeExtensionContext {
	var p = new(UnionTypeExtensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_unionTypeExtension
	return p
}

func (*UnionTypeExtensionContext) IsUnionTypeExtensionContext() {}

func NewUnionTypeExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionTypeExtensionContext {
	var p = new(UnionTypeExtensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_unionTypeExtension

	return p
}

func (s *UnionTypeExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionTypeExtensionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *UnionTypeExtensionContext) UnionMemberTypes() IUnionMemberTypesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionMemberTypesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionMemberTypesContext)
}

func (s *UnionTypeExtensionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *UnionTypeExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionTypeExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionTypeExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterUnionTypeExtension(s)
	}
}

func (s *UnionTypeExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitUnionTypeExtension(s)
	}
}

func (p *EGraphQLParser) UnionTypeExtension() (localctx IUnionTypeExtensionContext) {
	localctx = NewUnionTypeExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, EGraphQLParserRULE_unionTypeExtension)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(583)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 70, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(570)
			p.Match(EGraphQLParserT__21)
		}
		{
			p.SetState(571)
			p.Match(EGraphQLParserT__28)
		}
		{
			p.SetState(572)
			p.Name()
		}
		p.SetState(574)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EGraphQLParserT__19 {
			{
				p.SetState(573)
				p.Directives()
			}

		}
		{
			p.SetState(576)
			p.UnionMemberTypes()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(578)
			p.Match(EGraphQLParserT__21)
		}
		{
			p.SetState(579)
			p.Match(EGraphQLParserT__28)
		}
		{
			p.SetState(580)
			p.Name()
		}
		{
			p.SetState(581)
			p.Directives()
		}

	}

	return localctx
}

// IEnumTypeDefinitionContext is an interface to support dynamic dispatch.
type IEnumTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumTypeDefinitionContext differentiates from other interfaces.
	IsEnumTypeDefinitionContext()
}

type EnumTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumTypeDefinitionContext() *EnumTypeDefinitionContext {
	var p = new(EnumTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_enumTypeDefinition
	return p
}

func (*EnumTypeDefinitionContext) IsEnumTypeDefinitionContext() {}

func NewEnumTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumTypeDefinitionContext {
	var p = new(EnumTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_enumTypeDefinition

	return p
}

func (s *EnumTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumTypeDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *EnumTypeDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *EnumTypeDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *EnumTypeDefinitionContext) EnumValuesDefinition() IEnumValuesDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValuesDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumValuesDefinitionContext)
}

func (s *EnumTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterEnumTypeDefinition(s)
	}
}

func (s *EnumTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitEnumTypeDefinition(s)
	}
}

func (p *EGraphQLParser) EnumTypeDefinition() (localctx IEnumTypeDefinitionContext) {
	localctx = NewEnumTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, EGraphQLParserRULE_enumTypeDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(586)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserSTRING || _la == EGraphQLParserBLOCK_STRING {
		{
			p.SetState(585)
			p.Description()
		}

	}
	{
		p.SetState(588)
		p.Match(EGraphQLParserT__30)
	}
	{
		p.SetState(589)
		p.Name()
	}
	p.SetState(591)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__19 {
		{
			p.SetState(590)
			p.Directives()
		}

	}
	p.SetState(594)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 73, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(593)
			p.EnumValuesDefinition()
		}

	}

	return localctx
}

// IEnumValuesDefinitionContext is an interface to support dynamic dispatch.
type IEnumValuesDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumValuesDefinitionContext differentiates from other interfaces.
	IsEnumValuesDefinitionContext()
}

type EnumValuesDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValuesDefinitionContext() *EnumValuesDefinitionContext {
	var p = new(EnumValuesDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_enumValuesDefinition
	return p
}

func (*EnumValuesDefinitionContext) IsEnumValuesDefinitionContext() {}

func NewEnumValuesDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValuesDefinitionContext {
	var p = new(EnumValuesDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_enumValuesDefinition

	return p
}

func (s *EnumValuesDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValuesDefinitionContext) AllEnumValueDefinition() []IEnumValueDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumValueDefinitionContext)(nil)).Elem())
	var tst = make([]IEnumValueDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumValueDefinitionContext)
		}
	}

	return tst
}

func (s *EnumValuesDefinitionContext) EnumValueDefinition(i int) IEnumValueDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValueDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumValueDefinitionContext)
}

func (s *EnumValuesDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValuesDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValuesDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterEnumValuesDefinition(s)
	}
}

func (s *EnumValuesDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitEnumValuesDefinition(s)
	}
}

func (p *EGraphQLParser) EnumValuesDefinition() (localctx IEnumValuesDefinitionContext) {
	localctx = NewEnumValuesDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, EGraphQLParserRULE_enumValuesDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(596)
		p.Match(EGraphQLParserT__3)
	}
	p.SetState(598)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-57)&-(0x1f+1)) == 0 && ((1<<uint((_la-57)))&((1<<(EGraphQLParserNAME-57))|(1<<(EGraphQLParserSTRING-57))|(1<<(EGraphQLParserBLOCK_STRING-57)))) != 0) {
		{
			p.SetState(597)
			p.EnumValueDefinition()
		}

		p.SetState(600)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(602)
		p.Match(EGraphQLParserT__4)
	}

	return localctx
}

// IEnumValueDefinitionContext is an interface to support dynamic dispatch.
type IEnumValueDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumValueDefinitionContext differentiates from other interfaces.
	IsEnumValueDefinitionContext()
}

type EnumValueDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueDefinitionContext() *EnumValueDefinitionContext {
	var p = new(EnumValueDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_enumValueDefinition
	return p
}

func (*EnumValueDefinitionContext) IsEnumValueDefinitionContext() {}

func NewEnumValueDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueDefinitionContext {
	var p = new(EnumValueDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_enumValueDefinition

	return p
}

func (s *EnumValueDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueDefinitionContext) EnumValue() IEnumValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumValueContext)
}

func (s *EnumValueDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *EnumValueDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *EnumValueDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterEnumValueDefinition(s)
	}
}

func (s *EnumValueDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitEnumValueDefinition(s)
	}
}

func (p *EGraphQLParser) EnumValueDefinition() (localctx IEnumValueDefinitionContext) {
	localctx = NewEnumValueDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, EGraphQLParserRULE_enumValueDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(605)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserSTRING || _la == EGraphQLParserBLOCK_STRING {
		{
			p.SetState(604)
			p.Description()
		}

	}
	{
		p.SetState(607)
		p.EnumValue()
	}
	p.SetState(609)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__19 {
		{
			p.SetState(608)
			p.Directives()
		}

	}

	return localctx
}

// IEnumTypeExtensionContext is an interface to support dynamic dispatch.
type IEnumTypeExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumTypeExtensionContext differentiates from other interfaces.
	IsEnumTypeExtensionContext()
}

type EnumTypeExtensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumTypeExtensionContext() *EnumTypeExtensionContext {
	var p = new(EnumTypeExtensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_enumTypeExtension
	return p
}

func (*EnumTypeExtensionContext) IsEnumTypeExtensionContext() {}

func NewEnumTypeExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumTypeExtensionContext {
	var p = new(EnumTypeExtensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_enumTypeExtension

	return p
}

func (s *EnumTypeExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumTypeExtensionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *EnumTypeExtensionContext) EnumValuesDefinition() IEnumValuesDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValuesDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumValuesDefinitionContext)
}

func (s *EnumTypeExtensionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *EnumTypeExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumTypeExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumTypeExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterEnumTypeExtension(s)
	}
}

func (s *EnumTypeExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitEnumTypeExtension(s)
	}
}

func (p *EGraphQLParser) EnumTypeExtension() (localctx IEnumTypeExtensionContext) {
	localctx = NewEnumTypeExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, EGraphQLParserRULE_enumTypeExtension)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(624)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 78, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(611)
			p.Match(EGraphQLParserT__21)
		}
		{
			p.SetState(612)
			p.Match(EGraphQLParserT__30)
		}
		{
			p.SetState(613)
			p.Name()
		}
		p.SetState(615)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EGraphQLParserT__19 {
			{
				p.SetState(614)
				p.Directives()
			}

		}
		{
			p.SetState(617)
			p.EnumValuesDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(619)
			p.Match(EGraphQLParserT__21)
		}
		{
			p.SetState(620)
			p.Match(EGraphQLParserT__30)
		}
		{
			p.SetState(621)
			p.Name()
		}
		{
			p.SetState(622)
			p.Directives()
		}

	}

	return localctx
}

// IInputObjectTypeDefinitionContext is an interface to support dynamic dispatch.
type IInputObjectTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInputObjectTypeDefinitionContext differentiates from other interfaces.
	IsInputObjectTypeDefinitionContext()
}

type InputObjectTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputObjectTypeDefinitionContext() *InputObjectTypeDefinitionContext {
	var p = new(InputObjectTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_inputObjectTypeDefinition
	return p
}

func (*InputObjectTypeDefinitionContext) IsInputObjectTypeDefinitionContext() {}

func NewInputObjectTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputObjectTypeDefinitionContext {
	var p = new(InputObjectTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_inputObjectTypeDefinition

	return p
}

func (s *InputObjectTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *InputObjectTypeDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *InputObjectTypeDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *InputObjectTypeDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InputObjectTypeDefinitionContext) InputFieldsDefinition() IInputFieldsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputFieldsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInputFieldsDefinitionContext)
}

func (s *InputObjectTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputObjectTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputObjectTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterInputObjectTypeDefinition(s)
	}
}

func (s *InputObjectTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitInputObjectTypeDefinition(s)
	}
}

func (p *EGraphQLParser) InputObjectTypeDefinition() (localctx IInputObjectTypeDefinitionContext) {
	localctx = NewInputObjectTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, EGraphQLParserRULE_inputObjectTypeDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(627)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserSTRING || _la == EGraphQLParserBLOCK_STRING {
		{
			p.SetState(626)
			p.Description()
		}

	}
	{
		p.SetState(629)
		p.Match(EGraphQLParserT__31)
	}
	{
		p.SetState(630)
		p.Name()
	}
	p.SetState(632)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__19 {
		{
			p.SetState(631)
			p.Directives()
		}

	}
	p.SetState(635)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 81, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(634)
			p.InputFieldsDefinition()
		}

	}

	return localctx
}

// IInputFieldsDefinitionContext is an interface to support dynamic dispatch.
type IInputFieldsDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInputFieldsDefinitionContext differentiates from other interfaces.
	IsInputFieldsDefinitionContext()
}

type InputFieldsDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputFieldsDefinitionContext() *InputFieldsDefinitionContext {
	var p = new(InputFieldsDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_inputFieldsDefinition
	return p
}

func (*InputFieldsDefinitionContext) IsInputFieldsDefinitionContext() {}

func NewInputFieldsDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputFieldsDefinitionContext {
	var p = new(InputFieldsDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_inputFieldsDefinition

	return p
}

func (s *InputFieldsDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *InputFieldsDefinitionContext) AllInputValueDefinition() []IInputValueDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInputValueDefinitionContext)(nil)).Elem())
	var tst = make([]IInputValueDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInputValueDefinitionContext)
		}
	}

	return tst
}

func (s *InputFieldsDefinitionContext) InputValueDefinition(i int) IInputValueDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputValueDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInputValueDefinitionContext)
}

func (s *InputFieldsDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputFieldsDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputFieldsDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterInputFieldsDefinition(s)
	}
}

func (s *InputFieldsDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitInputFieldsDefinition(s)
	}
}

func (p *EGraphQLParser) InputFieldsDefinition() (localctx IInputFieldsDefinitionContext) {
	localctx = NewInputFieldsDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, EGraphQLParserRULE_inputFieldsDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(637)
		p.Match(EGraphQLParserT__3)
	}
	p.SetState(639)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-57)&-(0x1f+1)) == 0 && ((1<<uint((_la-57)))&((1<<(EGraphQLParserNAME-57))|(1<<(EGraphQLParserSTRING-57))|(1<<(EGraphQLParserBLOCK_STRING-57)))) != 0) {
		{
			p.SetState(638)
			p.InputValueDefinition()
		}

		p.SetState(641)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(643)
		p.Match(EGraphQLParserT__4)
	}

	return localctx
}

// IInputObjectTypeExtensionContext is an interface to support dynamic dispatch.
type IInputObjectTypeExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInputObjectTypeExtensionContext differentiates from other interfaces.
	IsInputObjectTypeExtensionContext()
}

type InputObjectTypeExtensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputObjectTypeExtensionContext() *InputObjectTypeExtensionContext {
	var p = new(InputObjectTypeExtensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_inputObjectTypeExtension
	return p
}

func (*InputObjectTypeExtensionContext) IsInputObjectTypeExtensionContext() {}

func NewInputObjectTypeExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputObjectTypeExtensionContext {
	var p = new(InputObjectTypeExtensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_inputObjectTypeExtension

	return p
}

func (s *InputObjectTypeExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *InputObjectTypeExtensionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *InputObjectTypeExtensionContext) InputFieldsDefinition() IInputFieldsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputFieldsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInputFieldsDefinitionContext)
}

func (s *InputObjectTypeExtensionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InputObjectTypeExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputObjectTypeExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputObjectTypeExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterInputObjectTypeExtension(s)
	}
}

func (s *InputObjectTypeExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitInputObjectTypeExtension(s)
	}
}

func (p *EGraphQLParser) InputObjectTypeExtension() (localctx IInputObjectTypeExtensionContext) {
	localctx = NewInputObjectTypeExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, EGraphQLParserRULE_inputObjectTypeExtension)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(658)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 84, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(645)
			p.Match(EGraphQLParserT__21)
		}
		{
			p.SetState(646)
			p.Match(EGraphQLParserT__31)
		}
		{
			p.SetState(647)
			p.Name()
		}
		p.SetState(649)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EGraphQLParserT__19 {
			{
				p.SetState(648)
				p.Directives()
			}

		}
		{
			p.SetState(651)
			p.InputFieldsDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(653)
			p.Match(EGraphQLParserT__21)
		}
		{
			p.SetState(654)
			p.Match(EGraphQLParserT__31)
		}
		{
			p.SetState(655)
			p.Name()
		}
		{
			p.SetState(656)
			p.Directives()
		}

	}

	return localctx
}

// IDirectiveDefinitionContext is an interface to support dynamic dispatch.
type IDirectiveDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectiveDefinitionContext differentiates from other interfaces.
	IsDirectiveDefinitionContext()
}

type DirectiveDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveDefinitionContext() *DirectiveDefinitionContext {
	var p = new(DirectiveDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_directiveDefinition
	return p
}

func (*DirectiveDefinitionContext) IsDirectiveDefinitionContext() {}

func NewDirectiveDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveDefinitionContext {
	var p = new(DirectiveDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_directiveDefinition

	return p
}

func (s *DirectiveDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *DirectiveDefinitionContext) DirectiveLocations() IDirectiveLocationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveLocationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectiveLocationsContext)
}

func (s *DirectiveDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *DirectiveDefinitionContext) ArgumentsDefinition() IArgumentsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsDefinitionContext)
}

func (s *DirectiveDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterDirectiveDefinition(s)
	}
}

func (s *DirectiveDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitDirectiveDefinition(s)
	}
}

func (p *EGraphQLParser) DirectiveDefinition() (localctx IDirectiveDefinitionContext) {
	localctx = NewDirectiveDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, EGraphQLParserRULE_directiveDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(661)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserSTRING || _la == EGraphQLParserBLOCK_STRING {
		{
			p.SetState(660)
			p.Description()
		}

	}
	{
		p.SetState(663)
		p.Match(EGraphQLParserT__32)
	}
	{
		p.SetState(664)
		p.Match(EGraphQLParserT__19)
	}
	{
		p.SetState(665)
		p.Name()
	}
	p.SetState(667)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__5 {
		{
			p.SetState(666)
			p.ArgumentsDefinition()
		}

	}
	{
		p.SetState(669)
		p.Match(EGraphQLParserT__10)
	}
	{
		p.SetState(670)
		p.DirectiveLocations()
	}

	return localctx
}

// IDirectiveLocationsContext is an interface to support dynamic dispatch.
type IDirectiveLocationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectiveLocationsContext differentiates from other interfaces.
	IsDirectiveLocationsContext()
}

type DirectiveLocationsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveLocationsContext() *DirectiveLocationsContext {
	var p = new(DirectiveLocationsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_directiveLocations
	return p
}

func (*DirectiveLocationsContext) IsDirectiveLocationsContext() {}

func NewDirectiveLocationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveLocationsContext {
	var p = new(DirectiveLocationsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_directiveLocations

	return p
}

func (s *DirectiveLocationsContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveLocationsContext) AllDirectiveLocation() []IDirectiveLocationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDirectiveLocationContext)(nil)).Elem())
	var tst = make([]IDirectiveLocationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDirectiveLocationContext)
		}
	}

	return tst
}

func (s *DirectiveLocationsContext) DirectiveLocation(i int) IDirectiveLocationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveLocationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDirectiveLocationContext)
}

func (s *DirectiveLocationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveLocationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveLocationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterDirectiveLocations(s)
	}
}

func (s *DirectiveLocationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitDirectiveLocations(s)
	}
}

func (p *EGraphQLParser) DirectiveLocations() (localctx IDirectiveLocationsContext) {
	localctx = NewDirectiveLocationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 134, EGraphQLParserRULE_directiveLocations)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(672)
		p.DirectiveLocation()
	}
	p.SetState(677)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == EGraphQLParserT__29 {
		{
			p.SetState(673)
			p.Match(EGraphQLParserT__29)
		}
		{
			p.SetState(674)
			p.DirectiveLocation()
		}

		p.SetState(679)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDirectiveLocationContext is an interface to support dynamic dispatch.
type IDirectiveLocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectiveLocationContext differentiates from other interfaces.
	IsDirectiveLocationContext()
}

type DirectiveLocationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveLocationContext() *DirectiveLocationContext {
	var p = new(DirectiveLocationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_directiveLocation
	return p
}

func (*DirectiveLocationContext) IsDirectiveLocationContext() {}

func NewDirectiveLocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveLocationContext {
	var p = new(DirectiveLocationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_directiveLocation

	return p
}

func (s *DirectiveLocationContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveLocationContext) ExecutableDirectiveLocation() IExecutableDirectiveLocationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExecutableDirectiveLocationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExecutableDirectiveLocationContext)
}

func (s *DirectiveLocationContext) TypeSystemDirectiveLocation() ITypeSystemDirectiveLocationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSystemDirectiveLocationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSystemDirectiveLocationContext)
}

func (s *DirectiveLocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveLocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveLocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterDirectiveLocation(s)
	}
}

func (s *DirectiveLocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitDirectiveLocation(s)
	}
}

func (p *EGraphQLParser) DirectiveLocation() (localctx IDirectiveLocationContext) {
	localctx = NewDirectiveLocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 136, EGraphQLParserRULE_directiveLocation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(682)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case EGraphQLParserT__33, EGraphQLParserT__34, EGraphQLParserT__35, EGraphQLParserT__36, EGraphQLParserT__37, EGraphQLParserT__38, EGraphQLParserT__39:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(680)
			p.ExecutableDirectiveLocation()
		}

	case EGraphQLParserT__40, EGraphQLParserT__41, EGraphQLParserT__42, EGraphQLParserT__43, EGraphQLParserT__44, EGraphQLParserT__45, EGraphQLParserT__46, EGraphQLParserT__47, EGraphQLParserT__48, EGraphQLParserT__49, EGraphQLParserT__50:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(681)
			p.TypeSystemDirectiveLocation()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExecutableDirectiveLocationContext is an interface to support dynamic dispatch.
type IExecutableDirectiveLocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExecutableDirectiveLocationContext differentiates from other interfaces.
	IsExecutableDirectiveLocationContext()
}

type ExecutableDirectiveLocationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExecutableDirectiveLocationContext() *ExecutableDirectiveLocationContext {
	var p = new(ExecutableDirectiveLocationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_executableDirectiveLocation
	return p
}

func (*ExecutableDirectiveLocationContext) IsExecutableDirectiveLocationContext() {}

func NewExecutableDirectiveLocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExecutableDirectiveLocationContext {
	var p = new(ExecutableDirectiveLocationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_executableDirectiveLocation

	return p
}

func (s *ExecutableDirectiveLocationContext) GetParser() antlr.Parser { return s.parser }
func (s *ExecutableDirectiveLocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExecutableDirectiveLocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExecutableDirectiveLocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterExecutableDirectiveLocation(s)
	}
}

func (s *ExecutableDirectiveLocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitExecutableDirectiveLocation(s)
	}
}

func (p *EGraphQLParser) ExecutableDirectiveLocation() (localctx IExecutableDirectiveLocationContext) {
	localctx = NewExecutableDirectiveLocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 138, EGraphQLParserRULE_executableDirectiveLocation)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(684)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(EGraphQLParserT__33-34))|(1<<(EGraphQLParserT__34-34))|(1<<(EGraphQLParserT__35-34))|(1<<(EGraphQLParserT__36-34))|(1<<(EGraphQLParserT__37-34))|(1<<(EGraphQLParserT__38-34))|(1<<(EGraphQLParserT__39-34)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITypeSystemDirectiveLocationContext is an interface to support dynamic dispatch.
type ITypeSystemDirectiveLocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSystemDirectiveLocationContext differentiates from other interfaces.
	IsTypeSystemDirectiveLocationContext()
}

type TypeSystemDirectiveLocationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSystemDirectiveLocationContext() *TypeSystemDirectiveLocationContext {
	var p = new(TypeSystemDirectiveLocationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_typeSystemDirectiveLocation
	return p
}

func (*TypeSystemDirectiveLocationContext) IsTypeSystemDirectiveLocationContext() {}

func NewTypeSystemDirectiveLocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSystemDirectiveLocationContext {
	var p = new(TypeSystemDirectiveLocationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_typeSystemDirectiveLocation

	return p
}

func (s *TypeSystemDirectiveLocationContext) GetParser() antlr.Parser { return s.parser }
func (s *TypeSystemDirectiveLocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSystemDirectiveLocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSystemDirectiveLocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterTypeSystemDirectiveLocation(s)
	}
}

func (s *TypeSystemDirectiveLocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitTypeSystemDirectiveLocation(s)
	}
}

func (p *EGraphQLParser) TypeSystemDirectiveLocation() (localctx ITypeSystemDirectiveLocationContext) {
	localctx = NewTypeSystemDirectiveLocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 140, EGraphQLParserRULE_typeSystemDirectiveLocation)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(686)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(EGraphQLParserT__40-41))|(1<<(EGraphQLParserT__41-41))|(1<<(EGraphQLParserT__42-41))|(1<<(EGraphQLParserT__43-41))|(1<<(EGraphQLParserT__44-41))|(1<<(EGraphQLParserT__45-41))|(1<<(EGraphQLParserT__46-41))|(1<<(EGraphQLParserT__47-41))|(1<<(EGraphQLParserT__48-41))|(1<<(EGraphQLParserT__49-41))|(1<<(EGraphQLParserT__50-41)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) NAME() antlr.TerminalNode {
	return s.GetToken(EGraphQLParserNAME, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *EGraphQLParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 142, EGraphQLParserRULE_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(688)
		p.Match(EGraphQLParserNAME)
	}

	return localctx
}

// ITemplateParametersDefinitionContext is an interface to support dynamic dispatch.
type ITemplateParametersDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTemplateParametersDefinitionContext differentiates from other interfaces.
	IsTemplateParametersDefinitionContext()
}

type TemplateParametersDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemplateParametersDefinitionContext() *TemplateParametersDefinitionContext {
	var p = new(TemplateParametersDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_templateParametersDefinition
	return p
}

func (*TemplateParametersDefinitionContext) IsTemplateParametersDefinitionContext() {}

func NewTemplateParametersDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemplateParametersDefinitionContext {
	var p = new(TemplateParametersDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_templateParametersDefinition

	return p
}

func (s *TemplateParametersDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *TemplateParametersDefinitionContext) AllName() []INameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameContext)(nil)).Elem())
	var tst = make([]INameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameContext)
		}
	}

	return tst
}

func (s *TemplateParametersDefinitionContext) Name(i int) INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *TemplateParametersDefinitionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(EGraphQLParserCOMMA)
}

func (s *TemplateParametersDefinitionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(EGraphQLParserCOMMA, i)
}

func (s *TemplateParametersDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemplateParametersDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemplateParametersDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterTemplateParametersDefinition(s)
	}
}

func (s *TemplateParametersDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitTemplateParametersDefinition(s)
	}
}

func (p *EGraphQLParser) TemplateParametersDefinition() (localctx ITemplateParametersDefinitionContext) {
	localctx = NewTemplateParametersDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 144, EGraphQLParserRULE_templateParametersDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(690)
		p.Match(EGraphQLParserT__51)
	}
	{
		p.SetState(691)
		p.Name()
	}
	p.SetState(696)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == EGraphQLParserCOMMA {
		{
			p.SetState(692)
			p.Match(EGraphQLParserCOMMA)
		}
		{
			p.SetState(693)
			p.Name()
		}

		p.SetState(698)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(699)
		p.Match(EGraphQLParserT__52)
	}

	return localctx
}

// ITemplateTypeDefinitionContext is an interface to support dynamic dispatch.
type ITemplateTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTemplateTypeDefinitionContext differentiates from other interfaces.
	IsTemplateTypeDefinitionContext()
}

type TemplateTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemplateTypeDefinitionContext() *TemplateTypeDefinitionContext {
	var p = new(TemplateTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_templateTypeDefinition
	return p
}

func (*TemplateTypeDefinitionContext) IsTemplateTypeDefinitionContext() {}

func NewTemplateTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemplateTypeDefinitionContext {
	var p = new(TemplateTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_templateTypeDefinition

	return p
}

func (s *TemplateTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *TemplateTypeDefinitionContext) TempletableTypeDefinition() ITempletableTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITempletableTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITempletableTypeDefinitionContext)
}

func (s *TemplateTypeDefinitionContext) TemplateParametersDefinition() ITemplateParametersDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemplateParametersDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITemplateParametersDefinitionContext)
}

func (s *TemplateTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemplateTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemplateTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterTemplateTypeDefinition(s)
	}
}

func (s *TemplateTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitTemplateTypeDefinition(s)
	}
}

func (p *EGraphQLParser) TemplateTypeDefinition() (localctx ITemplateTypeDefinitionContext) {
	localctx = NewTemplateTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 146, EGraphQLParserRULE_templateTypeDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(701)
		p.Match(EGraphQLParserT__53)
	}
	p.SetState(703)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__51 {
		{
			p.SetState(702)
			p.TemplateParametersDefinition()
		}

	}
	{
		p.SetState(705)
		p.TempletableTypeDefinition()
	}

	return localctx
}

// ITemplateImplementedTypeDefinitionContext is an interface to support dynamic dispatch.
type ITemplateImplementedTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTemplateImplementedTypeDefinitionContext differentiates from other interfaces.
	IsTemplateImplementedTypeDefinitionContext()
}

type TemplateImplementedTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemplateImplementedTypeDefinitionContext() *TemplateImplementedTypeDefinitionContext {
	var p = new(TemplateImplementedTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EGraphQLParserRULE_templateImplementedTypeDefinition
	return p
}

func (*TemplateImplementedTypeDefinitionContext) IsTemplateImplementedTypeDefinitionContext() {}

func NewTemplateImplementedTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemplateImplementedTypeDefinitionContext {
	var p = new(TemplateImplementedTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EGraphQLParserRULE_templateImplementedTypeDefinition

	return p
}

func (s *TemplateImplementedTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *TemplateImplementedTypeDefinitionContext) AllName() []INameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameContext)(nil)).Elem())
	var tst = make([]INameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameContext)
		}
	}

	return tst
}

func (s *TemplateImplementedTypeDefinitionContext) Name(i int) INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *TemplateImplementedTypeDefinitionContext) TempletableTypeDefinition() ITempletableTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITempletableTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITempletableTypeDefinitionContext)
}

func (s *TemplateImplementedTypeDefinitionContext) AllTemplateParametersDefinition() []ITemplateParametersDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITemplateParametersDefinitionContext)(nil)).Elem())
	var tst = make([]ITemplateParametersDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITemplateParametersDefinitionContext)
		}
	}

	return tst
}

func (s *TemplateImplementedTypeDefinitionContext) TemplateParametersDefinition(i int) ITemplateParametersDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemplateParametersDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITemplateParametersDefinitionContext)
}

func (s *TemplateImplementedTypeDefinitionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(EGraphQLParserCOMMA)
}

func (s *TemplateImplementedTypeDefinitionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(EGraphQLParserCOMMA, i)
}

func (s *TemplateImplementedTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemplateImplementedTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemplateImplementedTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.EnterTemplateImplementedTypeDefinition(s)
	}
}

func (s *TemplateImplementedTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EGraphQLListener); ok {
		listenerT.ExitTemplateImplementedTypeDefinition(s)
	}
}

func (p *EGraphQLParser) TemplateImplementedTypeDefinition() (localctx ITemplateImplementedTypeDefinitionContext) {
	localctx = NewTemplateImplementedTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 148, EGraphQLParserRULE_templateImplementedTypeDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(707)
		p.Match(EGraphQLParserT__54)
	}
	{
		p.SetState(708)
		p.Name()
	}
	p.SetState(710)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__51 {
		{
			p.SetState(709)
			p.TemplateParametersDefinition()
		}

	}
	p.SetState(719)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == EGraphQLParserCOMMA {
		{
			p.SetState(712)
			p.Match(EGraphQLParserCOMMA)
		}
		{
			p.SetState(713)
			p.Name()
		}
		p.SetState(715)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EGraphQLParserT__51 {
			{
				p.SetState(714)
				p.TemplateParametersDefinition()
			}

		}

		p.SetState(721)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(722)
		p.Match(EGraphQLParserT__55)
	}
	{
		p.SetState(723)
		p.TempletableTypeDefinition()
	}

	return localctx
}

func (p *EGraphQLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 48:
		var t *ImplementsInterfacesContext = nil
		if localctx != nil {
			t = localctx.(*ImplementsInterfacesContext)
		}
		return p.ImplementsInterfaces_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *EGraphQLParser) ImplementsInterfaces_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
