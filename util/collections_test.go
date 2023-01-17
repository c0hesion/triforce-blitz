package util

import "testing"

// 10,000 random integers between 1 and 100
var TestData = []int{
	50, 90, 47, 42, 80, 100, 59, 94, 69, 96, 31, 63, 60, 80, 83, 88, 32, 27, 91, 32, 89, 98, 68, 48, 56,
	65, 20, 58, 75, 48, 37, 3, 39, 96, 23, 68, 35, 83, 77, 87, 12, 49, 7, 84, 81, 59, 34, 46, 71, 65,
	81, 98, 55, 37, 94, 50, 32, 90, 98, 19, 2, 72, 68, 80, 62, 78, 96, 29, 23, 69, 98, 4, 22, 76, 47,
	61, 94, 12, 26, 43, 84, 79, 27, 96, 43, 12, 65, 8, 27, 2, 35, 9, 98, 43, 26, 1, 69, 22, 57, 80,
	70, 97, 42, 14, 96, 46, 97, 70, 91, 27, 23, 77, 2, 93, 10, 35, 59, 71, 38, 21, 18, 28, 55, 40, 96,
	61, 10, 99, 48, 55, 87, 35, 77, 87, 44, 17, 81, 7, 11, 89, 17, 73, 10, 81, 41, 84, 23, 69, 50, 1,
	78, 35, 67, 4, 15, 46, 74, 82, 77, 80, 31, 82, 42, 63, 58, 23, 37, 28, 80, 44, 75, 70, 89, 3, 71,
	31, 63, 78, 75, 76, 41, 55, 48, 32, 67, 47, 53, 79, 24, 16, 28, 60, 86, 27, 33, 20, 59, 91, 70, 43,
	30, 46, 39, 77, 35, 85, 36, 87, 78, 50, 63, 90, 28, 68, 23, 47, 29, 87, 68, 57, 74, 62, 93, 32, 42,
	21, 85, 84, 47, 23, 1, 30, 60, 55, 46, 38, 70, 97, 2, 27, 63, 94, 6, 45, 92, 83, 4, 41, 45, 54,
	56, 59, 7, 34, 66, 52, 92, 73, 44, 62, 6, 54, 55, 61, 19, 41, 21, 81, 55, 4, 5, 61, 62, 38, 80,
	13, 27, 79, 23, 50, 43, 66, 94, 22, 76, 20, 30, 78, 23, 16, 66, 90, 1, 24, 39, 38, 37, 74, 99, 54,
	65, 84, 49, 25, 55, 52, 40, 70, 14, 98, 69, 95, 48, 60, 56, 64, 82, 71, 60, 70, 67, 12, 92, 99, 70,
	72, 56, 91, 36, 49, 82, 80, 79, 75, 43, 94, 72, 53, 91, 87, 73, 63, 43, 89, 9, 13, 94, 59, 8, 84,
	42, 45, 17, 53, 69, 85, 24, 95, 3, 45, 22, 49, 62, 26, 90, 81, 51, 7, 39, 64, 16, 38, 85, 79, 100,
	60, 13, 92, 38, 97, 37, 23, 44, 38, 41, 55, 59, 68, 3, 91, 6, 12, 80, 48, 66, 60, 68, 54, 81, 75,
	10, 97, 84, 53, 2, 80, 92, 67, 48, 5, 98, 67, 31, 60, 100, 96, 90, 13, 20, 23, 96, 75, 69, 7, 54,
	88, 77, 72, 6, 27, 35, 60, 93, 84, 99, 84, 48, 99, 76, 83, 47, 88, 59, 31, 38, 71, 61, 48, 54, 87,
	51, 22, 41, 21, 13, 87, 9, 31, 81, 88, 42, 87, 85, 2, 23, 54, 100, 37, 31, 15, 19, 21, 31, 51, 22,
	98, 38, 89, 41, 100, 18, 15, 74, 7, 65, 9, 91, 76, 47, 70, 37, 46, 97, 86, 71, 49, 64, 39, 49, 25,
	41, 50, 35, 60, 9, 89, 97, 97, 60, 79, 40, 87, 66, 42, 95, 13, 69, 80, 24, 61, 95, 85, 9, 21, 24,
	37, 93, 37, 14, 98, 66, 53, 37, 20, 32, 95, 54, 28, 72, 51, 54, 16, 72, 4, 29, 53, 10, 75, 15, 36,
	65, 47, 82, 87, 57, 16, 8, 20, 39, 23, 33, 50, 21, 85, 60, 13, 75, 72, 1, 24, 51, 14, 79, 66, 2,
	94, 86, 20, 54, 4, 93, 8, 39, 66, 14, 97, 74, 47, 43, 11, 34, 29, 47, 45, 24, 25, 99, 41, 45, 18,
	21, 43, 22, 3, 100, 53, 24, 65, 31, 46, 9, 57, 88, 55, 47, 28, 75, 6, 81, 10, 66, 98, 96, 20, 64,
	70, 32, 7, 39, 28, 10, 54, 55, 33, 13, 60, 85, 15, 81, 65, 38, 64, 90, 56, 66, 60, 19, 98, 30, 47,
	34, 71, 77, 92, 81, 81, 96, 35, 82, 66, 91, 85, 19, 66, 77, 32, 40, 75, 65, 90, 17, 60, 81, 74, 84,
	55, 57, 36, 92, 88, 80, 84, 56, 57, 78, 96, 59, 35, 23, 29, 68, 18, 43, 65, 84, 90, 27, 5, 86, 26,
	13, 3, 58, 51, 26, 58, 37, 54, 39, 30, 30, 36, 67, 23, 6, 89, 23, 16, 17, 7, 1, 59, 34, 45, 49,
	78, 58, 19, 59, 83, 29, 32, 35, 68, 19, 39, 24, 23, 55, 66, 30, 54, 72, 2, 100, 60, 41, 36, 10, 18,
	12, 100, 42, 38, 23, 2, 40, 76, 99, 23, 41, 73, 41, 29, 86, 54, 33, 74, 11, 57, 89, 37, 38, 41, 62,
	11, 42, 40, 37, 21, 48, 9, 7, 57, 81, 82, 12, 78, 71, 93, 49, 58, 98, 37, 45, 36, 77, 51, 29, 15,
	61, 44, 95, 37, 51, 45, 37, 40, 36, 54, 13, 81, 48, 45, 76, 72, 75, 18, 18, 51, 58, 72, 5, 33, 93,
	57, 88, 17, 11, 73, 58, 27, 13, 23, 56, 76, 32, 26, 62, 53, 4, 4, 39, 78, 39, 1, 27, 44, 45, 23,
	27, 44, 46, 75, 59, 58, 12, 85, 22, 30, 16, 51, 91, 24, 85, 73, 88, 71, 71, 24, 57, 48, 95, 39, 45,
	77, 14, 98, 10, 6, 81, 76, 84, 7, 21, 13, 18, 45, 20, 24, 18, 10, 83, 39, 17, 96, 44, 59, 48, 38,
	5, 23, 19, 98, 99, 88, 92, 11, 11, 30, 9, 48, 96, 49, 29, 52, 23, 2, 63, 89, 11, 58, 31, 51, 77,
	9, 5, 91, 46, 67, 88, 89, 99, 56, 65, 6, 12, 75, 36, 77, 88, 47, 77, 34, 77, 30, 33, 39, 12, 58,
	24, 71, 11, 93, 4, 69, 50, 98, 31, 80, 54, 67, 46, 3, 41, 87, 52, 78, 34, 50, 20, 3, 19, 59, 4,
	36, 17, 47, 29, 6, 100, 85, 51, 80, 78, 99, 23, 95, 100, 72, 84, 64, 35, 73, 28, 99, 33, 57, 1, 2,
	100, 16, 3, 60, 62, 83, 39, 96, 43, 93, 30, 24, 34, 65, 23, 5, 9, 23, 91, 96, 24, 24, 60, 97, 75,
	36, 85, 90, 2, 62, 99, 18, 59, 46, 94, 76, 17, 91, 68, 15, 25, 22, 35, 89, 15, 90, 71, 32, 10, 40,
	78, 38, 97, 62, 77, 18, 42, 73, 39, 17, 80, 2, 15, 97, 94, 51, 75, 24, 70, 6, 98, 55, 100, 57, 36,
	91, 59, 54, 83, 58, 90, 23, 78, 78, 94, 42, 7, 29, 20, 20, 92, 5, 69, 30, 96, 67, 22, 86, 11, 51,
	76, 91, 58, 97, 73, 37, 89, 49, 11, 93, 11, 31, 25, 5, 1, 56, 100, 52, 87, 88, 40, 71, 39, 65, 58,
	59, 21, 62, 34, 84, 51, 2, 75, 50, 100, 94, 17, 44, 84, 34, 87, 5, 72, 59, 21, 32, 84, 43, 42, 20,
	4, 1, 87, 55, 94, 71, 45, 15, 60, 65, 92, 47, 1, 99, 7, 94, 66, 87, 22, 30, 97, 64, 34, 20, 5,
	59, 94, 20, 9, 94, 43, 39, 56, 35, 93, 87, 99, 57, 39, 61, 22, 85, 88, 61, 60, 71, 45, 25, 70, 6,
	94, 88, 17, 85, 31, 78, 7, 72, 22, 80, 18, 25, 44, 90, 81, 96, 82, 2, 71, 12, 41, 71, 78, 70, 65,
	22, 70, 19, 64, 39, 96, 90, 54, 65, 75, 73, 57, 43, 6, 89, 21, 12, 8, 36, 21, 61, 36, 18, 91, 96,
	86, 32, 45, 32, 36, 3, 49, 79, 38, 7, 2, 9, 30, 20, 75, 64, 30, 76, 50, 77, 16, 22, 28, 12, 100,
	23, 10, 70, 49, 21, 24, 38, 77, 100, 50, 39, 37, 94, 13, 95, 48, 86, 38, 44, 22, 11, 26, 23, 71, 90,
	72, 63, 70, 20, 5, 53, 41, 47, 80, 1, 73, 88, 5, 80, 66, 88, 73, 71, 27, 23, 24, 100, 31, 57, 43,
	50, 47, 51, 85, 46, 52, 33, 83, 12, 65, 59, 75, 13, 74, 95, 9, 76, 20, 49, 10, 5, 36, 22, 82, 36,
	59, 51, 45, 1, 52, 35, 41, 18, 72, 89, 85, 47, 19, 59, 57, 65, 13, 41, 24, 17, 39, 26, 3, 95, 15,
	13, 56, 91, 82, 41, 20, 60, 52, 45, 90, 54, 90, 17, 58, 47, 54, 49, 35, 93, 90, 26, 89, 12, 29, 76,
	64, 90, 66, 4, 46, 70, 14, 97, 91, 77, 64, 69, 8, 21, 99, 52, 43, 66, 28, 98, 9, 49, 56, 34, 46,
	18, 96, 3, 46, 19, 65, 55, 33, 81, 74, 56, 46, 41, 17, 33, 9, 23, 62, 68, 55, 48, 29, 36, 22, 64,
	87, 82, 57, 67, 84, 49, 37, 12, 41, 80, 30, 78, 73, 49, 25, 6, 65, 67, 42, 58, 31, 58, 57, 91, 98,
	14, 99, 95, 82, 18, 49, 91, 75, 55, 48, 8, 61, 66, 98, 51, 3, 39, 84, 82, 50, 9, 40, 72, 31, 94,
	46, 69, 35, 34, 67, 26, 45, 31, 11, 7, 87, 100, 73, 72, 100, 31, 98, 6, 83, 83, 17, 72, 53, 48, 32,
	17, 33, 64, 18, 59, 92, 94, 96, 36, 49, 15, 90, 63, 91, 51, 36, 18, 93, 87, 81, 43, 19, 61, 75, 3,
	78, 60, 75, 74, 37, 89, 61, 89, 29, 92, 85, 92, 80, 49, 4, 5, 45, 72, 40, 42, 3, 18, 11, 76, 37,
	68, 31, 84, 26, 50, 40, 54, 30, 83, 34, 71, 3, 68, 25, 82, 45, 32, 93, 1, 52, 72, 6, 81, 27, 8,
	82, 38, 82, 52, 38, 79, 69, 51, 9, 96, 73, 44, 86, 87, 54, 4, 54, 24, 10, 66, 25, 28, 12, 29, 66,
	58, 13, 34, 38, 44, 58, 43, 16, 24, 24, 80, 75, 79, 46, 29, 22, 83, 7, 7, 87, 16, 88, 95, 90, 47,
	21, 54, 21, 45, 88, 45, 4, 48, 11, 22, 91, 62, 52, 19, 15, 2, 96, 64, 57, 34, 85, 52, 32, 88, 65,
	80, 54, 99, 38, 32, 54, 82, 71, 12, 40, 37, 19, 85, 78, 70, 28, 5, 6, 23, 35, 52, 97, 14, 88, 34,
	44, 55, 11, 45, 51, 44, 37, 92, 10, 44, 98, 7, 1, 1, 31, 82, 14, 24, 50, 37, 16, 11, 36, 5, 11,
	91, 80, 69, 81, 40, 94, 76, 17, 69, 79, 11, 80, 18, 13, 94, 3, 55, 12, 45, 41, 47, 78, 47, 10, 70,
	13, 78, 11, 7, 26, 45, 70, 99, 54, 84, 16, 43, 61, 55, 56, 71, 41, 93, 67, 28, 71, 16, 85, 30, 58,
	46, 6, 82, 75, 53, 93, 27, 56, 2, 59, 71, 12, 21, 32, 54, 31, 5, 46, 62, 1, 22, 22, 7, 42, 97,
	37, 28, 67, 46, 24, 13, 49, 56, 83, 92, 44, 98, 5, 55, 13, 51, 45, 69, 23, 58, 35, 69, 62, 50, 86,
	77, 63, 73, 7, 69, 78, 13, 19, 23, 30, 74, 27, 51, 91, 22, 12, 67, 96, 41, 14, 73, 46, 46, 81, 16,
	44, 9, 34, 16, 46, 87, 91, 89, 80, 46, 80, 44, 26, 36, 60, 82, 50, 29, 6, 48, 8, 3, 17, 26, 59,
	88, 60, 71, 6, 82, 98, 5, 71, 65, 92, 61, 79, 42, 45, 21, 34, 56, 94, 36, 32, 28, 69, 12, 83, 2,
	75, 78, 32, 49, 5, 85, 31, 30, 92, 80, 39, 57, 93, 8, 59, 7, 13, 88, 100, 51, 25, 75, 57, 45, 61,
	22, 48, 24, 30, 35, 11, 56, 83, 76, 22, 53, 12, 64, 81, 92, 26, 45, 11, 90, 99, 55, 47, 4, 71, 47,
	97, 16, 54, 63, 30, 19, 65, 64, 55, 20, 26, 13, 23, 44, 62, 74, 9, 4, 40, 92, 86, 89, 16, 63, 74,
	37, 76, 8, 22, 94, 20, 61, 93, 20, 11, 81, 26, 78, 10, 49, 76, 24, 20, 46, 70, 7, 42, 65, 96, 7,
	11, 33, 77, 49, 48, 44, 89, 6, 27, 9, 56, 3, 34, 53, 25, 18, 47, 3, 63, 73, 38, 55, 90, 68, 66,
	52, 35, 56, 98, 2, 98, 10, 43, 62, 94, 94, 34, 37, 63, 10, 74, 85, 26, 51, 90, 62, 9, 84, 1, 76,
	60, 34, 78, 84, 59, 94, 20, 15, 65, 78, 94, 64, 75, 92, 38, 49, 95, 8, 27, 20, 61, 35, 17, 39, 27,
	55, 21, 30, 30, 87, 93, 49, 57, 15, 81, 77, 46, 27, 84, 96, 86, 12, 67, 27, 26, 3, 96, 64, 11, 21,
	6, 87, 100, 42, 65, 34, 91, 90, 96, 39, 74, 77, 98, 94, 4, 80, 68, 66, 55, 12, 46, 13, 28, 52, 24,
	15, 63, 49, 83, 35, 61, 21, 66, 25, 58, 63, 70, 21, 97, 22, 69, 78, 88, 52, 30, 79, 35, 100, 56, 79,
	82, 55, 42, 34, 99, 2, 14, 20, 78, 30, 83, 7, 78, 67, 96, 28, 34, 74, 78, 41, 68, 41, 68, 87, 84,
	10, 12, 71, 18, 56, 11, 64, 75, 13, 8, 23, 21, 5, 77, 29, 46, 40, 79, 51, 57, 87, 17, 35, 55, 92,
	54, 11, 45, 88, 63, 51, 95, 20, 78, 32, 98, 37, 12, 17, 88, 4, 23, 91, 98, 51, 53, 61, 34, 48, 50,
	66, 25, 55, 46, 98, 65, 33, 81, 37, 19, 39, 26, 15, 61, 97, 71, 29, 57, 100, 99, 71, 40, 63, 65, 75,
	92, 95, 4, 71, 39, 16, 12, 47, 17, 72, 22, 96, 46, 28, 86, 8, 39, 98, 39, 15, 47, 77, 25, 87, 86,
	20, 59, 10, 99, 82, 66, 19, 42, 73, 94, 83, 74, 52, 11, 28, 32, 49, 10, 74, 78, 59, 21, 61, 38, 96,
	26, 76, 75, 25, 23, 97, 76, 36, 61, 52, 20, 16, 2, 8, 40, 37, 63, 66, 66, 54, 66, 17, 8, 38, 98,
	36, 42, 79, 10, 19, 15, 80, 17, 16, 2, 19, 61, 10, 18, 20, 26, 58, 52, 3, 86, 80, 80, 33, 61, 11,
	43, 60, 27, 79, 58, 60, 22, 71, 97, 31, 9, 96, 37, 73, 26, 69, 65, 16, 62, 98, 92, 54, 76, 60, 84,
	12, 94, 77, 79, 70, 76, 97, 26, 12, 8, 75, 53, 14, 4, 73, 20, 37, 100, 20, 59, 52, 27, 54, 82, 7,
	15, 85, 71, 41, 39, 21, 72, 96, 72, 98, 6, 81, 2, 14, 5, 20, 19, 85, 37, 100, 56, 31, 95, 84, 37,
	31, 1, 59, 2, 49, 50, 62, 94, 84, 1, 88, 28, 81, 78, 29, 62, 41, 22, 94, 77, 80, 8, 14, 9, 83,
	48, 57, 65, 21, 16, 23, 90, 70, 86, 6, 50, 75, 49, 31, 28, 53, 49, 96, 40, 69, 86, 64, 37, 4, 2,
	62, 76, 16, 99, 79, 57, 51, 24, 63, 13, 63, 84, 20, 30, 12, 36, 26, 59, 81, 62, 65, 10, 77, 14, 42,
	30, 77, 14, 82, 8, 78, 87, 52, 32, 2, 67, 25, 13, 66, 99, 11, 38, 24, 96, 31, 64, 24, 16, 12, 45,
	69, 99, 29, 56, 5, 68, 72, 71, 40, 12, 62, 15, 37, 95, 2, 52, 16, 58, 31, 3, 15, 63, 51, 29, 91,
	82, 71, 61, 84, 32, 68, 84, 33, 4, 69, 96, 15, 87, 18, 30, 18, 23, 44, 6, 71, 64, 82, 81, 39, 23,
	15, 43, 80, 70, 50, 6, 99, 48, 21, 15, 10, 7, 95, 89, 95, 45, 34, 68, 28, 68, 43, 33, 32, 3, 32,
	25, 89, 62, 57, 59, 8, 26, 78, 44, 99, 49, 36, 63, 9, 23, 75, 87, 51, 57, 49, 79, 36, 73, 64, 1,
	42, 30, 84, 99, 100, 20, 17, 12, 85, 94, 39, 49, 64, 75, 38, 44, 84, 53, 24, 99, 74, 57, 51, 99, 8,
	41, 88, 88, 51, 22, 80, 29, 9, 32, 93, 90, 75, 75, 37, 4, 59, 78, 3, 10, 76, 95, 62, 14, 25, 57,
	1, 32, 55, 71, 3, 61, 12, 47, 67, 25, 3, 90, 88, 65, 67, 39, 68, 63, 50, 84, 75, 52, 87, 13, 8,
	57, 61, 2, 72, 2, 38, 23, 92, 54, 33, 55, 96, 29, 77, 74, 34, 3, 89, 57, 23, 47, 37, 53, 34, 98,
	37, 80, 7, 47, 24, 48, 55, 80, 23, 55, 7, 25, 35, 18, 66, 98, 100, 35, 24, 21, 23, 8, 48, 86, 58,
	22, 81, 37, 95, 11, 34, 26, 81, 68, 11, 40, 66, 14, 12, 27, 85, 89, 10, 77, 42, 53, 84, 31, 67, 50,
	96, 25, 36, 72, 65, 13, 29, 13, 4, 42, 12, 49, 32, 30, 53, 59, 94, 77, 16, 50, 86, 27, 82, 97, 3,
	67, 67, 50, 24, 2, 95, 11, 61, 61, 92, 84, 53, 39, 10, 31, 45, 23, 52, 19, 88, 26, 45, 59, 83, 20,
	16, 16, 29, 63, 30, 43, 87, 27, 9, 83, 97, 31, 31, 82, 36, 37, 69, 32, 29, 86, 88, 70, 37, 15, 35,
	36, 88, 84, 89, 43, 86, 59, 66, 88, 25, 71, 39, 90, 50, 28, 86, 34, 81, 39, 53, 37, 97, 9, 5, 32,
	78, 83, 99, 15, 11, 95, 67, 95, 81, 28, 70, 98, 29, 94, 34, 81, 13, 45, 96, 99, 31, 50, 36, 24, 84,
	43, 63, 44, 75, 10, 65, 31, 7, 16, 44, 14, 24, 40, 3, 52, 98, 83, 65, 79, 31, 18, 31, 39, 100, 84,
	21, 87, 95, 96, 10, 52, 59, 49, 90, 15, 40, 28, 53, 44, 84, 26, 35, 43, 91, 16, 91, 3, 33, 24, 40,
	67, 45, 58, 62, 50, 23, 1, 43, 99, 13, 16, 4, 91, 15, 29, 19, 78, 66, 40, 50, 80, 82, 68, 99, 75,
	2, 39, 86, 49, 72, 74, 45, 93, 48, 70, 41, 70, 97, 46, 7, 79, 63, 42, 88, 86, 99, 4, 29, 54, 29,
	15, 69, 77, 87, 59, 2, 24, 33, 61, 4, 36, 45, 86, 13, 1, 37, 88, 71, 90, 23, 20, 8, 74, 33, 43,
	4, 40, 9, 34, 10, 78, 69, 24, 72, 12, 61, 8, 56, 91, 84, 100, 88, 80, 98, 3, 97, 59, 94, 6, 92,
	74, 93, 22, 40, 44, 77, 4, 36, 95, 82, 90, 28, 98, 88, 66, 43, 64, 87, 38, 16, 24, 16, 64, 57, 41,
	34, 62, 2, 7, 76, 3, 30, 38, 89, 17, 15, 86, 8, 89, 73, 23, 48, 63, 12, 80, 58, 84, 60, 16, 53,
	66, 8, 23, 96, 9, 3, 99, 78, 5, 32, 72, 5, 14, 60, 49, 9, 93, 92, 1, 57, 33, 55, 32, 15, 10,
	68, 82, 6, 10, 25, 78, 85, 91, 60, 100, 94, 81, 13, 40, 63, 96, 43, 92, 29, 12, 51, 69, 66, 16, 71,
	84, 90, 69, 88, 97, 31, 26, 9, 31, 76, 18, 54, 79, 71, 84, 67, 74, 6, 22, 51, 11, 34, 24, 40, 64,
	57, 95, 88, 77, 80, 75, 54, 68, 16, 45, 4, 7, 84, 50, 76, 59, 36, 45, 28, 77, 68, 80, 96, 40, 17,
	75, 70, 18, 80, 85, 97, 56, 62, 24, 77, 72, 2, 13, 17, 72, 23, 98, 63, 15, 47, 5, 99, 37, 62, 38,
	41, 60, 99, 43, 99, 50, 64, 62, 95, 3, 72, 14, 11, 36, 86, 44, 16, 41, 9, 69, 27, 75, 96, 77, 4,
	66, 84, 81, 15, 36, 50, 39, 7, 78, 61, 25, 80, 37, 24, 19, 87, 63, 55, 8, 28, 17, 55, 97, 49, 90,
	40, 16, 44, 89, 52, 19, 31, 30, 69, 70, 12, 38, 24, 43, 90, 9, 12, 27, 28, 35, 41, 8, 61, 61, 50,
	6, 46, 53, 7, 99, 20, 99, 59, 8, 52, 44, 27, 92, 17, 1, 5, 29, 9, 43, 4, 66, 90, 9, 2, 92,
	34, 99, 51, 69, 80, 78, 94, 88, 91, 1, 51, 77, 36, 99, 92, 35, 24, 86, 56, 59, 33, 90, 4, 24, 17,
	16, 38, 98, 36, 31, 53, 88, 12, 87, 50, 78, 38, 75, 82, 50, 14, 44, 25, 40, 82, 22, 33, 14, 92, 81,
	63, 77, 54, 68, 60, 45, 79, 59, 69, 75, 52, 57, 41, 55, 79, 85, 43, 39, 1, 77, 2, 79, 35, 30, 98,
	8, 9, 4, 91, 88, 69, 66, 35, 100, 44, 40, 51, 7, 91, 81, 93, 45, 59, 51, 9, 50, 75, 29, 87, 23,
	92, 11, 18, 77, 41, 53, 50, 34, 10, 63, 67, 21, 95, 94, 96, 33, 29, 79, 88, 47, 93, 48, 91, 20, 31,
	9, 44, 33, 51, 67, 75, 48, 84, 43, 64, 62, 79, 63, 82, 89, 13, 3, 64, 34, 43, 91, 62, 58, 29, 61,
	17, 63, 72, 77, 27, 9, 20, 53, 38, 27, 57, 67, 59, 64, 32, 13, 76, 60, 77, 63, 11, 25, 88, 12, 55,
	2, 24, 70, 56, 82, 65, 31, 65, 83, 42, 84, 85, 35, 39, 50, 2, 50, 32, 100, 84, 32, 94, 42, 95, 75,
	84, 24, 18, 2, 60, 57, 88, 70, 30, 81, 93, 7, 18, 73, 31, 70, 45, 71, 75, 4, 84, 58, 43, 63, 1,
	18, 6, 23, 61, 78, 64, 30, 98, 23, 41, 10, 10, 6, 66, 27, 96, 58, 77, 100, 32, 56, 42, 91, 52, 74,
	17, 43, 14, 90, 11, 74, 100, 24, 68, 90, 6, 17, 41, 12, 34, 82, 58, 79, 20, 74, 31, 33, 84, 86, 7,
	70, 91, 20, 95, 51, 99, 78, 52, 36, 78, 70, 74, 91, 17, 80, 43, 89, 25, 56, 56, 23, 17, 23, 80, 40,
	84, 48, 64, 93, 46, 100, 16, 34, 50, 5, 74, 73, 55, 39, 41, 84, 77, 73, 33, 61, 85, 49, 71, 79, 15,
	40, 77, 40, 75, 38, 39, 85, 54, 52, 26, 12, 76, 42, 31, 39, 30, 65, 40, 95, 4, 36, 38, 62, 61, 45,
	73, 96, 3, 34, 24, 19, 3, 66, 35, 19, 53, 65, 10, 44, 84, 31, 67, 27, 95, 57, 2, 9, 33, 14, 6,
	3, 17, 15, 91, 85, 48, 21, 9, 98, 5, 68, 20, 29, 71, 18, 10, 93, 67, 55, 84, 47, 4, 79, 86, 100,
	7, 9, 3, 64, 88, 37, 54, 39, 17, 32, 81, 31, 59, 28, 90, 53, 100, 52, 3, 90, 56, 52, 77, 18, 63,
	83, 67, 58, 85, 86, 7, 61, 62, 1, 82, 77, 42, 83, 67, 33, 96, 83, 41, 49, 61, 94, 37, 72, 73, 30,
	22, 12, 100, 15, 13, 31, 4, 35, 89, 74, 44, 51, 98, 89, 94, 9, 14, 31, 12, 69, 65, 74, 13, 42, 23,
	24, 63, 86, 25, 13, 66, 47, 19, 46, 87, 13, 2, 48, 100, 93, 84, 100, 28, 13, 77, 70, 36, 72, 10, 61,
	96, 10, 89, 65, 37, 24, 50, 31, 20, 49, 8, 100, 100, 7, 33, 89, 46, 26, 41, 9, 1, 49, 47, 92, 69,
	16, 90, 11, 95, 54, 2, 65, 80, 39, 59, 15, 9, 81, 45, 42, 9, 8, 95, 49, 19, 22, 75, 78, 74, 93,
	42, 17, 64, 57, 44, 1, 44, 54, 96, 4, 36, 62, 64, 14, 19, 61, 27, 76, 91, 77, 51, 63, 79, 50, 79,
	24, 26, 5, 13, 53, 3, 43, 12, 28, 56, 61, 70, 84, 96, 36, 32, 39, 95, 66, 44, 19, 37, 51, 77, 78,
	56, 93, 80, 13, 93, 12, 56, 54, 77, 21, 99, 55, 37, 86, 10, 79, 68, 34, 91, 27, 27, 63, 38, 78, 64,
	75, 93, 72, 76, 33, 71, 88, 48, 92, 44, 57, 63, 66, 41, 31, 30, 95, 1, 36, 45, 99, 99, 44, 14, 1,
	9, 13, 44, 19, 54, 10, 34, 52, 71, 63, 100, 74, 12, 98, 68, 62, 11, 21, 53, 7, 47, 5, 52, 78, 57,
	21, 46, 87, 24, 10, 54, 44, 58, 79, 45, 5, 57, 19, 99, 38, 16, 35, 64, 94, 64, 16, 87, 90, 40, 3,
	39, 26, 62, 38, 93, 69, 91, 64, 76, 67, 37, 85, 34, 72, 22, 35, 94, 96, 25, 10, 28, 27, 13, 75, 42,
	5, 46, 61, 71, 56, 12, 92, 16, 82, 71, 92, 9, 5, 56, 65, 56, 41, 24, 53, 1, 31, 88, 91, 77, 66,
	7, 66, 32, 83, 70, 83, 16, 69, 61, 39, 36, 45, 89, 22, 60, 64, 98, 53, 2, 84, 4, 7, 94, 15, 77,
	44, 19, 58, 71, 54, 75, 3, 17, 45, 22, 67, 6, 30, 11, 71, 68, 72, 42, 88, 96, 78, 88, 50, 65, 60,
	23, 77, 76, 2, 43, 95, 55, 10, 2, 31, 83, 23, 33, 53, 89, 86, 89, 72, 88, 75, 65, 91, 71, 88, 95,
	43, 19, 25, 49, 54, 23, 48, 27, 7, 30, 79, 5, 99, 61, 12, 21, 73, 60, 2, 5, 45, 29, 40, 30, 47,
	63, 79, 43, 82, 61, 55, 21, 25, 81, 16, 59, 95, 14, 42, 26, 48, 16, 11, 12, 4, 98, 52, 88, 1, 53,
	83, 92, 22, 55, 36, 19, 22, 3, 58, 99, 89, 36, 74, 32, 79, 77, 100, 82, 6, 64, 95, 15, 86, 11, 11,
	64, 9, 61, 72, 41, 18, 43, 41, 10, 32, 81, 28, 6, 31, 90, 48, 68, 5, 53, 42, 95, 37, 37, 17, 8,
	36, 63, 91, 69, 10, 50, 8, 20, 77, 74, 100, 50, 79, 53, 42, 18, 63, 64, 57, 79, 33, 87, 95, 4, 33,
	96, 89, 80, 93, 85, 22, 33, 54, 8, 13, 94, 98, 15, 54, 27, 100, 45, 1, 54, 97, 24, 36, 49, 90, 86,
	23, 3, 69, 27, 88, 78, 60, 23, 36, 74, 4, 17, 23, 14, 74, 79, 62, 50, 12, 65, 78, 1, 74, 80, 2,
	4, 48, 93, 68, 61, 36, 93, 39, 24, 84, 79, 55, 54, 18, 78, 59, 79, 67, 54, 98, 39, 12, 38, 5, 45,
	39, 91, 21, 76, 82, 23, 56, 13, 8, 52, 94, 20, 75, 20, 52, 61, 77, 20, 30, 37, 75, 32, 73, 33, 18,
	22, 89, 29, 3, 81, 52, 79, 57, 13, 100, 40, 72, 54, 94, 56, 25, 32, 54, 36, 91, 89, 9, 20, 30, 2,
	41, 57, 3, 80, 30, 65, 92, 14, 33, 34, 77, 22, 97, 16, 33, 3, 2, 52, 31, 76, 19, 84, 59, 57, 15,
	8, 95, 65, 51, 36, 2, 14, 22, 42, 70, 22, 22, 96, 1, 51, 17, 5, 38, 40, 63, 19, 91, 62, 12, 38,
	88, 73, 88, 41, 93, 54, 1, 59, 90, 67, 24, 73, 9, 3, 56, 39, 96, 79, 3, 86, 7, 56, 74, 73, 9,
	69, 15, 74, 42, 17, 38, 10, 7, 37, 39, 26, 59, 25, 70, 15, 92, 19, 60, 62, 6, 59, 80, 7, 14, 43,
	13, 1, 97, 90, 34, 63, 87, 14, 79, 58, 18, 75, 64, 68, 58, 56, 6, 88, 33, 49, 35, 16, 67, 34, 5,
	96, 88, 58, 3, 49, 99, 99, 33, 5, 23, 59, 20, 88, 68, 12, 16, 7, 66, 21, 52, 95, 17, 57, 9, 43,
	30, 49, 81, 16, 14, 58, 7, 45, 10, 43, 79, 65, 60, 34, 98, 18, 77, 89, 35, 10, 85, 49, 43, 81, 41,
	87, 24, 10, 55, 23, 12, 61, 91, 74, 80, 22, 39, 79, 21, 56, 60, 52, 98, 55, 73, 12, 35, 55, 2, 3,
	19, 4, 49, 35, 64, 77, 87, 17, 20, 99, 59, 13, 66, 88, 90, 61, 39, 98, 20, 90, 36, 75, 65, 83, 14,
	21, 48, 60, 35, 81, 65, 93, 41, 31, 8, 93, 86, 29, 54, 39, 18, 5, 82, 63, 74, 8, 99, 33, 38, 45,
	89, 25, 94, 39, 13, 85, 49, 21, 78, 78, 45, 27, 15, 75, 30, 1, 58, 35, 57, 68, 32, 8, 54, 13, 29,
	88, 85, 29, 45, 90, 56, 72, 28, 53, 50, 34, 97, 76, 37, 87, 41, 65, 96, 46, 64, 99, 63, 97, 98, 58,
	48, 37, 75, 97, 17, 16, 56, 39, 72, 52, 82, 43, 84, 22, 85, 10, 76, 54, 18, 72, 9, 13, 2, 99, 26,
	35, 93, 83, 97, 4, 2, 66, 66, 10, 70, 71, 72, 82, 83, 6, 11, 34, 13, 20, 64, 6, 74, 18, 4, 15,
	69, 100, 50, 25, 64, 39, 52, 41, 33, 86, 62, 78, 25, 67, 45, 23, 27, 31, 36, 84, 62, 12, 94, 25, 91,
	41, 53, 72, 7, 33, 13, 4, 21, 92, 48, 40, 13, 42, 67, 20, 47, 42, 14, 22, 69, 33, 87, 87, 1, 22,
	27, 16, 88, 31, 11, 21, 79, 10, 20, 89, 72, 78, 89, 54, 43, 5, 99, 83, 29, 40, 15, 75, 81, 36, 23,
	70, 19, 20, 86, 49, 35, 78, 73, 41, 40, 91, 98, 85, 34, 90, 69, 38, 48, 27, 22, 6, 18, 38, 47, 23,
	47, 36, 5, 52, 54, 30, 10, 95, 53, 32, 48, 21, 43, 25, 24, 51, 49, 34, 18, 38, 32, 33, 54, 93, 42,
	19, 24, 90, 24, 7, 79, 31, 41, 95, 70, 46, 12, 6, 4, 90, 39, 53, 27, 11, 4, 41, 53, 45, 73, 83,
	88, 49, 96, 59, 97, 78, 12, 51, 76, 52, 61, 33, 11, 55, 81, 20, 21, 90, 70, 64, 97, 86, 56, 81, 76,
	47, 30, 80, 64, 8, 78, 45, 86, 6, 46, 39, 19, 34, 13, 80, 12, 32, 4, 33, 61, 38, 94, 100, 30, 13,
	16, 33, 93, 65, 26, 96, 50, 54, 6, 89, 58, 23, 53, 88, 62, 39, 63, 36, 43, 56, 43, 54, 93, 94, 79,
	50, 80, 5, 61, 7, 72, 6, 91, 48, 9, 55, 37, 88, 90, 6, 66, 52, 85, 90, 62, 68, 15, 38, 20, 54,
	1, 56, 36, 24, 22, 96, 27, 27, 45, 3, 57, 88, 17, 39, 56, 58, 42, 26, 48, 54, 92, 22, 5, 2, 8,
	37, 16, 98, 13, 90, 64, 26, 11, 95, 97, 47, 22, 6, 22, 56, 23, 14, 68, 54, 65, 21, 100, 92, 39, 10,
	82, 6, 56, 8, 82, 30, 2, 26, 90, 17, 62, 1, 98, 9, 73, 88, 88, 70, 52, 19, 52, 41, 30, 85, 26,
	18, 8, 19, 22, 68, 64, 59, 5, 46, 63, 56, 87, 18, 55, 9, 23, 79, 61, 56, 86, 6, 81, 78, 7, 2,
	29, 94, 8, 77, 96, 89, 83, 73, 9, 65, 6, 86, 52, 82, 6, 84, 86, 86, 43, 2, 22, 98, 51, 45, 61,
	79, 85, 60, 76, 23, 64, 48, 75, 43, 20, 24, 8, 98, 56, 30, 3, 86, 71, 64, 92, 41, 7, 16, 94, 79,
	11, 40, 49, 97, 96, 60, 38, 51, 3, 24, 25, 91, 42, 11, 17, 98, 30, 42, 55, 82, 71, 19, 99, 58, 59,
	25, 12, 6, 40, 85, 59, 21, 61, 54, 64, 70, 72, 93, 66, 49, 46, 65, 29, 97, 24, 72, 7, 33, 10, 50,
	71, 61, 79, 43, 63, 58, 82, 33, 48, 77, 28, 57, 96, 6, 35, 59, 80, 1, 42, 8, 67, 14, 51, 74, 57,
	91, 29, 8, 97, 59, 64, 99, 26, 13, 23, 50, 77, 84, 36, 22, 39, 58, 5, 95, 27, 89, 82, 58, 28, 73,
	44, 44, 28, 98, 13, 56, 100, 26, 47, 15, 43, 58, 100, 53, 61, 74, 82, 48, 72, 82, 20, 30, 72, 9, 9,
	55, 61, 14, 97, 24, 84, 34, 8, 6, 27, 45, 75, 61, 75, 81, 44, 17, 55, 18, 64, 99, 24, 15, 85, 53,
	93, 37, 30, 47, 10, 83, 24, 95, 8, 44, 16, 88, 1, 9, 24, 9, 77, 62, 11, 9, 16, 11, 42, 78, 58,
	17, 78, 96, 61, 28, 35, 19, 98, 20, 31, 56, 72, 7, 17, 99, 83, 42, 22, 56, 35, 92, 84, 96, 10, 97,
	99, 76, 91, 74, 68, 62, 4, 14, 49, 36, 60, 47, 16, 3, 57, 16, 76, 5, 36, 22, 75, 31, 18, 42, 16,
	98, 27, 93, 83, 77, 76, 75, 93, 1, 67, 2, 87, 67, 22, 27, 37, 48, 34, 32, 4, 44, 30, 91, 30, 2,
	25, 79, 77, 3, 35, 91, 31, 44, 8, 84, 6, 82, 28, 60, 86, 55, 41, 53, 100, 75, 70, 99, 64, 29, 39,
	88, 64, 98, 75, 3, 37, 47, 32, 61, 77, 27, 50, 70, 52, 79, 24, 65, 31, 67, 56, 30, 44, 99, 93, 89,
	6, 2, 30, 10, 33, 68, 78, 36, 49, 73, 23, 3, 9, 84, 35, 98, 28, 15, 41, 100, 71, 26, 22, 35, 45,
	97, 63, 99, 36, 25, 71, 84, 66, 15, 17, 98, 47, 53, 85, 34, 3, 45, 30, 20, 31, 85, 49, 19, 50, 10,
	70, 14, 30, 17, 50, 24, 19, 1, 53, 8, 34, 53, 31, 83, 86, 24, 41, 90, 98, 38, 80, 83, 41, 88, 88,
	62, 56, 53, 9, 57, 80, 61, 49, 33, 58, 5, 67, 10, 34, 5, 65, 52, 33, 55, 47, 77, 68, 68, 3, 74,
	95, 60, 85, 72, 70, 78, 96, 8, 78, 60, 10, 30, 44, 95, 14, 90, 47, 38, 99, 88, 55, 77, 69, 74, 22,
	63, 79, 99, 22, 79, 3, 32, 90, 87, 19, 15, 63, 64, 47, 28, 18, 54, 65, 45, 24, 97, 42, 29, 57, 53,
	98, 16, 15, 8, 52, 43, 74, 14, 67, 34, 41, 43, 24, 21, 75, 58, 59, 76, 58, 96, 72, 56, 77, 5, 5,
	31, 92, 95, 63, 75, 43, 47, 77, 7, 54, 62, 31, 40, 91, 75, 67, 43, 71, 45, 33, 47, 28, 52, 29, 54,
	67, 1, 57, 70, 98, 92, 97, 17, 28, 23, 84, 64, 65, 18, 90, 15, 92, 47, 5, 7, 38, 76, 37, 87, 64,
	65, 41, 56, 25, 36, 1, 75, 23, 29, 41, 63, 52, 13, 91, 68, 76, 64, 90, 84, 26, 62, 52, 59, 19, 12,
	55, 90, 84, 50, 30, 68, 60, 55, 12, 54, 64, 60, 98, 25, 11, 21, 86, 45, 97, 91, 32, 56, 59, 54, 54,
	24, 73, 31, 39, 13, 77, 4, 1, 71, 9, 65, 79, 94, 70, 28, 93, 38, 25, 40, 43, 62, 7, 26, 98, 39,
	50, 94, 67, 34, 83, 58, 8, 54, 31, 60, 78, 27, 21, 84, 22, 86, 87, 16, 31, 4, 35, 43, 54, 72, 2,
	61, 79, 72, 67, 9, 78, 87, 4, 51, 21, 83, 36, 87, 74, 4, 61, 89, 96, 75, 10, 19, 5, 37, 8, 28,
	44, 11, 36, 57, 10, 38, 22, 69, 13, 40, 63, 44, 27, 17, 84, 84, 31, 60, 6, 13, 75, 38, 44, 5, 53,
	93, 45, 76, 57, 79, 1, 72, 46, 77, 38, 53, 17, 24, 43, 51, 75, 24, 3, 82, 47, 85, 43, 3, 70, 30,
	19, 30, 23, 47, 52, 6, 55, 45, 52, 90, 49, 69, 99, 84, 23, 47, 7, 40, 41, 49, 56, 38, 35, 99, 7,
	54, 70, 17, 71, 83, 63, 34, 90, 55, 37, 87, 34, 83, 18, 78, 1, 74, 79, 59, 51, 96, 19, 36, 80, 7,
	48, 70, 95, 57, 40, 53, 33, 85, 42, 94, 46, 23, 34, 96, 50, 24, 73, 60, 66, 9, 31, 19, 9, 10, 55,
	6, 72, 3, 45, 8, 65, 38, 89, 30, 79, 22, 82, 59, 62, 22, 13, 26, 4, 78, 4, 46, 33, 54, 15, 24,
	76, 70, 93, 71, 100, 47, 50, 35, 61, 2, 58, 69, 83, 91, 44, 50, 1, 98, 46, 100, 10, 18, 69, 54, 24,
	30, 30, 76, 79, 36, 47, 14, 17, 71, 64, 40, 9, 75, 81, 85, 25, 74, 75, 84, 79, 27, 59, 15, 100, 65,
	37, 23, 9, 60, 27, 63, 49, 74, 99, 60, 16, 40, 7, 57, 48, 99, 73, 98, 80, 19, 89, 60, 49, 17, 62,
	83, 98, 96, 33, 28, 56, 85, 76, 88, 86, 19, 87, 7, 8, 14, 57, 92, 36, 99, 82, 88, 66, 56, 46, 26,
	31, 35, 33, 41, 70, 95, 21, 44, 64, 26, 26, 38, 26, 90, 40, 15, 70, 67, 29, 70, 72, 57, 69, 51, 94,
	37, 93, 88, 87, 12, 81, 91, 83, 77, 92, 58, 76, 4, 12, 49, 96, 96, 8, 88, 16, 48, 82, 64, 63, 31,
	4, 33, 97, 2, 53, 34, 52, 40, 31, 43, 92, 36, 67, 17, 38, 81, 95, 16, 58, 100, 93, 7, 60, 20, 92,
	21, 28, 6, 12, 89, 50, 52, 19, 95, 22, 76, 9, 33, 17, 44, 26, 96, 74, 97, 9, 56, 11, 52, 31, 71,
	4, 33, 48, 40, 22, 68, 37, 21, 98, 43, 17, 63, 32, 97, 37, 60, 62, 49, 85, 81, 36, 51, 37, 87, 20,
	23, 88, 37, 33, 82, 35, 60, 19, 99, 28, 42, 4, 40, 96, 11, 33, 31, 57, 17, 19, 100, 46, 91, 34, 33,
	69, 7, 79, 61, 8, 91, 60, 2, 96, 2, 78, 31, 25, 49, 13, 41, 43, 42, 96, 90, 7, 79, 24, 66, 7,
	81, 93, 87, 31, 64, 72, 46, 55, 38, 6, 22, 66, 96, 82, 18, 73, 100, 71, 99, 37, 50, 59, 88, 40, 21,
	98, 52, 60, 9, 10, 62, 64, 46, 16, 29, 36, 56, 82, 5, 24, 28, 57, 61, 70, 88, 19, 71, 73, 20, 98,
	72, 98, 33, 62, 25, 85, 89, 85, 31, 51, 73, 30, 91, 54, 77, 68, 96, 45, 6, 44, 59, 30, 53, 79, 18,
	100, 4, 22, 80, 49, 37, 79, 14, 24, 54, 41, 61, 46, 81, 26, 60, 73, 64, 65, 1, 16, 60, 85, 32, 28,
	10, 60, 45, 41, 52, 22, 98, 78, 93, 10, 9, 87, 98, 78, 31, 96, 1, 8, 88, 92, 48, 26, 60, 94, 55,
	87, 67, 80, 40, 6, 80, 99, 20, 92, 59, 10, 45, 50, 73, 64, 33, 12, 9, 84, 86, 82, 83, 76, 74, 16,
	68, 87, 100, 23, 43, 75, 7, 42, 45, 37, 20, 45, 1, 46, 4, 70, 31, 73, 37, 91, 4, 67, 43, 84, 81,
	49, 28, 92, 19, 30, 41, 7, 100, 52, 75, 69, 26, 76, 60, 53, 12, 69, 23, 47, 47, 13, 47, 50, 10, 17,
	50, 32, 71, 56, 50, 21, 60, 35, 50, 95, 92, 75, 17, 42, 38, 58, 27, 98, 91, 38, 5, 73, 12, 78, 44,
	51, 40, 99, 33, 63, 73, 37, 84, 85, 23, 68, 22, 2, 62, 49, 8, 77, 54, 74, 28, 73, 6, 89, 96, 80,
	79, 51, 32, 24, 39, 23, 43, 47, 95, 12, 70, 79, 7, 63, 74, 99, 14, 27, 87, 53, 15, 35, 23, 9, 68,
	17, 20, 19, 72, 55, 63, 93, 26, 99, 2, 51, 27, 80, 98, 38, 2, 86, 83, 53, 72, 88, 40, 40, 13, 48,
	7, 76, 50, 83, 68, 79, 44, 36, 45, 11, 68, 40, 60, 71, 36, 19, 67, 92, 82, 57, 5, 77, 98, 100, 25,
	52, 91, 100, 18, 88, 92, 91, 64, 59, 40, 21, 38, 58, 100, 64, 66, 81, 69, 64, 20, 45, 59, 75, 86, 15,
	44, 45, 66, 70, 87, 57, 63, 1, 42, 60, 64, 32, 57, 67, 52, 60, 22, 44, 60, 11, 100, 23, 81, 40, 64,
	74, 26, 90, 14, 66, 79, 11, 14, 1, 24, 89, 39, 84, 45, 74, 16, 23, 78, 6, 5, 4, 34, 21, 19, 20,
	65, 45, 9, 88, 47, 9, 67, 63, 79, 35, 36, 35, 41, 98, 18, 89, 41, 64, 94, 62, 56, 3, 40, 2, 49,
	66, 98, 17, 74, 8, 46, 69, 83, 98, 52, 79, 18, 85, 68, 61, 64, 84, 23, 17, 13, 39, 82, 48, 24, 24,
	70, 33, 28, 74, 16, 20, 71, 83, 35, 68, 58, 64, 30, 92, 24, 52, 94, 61, 2, 85, 8, 22, 35, 68, 88,
	57, 2, 24, 95, 43, 100, 94, 8, 95, 83, 12, 18, 87, 10, 24, 22, 79, 23, 80, 81, 36, 7, 66, 5, 54,
	75, 29, 88, 80, 44, 13, 14, 66, 6, 88, 87, 55, 91, 42, 42, 34, 39, 91, 80, 26, 31, 82, 8, 6, 90,
	95, 13, 56, 83, 74, 95, 25, 22, 10, 64, 15, 11, 31, 43, 55, 95, 59, 22, 89, 19, 6, 23, 68, 39, 47,
	31, 10, 47, 95, 25, 78, 86, 70, 84, 19, 36, 10, 32, 18, 43, 27, 26, 95, 75, 87, 38, 52, 41, 38, 25,
	22, 100, 53, 39, 35, 16, 51, 59, 5, 6, 28, 38, 20, 89, 1, 77, 12, 95, 49, 56, 67, 22, 97, 21, 88,
	85, 90, 44, 35, 54, 90, 84, 14, 75, 75, 52, 68, 37, 9, 22, 77, 33, 43, 55, 1, 99, 85, 40, 96, 88,
	72, 46, 78, 80, 10, 66, 34, 76, 45, 21, 12, 73, 50, 28, 63, 99, 84, 89, 94, 48, 63, 88, 22, 63, 60,
	54, 31, 53, 52, 5, 66, 48, 25, 6, 84, 95, 35, 31, 86, 42, 31, 77, 69, 95, 71, 99, 36, 52, 24, 45,
	44, 43, 17, 63, 97, 88, 70, 85, 31, 7, 50, 16, 13, 21, 73, 21, 15, 40, 95, 100, 49, 33, 41, 58, 99,
	93, 10, 35, 14, 95, 26, 35, 86, 63, 72, 36, 64, 29, 17, 69, 26, 94, 10, 79, 93, 42, 27, 4, 64, 34,
	54, 72, 54, 34, 42, 75, 94, 37, 87, 61, 74, 48, 5, 12, 69, 79, 60, 53, 12, 40, 22, 27, 52, 75, 20,
	38, 4, 1, 81, 37, 41, 67, 70, 40, 39, 71, 93, 43, 4, 27, 79, 84, 8, 96, 38, 49, 64, 42, 10, 53,
	67, 50, 40, 53, 69, 30, 90, 68, 18, 82, 45, 48, 79, 51, 87, 28, 4, 84, 49, 1, 9, 89, 62, 81, 2,
	34, 87, 89, 62, 68, 28, 95, 22, 56, 11, 15, 43, 68, 94, 62, 97, 35, 46, 15, 55, 81, 67, 51, 1, 65,
	85, 42, 8, 35, 65, 98, 3, 81, 55, 96, 1, 43, 79, 67, 66, 55, 47, 65, 81, 4, 34, 49, 100, 64, 92,
	70, 15, 51, 57, 6, 62, 70, 38, 42, 86, 32, 82, 68, 70, 88, 76, 73, 2, 32, 49, 80, 58, 78, 90, 35,
	54, 74, 90, 21, 64, 54, 78, 54, 65, 81, 43, 39, 12, 15, 27, 20, 28, 46, 65, 32, 49, 26, 4, 53, 44,
	57, 71, 74, 9, 84, 94, 93, 80, 34, 34, 13, 19, 84, 72, 70, 95, 5, 55, 2, 28, 70, 11, 30, 23, 71,
	86, 37, 94, 7, 66, 73, 40, 35, 70, 1, 32, 94, 78, 79, 56, 78, 48, 15, 36, 64, 92, 41, 10, 96, 40,
	53, 10, 10, 81, 1, 16, 22, 5, 60, 52, 64, 86, 46, 34, 1, 59, 95, 11, 43, 79, 78, 87, 37, 42, 11,
	66, 8, 11, 97, 89, 98, 34, 94, 25, 69, 2, 8, 26, 27, 21, 12, 29, 39, 25, 39, 47, 94, 56, 61, 28,
	73, 71, 15, 60, 67, 37, 64, 14, 32, 42, 83, 41, 50, 57, 7, 59, 32, 56, 21, 69, 18, 24, 73, 93, 66,
	25, 58, 36, 36, 45, 12, 84, 80, 20, 98, 58, 32, 96, 76, 65, 42, 38, 28, 8, 8, 46, 82, 78, 92, 46,
	64, 2, 78, 24, 52, 14, 72, 79, 49, 74, 93, 25, 23, 23, 23, 22, 80, 50, 44, 42, 98, 18, 77, 71, 55,
	32, 83, 90, 32, 49, 57, 22, 54, 5, 6, 29, 23, 94, 93, 76, 45, 24, 80, 3, 25, 75, 100, 77, 91, 62,
	48, 54, 28, 96, 93, 56, 100, 60, 99, 83, 61, 12, 64, 88, 39, 90, 8, 48, 44, 86, 69, 47, 91, 80, 6,
	26, 51, 31, 43, 47, 8, 38, 23, 27, 73, 62, 77, 57, 62, 68, 65, 56, 70, 23, 2, 92, 96, 12, 99, 31,
	3, 16, 75, 21, 67, 94, 41, 4, 66, 40, 37, 33, 51, 7, 54, 31, 8, 91, 16, 65, 69, 75, 89, 40, 69,
	99, 49, 70, 33, 78, 74, 77, 78, 66, 8, 6, 18, 57, 39, 65, 65, 9, 69, 70, 89, 18, 90, 52, 18, 4,
	9, 5, 99, 53, 26, 19, 34, 33, 73, 10, 21, 28, 82, 36, 64, 55, 46, 68, 75, 7, 77, 40, 5, 73, 80,
	30, 59, 38, 61, 57, 20, 91, 46, 64, 61, 35, 40, 24, 70, 79, 73, 67, 72, 77, 42, 8, 68, 7, 58, 21,
	14, 68, 57, 54, 74, 69, 87, 63, 9, 83, 16, 4, 28, 20, 53, 38, 38, 45, 100, 56, 11, 76, 60, 85, 60,
	57, 52, 36, 87, 22, 42, 95, 58, 27, 84, 34, 35, 31, 97, 82, 45, 4, 72, 81, 8, 61, 60, 11, 73, 49,
	7, 60, 93, 46, 98, 22, 12, 54, 39, 44, 91, 3, 3, 86, 25, 20, 91, 4, 73, 17, 61, 43, 85, 20, 6,
	1, 74, 33, 47, 33, 14, 57, 91, 47, 39, 45, 42, 92, 7, 82, 2, 82, 49, 26, 50, 73, 66, 8, 83, 2,
	78, 93, 17, 54, 26, 69, 69, 25, 37, 26, 11, 65, 96, 25, 44, 10, 58, 64, 83, 62, 7, 59, 49, 98, 82,
	44, 84, 75, 91, 89, 94, 75, 15, 96, 68, 85, 83, 30, 65, 51, 21, 2, 45, 25, 79, 23, 74, 78, 31, 68,
	63, 73, 38, 36, 70, 87, 12, 54, 97, 47, 8, 15, 99, 68, 93, 30, 9, 89, 9, 9, 52, 45, 95, 78, 37,
	77, 19, 3, 18, 27, 81, 4, 71, 31, 40, 21, 17, 68, 61, 27, 51, 22, 27, 14, 23, 97, 10, 59, 15, 98,
	60, 68, 75, 50, 63, 41, 21, 92, 53, 7, 81, 47, 45, 40, 64, 72, 7, 3, 39, 34, 34, 87, 82, 82, 89,
	89, 68, 30, 64, 97, 87, 20, 74, 92, 44, 67, 18, 72, 21, 97, 67, 20, 17, 75, 91, 4, 56, 77, 36, 60,
	48, 31, 28, 86, 82, 99, 55, 57, 26, 7, 74, 41, 99, 47, 29, 64, 31, 34, 42, 71, 22, 93, 66, 77, 17,
	29, 91, 84, 52, 31, 6, 82, 34, 17, 92, 28, 61, 66, 64, 36, 90, 41, 18, 46, 81, 22, 11, 53, 62, 79,
	76, 99, 74, 7, 74, 4, 12, 37, 9, 93, 56, 6, 61, 79, 63, 8, 24, 85, 87, 11, 4, 62, 15, 70, 3,
	19, 85, 3, 71, 16, 91, 53, 41, 27, 46, 33, 48, 75, 42, 35, 91, 66, 80, 79, 84, 48, 57, 77, 59, 63,
	22, 23, 28, 57, 24, 85, 30, 70, 97, 80, 35, 85, 78, 12, 96, 72, 73, 42, 68, 60, 39, 41, 5, 91, 33,
	52, 39, 51, 77, 97, 96, 38, 80, 7, 81, 35, 66, 43, 39, 13, 38, 81, 95, 59, 27, 77, 28, 98, 72, 94,
	86, 60, 61, 38, 51, 43, 95, 70, 100, 56, 8, 11, 98, 74, 21, 85, 58, 39, 74, 20, 97, 51, 10, 87, 86,
	84, 80, 50, 78, 81, 93, 65, 15, 5, 47, 67, 11, 83, 57, 53, 22, 85, 68, 69, 99, 68, 58, 40, 9, 87,
	34, 16, 62, 66, 69, 63, 74, 17, 47, 29, 3, 47, 30, 23, 13, 34, 98, 21, 30, 33, 59, 21, 43, 31, 99,
	48, 25, 32, 20, 68, 74, 65, 16, 9, 94, 38, 77, 73, 66, 48, 58, 9, 43, 88, 80, 23, 21, 96, 76, 92,
	90, 40, 100, 48, 68, 70, 71, 77, 70, 6, 86, 35, 2, 79, 98, 38, 56, 18, 86, 78, 71, 10, 58, 47, 24,
	56, 19, 74, 84, 93, 30, 21, 27, 84, 39, 17, 5, 38, 33, 100, 33, 53, 43, 73, 61, 96, 95, 78, 53, 85,
	7, 18, 36, 42, 70, 5, 99, 32, 34, 44, 82, 20, 83, 17, 48, 29, 66, 12, 72, 19, 38, 11, 11, 94, 5,
	23, 38, 41, 54, 37, 46, 21, 55, 79, 32, 47, 75, 69, 35, 88, 62, 19, 60, 9, 53, 90, 79, 67, 64, 14,
	82, 19, 63, 28, 35, 68, 25, 54, 24, 51, 40, 23, 54, 49, 1, 37, 47, 3, 82, 58, 76, 63, 24, 5, 15,
	55, 30, 86, 18, 39, 83, 30, 79, 30, 5, 87, 77, 12, 68, 16, 44, 12, 80, 66, 97, 8, 85, 9, 69, 8,
	88, 76, 26, 20, 42, 81, 63, 68, 63, 31, 28, 97, 58, 4, 73, 63, 99, 29, 86, 89, 65, 83, 21, 3, 61,
	78, 38, 95, 46, 72, 36, 94, 24, 32, 2, 85, 68, 31, 54, 71, 12, 40, 92, 73, 1, 83, 55, 33, 70, 30,
	41, 96, 86, 14, 78, 88, 5, 6, 93, 22, 7, 86, 30, 28, 9, 11, 83, 29, 50, 28, 79, 89, 4, 65, 61,
	56, 28, 22, 74, 85, 24, 38, 7, 56, 97, 36, 46, 68, 61, 91, 54, 47, 42, 56, 1, 96, 44, 53, 70, 43,
	64, 30, 38, 61, 64, 75, 19, 12, 21, 36, 79, 44, 76, 28, 51, 24, 34, 11, 70, 60, 75, 34, 84, 16, 74,
	21, 58, 56, 15, 20, 60, 100, 1, 7, 29, 96, 22, 85, 68, 63, 2, 15, 38, 28, 69, 38, 60, 60, 86, 10,
	28, 11, 82, 5, 73, 49, 13, 66, 97, 32, 95, 3, 7, 60, 85, 70, 56, 78, 59, 77, 23, 72, 57, 87, 94,
	41, 32, 96, 11, 99, 24, 21, 87, 68, 15, 93, 92, 48, 73, 61, 81, 92, 39, 92, 83, 37, 47, 29, 44, 78,
	99, 68, 33, 35, 82, 48, 25, 51, 36, 79, 85, 53, 53, 94, 47, 48, 1, 48, 75, 8, 50, 77, 21, 74, 37,
	87, 69, 3, 20, 32, 32, 97, 91, 47, 38, 80, 58, 86, 77, 23, 92, 83, 78, 100, 87, 81, 5, 24, 51, 87,
	8, 37, 22, 21, 42, 46, 31, 28, 95, 21, 68, 93, 91, 21, 31, 56, 21, 42, 15, 97, 92, 54, 60, 96, 23,
	3, 76, 99, 23, 25, 24, 33, 95, 99, 40, 74, 94, 40, 2, 54, 80, 87, 86, 88, 36, 1, 51, 7, 21, 34,
	11, 54, 2, 23, 27, 60, 29, 5, 32, 50, 11, 4, 44, 28, 61, 100, 19, 10, 77, 21, 35, 98, 83, 81, 35,
	32, 85, 50, 25, 77, 97, 28, 98, 61, 80, 83, 22, 27, 36, 76, 22, 88, 54, 87, 95, 24, 99, 39, 52, 58,
	70, 69, 44, 100, 37, 21, 75, 71, 1, 1, 98, 2, 92, 57, 3, 97, 98, 87, 30, 32, 3, 73, 25, 3, 91,
	24, 75, 51, 99, 63, 51, 93, 33, 79, 73, 62, 21, 60, 93, 49, 56, 26, 92, 58, 87, 99, 66, 67, 2, 54,
	89, 55, 34, 47, 30, 13, 75, 30, 47, 60, 31, 79, 75, 13, 35, 42, 94, 5, 67, 54, 46, 52, 35, 11, 13,
	98, 8, 51, 26, 81, 54, 22, 42, 4, 94, 79, 15, 51, 8, 100, 81, 40, 4, 97, 88, 35, 59, 1, 50, 24,
	6, 16, 19, 89, 12, 88, 21, 100, 41, 84, 48, 46, 16, 52, 8, 19, 30, 59, 79, 26, 1, 1, 78, 1, 42,
	18, 14, 99, 6, 17, 62, 25, 12, 5, 19, 58, 23, 90, 10, 70, 64, 21, 7, 70, 85, 97, 31, 20, 96, 25,
	94, 89, 34, 53, 1, 82, 2, 27, 14, 79, 8, 29, 97, 41, 39, 1, 20, 64, 54, 41, 53, 31, 25, 12, 23,
	38, 4, 92, 58, 79, 91, 66, 21, 22, 33, 86, 45, 1, 13, 47, 26, 84, 3, 35, 91, 78, 21, 44, 34, 4,
	12, 100, 4, 50, 46, 39, 19, 87, 41, 87, 90, 64, 36, 92, 89, 4, 95, 67, 78, 33, 63, 77, 71, 10, 85,
	23, 35, 93, 80, 7, 26, 76, 22, 2, 7, 58, 90, 16, 40, 16, 10, 41, 32, 81, 14, 28, 34, 94, 49, 1,
	92, 88, 24, 56, 29, 7, 9, 92, 58, 66, 27, 77, 58, 59, 20, 50, 21, 40, 68, 73, 95, 81, 83, 20, 73,
	87, 82, 94, 84, 34, 65, 20, 27, 85, 49, 74, 99, 78, 74, 66, 88, 39, 3, 91, 35, 47, 34, 96, 40, 32,
	74, 22, 45, 33, 37, 59, 52, 96, 100, 99, 44, 22, 86, 40, 37, 21, 15, 90, 66, 22, 94, 49, 75, 46, 97,
	33, 67, 25, 19, 24, 55, 49, 21, 67, 60, 95, 37, 2, 71, 5, 95, 93, 19, 51, 63, 75, 57, 35, 27, 63,
	7, 39, 43, 18, 55, 88, 11, 99, 68, 2, 99, 80, 26, 39, 65, 88, 9, 56, 7, 5, 8, 25, 29, 73, 46,
	78, 90, 60, 51, 4, 93, 72, 21, 13, 12, 37, 67, 23, 84, 46, 27, 74, 7, 6, 100, 49, 48, 12, 100, 60,
	3, 13, 53, 67, 17, 82, 42, 73, 97, 4, 41, 40, 34, 59, 100, 47, 56, 40, 15, 5, 75, 46, 85, 22, 15,
	78, 56, 88, 69, 25, 95, 16, 75, 49, 93, 31, 44, 75, 88, 84, 87, 51, 29, 4, 15, 53, 46, 68, 18, 61,
	100, 48, 48, 5, 21, 19, 92, 24, 57, 2, 73, 37, 97, 62, 49, 6, 15, 54, 93, 16, 41, 17, 75, 49, 10,
	33, 52, 92, 94, 89, 3, 57, 35, 72, 42, 71, 82, 74, 18, 79, 63, 93, 57, 15, 96, 82, 55, 66, 37, 27,
	20, 61, 6, 60, 10, 5, 56, 3, 48, 81, 24, 45, 62, 57, 45, 86, 92, 87, 18, 89, 74, 76, 74, 60, 53,
	48, 59, 67, 40, 64, 34, 54, 55, 90, 5, 66, 52, 94, 80, 55, 56, 53, 67, 60, 45, 50, 68, 77, 56, 63,
	97, 74, 12, 67, 76, 56, 27, 31, 82, 43, 14, 76, 60, 87, 89, 12, 79, 46, 34, 62, 77, 5, 39, 56, 64,
	32, 37, 40, 98, 76, 81, 95, 36, 31, 48, 95, 43, 71, 34, 86, 52, 47, 65, 70, 29, 31, 48, 4, 5, 95,
	28, 72, 62, 7, 58, 91, 46, 8, 14, 12, 52, 25, 84, 66, 2, 30, 7, 72, 84, 90, 39, 1, 76, 77, 9,
	80, 48, 70, 60, 74, 34, 29, 23, 16, 79, 44, 53, 23, 13, 17, 55, 41, 12, 82, 68, 13, 31, 85, 96, 35,
	7, 31, 21, 30, 9, 99, 92, 78, 76, 80, 45, 33, 36, 16, 73, 2, 8, 46, 33, 93, 78, 84, 58, 73, 95,
	88, 82, 59, 75, 23, 85, 66, 33, 69, 60, 39, 66, 32, 2, 32, 16, 96, 27, 37, 1, 62, 44, 50, 31, 68,
	30, 93, 6, 64, 89, 39, 92, 21, 59, 64, 92, 7, 70, 75, 72, 37, 42, 56, 44, 54, 26, 35, 75, 17, 85,
	4, 51, 32, 89, 93, 31, 52, 89, 40, 92, 99, 80, 33, 99, 38, 13, 6, 33, 49, 49, 38, 53, 31, 5, 36,
	63, 10, 21, 32, 55, 88, 64, 37, 34, 37, 69, 31, 81, 92, 26, 34, 90, 40, 57, 64, 24, 47, 49, 31, 39,
	60, 37, 82, 78, 28, 79, 15, 4, 32, 54, 5, 93, 34, 23, 26, 4, 7, 58, 96, 10, 63, 95, 50, 67, 25,
	19, 56, 59, 95, 54, 53, 64, 99, 54, 26, 15, 93, 65, 91, 15, 68, 5, 81, 95, 100, 46, 18, 12, 70, 96,
	4, 44, 68, 94, 40, 20, 58, 77, 18, 51, 56, 70, 70, 67, 48, 51, 15, 18, 31, 26, 15, 4, 66, 63, 34,
	70, 2, 70, 13, 2, 4, 61, 59, 5, 30, 25, 65, 2, 2, 46, 12, 9, 66, 52, 93, 60, 75, 64, 7, 42,
	41, 43, 69, 48, 5, 13, 22, 71, 13, 13, 81, 100, 77, 76, 41, 37, 45, 90, 8, 46, 73, 28, 60, 4, 65,
	68, 73, 36, 58, 2, 2, 38, 14, 71, 32, 53, 73, 61, 74, 91, 27, 17, 19, 19, 24, 27, 98, 90, 32, 51,
	96, 26, 64, 83, 76, 96, 2, 1, 97, 58, 84, 44, 51, 69, 32, 39, 71, 1, 16, 9, 91, 45, 43, 68, 46,
	26, 85, 25, 65, 31, 70, 43, 79, 42, 65, 4, 53, 29, 64, 17, 96, 5, 53, 45, 88, 18, 55, 41, 71, 67,
	80, 14, 78, 12, 99, 63, 24, 20, 81, 87, 12, 36, 26, 85, 51, 26, 73, 31, 87, 71, 69, 42, 86, 54, 38,
	73, 3, 88, 37, 26, 82, 37, 65, 53, 31, 12, 80, 52, 11, 19, 86, 83, 17, 93, 90, 60, 96, 99, 46, 56,
	76, 68, 35, 52, 23, 100, 51, 22, 52, 83, 14, 95, 20, 85, 53, 21, 47, 77, 26, 26, 94, 9, 27, 49, 49,
	98, 4, 67, 67, 55, 36, 54, 59, 60, 3, 92, 91, 58, 28, 46, 33, 92, 100, 45, 97, 26, 24, 88, 55, 90,
	50, 39, 40, 26, 42, 1, 17, 82, 68, 10, 41, 66, 76, 69, 4, 83, 51, 39, 33, 83, 44, 3, 93, 30, 34,
	42, 75, 57, 32, 58, 17, 9, 52, 21, 13, 70, 100, 66, 16, 81, 54, 70, 92, 16, 43, 31, 27, 56, 22, 57,
	50, 25, 55, 61, 41, 99, 53, 79, 67, 17, 68, 37, 56, 22, 8, 5, 57, 85, 73, 6, 92, 81, 44, 97, 93,
	33, 5, 82, 41, 82, 90, 19, 44, 96, 63, 24, 58, 11, 84, 43, 46, 12, 2, 15, 88, 69, 9, 62, 38, 98,
	67, 18, 46, 22, 2, 88, 18, 41, 77, 83, 52, 39, 43, 96, 92, 80, 89, 35, 46, 58, 19, 56, 35, 27, 96,
	35, 11, 67, 100, 40, 45, 44, 6, 100, 27, 86, 82, 45, 10, 18, 36, 31, 80, 30, 59, 6, 53, 44, 71, 32,
	50, 77, 64, 70, 39, 12, 87, 83, 66, 81, 26, 9, 33, 51, 53, 85, 45, 81, 6, 82, 30, 12, 6, 12, 4,
	83, 58, 82, 86, 73, 59, 55, 51, 47, 24, 85, 28, 71, 1, 43, 23, 32, 82, 69, 49, 43, 74, 76, 93, 17,
	95, 71, 84, 22, 94, 36, 94, 86, 77, 31, 9, 20, 35, 67, 81, 26, 81, 90, 65, 16, 32, 60, 83, 33, 8,
	71, 21, 45, 99, 41, 9, 61, 13, 16, 76, 85, 99, 48, 88, 53, 15, 51, 84, 75, 3, 54, 51, 47, 20, 55,
	70, 54, 68, 24, 16, 12, 89, 88, 66, 83, 17, 68, 75, 75, 84, 86, 7, 16, 94, 63, 77, 32, 53, 56, 46,
	60, 93, 10, 30, 29, 20, 85, 71, 88, 71, 43, 3, 32, 3, 14, 40, 27, 54, 69, 63, 10, 13, 29, 43, 49,
	86, 12, 29, 29, 73, 99, 13, 55, 43, 23, 90, 27, 47, 13, 88, 51, 8, 74, 37, 19, 83, 40, 14, 85, 52,
	61, 83, 59, 93, 18, 83, 41, 74, 13, 100, 59, 85, 17, 77, 39, 92, 21, 16, 71, 63, 59, 94, 4, 70, 8,
	1, 19, 88, 26, 26, 19, 50, 35, 31, 99, 62, 28, 45, 68, 41, 60, 3, 13, 12, 27, 81, 36, 75, 9, 90,
	72, 23, 97, 42, 1, 25, 17, 99, 74, 13, 13, 90, 82, 31, 74, 76, 85, 72, 71, 49, 66, 45, 27, 34, 29,
	97, 96, 36, 31, 77, 100, 15, 14, 26, 53, 100, 36, 5, 5, 48, 78, 74, 98, 29, 22, 80, 23, 99, 84, 64,
	30, 58, 63, 34, 57, 50, 50, 100, 98, 72, 47, 31, 100, 45, 1, 98, 76, 5, 55, 7, 3, 55, 99, 28, 90,
}

func BenchmarkFilter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Filter(TestData, func(v int) bool {
			return v == 42
		})
	}
}
