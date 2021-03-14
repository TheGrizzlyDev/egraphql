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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 70, 701,
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
	71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 3, 2, 6, 2, 150, 10,
	2, 13, 2, 14, 2, 151, 3, 3, 3, 3, 3, 3, 5, 3, 157, 10, 3, 3, 4, 3, 4, 5,
	4, 161, 10, 4, 3, 5, 3, 5, 5, 5, 165, 10, 5, 3, 5, 5, 5, 168, 10, 5, 3,
	5, 5, 5, 171, 10, 5, 3, 5, 3, 5, 3, 5, 5, 5, 176, 10, 5, 3, 6, 3, 6, 3,
	7, 3, 7, 6, 7, 182, 10, 7, 13, 7, 14, 7, 183, 3, 7, 3, 7, 3, 8, 3, 8, 3,
	8, 5, 8, 191, 10, 8, 3, 9, 5, 9, 194, 10, 9, 3, 9, 3, 9, 5, 9, 198, 10,
	9, 3, 9, 5, 9, 201, 10, 9, 3, 9, 5, 9, 204, 10, 9, 3, 10, 3, 10, 6, 10,
	208, 10, 10, 13, 10, 14, 10, 209, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 5, 13, 224, 10, 13, 3, 14,
	3, 14, 3, 14, 3, 14, 5, 14, 230, 10, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3,
	16, 3, 16, 3, 16, 3, 17, 3, 17, 5, 17, 241, 10, 17, 3, 17, 5, 17, 244,
	10, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 5, 18, 257, 10, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3,
	21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25,
	6, 25, 275, 10, 25, 13, 25, 14, 25, 276, 3, 25, 3, 25, 5, 25, 281, 10,
	25, 3, 26, 3, 26, 7, 26, 285, 10, 26, 12, 26, 14, 26, 288, 11, 26, 3, 26,
	3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 6,
	29, 301, 10, 29, 13, 29, 14, 29, 302, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30,
	3, 30, 5, 30, 311, 10, 30, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 5, 32, 318,
	10, 32, 3, 32, 3, 32, 5, 32, 322, 10, 32, 5, 32, 324, 10, 32, 3, 33, 3,
	33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 6, 35, 333, 10, 35, 13, 35, 14,
	35, 334, 3, 36, 3, 36, 3, 36, 5, 36, 340, 10, 36, 3, 37, 3, 37, 3, 37,
	5, 37, 345, 10, 37, 3, 38, 3, 38, 5, 38, 349, 10, 38, 3, 39, 3, 39, 5,
	39, 353, 10, 39, 3, 39, 3, 39, 6, 39, 357, 10, 39, 13, 39, 14, 39, 358,
	3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 5, 41, 370,
	10, 41, 3, 41, 3, 41, 6, 41, 374, 10, 41, 13, 41, 14, 41, 375, 3, 41, 3,
	41, 3, 41, 3, 41, 3, 41, 5, 41, 383, 10, 41, 3, 42, 3, 42, 3, 42, 3, 42,
	3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 5, 44, 396, 10, 44, 3,
	45, 3, 45, 3, 45, 5, 45, 401, 10, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46,
	3, 46, 5, 46, 409, 10, 46, 3, 47, 5, 47, 412, 10, 47, 3, 47, 3, 47, 3,
	47, 5, 47, 417, 10, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 49, 5, 49,
	425, 10, 49, 3, 49, 3, 49, 3, 49, 5, 49, 430, 10, 49, 3, 49, 5, 49, 433,
	10, 49, 3, 49, 5, 49, 436, 10, 49, 3, 50, 3, 50, 3, 50, 5, 50, 441, 10,
	50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 7, 50, 448, 10, 50, 12, 50, 14,
	50, 451, 11, 50, 3, 51, 3, 51, 6, 51, 455, 10, 51, 13, 51, 14, 51, 456,
	3, 51, 3, 51, 3, 52, 5, 52, 462, 10, 52, 3, 52, 3, 52, 5, 52, 466, 10,
	52, 3, 52, 3, 52, 3, 52, 5, 52, 471, 10, 52, 3, 53, 3, 53, 6, 53, 475,
	10, 53, 13, 53, 14, 53, 476, 3, 53, 3, 53, 3, 54, 5, 54, 482, 10, 54, 3,
	54, 3, 54, 3, 54, 3, 54, 5, 54, 488, 10, 54, 3, 54, 5, 54, 491, 10, 54,
	3, 55, 3, 55, 3, 55, 3, 55, 5, 55, 497, 10, 55, 3, 55, 5, 55, 500, 10,
	55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 5, 55, 508, 10, 55, 3, 55,
	3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 5, 55, 517, 10, 55, 3, 56, 5,
	56, 520, 10, 56, 3, 56, 3, 56, 3, 56, 5, 56, 525, 10, 56, 3, 56, 5, 56,
	528, 10, 56, 3, 57, 3, 57, 3, 57, 3, 57, 5, 57, 534, 10, 57, 3, 57, 3,
	57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 5, 57, 543, 10, 57, 3, 58, 5, 58,
	546, 10, 58, 3, 58, 3, 58, 3, 58, 5, 58, 551, 10, 58, 3, 58, 5, 58, 554,
	10, 58, 3, 59, 3, 59, 5, 59, 558, 10, 59, 3, 59, 3, 59, 3, 59, 7, 59, 563,
	10, 59, 12, 59, 14, 59, 566, 11, 59, 3, 60, 3, 60, 3, 60, 3, 60, 5, 60,
	572, 10, 60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 60, 5, 60, 581,
	10, 60, 3, 61, 5, 61, 584, 10, 61, 3, 61, 3, 61, 3, 61, 5, 61, 589, 10,
	61, 3, 61, 5, 61, 592, 10, 61, 3, 62, 3, 62, 6, 62, 596, 10, 62, 13, 62,
	14, 62, 597, 3, 62, 3, 62, 3, 63, 5, 63, 603, 10, 63, 3, 63, 3, 63, 5,
	63, 607, 10, 63, 3, 64, 3, 64, 3, 64, 3, 64, 5, 64, 613, 10, 64, 3, 64,
	3, 64, 3, 64, 3, 64, 3, 64, 3, 64, 3, 64, 5, 64, 622, 10, 64, 3, 65, 5,
	65, 625, 10, 65, 3, 65, 3, 65, 3, 65, 5, 65, 630, 10, 65, 3, 65, 5, 65,
	633, 10, 65, 3, 66, 3, 66, 6, 66, 637, 10, 66, 13, 66, 14, 66, 638, 3,
	66, 3, 66, 3, 67, 3, 67, 3, 67, 3, 67, 5, 67, 647, 10, 67, 3, 67, 3, 67,
	3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 5, 67, 656, 10, 67, 3, 68, 5, 68, 659,
	10, 68, 3, 68, 3, 68, 3, 68, 3, 68, 5, 68, 665, 10, 68, 3, 68, 3, 68, 3,
	68, 3, 69, 3, 69, 3, 69, 7, 69, 673, 10, 69, 12, 69, 14, 69, 676, 11, 69,
	3, 70, 3, 70, 5, 70, 680, 10, 70, 3, 71, 3, 71, 3, 72, 3, 72, 3, 73, 3,
	73, 3, 74, 3, 74, 3, 74, 3, 74, 3, 74, 7, 74, 693, 10, 74, 12, 74, 14,
	74, 696, 11, 74, 3, 74, 3, 74, 3, 74, 3, 74, 2, 3, 98, 75, 2, 4, 6, 8,
	10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
	46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80,
	82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112,
	114, 116, 118, 120, 122, 124, 126, 128, 130, 132, 134, 136, 138, 140, 142,
	144, 146, 2, 7, 3, 2, 3, 5, 3, 2, 14, 15, 3, 2, 58, 59, 3, 2, 36, 42, 3,
	2, 43, 53, 2, 736, 2, 149, 3, 2, 2, 2, 4, 156, 3, 2, 2, 2, 6, 160, 3, 2,
	2, 2, 8, 175, 3, 2, 2, 2, 10, 177, 3, 2, 2, 2, 12, 179, 3, 2, 2, 2, 14,
	190, 3, 2, 2, 2, 16, 193, 3, 2, 2, 2, 18, 205, 3, 2, 2, 2, 20, 213, 3,
	2, 2, 2, 22, 217, 3, 2, 2, 2, 24, 220, 3, 2, 2, 2, 26, 225, 3, 2, 2, 2,
	28, 233, 3, 2, 2, 2, 30, 235, 3, 2, 2, 2, 32, 238, 3, 2, 2, 2, 34, 256,
	3, 2, 2, 2, 36, 258, 3, 2, 2, 2, 38, 260, 3, 2, 2, 2, 40, 262, 3, 2, 2,
	2, 42, 264, 3, 2, 2, 2, 44, 266, 3, 2, 2, 2, 46, 268, 3, 2, 2, 2, 48, 280,
	3, 2, 2, 2, 50, 282, 3, 2, 2, 2, 52, 291, 3, 2, 2, 2, 54, 295, 3, 2, 2,
	2, 56, 298, 3, 2, 2, 2, 58, 306, 3, 2, 2, 2, 60, 312, 3, 2, 2, 2, 62, 323,
	3, 2, 2, 2, 64, 325, 3, 2, 2, 2, 66, 327, 3, 2, 2, 2, 68, 332, 3, 2, 2,
	2, 70, 336, 3, 2, 2, 2, 72, 344, 3, 2, 2, 2, 74, 348, 3, 2, 2, 2, 76, 350,
	3, 2, 2, 2, 78, 362, 3, 2, 2, 2, 80, 382, 3, 2, 2, 2, 82, 384, 3, 2, 2,
	2, 84, 388, 3, 2, 2, 2, 86, 395, 3, 2, 2, 2, 88, 400, 3, 2, 2, 2, 90, 408,
	3, 2, 2, 2, 92, 411, 3, 2, 2, 2, 94, 418, 3, 2, 2, 2, 96, 424, 3, 2, 2,
	2, 98, 437, 3, 2, 2, 2, 100, 452, 3, 2, 2, 2, 102, 461, 3, 2, 2, 2, 104,
	472, 3, 2, 2, 2, 106, 481, 3, 2, 2, 2, 108, 516, 3, 2, 2, 2, 110, 519,
	3, 2, 2, 2, 112, 542, 3, 2, 2, 2, 114, 545, 3, 2, 2, 2, 116, 555, 3, 2,
	2, 2, 118, 580, 3, 2, 2, 2, 120, 583, 3, 2, 2, 2, 122, 593, 3, 2, 2, 2,
	124, 602, 3, 2, 2, 2, 126, 621, 3, 2, 2, 2, 128, 624, 3, 2, 2, 2, 130,
	634, 3, 2, 2, 2, 132, 655, 3, 2, 2, 2, 134, 658, 3, 2, 2, 2, 136, 669,
	3, 2, 2, 2, 138, 679, 3, 2, 2, 2, 140, 681, 3, 2, 2, 2, 142, 683, 3, 2,
	2, 2, 144, 685, 3, 2, 2, 2, 146, 687, 3, 2, 2, 2, 148, 150, 5, 4, 3, 2,
	149, 148, 3, 2, 2, 2, 150, 151, 3, 2, 2, 2, 151, 149, 3, 2, 2, 2, 151,
	152, 3, 2, 2, 2, 152, 3, 3, 2, 2, 2, 153, 157, 5, 6, 4, 2, 154, 157, 5,
	72, 37, 2, 155, 157, 5, 74, 38, 2, 156, 153, 3, 2, 2, 2, 156, 154, 3, 2,
	2, 2, 156, 155, 3, 2, 2, 2, 157, 5, 3, 2, 2, 2, 158, 161, 5, 8, 5, 2, 159,
	161, 5, 26, 14, 2, 160, 158, 3, 2, 2, 2, 160, 159, 3, 2, 2, 2, 161, 7,
	3, 2, 2, 2, 162, 164, 5, 10, 6, 2, 163, 165, 5, 144, 73, 2, 164, 163, 3,
	2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 167, 3, 2, 2, 2, 166, 168, 5, 56, 29,
	2, 167, 166, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168, 170, 3, 2, 2, 2, 169,
	171, 5, 68, 35, 2, 170, 169, 3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171, 172,
	3, 2, 2, 2, 172, 173, 5, 12, 7, 2, 173, 176, 3, 2, 2, 2, 174, 176, 5, 12,
	7, 2, 175, 162, 3, 2, 2, 2, 175, 174, 3, 2, 2, 2, 176, 9, 3, 2, 2, 2, 177,
	178, 9, 2, 2, 2, 178, 11, 3, 2, 2, 2, 179, 181, 7, 6, 2, 2, 180, 182, 5,
	14, 8, 2, 181, 180, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183, 181, 3, 2, 2,
	2, 183, 184, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 186, 7, 7, 2, 2, 186,
	13, 3, 2, 2, 2, 187, 191, 5, 16, 9, 2, 188, 191, 5, 24, 13, 2, 189, 191,
	5, 32, 17, 2, 190, 187, 3, 2, 2, 2, 190, 188, 3, 2, 2, 2, 190, 189, 3,
	2, 2, 2, 191, 15, 3, 2, 2, 2, 192, 194, 5, 22, 12, 2, 193, 192, 3, 2, 2,
	2, 193, 194, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195, 197, 5, 144, 73, 2,
	196, 198, 5, 18, 10, 2, 197, 196, 3, 2, 2, 2, 197, 198, 3, 2, 2, 2, 198,
	200, 3, 2, 2, 2, 199, 201, 5, 68, 35, 2, 200, 199, 3, 2, 2, 2, 200, 201,
	3, 2, 2, 2, 201, 203, 3, 2, 2, 2, 202, 204, 5, 12, 7, 2, 203, 202, 3, 2,
	2, 2, 203, 204, 3, 2, 2, 2, 204, 17, 3, 2, 2, 2, 205, 207, 7, 8, 2, 2,
	206, 208, 5, 20, 11, 2, 207, 206, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209,
	207, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 211, 3, 2, 2, 2, 211, 212,
	7, 9, 2, 2, 212, 19, 3, 2, 2, 2, 213, 214, 5, 144, 73, 2, 214, 215, 7,
	10, 2, 2, 215, 216, 5, 34, 18, 2, 216, 21, 3, 2, 2, 2, 217, 218, 5, 144,
	73, 2, 218, 219, 7, 10, 2, 2, 219, 23, 3, 2, 2, 2, 220, 221, 7, 11, 2,
	2, 221, 223, 5, 28, 15, 2, 222, 224, 5, 68, 35, 2, 223, 222, 3, 2, 2, 2,
	223, 224, 3, 2, 2, 2, 224, 25, 3, 2, 2, 2, 225, 226, 7, 12, 2, 2, 226,
	227, 5, 28, 15, 2, 227, 229, 5, 30, 16, 2, 228, 230, 5, 68, 35, 2, 229,
	228, 3, 2, 2, 2, 229, 230, 3, 2, 2, 2, 230, 231, 3, 2, 2, 2, 231, 232,
	5, 12, 7, 2, 232, 27, 3, 2, 2, 2, 233, 234, 5, 144, 73, 2, 234, 29, 3,
	2, 2, 2, 235, 236, 7, 13, 2, 2, 236, 237, 5, 64, 33, 2, 237, 31, 3, 2,
	2, 2, 238, 240, 7, 11, 2, 2, 239, 241, 5, 30, 16, 2, 240, 239, 3, 2, 2,
	2, 240, 241, 3, 2, 2, 2, 241, 243, 3, 2, 2, 2, 242, 244, 5, 68, 35, 2,
	243, 242, 3, 2, 2, 2, 243, 244, 3, 2, 2, 2, 244, 245, 3, 2, 2, 2, 245,
	246, 5, 12, 7, 2, 246, 33, 3, 2, 2, 2, 247, 257, 5, 54, 28, 2, 248, 257,
	5, 36, 19, 2, 249, 257, 5, 38, 20, 2, 250, 257, 5, 42, 22, 2, 251, 257,
	5, 40, 21, 2, 252, 257, 5, 44, 23, 2, 253, 257, 5, 46, 24, 2, 254, 257,
	5, 48, 25, 2, 255, 257, 5, 50, 26, 2, 256, 247, 3, 2, 2, 2, 256, 248, 3,
	2, 2, 2, 256, 249, 3, 2, 2, 2, 256, 250, 3, 2, 2, 2, 256, 251, 3, 2, 2,
	2, 256, 252, 3, 2, 2, 2, 256, 253, 3, 2, 2, 2, 256, 254, 3, 2, 2, 2, 256,
	255, 3, 2, 2, 2, 257, 35, 3, 2, 2, 2, 258, 259, 7, 62, 2, 2, 259, 37, 3,
	2, 2, 2, 260, 261, 7, 61, 2, 2, 261, 39, 3, 2, 2, 2, 262, 263, 9, 3, 2,
	2, 263, 41, 3, 2, 2, 2, 264, 265, 9, 4, 2, 2, 265, 43, 3, 2, 2, 2, 266,
	267, 7, 16, 2, 2, 267, 45, 3, 2, 2, 2, 268, 269, 5, 144, 73, 2, 269, 47,
	3, 2, 2, 2, 270, 271, 7, 17, 2, 2, 271, 281, 7, 18, 2, 2, 272, 274, 7,
	17, 2, 2, 273, 275, 5, 34, 18, 2, 274, 273, 3, 2, 2, 2, 275, 276, 3, 2,
	2, 2, 276, 274, 3, 2, 2, 2, 276, 277, 3, 2, 2, 2, 277, 278, 3, 2, 2, 2,
	278, 279, 7, 18, 2, 2, 279, 281, 3, 2, 2, 2, 280, 270, 3, 2, 2, 2, 280,
	272, 3, 2, 2, 2, 281, 49, 3, 2, 2, 2, 282, 286, 7, 6, 2, 2, 283, 285, 5,
	52, 27, 2, 284, 283, 3, 2, 2, 2, 285, 288, 3, 2, 2, 2, 286, 284, 3, 2,
	2, 2, 286, 287, 3, 2, 2, 2, 287, 289, 3, 2, 2, 2, 288, 286, 3, 2, 2, 2,
	289, 290, 7, 7, 2, 2, 290, 51, 3, 2, 2, 2, 291, 292, 5, 144, 73, 2, 292,
	293, 7, 10, 2, 2, 293, 294, 5, 34, 18, 2, 294, 53, 3, 2, 2, 2, 295, 296,
	7, 19, 2, 2, 296, 297, 5, 144, 73, 2, 297, 55, 3, 2, 2, 2, 298, 300, 7,
	8, 2, 2, 299, 301, 5, 58, 30, 2, 300, 299, 3, 2, 2, 2, 301, 302, 3, 2,
	2, 2, 302, 300, 3, 2, 2, 2, 302, 303, 3, 2, 2, 2, 303, 304, 3, 2, 2, 2,
	304, 305, 7, 9, 2, 2, 305, 57, 3, 2, 2, 2, 306, 307, 5, 54, 28, 2, 307,
	308, 7, 10, 2, 2, 308, 310, 5, 62, 32, 2, 309, 311, 5, 60, 31, 2, 310,
	309, 3, 2, 2, 2, 310, 311, 3, 2, 2, 2, 311, 59, 3, 2, 2, 2, 312, 313, 7,
	20, 2, 2, 313, 314, 5, 34, 18, 2, 314, 61, 3, 2, 2, 2, 315, 317, 5, 64,
	33, 2, 316, 318, 7, 21, 2, 2, 317, 316, 3, 2, 2, 2, 317, 318, 3, 2, 2,
	2, 318, 324, 3, 2, 2, 2, 319, 321, 5, 66, 34, 2, 320, 322, 7, 21, 2, 2,
	321, 320, 3, 2, 2, 2, 321, 322, 3, 2, 2, 2, 322, 324, 3, 2, 2, 2, 323,
	315, 3, 2, 2, 2, 323, 319, 3, 2, 2, 2, 324, 63, 3, 2, 2, 2, 325, 326, 5,
	144, 73, 2, 326, 65, 3, 2, 2, 2, 327, 328, 7, 17, 2, 2, 328, 329, 5, 62,
	32, 2, 329, 330, 7, 18, 2, 2, 330, 67, 3, 2, 2, 2, 331, 333, 5, 70, 36,
	2, 332, 331, 3, 2, 2, 2, 333, 334, 3, 2, 2, 2, 334, 332, 3, 2, 2, 2, 334,
	335, 3, 2, 2, 2, 335, 69, 3, 2, 2, 2, 336, 337, 7, 22, 2, 2, 337, 339,
	5, 144, 73, 2, 338, 340, 5, 18, 10, 2, 339, 338, 3, 2, 2, 2, 339, 340,
	3, 2, 2, 2, 340, 71, 3, 2, 2, 2, 341, 345, 5, 76, 39, 2, 342, 345, 5, 86,
	44, 2, 343, 345, 5, 134, 68, 2, 344, 341, 3, 2, 2, 2, 344, 342, 3, 2, 2,
	2, 344, 343, 3, 2, 2, 2, 345, 73, 3, 2, 2, 2, 346, 349, 5, 80, 41, 2, 347,
	349, 5, 90, 46, 2, 348, 346, 3, 2, 2, 2, 348, 347, 3, 2, 2, 2, 349, 75,
	3, 2, 2, 2, 350, 352, 7, 23, 2, 2, 351, 353, 5, 68, 35, 2, 352, 351, 3,
	2, 2, 2, 352, 353, 3, 2, 2, 2, 353, 354, 3, 2, 2, 2, 354, 356, 7, 6, 2,
	2, 355, 357, 5, 78, 40, 2, 356, 355, 3, 2, 2, 2, 357, 358, 3, 2, 2, 2,
	358, 356, 3, 2, 2, 2, 358, 359, 3, 2, 2, 2, 359, 360, 3, 2, 2, 2, 360,
	361, 7, 7, 2, 2, 361, 77, 3, 2, 2, 2, 362, 363, 5, 10, 6, 2, 363, 364,
	7, 10, 2, 2, 364, 365, 5, 64, 33, 2, 365, 79, 3, 2, 2, 2, 366, 367, 7,
	24, 2, 2, 367, 369, 7, 23, 2, 2, 368, 370, 5, 68, 35, 2, 369, 368, 3, 2,
	2, 2, 369, 370, 3, 2, 2, 2, 370, 371, 3, 2, 2, 2, 371, 373, 7, 6, 2, 2,
	372, 374, 5, 82, 42, 2, 373, 372, 3, 2, 2, 2, 374, 375, 3, 2, 2, 2, 375,
	373, 3, 2, 2, 2, 375, 376, 3, 2, 2, 2, 376, 377, 3, 2, 2, 2, 377, 378,
	7, 7, 2, 2, 378, 383, 3, 2, 2, 2, 379, 380, 7, 24, 2, 2, 380, 381, 7, 23,
	2, 2, 381, 383, 5, 68, 35, 2, 382, 366, 3, 2, 2, 2, 382, 379, 3, 2, 2,
	2, 383, 81, 3, 2, 2, 2, 384, 385, 5, 10, 6, 2, 385, 386, 7, 10, 2, 2, 386,
	387, 5, 64, 33, 2, 387, 83, 3, 2, 2, 2, 388, 389, 5, 42, 22, 2, 389, 85,
	3, 2, 2, 2, 390, 396, 5, 92, 47, 2, 391, 396, 5, 114, 58, 2, 392, 396,
	5, 120, 61, 2, 393, 396, 5, 88, 45, 2, 394, 396, 5, 146, 74, 2, 395, 390,
	3, 2, 2, 2, 395, 391, 3, 2, 2, 2, 395, 392, 3, 2, 2, 2, 395, 393, 3, 2,
	2, 2, 395, 394, 3, 2, 2, 2, 396, 87, 3, 2, 2, 2, 397, 401, 5, 96, 49, 2,
	398, 401, 5, 110, 56, 2, 399, 401, 5, 128, 65, 2, 400, 397, 3, 2, 2, 2,
	400, 398, 3, 2, 2, 2, 400, 399, 3, 2, 2, 2, 401, 89, 3, 2, 2, 2, 402, 409,
	5, 94, 48, 2, 403, 409, 5, 108, 55, 2, 404, 409, 5, 112, 57, 2, 405, 409,
	5, 118, 60, 2, 406, 409, 5, 126, 64, 2, 407, 409, 5, 132, 67, 2, 408, 402,
	3, 2, 2, 2, 408, 403, 3, 2, 2, 2, 408, 404, 3, 2, 2, 2, 408, 405, 3, 2,
	2, 2, 408, 406, 3, 2, 2, 2, 408, 407, 3, 2, 2, 2, 409, 91, 3, 2, 2, 2,
	410, 412, 5, 84, 43, 2, 411, 410, 3, 2, 2, 2, 411, 412, 3, 2, 2, 2, 412,
	413, 3, 2, 2, 2, 413, 414, 7, 25, 2, 2, 414, 416, 5, 144, 73, 2, 415, 417,
	5, 68, 35, 2, 416, 415, 3, 2, 2, 2, 416, 417, 3, 2, 2, 2, 417, 93, 3, 2,
	2, 2, 418, 419, 7, 26, 2, 2, 419, 420, 7, 25, 2, 2, 420, 421, 5, 144, 73,
	2, 421, 422, 5, 68, 35, 2, 422, 95, 3, 2, 2, 2, 423, 425, 5, 84, 43, 2,
	424, 423, 3, 2, 2, 2, 424, 425, 3, 2, 2, 2, 425, 426, 3, 2, 2, 2, 426,
	427, 7, 27, 2, 2, 427, 429, 5, 144, 73, 2, 428, 430, 5, 98, 50, 2, 429,
	428, 3, 2, 2, 2, 429, 430, 3, 2, 2, 2, 430, 432, 3, 2, 2, 2, 431, 433,
	5, 68, 35, 2, 432, 431, 3, 2, 2, 2, 432, 433, 3, 2, 2, 2, 433, 435, 3,
	2, 2, 2, 434, 436, 5, 100, 51, 2, 435, 434, 3, 2, 2, 2, 435, 436, 3, 2,
	2, 2, 436, 97, 3, 2, 2, 2, 437, 438, 8, 50, 1, 2, 438, 440, 7, 28, 2, 2,
	439, 441, 7, 29, 2, 2, 440, 439, 3, 2, 2, 2, 440, 441, 3, 2, 2, 2, 441,
	442, 3, 2, 2, 2, 442, 443, 5, 64, 33, 2, 443, 449, 3, 2, 2, 2, 444, 445,
	12, 3, 2, 2, 445, 446, 7, 29, 2, 2, 446, 448, 5, 64, 33, 2, 447, 444, 3,
	2, 2, 2, 448, 451, 3, 2, 2, 2, 449, 447, 3, 2, 2, 2, 449, 450, 3, 2, 2,
	2, 450, 99, 3, 2, 2, 2, 451, 449, 3, 2, 2, 2, 452, 454, 7, 6, 2, 2, 453,
	455, 5, 102, 52, 2, 454, 453, 3, 2, 2, 2, 455, 456, 3, 2, 2, 2, 456, 454,
	3, 2, 2, 2, 456, 457, 3, 2, 2, 2, 457, 458, 3, 2, 2, 2, 458, 459, 7, 7,
	2, 2, 459, 101, 3, 2, 2, 2, 460, 462, 5, 84, 43, 2, 461, 460, 3, 2, 2,
	2, 461, 462, 3, 2, 2, 2, 462, 463, 3, 2, 2, 2, 463, 465, 5, 144, 73, 2,
	464, 466, 5, 104, 53, 2, 465, 464, 3, 2, 2, 2, 465, 466, 3, 2, 2, 2, 466,
	467, 3, 2, 2, 2, 467, 468, 7, 10, 2, 2, 468, 470, 5, 62, 32, 2, 469, 471,
	5, 68, 35, 2, 470, 469, 3, 2, 2, 2, 470, 471, 3, 2, 2, 2, 471, 103, 3,
	2, 2, 2, 472, 474, 7, 8, 2, 2, 473, 475, 5, 106, 54, 2, 474, 473, 3, 2,
	2, 2, 475, 476, 3, 2, 2, 2, 476, 474, 3, 2, 2, 2, 476, 477, 3, 2, 2, 2,
	477, 478, 3, 2, 2, 2, 478, 479, 7, 9, 2, 2, 479, 105, 3, 2, 2, 2, 480,
	482, 5, 84, 43, 2, 481, 480, 3, 2, 2, 2, 481, 482, 3, 2, 2, 2, 482, 483,
	3, 2, 2, 2, 483, 484, 5, 144, 73, 2, 484, 485, 7, 10, 2, 2, 485, 487, 5,
	62, 32, 2, 486, 488, 5, 60, 31, 2, 487, 486, 3, 2, 2, 2, 487, 488, 3, 2,
	2, 2, 488, 490, 3, 2, 2, 2, 489, 491, 5, 68, 35, 2, 490, 489, 3, 2, 2,
	2, 490, 491, 3, 2, 2, 2, 491, 107, 3, 2, 2, 2, 492, 493, 7, 24, 2, 2, 493,
	494, 7, 27, 2, 2, 494, 496, 5, 144, 73, 2, 495, 497, 5, 98, 50, 2, 496,
	495, 3, 2, 2, 2, 496, 497, 3, 2, 2, 2, 497, 499, 3, 2, 2, 2, 498, 500,
	5, 68, 35, 2, 499, 498, 3, 2, 2, 2, 499, 500, 3, 2, 2, 2, 500, 501, 3,
	2, 2, 2, 501, 502, 5, 100, 51, 2, 502, 517, 3, 2, 2, 2, 503, 504, 7, 24,
	2, 2, 504, 505, 7, 27, 2, 2, 505, 507, 5, 144, 73, 2, 506, 508, 5, 98,
	50, 2, 507, 506, 3, 2, 2, 2, 507, 508, 3, 2, 2, 2, 508, 509, 3, 2, 2, 2,
	509, 510, 5, 68, 35, 2, 510, 517, 3, 2, 2, 2, 511, 512, 7, 24, 2, 2, 512,
	513, 7, 27, 2, 2, 513, 514, 5, 144, 73, 2, 514, 515, 5, 98, 50, 2, 515,
	517, 3, 2, 2, 2, 516, 492, 3, 2, 2, 2, 516, 503, 3, 2, 2, 2, 516, 511,
	3, 2, 2, 2, 517, 109, 3, 2, 2, 2, 518, 520, 5, 84, 43, 2, 519, 518, 3,
	2, 2, 2, 519, 520, 3, 2, 2, 2, 520, 521, 3, 2, 2, 2, 521, 522, 7, 30, 2,
	2, 522, 524, 5, 144, 73, 2, 523, 525, 5, 68, 35, 2, 524, 523, 3, 2, 2,
	2, 524, 525, 3, 2, 2, 2, 525, 527, 3, 2, 2, 2, 526, 528, 5, 100, 51, 2,
	527, 526, 3, 2, 2, 2, 527, 528, 3, 2, 2, 2, 528, 111, 3, 2, 2, 2, 529,
	530, 7, 24, 2, 2, 530, 531, 7, 30, 2, 2, 531, 533, 5, 144, 73, 2, 532,
	534, 5, 68, 35, 2, 533, 532, 3, 2, 2, 2, 533, 534, 3, 2, 2, 2, 534, 535,
	3, 2, 2, 2, 535, 536, 5, 100, 51, 2, 536, 543, 3, 2, 2, 2, 537, 538, 7,
	24, 2, 2, 538, 539, 7, 30, 2, 2, 539, 540, 5, 144, 73, 2, 540, 541, 5,
	68, 35, 2, 541, 543, 3, 2, 2, 2, 542, 529, 3, 2, 2, 2, 542, 537, 3, 2,
	2, 2, 543, 113, 3, 2, 2, 2, 544, 546, 5, 84, 43, 2, 545, 544, 3, 2, 2,
	2, 545, 546, 3, 2, 2, 2, 546, 547, 3, 2, 2, 2, 547, 548, 7, 31, 2, 2, 548,
	550, 5, 144, 73, 2, 549, 551, 5, 68, 35, 2, 550, 549, 3, 2, 2, 2, 550,
	551, 3, 2, 2, 2, 551, 553, 3, 2, 2, 2, 552, 554, 5, 116, 59, 2, 553, 552,
	3, 2, 2, 2, 553, 554, 3, 2, 2, 2, 554, 115, 3, 2, 2, 2, 555, 557, 7, 20,
	2, 2, 556, 558, 7, 32, 2, 2, 557, 556, 3, 2, 2, 2, 557, 558, 3, 2, 2, 2,
	558, 559, 3, 2, 2, 2, 559, 564, 5, 64, 33, 2, 560, 561, 7, 32, 2, 2, 561,
	563, 5, 64, 33, 2, 562, 560, 3, 2, 2, 2, 563, 566, 3, 2, 2, 2, 564, 562,
	3, 2, 2, 2, 564, 565, 3, 2, 2, 2, 565, 117, 3, 2, 2, 2, 566, 564, 3, 2,
	2, 2, 567, 568, 7, 24, 2, 2, 568, 569, 7, 31, 2, 2, 569, 571, 5, 144, 73,
	2, 570, 572, 5, 68, 35, 2, 571, 570, 3, 2, 2, 2, 571, 572, 3, 2, 2, 2,
	572, 573, 3, 2, 2, 2, 573, 574, 5, 116, 59, 2, 574, 581, 3, 2, 2, 2, 575,
	576, 7, 24, 2, 2, 576, 577, 7, 31, 2, 2, 577, 578, 5, 144, 73, 2, 578,
	579, 5, 68, 35, 2, 579, 581, 3, 2, 2, 2, 580, 567, 3, 2, 2, 2, 580, 575,
	3, 2, 2, 2, 581, 119, 3, 2, 2, 2, 582, 584, 5, 84, 43, 2, 583, 582, 3,
	2, 2, 2, 583, 584, 3, 2, 2, 2, 584, 585, 3, 2, 2, 2, 585, 586, 7, 33, 2,
	2, 586, 588, 5, 144, 73, 2, 587, 589, 5, 68, 35, 2, 588, 587, 3, 2, 2,
	2, 588, 589, 3, 2, 2, 2, 589, 591, 3, 2, 2, 2, 590, 592, 5, 122, 62, 2,
	591, 590, 3, 2, 2, 2, 591, 592, 3, 2, 2, 2, 592, 121, 3, 2, 2, 2, 593,
	595, 7, 6, 2, 2, 594, 596, 5, 124, 63, 2, 595, 594, 3, 2, 2, 2, 596, 597,
	3, 2, 2, 2, 597, 595, 3, 2, 2, 2, 597, 598, 3, 2, 2, 2, 598, 599, 3, 2,
	2, 2, 599, 600, 7, 7, 2, 2, 600, 123, 3, 2, 2, 2, 601, 603, 5, 84, 43,
	2, 602, 601, 3, 2, 2, 2, 602, 603, 3, 2, 2, 2, 603, 604, 3, 2, 2, 2, 604,
	606, 5, 46, 24, 2, 605, 607, 5, 68, 35, 2, 606, 605, 3, 2, 2, 2, 606, 607,
	3, 2, 2, 2, 607, 125, 3, 2, 2, 2, 608, 609, 7, 24, 2, 2, 609, 610, 7, 33,
	2, 2, 610, 612, 5, 144, 73, 2, 611, 613, 5, 68, 35, 2, 612, 611, 3, 2,
	2, 2, 612, 613, 3, 2, 2, 2, 613, 614, 3, 2, 2, 2, 614, 615, 5, 122, 62,
	2, 615, 622, 3, 2, 2, 2, 616, 617, 7, 24, 2, 2, 617, 618, 7, 33, 2, 2,
	618, 619, 5, 144, 73, 2, 619, 620, 5, 68, 35, 2, 620, 622, 3, 2, 2, 2,
	621, 608, 3, 2, 2, 2, 621, 616, 3, 2, 2, 2, 622, 127, 3, 2, 2, 2, 623,
	625, 5, 84, 43, 2, 624, 623, 3, 2, 2, 2, 624, 625, 3, 2, 2, 2, 625, 626,
	3, 2, 2, 2, 626, 627, 7, 34, 2, 2, 627, 629, 5, 144, 73, 2, 628, 630, 5,
	68, 35, 2, 629, 628, 3, 2, 2, 2, 629, 630, 3, 2, 2, 2, 630, 632, 3, 2,
	2, 2, 631, 633, 5, 130, 66, 2, 632, 631, 3, 2, 2, 2, 632, 633, 3, 2, 2,
	2, 633, 129, 3, 2, 2, 2, 634, 636, 7, 6, 2, 2, 635, 637, 5, 106, 54, 2,
	636, 635, 3, 2, 2, 2, 637, 638, 3, 2, 2, 2, 638, 636, 3, 2, 2, 2, 638,
	639, 3, 2, 2, 2, 639, 640, 3, 2, 2, 2, 640, 641, 7, 7, 2, 2, 641, 131,
	3, 2, 2, 2, 642, 643, 7, 24, 2, 2, 643, 644, 7, 34, 2, 2, 644, 646, 5,
	144, 73, 2, 645, 647, 5, 68, 35, 2, 646, 645, 3, 2, 2, 2, 646, 647, 3,
	2, 2, 2, 647, 648, 3, 2, 2, 2, 648, 649, 5, 130, 66, 2, 649, 656, 3, 2,
	2, 2, 650, 651, 7, 24, 2, 2, 651, 652, 7, 34, 2, 2, 652, 653, 5, 144, 73,
	2, 653, 654, 5, 68, 35, 2, 654, 656, 3, 2, 2, 2, 655, 642, 3, 2, 2, 2,
	655, 650, 3, 2, 2, 2, 656, 133, 3, 2, 2, 2, 657, 659, 5, 84, 43, 2, 658,
	657, 3, 2, 2, 2, 658, 659, 3, 2, 2, 2, 659, 660, 3, 2, 2, 2, 660, 661,
	7, 35, 2, 2, 661, 662, 7, 22, 2, 2, 662, 664, 5, 144, 73, 2, 663, 665,
	5, 104, 53, 2, 664, 663, 3, 2, 2, 2, 664, 665, 3, 2, 2, 2, 665, 666, 3,
	2, 2, 2, 666, 667, 7, 13, 2, 2, 667, 668, 5, 136, 69, 2, 668, 135, 3, 2,
	2, 2, 669, 674, 5, 138, 70, 2, 670, 671, 7, 32, 2, 2, 671, 673, 5, 138,
	70, 2, 672, 670, 3, 2, 2, 2, 673, 676, 3, 2, 2, 2, 674, 672, 3, 2, 2, 2,
	674, 675, 3, 2, 2, 2, 675, 137, 3, 2, 2, 2, 676, 674, 3, 2, 2, 2, 677,
	680, 5, 140, 71, 2, 678, 680, 5, 142, 72, 2, 679, 677, 3, 2, 2, 2, 679,
	678, 3, 2, 2, 2, 680, 139, 3, 2, 2, 2, 681, 682, 9, 5, 2, 2, 682, 141,
	3, 2, 2, 2, 683, 684, 9, 6, 2, 2, 684, 143, 3, 2, 2, 2, 685, 686, 7, 57,
	2, 2, 686, 145, 3, 2, 2, 2, 687, 688, 7, 54, 2, 2, 688, 689, 7, 55, 2,
	2, 689, 694, 5, 144, 73, 2, 690, 691, 7, 65, 2, 2, 691, 693, 5, 144, 73,
	2, 692, 690, 3, 2, 2, 2, 693, 696, 3, 2, 2, 2, 694, 692, 3, 2, 2, 2, 694,
	695, 3, 2, 2, 2, 695, 697, 3, 2, 2, 2, 696, 694, 3, 2, 2, 2, 697, 698,
	7, 56, 2, 2, 698, 699, 5, 88, 45, 2, 699, 147, 3, 2, 2, 2, 92, 151, 156,
	160, 164, 167, 170, 175, 183, 190, 193, 197, 200, 203, 209, 223, 229, 240,
	243, 256, 276, 280, 286, 302, 310, 317, 321, 323, 334, 339, 344, 348, 352,
	358, 369, 375, 382, 395, 400, 408, 411, 416, 424, 429, 432, 435, 440, 449,
	456, 461, 465, 470, 476, 481, 487, 490, 496, 499, 507, 516, 519, 524, 527,
	533, 542, 545, 550, 553, 557, 564, 571, 580, 583, 588, 591, 597, 602, 606,
	612, 621, 624, 629, 632, 638, 646, 655, 658, 664, 674, 679, 694,
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
	"'template'", "'<'", "'>'", "", "", "", "", "", "", "", "", "','", "",
	"", "'\uEFBBBF'", "'\uFEFF'", "'\u0000FEFF'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "NAME", "STRING", "BLOCK_STRING", "ID", "FLOAT", "INT", "PUNCTUATOR",
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
	"typeSystemDirectiveLocation", "name", "templateTypeDefinition",
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
	EGraphQLParserNAME         = 55
	EGraphQLParserSTRING       = 56
	EGraphQLParserBLOCK_STRING = 57
	EGraphQLParserID           = 58
	EGraphQLParserFLOAT        = 59
	EGraphQLParserINT          = 60
	EGraphQLParserPUNCTUATOR   = 61
	EGraphQLParserWS           = 62
	EGraphQLParserCOMMA        = 63
	EGraphQLParserLineComment  = 64
	EGraphQLParserUNICODE_BOM  = 65
	EGraphQLParserUTF8_BOM     = 66
	EGraphQLParserUTF16_BOM    = 67
	EGraphQLParserUTF32_BOM    = 68
)

// EGraphQLParser rules.
const (
	EGraphQLParserRULE_document                    = 0
	EGraphQLParserRULE_definition                  = 1
	EGraphQLParserRULE_executableDefinition        = 2
	EGraphQLParserRULE_operationDefinition         = 3
	EGraphQLParserRULE_operationType               = 4
	EGraphQLParserRULE_selectionSet                = 5
	EGraphQLParserRULE_selection                   = 6
	EGraphQLParserRULE_field                       = 7
	EGraphQLParserRULE_arguments                   = 8
	EGraphQLParserRULE_argument                    = 9
	EGraphQLParserRULE_alias                       = 10
	EGraphQLParserRULE_fragmentSpread              = 11
	EGraphQLParserRULE_fragmentDefinition          = 12
	EGraphQLParserRULE_fragmentName                = 13
	EGraphQLParserRULE_typeCondition               = 14
	EGraphQLParserRULE_inlineFragment              = 15
	EGraphQLParserRULE_value                       = 16
	EGraphQLParserRULE_intValue                    = 17
	EGraphQLParserRULE_floatValue                  = 18
	EGraphQLParserRULE_booleanValue                = 19
	EGraphQLParserRULE_stringValue                 = 20
	EGraphQLParserRULE_nullValue                   = 21
	EGraphQLParserRULE_enumValue                   = 22
	EGraphQLParserRULE_listValue                   = 23
	EGraphQLParserRULE_objectValue                 = 24
	EGraphQLParserRULE_objectField                 = 25
	EGraphQLParserRULE_variable                    = 26
	EGraphQLParserRULE_variableDefinitions         = 27
	EGraphQLParserRULE_variableDefinition          = 28
	EGraphQLParserRULE_defaultValue                = 29
	EGraphQLParserRULE_type_                       = 30
	EGraphQLParserRULE_namedType                   = 31
	EGraphQLParserRULE_listType                    = 32
	EGraphQLParserRULE_directives                  = 33
	EGraphQLParserRULE_directive                   = 34
	EGraphQLParserRULE_typeSystemDefinition        = 35
	EGraphQLParserRULE_typeSystemExtension         = 36
	EGraphQLParserRULE_schemaDefinition            = 37
	EGraphQLParserRULE_rootOperationTypeDefinition = 38
	EGraphQLParserRULE_schemaExtension             = 39
	EGraphQLParserRULE_operationTypeDefinition     = 40
	EGraphQLParserRULE_description                 = 41
	EGraphQLParserRULE_typeDefinition              = 42
	EGraphQLParserRULE_templetableTypeDefinition   = 43
	EGraphQLParserRULE_typeExtension               = 44
	EGraphQLParserRULE_scalarTypeDefinition        = 45
	EGraphQLParserRULE_scalarTypeExtension         = 46
	EGraphQLParserRULE_objectTypeDefinition        = 47
	EGraphQLParserRULE_implementsInterfaces        = 48
	EGraphQLParserRULE_fieldsDefinition            = 49
	EGraphQLParserRULE_fieldDefinition             = 50
	EGraphQLParserRULE_argumentsDefinition         = 51
	EGraphQLParserRULE_inputValueDefinition        = 52
	EGraphQLParserRULE_objectTypeExtension         = 53
	EGraphQLParserRULE_interfaceTypeDefinition     = 54
	EGraphQLParserRULE_interfaceTypeExtension      = 55
	EGraphQLParserRULE_unionTypeDefinition         = 56
	EGraphQLParserRULE_unionMemberTypes            = 57
	EGraphQLParserRULE_unionTypeExtension          = 58
	EGraphQLParserRULE_enumTypeDefinition          = 59
	EGraphQLParserRULE_enumValuesDefinition        = 60
	EGraphQLParserRULE_enumValueDefinition         = 61
	EGraphQLParserRULE_enumTypeExtension           = 62
	EGraphQLParserRULE_inputObjectTypeDefinition   = 63
	EGraphQLParserRULE_inputFieldsDefinition       = 64
	EGraphQLParserRULE_inputObjectTypeExtension    = 65
	EGraphQLParserRULE_directiveDefinition         = 66
	EGraphQLParserRULE_directiveLocations          = 67
	EGraphQLParserRULE_directiveLocation           = 68
	EGraphQLParserRULE_executableDirectiveLocation = 69
	EGraphQLParserRULE_typeSystemDirectiveLocation = 70
	EGraphQLParserRULE_name                        = 71
	EGraphQLParserRULE_templateTypeDefinition      = 72
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
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<EGraphQLParserT__0)|(1<<EGraphQLParserT__1)|(1<<EGraphQLParserT__2)|(1<<EGraphQLParserT__3)|(1<<EGraphQLParserT__9)|(1<<EGraphQLParserT__20)|(1<<EGraphQLParserT__21)|(1<<EGraphQLParserT__22)|(1<<EGraphQLParserT__23)|(1<<EGraphQLParserT__24)|(1<<EGraphQLParserT__27)|(1<<EGraphQLParserT__28)|(1<<EGraphQLParserT__30))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(EGraphQLParserT__31-32))|(1<<(EGraphQLParserT__32-32))|(1<<(EGraphQLParserT__51-32))|(1<<(EGraphQLParserSTRING-32))|(1<<(EGraphQLParserBLOCK_STRING-32)))) != 0) {
		{
			p.SetState(146)
			p.Definition()
		}

		p.SetState(149)
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

	p.SetState(154)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case EGraphQLParserT__0, EGraphQLParserT__1, EGraphQLParserT__2, EGraphQLParserT__3, EGraphQLParserT__9:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(151)
			p.ExecutableDefinition()
		}

	case EGraphQLParserT__20, EGraphQLParserT__22, EGraphQLParserT__24, EGraphQLParserT__27, EGraphQLParserT__28, EGraphQLParserT__30, EGraphQLParserT__31, EGraphQLParserT__32, EGraphQLParserT__51, EGraphQLParserSTRING, EGraphQLParserBLOCK_STRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(152)
			p.TypeSystemDefinition()
		}

	case EGraphQLParserT__21, EGraphQLParserT__23:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(153)
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

	p.SetState(158)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case EGraphQLParserT__0, EGraphQLParserT__1, EGraphQLParserT__2, EGraphQLParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(156)
			p.OperationDefinition()
		}

	case EGraphQLParserT__9:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(157)
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

	p.SetState(173)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case EGraphQLParserT__0, EGraphQLParserT__1, EGraphQLParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(160)
			p.OperationType()
		}
		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EGraphQLParserNAME {
			{
				p.SetState(161)
				p.Name()
			}

		}
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EGraphQLParserT__5 {
			{
				p.SetState(164)
				p.VariableDefinitions()
			}

		}
		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EGraphQLParserT__19 {
			{
				p.SetState(167)
				p.Directives()
			}

		}
		{
			p.SetState(170)
			p.SelectionSet()
		}

	case EGraphQLParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(172)
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
		p.SetState(175)
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
		p.SetState(177)
		p.Match(EGraphQLParserT__3)
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == EGraphQLParserT__8 || _la == EGraphQLParserNAME {
		{
			p.SetState(178)
			p.Selection()
		}

		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(183)
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

	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(185)
			p.Field()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(186)
			p.FragmentSpread()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(187)
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
	p.SetState(191)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(190)
			p.Alias()
		}

	}
	{
		p.SetState(193)
		p.Name()
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__5 {
		{
			p.SetState(194)
			p.Arguments()
		}

	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__19 {
		{
			p.SetState(197)
			p.Directives()
		}

	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__3 {
		{
			p.SetState(200)
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
		p.SetState(203)
		p.Match(EGraphQLParserT__5)
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == EGraphQLParserNAME {
		{
			p.SetState(204)
			p.Argument()
		}

		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(209)
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
		p.SetState(211)
		p.Name()
	}
	{
		p.SetState(212)
		p.Match(EGraphQLParserT__7)
	}
	{
		p.SetState(213)
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
		p.SetState(215)
		p.Name()
	}
	{
		p.SetState(216)
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
		p.SetState(218)
		p.Match(EGraphQLParserT__8)
	}
	{
		p.SetState(219)
		p.FragmentName()
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__19 {
		{
			p.SetState(220)
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
		p.SetState(223)
		p.Match(EGraphQLParserT__9)
	}
	{
		p.SetState(224)
		p.FragmentName()
	}
	{
		p.SetState(225)
		p.TypeCondition()
	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__19 {
		{
			p.SetState(226)
			p.Directives()
		}

	}
	{
		p.SetState(229)
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
		p.SetState(231)
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
		p.SetState(233)
		p.Match(EGraphQLParserT__10)
	}
	{
		p.SetState(234)
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
		p.SetState(236)
		p.Match(EGraphQLParserT__8)
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__10 {
		{
			p.SetState(237)
			p.TypeCondition()
		}

	}
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__19 {
		{
			p.SetState(240)
			p.Directives()
		}

	}
	{
		p.SetState(243)
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

	p.SetState(254)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case EGraphQLParserT__16:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(245)
			p.Variable()
		}

	case EGraphQLParserINT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(246)
			p.IntValue()
		}

	case EGraphQLParserFLOAT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(247)
			p.FloatValue()
		}

	case EGraphQLParserSTRING, EGraphQLParserBLOCK_STRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(248)
			p.StringValue()
		}

	case EGraphQLParserT__11, EGraphQLParserT__12:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(249)
			p.BooleanValue()
		}

	case EGraphQLParserT__13:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(250)
			p.NullValue()
		}

	case EGraphQLParserNAME:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(251)
			p.EnumValue()
		}

	case EGraphQLParserT__14:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(252)
			p.ListValue()
		}

	case EGraphQLParserT__3:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(253)
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
		p.SetState(256)
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
		p.SetState(258)
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
		p.SetState(260)
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
		p.SetState(262)
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
		p.SetState(264)
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
		p.SetState(266)
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

	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(268)
			p.Match(EGraphQLParserT__14)
		}
		{
			p.SetState(269)
			p.Match(EGraphQLParserT__15)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(270)
			p.Match(EGraphQLParserT__14)
		}
		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<EGraphQLParserT__3)|(1<<EGraphQLParserT__11)|(1<<EGraphQLParserT__12)|(1<<EGraphQLParserT__13)|(1<<EGraphQLParserT__14)|(1<<EGraphQLParserT__16))) != 0) || (((_la-55)&-(0x1f+1)) == 0 && ((1<<uint((_la-55)))&((1<<(EGraphQLParserNAME-55))|(1<<(EGraphQLParserSTRING-55))|(1<<(EGraphQLParserBLOCK_STRING-55))|(1<<(EGraphQLParserFLOAT-55))|(1<<(EGraphQLParserINT-55)))) != 0) {
			{
				p.SetState(271)
				p.Value()
			}

			p.SetState(274)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(276)
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
		p.SetState(280)
		p.Match(EGraphQLParserT__3)
	}
	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == EGraphQLParserNAME {
		{
			p.SetState(281)
			p.ObjectField()
		}

		p.SetState(286)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(287)
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
		p.SetState(289)
		p.Name()
	}
	{
		p.SetState(290)
		p.Match(EGraphQLParserT__7)
	}
	{
		p.SetState(291)
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
		p.SetState(293)
		p.Match(EGraphQLParserT__16)
	}
	{
		p.SetState(294)
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
		p.SetState(296)
		p.Match(EGraphQLParserT__5)
	}
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == EGraphQLParserT__16 {
		{
			p.SetState(297)
			p.VariableDefinition()
		}

		p.SetState(300)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(302)
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
		p.SetState(304)
		p.Variable()
	}
	{
		p.SetState(305)
		p.Match(EGraphQLParserT__7)
	}
	{
		p.SetState(306)
		p.Type_()
	}
	p.SetState(308)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__17 {
		{
			p.SetState(307)
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
		p.SetState(310)
		p.Match(EGraphQLParserT__17)
	}
	{
		p.SetState(311)
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

	p.SetState(321)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case EGraphQLParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(313)
			p.NamedType()
		}
		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EGraphQLParserT__18 {
			{
				p.SetState(314)
				p.Match(EGraphQLParserT__18)
			}

		}

	case EGraphQLParserT__14:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(317)
			p.ListType()
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
		p.SetState(323)
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
		p.SetState(325)
		p.Match(EGraphQLParserT__14)
	}
	{
		p.SetState(326)
		p.Type_()
	}
	{
		p.SetState(327)
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
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == EGraphQLParserT__19 {
		{
			p.SetState(329)
			p.Directive()
		}

		p.SetState(332)
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
		p.SetState(334)
		p.Match(EGraphQLParserT__19)
	}
	{
		p.SetState(335)
		p.Name()
	}
	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__5 {
		{
			p.SetState(336)
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

	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(339)
			p.SchemaDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(340)
			p.TypeDefinition()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(341)
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

	p.SetState(346)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(344)
			p.SchemaExtension()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(345)
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
		p.SetState(348)
		p.Match(EGraphQLParserT__20)
	}
	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__19 {
		{
			p.SetState(349)
			p.Directives()
		}

	}
	{
		p.SetState(352)
		p.Match(EGraphQLParserT__3)
	}
	p.SetState(354)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<EGraphQLParserT__0)|(1<<EGraphQLParserT__1)|(1<<EGraphQLParserT__2))) != 0) {
		{
			p.SetState(353)
			p.RootOperationTypeDefinition()
		}

		p.SetState(356)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(358)
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
		p.SetState(360)
		p.OperationType()
	}
	{
		p.SetState(361)
		p.Match(EGraphQLParserT__7)
	}
	{
		p.SetState(362)
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

	p.SetState(380)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(364)
			p.Match(EGraphQLParserT__21)
		}
		{
			p.SetState(365)
			p.Match(EGraphQLParserT__20)
		}
		p.SetState(367)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EGraphQLParserT__19 {
			{
				p.SetState(366)
				p.Directives()
			}

		}
		{
			p.SetState(369)
			p.Match(EGraphQLParserT__3)
		}
		p.SetState(371)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<EGraphQLParserT__0)|(1<<EGraphQLParserT__1)|(1<<EGraphQLParserT__2))) != 0) {
			{
				p.SetState(370)
				p.OperationTypeDefinition()
			}

			p.SetState(373)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(375)
			p.Match(EGraphQLParserT__4)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(377)
			p.Match(EGraphQLParserT__21)
		}
		{
			p.SetState(378)
			p.Match(EGraphQLParserT__20)
		}
		{
			p.SetState(379)
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
		p.SetState(382)
		p.OperationType()
	}
	{
		p.SetState(383)
		p.Match(EGraphQLParserT__7)
	}
	{
		p.SetState(384)
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
		p.SetState(386)
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

	p.SetState(393)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(388)
			p.ScalarTypeDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(389)
			p.UnionTypeDefinition()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(390)
			p.EnumTypeDefinition()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(391)
			p.TempletableTypeDefinition()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(392)
			p.TemplateTypeDefinition()
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

	p.SetState(398)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(395)
			p.ObjectTypeDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(396)
			p.InterfaceTypeDefinition()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(397)
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

	p.SetState(406)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(400)
			p.ScalarTypeExtension()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(401)
			p.ObjectTypeExtension()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(402)
			p.InterfaceTypeExtension()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(403)
			p.UnionTypeExtension()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(404)
			p.EnumTypeExtension()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(405)
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
	p.SetState(409)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserSTRING || _la == EGraphQLParserBLOCK_STRING {
		{
			p.SetState(408)
			p.Description()
		}

	}
	{
		p.SetState(411)
		p.Match(EGraphQLParserT__22)
	}
	{
		p.SetState(412)
		p.Name()
	}
	p.SetState(414)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__19 {
		{
			p.SetState(413)
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
		p.SetState(416)
		p.Match(EGraphQLParserT__23)
	}
	{
		p.SetState(417)
		p.Match(EGraphQLParserT__22)
	}
	{
		p.SetState(418)
		p.Name()
	}
	{
		p.SetState(419)
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
	p.SetState(422)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserSTRING || _la == EGraphQLParserBLOCK_STRING {
		{
			p.SetState(421)
			p.Description()
		}

	}
	{
		p.SetState(424)
		p.Match(EGraphQLParserT__24)
	}
	{
		p.SetState(425)
		p.Name()
	}
	p.SetState(427)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__25 {
		{
			p.SetState(426)
			p.implementsInterfaces(0)
		}

	}
	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__19 {
		{
			p.SetState(429)
			p.Directives()
		}

	}
	p.SetState(433)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(432)
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
		p.SetState(436)
		p.Match(EGraphQLParserT__25)
	}
	p.SetState(438)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__26 {
		{
			p.SetState(437)
			p.Match(EGraphQLParserT__26)
		}

	}
	{
		p.SetState(440)
		p.NamedType()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(447)
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
			p.SetState(442)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(443)
				p.Match(EGraphQLParserT__26)
			}
			{
				p.SetState(444)
				p.NamedType()
			}

		}
		p.SetState(449)
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
		p.SetState(450)
		p.Match(EGraphQLParserT__3)
	}
	p.SetState(452)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-55)&-(0x1f+1)) == 0 && ((1<<uint((_la-55)))&((1<<(EGraphQLParserNAME-55))|(1<<(EGraphQLParserSTRING-55))|(1<<(EGraphQLParserBLOCK_STRING-55)))) != 0) {
		{
			p.SetState(451)
			p.FieldDefinition()
		}

		p.SetState(454)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(456)
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
	p.SetState(459)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserSTRING || _la == EGraphQLParserBLOCK_STRING {
		{
			p.SetState(458)
			p.Description()
		}

	}
	{
		p.SetState(461)
		p.Name()
	}
	p.SetState(463)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__5 {
		{
			p.SetState(462)
			p.ArgumentsDefinition()
		}

	}
	{
		p.SetState(465)
		p.Match(EGraphQLParserT__7)
	}
	{
		p.SetState(466)
		p.Type_()
	}
	p.SetState(468)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__19 {
		{
			p.SetState(467)
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
		p.SetState(470)
		p.Match(EGraphQLParserT__5)
	}
	p.SetState(472)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-55)&-(0x1f+1)) == 0 && ((1<<uint((_la-55)))&((1<<(EGraphQLParserNAME-55))|(1<<(EGraphQLParserSTRING-55))|(1<<(EGraphQLParserBLOCK_STRING-55)))) != 0) {
		{
			p.SetState(471)
			p.InputValueDefinition()
		}

		p.SetState(474)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(476)
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
	p.SetState(479)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserSTRING || _la == EGraphQLParserBLOCK_STRING {
		{
			p.SetState(478)
			p.Description()
		}

	}
	{
		p.SetState(481)
		p.Name()
	}
	{
		p.SetState(482)
		p.Match(EGraphQLParserT__7)
	}
	{
		p.SetState(483)
		p.Type_()
	}
	p.SetState(485)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__17 {
		{
			p.SetState(484)
			p.DefaultValue()
		}

	}
	p.SetState(488)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__19 {
		{
			p.SetState(487)
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

	p.SetState(514)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 58, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(490)
			p.Match(EGraphQLParserT__21)
		}
		{
			p.SetState(491)
			p.Match(EGraphQLParserT__24)
		}
		{
			p.SetState(492)
			p.Name()
		}
		p.SetState(494)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EGraphQLParserT__25 {
			{
				p.SetState(493)
				p.implementsInterfaces(0)
			}

		}
		p.SetState(497)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EGraphQLParserT__19 {
			{
				p.SetState(496)
				p.Directives()
			}

		}
		{
			p.SetState(499)
			p.FieldsDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(501)
			p.Match(EGraphQLParserT__21)
		}
		{
			p.SetState(502)
			p.Match(EGraphQLParserT__24)
		}
		{
			p.SetState(503)
			p.Name()
		}
		p.SetState(505)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EGraphQLParserT__25 {
			{
				p.SetState(504)
				p.implementsInterfaces(0)
			}

		}
		{
			p.SetState(507)
			p.Directives()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(509)
			p.Match(EGraphQLParserT__21)
		}
		{
			p.SetState(510)
			p.Match(EGraphQLParserT__24)
		}
		{
			p.SetState(511)
			p.Name()
		}
		{
			p.SetState(512)
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
	p.SetState(517)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserSTRING || _la == EGraphQLParserBLOCK_STRING {
		{
			p.SetState(516)
			p.Description()
		}

	}
	{
		p.SetState(519)
		p.Match(EGraphQLParserT__27)
	}
	{
		p.SetState(520)
		p.Name()
	}
	p.SetState(522)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__19 {
		{
			p.SetState(521)
			p.Directives()
		}

	}
	p.SetState(525)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 61, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(524)
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

	p.SetState(540)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 63, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(527)
			p.Match(EGraphQLParserT__21)
		}
		{
			p.SetState(528)
			p.Match(EGraphQLParserT__27)
		}
		{
			p.SetState(529)
			p.Name()
		}
		p.SetState(531)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EGraphQLParserT__19 {
			{
				p.SetState(530)
				p.Directives()
			}

		}
		{
			p.SetState(533)
			p.FieldsDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(535)
			p.Match(EGraphQLParserT__21)
		}
		{
			p.SetState(536)
			p.Match(EGraphQLParserT__27)
		}
		{
			p.SetState(537)
			p.Name()
		}
		{
			p.SetState(538)
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
	p.SetState(543)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserSTRING || _la == EGraphQLParserBLOCK_STRING {
		{
			p.SetState(542)
			p.Description()
		}

	}
	{
		p.SetState(545)
		p.Match(EGraphQLParserT__28)
	}
	{
		p.SetState(546)
		p.Name()
	}
	p.SetState(548)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__19 {
		{
			p.SetState(547)
			p.Directives()
		}

	}
	p.SetState(551)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__17 {
		{
			p.SetState(550)
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
		p.SetState(553)
		p.Match(EGraphQLParserT__17)
	}
	p.SetState(555)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__29 {
		{
			p.SetState(554)
			p.Match(EGraphQLParserT__29)
		}

	}
	{
		p.SetState(557)
		p.NamedType()
	}
	p.SetState(562)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == EGraphQLParserT__29 {
		{
			p.SetState(558)
			p.Match(EGraphQLParserT__29)
		}
		{
			p.SetState(559)
			p.NamedType()
		}

		p.SetState(564)
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

	p.SetState(578)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 70, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(565)
			p.Match(EGraphQLParserT__21)
		}
		{
			p.SetState(566)
			p.Match(EGraphQLParserT__28)
		}
		{
			p.SetState(567)
			p.Name()
		}
		p.SetState(569)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EGraphQLParserT__19 {
			{
				p.SetState(568)
				p.Directives()
			}

		}
		{
			p.SetState(571)
			p.UnionMemberTypes()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(573)
			p.Match(EGraphQLParserT__21)
		}
		{
			p.SetState(574)
			p.Match(EGraphQLParserT__28)
		}
		{
			p.SetState(575)
			p.Name()
		}
		{
			p.SetState(576)
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
	p.SetState(581)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserSTRING || _la == EGraphQLParserBLOCK_STRING {
		{
			p.SetState(580)
			p.Description()
		}

	}
	{
		p.SetState(583)
		p.Match(EGraphQLParserT__30)
	}
	{
		p.SetState(584)
		p.Name()
	}
	p.SetState(586)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__19 {
		{
			p.SetState(585)
			p.Directives()
		}

	}
	p.SetState(589)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 73, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(588)
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
		p.SetState(591)
		p.Match(EGraphQLParserT__3)
	}
	p.SetState(593)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-55)&-(0x1f+1)) == 0 && ((1<<uint((_la-55)))&((1<<(EGraphQLParserNAME-55))|(1<<(EGraphQLParserSTRING-55))|(1<<(EGraphQLParserBLOCK_STRING-55)))) != 0) {
		{
			p.SetState(592)
			p.EnumValueDefinition()
		}

		p.SetState(595)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(597)
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
	p.SetState(600)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserSTRING || _la == EGraphQLParserBLOCK_STRING {
		{
			p.SetState(599)
			p.Description()
		}

	}
	{
		p.SetState(602)
		p.EnumValue()
	}
	p.SetState(604)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__19 {
		{
			p.SetState(603)
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

	p.SetState(619)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 78, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(606)
			p.Match(EGraphQLParserT__21)
		}
		{
			p.SetState(607)
			p.Match(EGraphQLParserT__30)
		}
		{
			p.SetState(608)
			p.Name()
		}
		p.SetState(610)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EGraphQLParserT__19 {
			{
				p.SetState(609)
				p.Directives()
			}

		}
		{
			p.SetState(612)
			p.EnumValuesDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(614)
			p.Match(EGraphQLParserT__21)
		}
		{
			p.SetState(615)
			p.Match(EGraphQLParserT__30)
		}
		{
			p.SetState(616)
			p.Name()
		}
		{
			p.SetState(617)
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
	p.SetState(622)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserSTRING || _la == EGraphQLParserBLOCK_STRING {
		{
			p.SetState(621)
			p.Description()
		}

	}
	{
		p.SetState(624)
		p.Match(EGraphQLParserT__31)
	}
	{
		p.SetState(625)
		p.Name()
	}
	p.SetState(627)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__19 {
		{
			p.SetState(626)
			p.Directives()
		}

	}
	p.SetState(630)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 81, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(629)
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
		p.SetState(632)
		p.Match(EGraphQLParserT__3)
	}
	p.SetState(634)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-55)&-(0x1f+1)) == 0 && ((1<<uint((_la-55)))&((1<<(EGraphQLParserNAME-55))|(1<<(EGraphQLParserSTRING-55))|(1<<(EGraphQLParserBLOCK_STRING-55)))) != 0) {
		{
			p.SetState(633)
			p.InputValueDefinition()
		}

		p.SetState(636)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(638)
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

	p.SetState(653)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 84, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(640)
			p.Match(EGraphQLParserT__21)
		}
		{
			p.SetState(641)
			p.Match(EGraphQLParserT__31)
		}
		{
			p.SetState(642)
			p.Name()
		}
		p.SetState(644)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == EGraphQLParserT__19 {
			{
				p.SetState(643)
				p.Directives()
			}

		}
		{
			p.SetState(646)
			p.InputFieldsDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(648)
			p.Match(EGraphQLParserT__21)
		}
		{
			p.SetState(649)
			p.Match(EGraphQLParserT__31)
		}
		{
			p.SetState(650)
			p.Name()
		}
		{
			p.SetState(651)
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
	p.SetState(656)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserSTRING || _la == EGraphQLParserBLOCK_STRING {
		{
			p.SetState(655)
			p.Description()
		}

	}
	{
		p.SetState(658)
		p.Match(EGraphQLParserT__32)
	}
	{
		p.SetState(659)
		p.Match(EGraphQLParserT__19)
	}
	{
		p.SetState(660)
		p.Name()
	}
	p.SetState(662)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == EGraphQLParserT__5 {
		{
			p.SetState(661)
			p.ArgumentsDefinition()
		}

	}
	{
		p.SetState(664)
		p.Match(EGraphQLParserT__10)
	}
	{
		p.SetState(665)
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
		p.SetState(667)
		p.DirectiveLocation()
	}
	p.SetState(672)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == EGraphQLParserT__29 {
		{
			p.SetState(668)
			p.Match(EGraphQLParserT__29)
		}
		{
			p.SetState(669)
			p.DirectiveLocation()
		}

		p.SetState(674)
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

	p.SetState(677)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case EGraphQLParserT__33, EGraphQLParserT__34, EGraphQLParserT__35, EGraphQLParserT__36, EGraphQLParserT__37, EGraphQLParserT__38, EGraphQLParserT__39:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(675)
			p.ExecutableDirectiveLocation()
		}

	case EGraphQLParserT__40, EGraphQLParserT__41, EGraphQLParserT__42, EGraphQLParserT__43, EGraphQLParserT__44, EGraphQLParserT__45, EGraphQLParserT__46, EGraphQLParserT__47, EGraphQLParserT__48, EGraphQLParserT__49, EGraphQLParserT__50:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(676)
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
		p.SetState(679)
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
		p.SetState(681)
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
		p.SetState(683)
		p.Match(EGraphQLParserNAME)
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

func (s *TemplateTypeDefinitionContext) AllName() []INameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameContext)(nil)).Elem())
	var tst = make([]INameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameContext)
		}
	}

	return tst
}

func (s *TemplateTypeDefinitionContext) Name(i int) INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *TemplateTypeDefinitionContext) TempletableTypeDefinition() ITempletableTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITempletableTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITempletableTypeDefinitionContext)
}

func (s *TemplateTypeDefinitionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(EGraphQLParserCOMMA)
}

func (s *TemplateTypeDefinitionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(EGraphQLParserCOMMA, i)
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
	p.EnterRule(localctx, 144, EGraphQLParserRULE_templateTypeDefinition)
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
		p.SetState(685)
		p.Match(EGraphQLParserT__51)
	}
	{
		p.SetState(686)
		p.Match(EGraphQLParserT__52)
	}
	{
		p.SetState(687)
		p.Name()
	}
	p.SetState(692)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == EGraphQLParserCOMMA {
		{
			p.SetState(688)
			p.Match(EGraphQLParserCOMMA)
		}
		{
			p.SetState(689)
			p.Name()
		}

		p.SetState(694)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(695)
		p.Match(EGraphQLParserT__53)
	}
	{
		p.SetState(696)
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
