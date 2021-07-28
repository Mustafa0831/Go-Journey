package main

// import (
// "fmt"
// "student"
// 	//"github.com/01-edu/z01"
// )

//SortListinsert--------------------------------------------------------------------------------------

// type List = student.List

// type Node = student.NodeL

// func PrintList(l *student.NodeI) {
// 	it := l
// 	for it != nil {
// 		fmt.Print(it.Data, " , ")
// 		it = it.Next
// 	}
// 	fmt.Print(nil, "\n")
// }

// func listPushBack(l *student.NodeI, data int) *student.NodeI {
// 	n := &student.NodeI{Data: data}
// 	fmt.Println(n)
// 	if l == nil {
// 		return n
// 	}
// 	iterator := l
// 	for iterator.Next != nil {
// 		iterator = iterator.Next
// 	}
// 	iterator.Next = n
// 	return l
// }

//Listremoveif----------------------------------------------------------------

// func PrintList(l *student.List) {
// 	it := l.Head
// 	for it != nil {
// 		fmt.Print(it.Data, " -> ")
// 		it = it.Next
// 	}
// 	fmt.Print(nil, "\n")
// }

func main() {

	//Atoi---------------------------------------

	// s9 := ""
	// s10 := "-"
	// s11 := "--123"
	// s12 := "1"
	// s13 := "-3"
	// s14 := "8292"
	// s15 := "9223372036854775807"
	// s16 := "-9223372036854775808"

	// n9 := student.Atoi(s9)
	// n10 := student.Atoi(s10)
	// n11 := student.Atoi(s11)
	// n12 := student.Atoi(s12)
	// n13 := student.Atoi(s13)
	// n14 := student.Atoi(s14)
	// n15 := student.Atoi(s15)
	// n16 := student.Atoi(s16)

	// fmt.Println(n9)
	// fmt.Println(n10)
	// fmt.Println(n11)
	// fmt.Println(n12)
	// fmt.Println(n13)
	// fmt.Println(n14)
	// fmt.Println(n15)
	// fmt.Println(n16)

	// 0
	// 0
	// 0
	// 1
	// -3
	// 8292
	// 9223372036854775807
	// -9223372036854775808

	//Atoibase-----------------------------------------------------------

	// fmt.Println(student.AtoiBase("bcbbbbaab", "abc"))
	// fmt.Println(student.AtoiBase("0001", "01"))
	// fmt.Println(student.AtoiBase("00", "01"))
	// fmt.Println(student.AtoiBase("saDt!I!sI", "CHOUMIisDAcat!"))
	// fmt.Println(student.AtoiBase("AAho?Ao", "WhoAmI?"))
	// fmt.Println(student.AtoiBase("thisinputshouldnotmatter", "abca"))
	// fmt.Println(student.AtoiBase("125", "0123456789"))
	// fmt.Println(student.AtoiBase("uoi", "choumi"))
	// fmt.Println(student.AtoiBase("bbbbbab", "-ab"))

	// 12016
	// 1
	// 0
	// 11557277891
	// 406772
	// 0
	// 125
	// 125
	// 0

	//Printcombn----------------------------------------------------------

	// student.PrintCombN(1)
	// student.PrintCombN(2)
	// student.PrintCombN(3)
	// student.PrintCombN(4)
	// student.PrintCombN(5)
	// student.PrintCombN(6)
	// student.PrintCombN(7)
	// student.PrintCombN(8)
	// student.PrintCombN(9)

	//Printnbrbase-------------------------------------------------------

	// student.PrintNbrBase(919617, "01")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(753639, "CHOUMIisDAcat!")
	// z01.PrintR	// str := "The earliest foundations of what would become   computer science predate the invention of the modern digital computer"
	// fmt.Println(student.SplitWhiteSpaces(str))
	// str = "Machines for calculating fixed numerical tasks such as the abacus have existed since antiquity,"
	// fmt.Println(student.SplitWhiteSpaces(str))
	// str = "aiding in  computations such as multiplication and division ."
	// fmt.Println(student.SplitWhiteSpaces(str))
	// str = "Algorithms for performing computations have existed since antiquity, even before the development of sophisticated computing equipment."
	// fmt.Println(student.SplitWhiteSpaces(str))
	// str = "Wilhelm Schickard designed and constructed the first working mechanical calculator in 1623.[4]"
	// fmt.Println(student.SplitWhiteSpaces(str))
	// str = "In 1673, Gottfried Leibniz demonstrated a digital mechanical calculator,"
	// fmt.Println(student.SplitWhiteSpaces(str))une('\n')
	// student.PrintNbrBase(-174336, "CHOUMIisDAcat!")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(-661165, "1")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(-861737, "Zone01Zone01")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(125, "0123456789ABCDEF")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(-125, "choumi")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(125, "-ab")
	// z01.PrintRune('\n')
	// student.PrintNbrBase(-9223372036854775808, "0123456789")
	// z01.PrintRune('\n')

	// 11100000100001000001
	// HIDAHI
	// -MssiD
	// NV
	// NV
	// 7D
	// -uoi
	// NV
	// -9223372036854775808

	//Recursive---------------------------------------------------------------

	// arg1 := -7
	// arg2 := -2
	// fmt.Println(student.RecursivePower(arg1, arg2))
	// arg1 = -8
	// arg2 = -7
	// fmt.Println(student.RecursivePower(arg1, arg2))
	// arg1 = 4
	// arg2 = 8
	// fmt.Println(student.RecursivePower(arg1, arg2))
	// arg1 = 1
	// arg2 = 3
	// fmt.Println(student.RecursivePower(arg1, arg2))
	// arg1 = -1
	// arg2 = 1
	// fmt.Println(student.RecursivePower(arg1, arg2))
	// arg1 = -6
	// arg2 = 5
	// fmt.Println(student.RecursivePower(arg1, arg2))

	// 0
	// 0
	// 65536
	// 1
	// -1
	// -7776

	//Split-------------------------------------------------------

	// str := "|=choumi=|which|=choumi=|itself|=choumi=|used|=choumi=|cards|=choumi=|and|=choumi=|a|=choumi=|central|=choumi=|computing|=choumi=|unit.|=choumi=|When|=choumi=|the|=choumi=|machine|=choumi=|was|=choumi=|finished,"
	// fmt.Println(student.Split(str, "|=choumi=|"))
	// str = "!==!which!==!was!==!making!==!all!==!kinds!==!of!==!punched!==!card!==!equipment!==!and!==!was!==!also!==!in!==!the!==!calculator!==!business[10]!==!to!==!develop!==!his!==!giant!==!programmable!==!calculator,"
	// fmt.Println(student.Split(str, "!==!"))
	// str = "AFJCharlesAFJBabbageAFJstartedAFJtheAFJdesignAFJofAFJtheAFJfirstAFJautomaticAFJmechanicalAFJcalculator,"
	// fmt.Println(student.Split(str, "AFJ"))
	// str = "<<==123==>>In<<==123==>>1820,<<==123==>>Thomas<<==123==>>de<<==123==>>Colmar<<==123==>>launched<<==123==>>the<<==123==>>mechanical<<==123==>>calculator<<==123==>>industry[note<<==123==>>1]<<==123==>>when<<==123==>>he<<==123==>>released<<==123==>>his<<==123==>>simplified<<==123==>>arithmometer,"
	// fmt.Println(student.Split(str, "<<==123==>>"))

	// [ which itself used cards and a central computing unit. When the machine was finished,]
	// [ which was making all kinds of punched card equipment and was also in the calculator business[10] to develop his giant programmable calculator,]
	// [ Charles Babbage started the design of the first automatic mechanical calculator,]
	// [ In 1820, Thomas de Colmar launched the mechanical calculator industry[note 1] when he released his simplified arithmometer,]

	//Splitwhitespaces------------------------------------------

	// str := "The earliest foundations of what would become   computer science predate the invention of the modern digital computer"
	// fmt.Println(student.SplitWhiteSpaces(str))
	// str = "Machines for calculating fixed numerical tasks such as the abacus have existed since antiquity,"
	// fmt.Println(student.SplitWhiteSpaces(str))
	// str = "aiding in  computations such as multiplication and division ."
	// fmt.Println(student.SplitWhiteSpaces(str))
	// str = "Algorithms for performing computations have existed since antiquity, even before the development of sophisticated computing equipment."
	// fmt.Println(student.SplitWhiteSpaces(str))
	// str = "Wilhelm Schickard designed and constructed the first working mechanical calculator in 1623.[4]"
	// fmt.Println(student.SplitWhiteSpaces(str))
	// str = "In 1673, Gottfried Leibniz demonstrated a digital mechanical calculator,"
	// fmt.Println(student.SplitWhiteSpaces(str))

	// [The earliest foundations of what would become computer science predate the invention of the modern digital computer]
	// [Machines for calculating fixed numerical tasks such as the abacus have existed since antiquity,]
	// [aiding in computations such as multiplication and division .]
	// [Algorithms for performing computations have existed since antiquity, even before the development of sophisticated computing equipment.]
	// [Wilhelm Schickard designed and constructed the first working mechanical calculator in 1623.[4]]
	// [In 1673, Gottfried Leibniz demonstrated a digital mechanical calculator,]

	//ConvertBase--------------------------------------------------

	// result := student.ConvertBase("4506C", "0123456789ABCDEF", "choumi")
	// fmt.Println(result)
	// result = student.ConvertBase("babcbaaaaabcb", "abc", "0123456789ABCDEF")
	// fmt.Println(result)
	// result = student.ConvertBase("256850", "0123456789", "01")
	// fmt.Println(result)
	// result = student.ConvertBase("uuhoumo", "choumi", "Zone01")
	// fmt.Println(result)
	// result = student.ConvertBase("683241", "0123456789", "0123456789")
	// fmt.Println(result)

	// "hccocimc"
	// "9B611"
	// "111110101101010010"
	// "eeone0n"
	// "683241"

	//Advansedsortwordarr-------------------------------------------------------

	// result := []string{"The", "earliest", "computing", "device", "undoubtedly", "consisted", "of", "the", "five", "fingers", "of", "each", "hand"}
	// student.AdvancedSortWordArr(result, strings.Compare)
	// fmt.Println(result)
	// result = []string{"The", "word", "digital", "comesfrom", "\"digits\"", "or", "fingers"}
	// student.AdvancedSortWordArr(result, strings.Compare)
	// fmt.Println(result)
	// result = []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	// student.AdvancedSortWordArr(result, Compare)
	// fmt.Println(result)
	// result = []string{"The", "computing", "consisted", "device", "each", "earliest", "fingers", "five", "hand", "of", "of", "the", "undoubtedly"}
	// student.AdvancedSortWordArr(result, student.Compare)
	// fmt.Println(result) // b, a
	// result = []string{"\"digits\"", "The", "comesfrom", "digital", "fingers", "or", "word"}
	// student.AdvancedSortWordArr(result, student.Compare)
	// fmt.Println(result) // b, a
	// result = []string{"1", "2", "3", "A", "B", "C", "a", "b", "c"}
	// student.AdvancedSortWordArr(result, student.Compare)
	// fmt.Println(result) // b, a

	// [The computing consisted device each earliest fingers five hand of of the undoubtedly]
	// ["digits" The comesfrom digital fingers or word]
	// [1 2 3 A B C a b c]
	// [undoubtedly the of of hand five fingers earliest each device consisted computing The]
	// [word or fingers digital comesfrom The "digits"]
	// [c b a C B A 3 2 1]

	// func Compare(a, b string) int {
	// 	if a == b {
	// 		return 0
	// 	}
	// 	if a < b {
	// 		return -1
	// 	}
	// 	return 1
	//}

	//ActiveBits+----------------------------------

	// nbits := student.ActiveBits(15)
	// fmt.Println(nbits)
	// nbits = student.ActiveBits(17)
	// fmt.Println(nbits)
	// nbits = student.ActiveBits(4)
	// fmt.Println(nbits)
	// nbits = student.ActiveBits(11)
	// fmt.Println(nbits)
	// nbits = student.ActiveBits(9)
	// fmt.Println(nbits)
	// nbits = student.ActiveBits(12)
	// fmt.Println(nbits)
	// nbits = student.ActiveBits(7)
	// fmt.Println(nbits)

	// 4
	// 2
	// 1
	// 3
	// 2
	// 2
	// 1

	//SortListinsert------------------------------------------------------------------

	// var link *student.NodeI

	//1

	// link = listPushBack(link, 0)
	// PrintList(link)
	// link = student.SortListInsert(link, 39)
	// PrintList(link)

	// (0, 39, <nil>)

	//2

	// link = listPushBack(link, 0)
	// link = listPushBack(link, 1)
	// link = listPushBack(link, 2)
	// link = listPushBack(link, 3)
	// link = listPushBack(link, 4)
	// link = listPushBack(link, 5)
	// link = listPushBack(link, 24)
	// link = listPushBack(link, 25)
	// link = listPushBack(link, 54)
	// PrintList(link)
	// link = student.SortListInsert(link, 33)
	// PrintList(link)

	// (0, 1, 2, 3, 4, 5, 24, 25, 33, 54, <nil>)

	//3

	// link = listPushBack(link, 0)
	// link = listPushBack(link, 2)
	// link = listPushBack(link, 18)
	// link = listPushBack(link, 33)
	// link = listPushBack(link, 37)
	// link = listPushBack(link, 37)
	// link = listPushBack(link, 39)
	// link = listPushBack(link, 52)
	// link = listPushBack(link, 53)
	// link = listPushBack(link, 57)
	// PrintList(link)
	// link = student.SortListInsert(link, 53)
	// PrintList(link)

	// (0, 2, 18, 33, 37, 37, 39, 52, 53, 53, 57, <nil>)

	//4

	// link = listPushBack(link, 0)
	// link = listPushBack(link, 5)
	// link = listPushBack(link, 18)
	// link = listPushBack(link, 24)
	// link = listPushBack(link, 28)
	// link = listPushBack(link, 35)
	// link = listPushBack(link, 42)
	// link = listPushBack(link, 45)
	// link = listPushBack(link, 52)
	// PrintList(link)
	// link = student.SortListInsert(link, 52)
	// PrintList(link)

	// (0, 5, 18, 24, 28, 35, 42, 45, 52, 52, <nil>)

	//5

	// link = listPushBack(link, 0)
	// link = listPushBack(link, 12)
	// link = listPushBack(link, 20)
	// link = listPushBack(link, 23)
	// link = listPushBack(link, 23)
	// link = listPushBack(link, 24)
	// link = listPushBack(link, 30)
	// link = listPushBack(link, 41)
	// link = listPushBack(link, 53)
	// link = listPushBack(link, 57)
	// link = listPushBack(link, 59)
	// PrintList(link)
	// link = student.SortListInsert(link, 38)
	// PrintList(link)

	// (0, 12, 20, 23, 23, 24, 30, 38, 41, 53, 57, 59, <nil>)

	//SortedlistMerge--------------------------------------------------------------------------

	//2

	// var link *student.NodeI
	// var link2 *student.NodeI

	// PrintList1(student.SortedListMerge(link2, link))

	//3

	// link2 = listPushBack(link2, 2)
	// link2 = listPushBack(link2, 2)
	// link2 = listPushBack(link2, 4)
	// link2 = listPushBack(link2, 9)
	// link2 = listPushBack(link2, 12)
	// link2 = listPushBack(link2, 12)
	// link2 = listPushBack(link2, 19)
	// link2 = listPushBack(link2, 20)
	// PrintList1(student.SortedListMerge(link2, link))

	// 2-> 2-> 4-> 9-> 12-> 12-> 19-> 20-> <nil>

	// 4

	// link = listPushBack(link, 4)
	// link = listPushBack(link, 4)
	// link = listPushBack(link, 6)
	// link = listPushBack(link, 9)
	// link = listPushBack(link, 13)
	// link = listPushBack(link, 18)
	// link = listPushBack(link, 20)
	// link = listPushBack(link, 20)
	// PrintList1(student.SortedListMerge(link, link2))

	// 4 -> 4 -> 6 -> 9 -> 13 -> 18 -> 20 -> 20 -> <nil>

	//5

	// link = listPushBack(link, 0)
	// link = listPushBack(link, 7)
	// link = listPushBack(link, 28)
	// link = listPushBack(link, 39)
	// link = listPushBack(link, 64)
	// link = listPushBack(link, 92)
	// link = listPushBack(link, 93)
	// link = listPushBack(link, 97)

	// link2 = listPushBack(link2, 23)
	// link2 = listPushBack(link2, 27)
	// link2 = listPushBack(link2, 30)
	// link2 = listPushBack(link2, 70)
	// link2 = listPushBack(link2, 75)
	// link2 = listPushBack(link2, 80)
	// link2 = listPushBack(link2, 81)
	// link2 = listPushBack(link2, 85)
	// PrintList1(student.SortedListMerge(link, link2))

	// 0-> 7-> 39-> 80-> 23-> 27-> 30-> 85-> 81-> 75-> 70-> 92-> 97-> 93-> 91-> 28-> 64-> <nil>

	// 6

	// link = listPushBack(link, 0)
	// link = listPushBack(link, 2)
	// link = listPushBack(link, 11)
	// link = listPushBack(link, 30)
	// link = listPushBack(link, 54)
	// link = listPushBack(link, 56)
	// link = listPushBack(link, 70)
	// link = listPushBack(link, 79)
	// link = listPushBack(link, 99)

	// link2 = listPushBack(link2, 23)
	// link2 = listPushBack(link2, 28)
	// link2 = listPushBack(link2, 38)
	// link2 = listPushBack(link2, 67)
	// link2 = listPushBack(link2, 67)
	// link2 = listPushBack(link2, 79)
	// link2 = listPushBack(link2, 95)
	// link2 = listPushBack(link2, 97)
	// PrintList1(student.SortedListMerge(link2, link))

	// 0-> 2-> 11-> 23-> 28-> 30-> 38-> 54-> 56-> 67-> 67-> 70-> 79-> 79-> 95-> 97-> 99-> <nil>

	// 7

	// link = listPushBack(link, 0)
	// link = listPushBack(link, 3)
	// link = listPushBack(link, 8)
	// link = listPushBack(link, 8)
	// link = listPushBack(link, 13)
	// link = listPushBack(link, 19)
	// link = listPushBack(link, 34)
	// link = listPushBack(link, 38)
	// link = listPushBack(link, 46)

	// link2 = listPushBack(link2, 7)
	// link2 = listPushBack(link2, 39)
	// link2 = listPushBack(link2, 45)
	// link2 = listPushBack(link2, 53)
	// link2 = listPushBack(link2, 59)
	// link2 = listPushBack(link2, 70)
	// link2 = listPushBack(link2, 76)
	// link2 = listPushBack(link2, 79)
	// PrintList(student.SortedListMerge(link2, link))

	// 0-> 3-> 7-> 8-> 8-> 13-> 19-> 34-> 38-> 39-> 45-> 46-> 53-> 59-> 70-> 76-> 79-> <nil>

	//listremoveif--------------------------------------------------------------------

	// a := &student.List{}
	// fmt.Println("----normal state----")
	// PrintList(a)
	// student.ListRemoveIf(a, 1)
	// fmt.Println("------answer-----")
	// PrintList(a)
	// fmt.Println()

	// b := &student.List{}
	// fmt.Println("----normal state----")
	// PrintList(b)
	// student.ListRemoveIf(b, 96)
	// fmt.Println("------answer-----")
	// PrintList(b)
	// fmt.Println()

	// c := &student.List{}
	// fmt.Println("----normal state----")
	// student.ListPushBack(c, 98)
	// student.ListPushBack(c, 98)
	// student.ListPushBack(c, 33)
	// student.ListPushBack(c, 34)
	// student.ListPushBack(c, 33)
	// student.ListPushBack(c, 34)
	// student.ListPushBack(c, 33)
	// student.ListPushBack(c, 89)
	// student.ListPushBack(c, 33)
	// PrintList(c)
	// student.ListRemoveIf(c, 34)
	// fmt.Println("------answer-----")
	// PrintList(c)
	// fmt.Println()

	// d := &student.List{}
	// fmt.Println("----normal state----")
	// student.ListPushBack(d, 79)
	// student.ListPushBack(d, 74)
	// student.ListPushBack(d, 99)
	// student.ListPushBack(d, 79)
	// student.ListPushBack(d, 7)
	// PrintList(d)
	// student.ListRemoveIf(d, 99)
	// fmt.Println("------answer-----")
	// PrintList(d)
	// fmt.Println()

	// e := &student.List{}
	// fmt.Println("----normal state----")
	// student.ListPushBack(e, 56)
	// student.ListPushBack(e, 93)
	// student.ListPushBack(e, 68)
	// student.ListPushBack(e, 56)
	// student.ListPushBack(e, 87)
	// student.ListPushBack(e, 68)
	// student.ListPushBack(e, 56)
	// student.ListPushBack(e, 68)
	// PrintList(e)
	// student.ListRemoveIf(e, 68)
	// fmt.Println("------answer-----")
	// PrintList(e)
	// fmt.Println()
	// f := &student.List{}
	// fmt.Println("----normal state----")
	// student.ListPushBack(f, "mvkUxbqhQve4l")
	// student.ListPushBack(f, "4Zc4t hnf SQ")
	// student.ListPushBack(f, "q2If E8BPuX")
	// PrintList(f)
	// student.ListRemoveIf(f, "4Zc4t hnf SQ")
	// fmt.Println("------answer-----")
	// PrintList(f)
	// fmt.Println()
	// link := &student.List{}
	// link2 := &student.List{}
	// fmt.Println("----normal state----")
	// student.ListPushBack(link2, 1)
	// PrintList(link2)
	// student.ListRemoveIf(link2, 1)
	// fmt.Println("------answer-----")
	// PrintList(link2)
	// fmt.Println()
	// fmt.Println("----normal state----")
	// student.ListPushBack(link, 1)
	// student.ListPushBack(link, "Hello")
	// student.ListPushBack(link, 1)
	// student.ListPushBack(link, "There")
	// student.ListPushBack(link, 1)
	// student.ListPushBack(link, 1)
	// student.ListPushBack(link, "How")
	// student.ListPushBack(link, 1)
	// student.ListPushBack(link, "are")
	// student.ListPushBack(link, "you")
	// student.ListPushBack(link, 1)
	// PrintList(link)
	// student.ListRemoveIf(link, 1)
	// fmt.Println("------answer-----")
	// PrintList(link)

	// (<nil>)
	// (<nil>)
	// (98-> 98-> 33-> 33-> 33-> 89-> 33-> <nil>)
	// (79-> 74-> 79-> 7-> <nil>)
	// (56-> 93-> 56-> 87-> 56-> <nil>)
	// (mvkUxbqhQve4l-> q2If E8BPuX -> <nil>)

	//BTreeTransPlant-------------------------------------------------------------------

	// root := &student.TreeNode{Data: "01"}
	// student.BTreeInsertData(root, "07")
	// student.BTreeInsertData(root, "05")
	// student.BTreeInsertData(root, "12")
	// student.BTreeInsertData(root, "10")
	// student.BTreeInsertData(root, "02")
	// student.BTreeInsertData(root, "03")
	// node := student.BTreeSearchItem(root, "12")
	// replacement := &student.TreeNode{Data: "55"}
	// student.BTreeInsertData(replacement, "60")
	// student.BTreeInsertData(replacement, "33")
	// student.BTreeInsertData(replacement, "12")
	// student.BTreeInsertData(replacement, "15")
	// root = student.BTreeTransplant(root, node, replacement)
	// student.BTreeApplyInorder(root, fmt.Println)

	// 01
	// 02
	// 03
	// 05
	// 07
	// 12
	// 15
	// 33
	// 55
	// 60

	// 	root := &student.TreeNode{Data: "33"}
	// student.BTreeInsertData(root, "5")
	// student.BTreeInsertData(root, "52")
	// student.BTreeInsertData(root, "20")
	// student.BTreeInsertData(root, "31")
	// student.BTreeInsertData(root, "13")
	// student.BTreeInsertData(root, "11")
	// node := student.BTreeSearchItem(root, "20")
	// replacement := &student.TreeNode{Data: "55"}
	// student.BTreeInsertData(replacement, "60")
	// student.BTreeInsertData(replacement, "33")
	// student.BTreeInsertData(replacement, "12")
	// student.BTreeInsertData(replacement, "15")
	// root = student.BTreeTransplant(root, node, replacement)
	// student.BTreeApplyInorder(root, fmt.Println)

	// 12
	// 15
	// 33
	// 55
	// 60
	// 33
	// 5
	// 52

	// root := &student.TreeNode{Data: "03"}
	// student.BTreeInsertData(root, "39")
	// student.BTreeInsertData(root, "99")
	// student.BTreeInsertData(root, "44")
	// student.BTreeInsertData(root, "11")
	// student.BTreeInsertData(root, "14")
	// student.BTreeInsertData(root, "11")
	// node := student.BTreeSearchItem(root, "11")
	// replacement := &student.TreeNode{Data: "55"}
	// student.BTreeInsertData(replacement, "60")
	// student.BTreeInsertData(replacement, "33")
	// student.BTreeInsertData(replacement, "12")
	// student.BTreeInsertData(replacement, "15")
	// root = student.BTreeTransplant(root, node, replacement)
	// student.BTreeApplyInorder(root, fmt.Println)

	// 03
	// 12
	// 15
	// 33
	// 55
	// 60
	// 39
	// 44
	// 99


	//BTreeApplyByLevel-----------------------------------------------------

	// root := &student.TreeNode{Data: "01"}
	// student.BTreeInsertData(root, "07")
	// student.BTreeInsertData(root, "12")
	// student.BTreeInsertData(root, "10")
	// student.BTreeInsertData(root, "05")
	// student.BTreeInsertData(root, "02")
	// student.BTreeInsertData(root, "03")
	// student.BTreeApplyByLevel(root, fmt.Println)

	// 01
	// 07
	// 05
	// 12
	// 02
	// 10
	// 03

	// root := &student.TreeNode{Data: "01"}
	// student.BTreeInsertData(root, "07")
	// student.BTreeInsertData(root, "12")
	// student.BTreeInsertData(root, "10")
	// student.BTreeInsertData(root, "05")
	// student.BTreeInsertData(root, "02")
	// student.BTreeInsertData(root, "03")
	// student.BTreeApplyByLevel(root, fmt.Print)

	// 01070512021003

	//  BTreeDeleteNode-----------------------------------------------------------------------------------

	// root := &student.TreeNode{Data: "01"}
	// student.BTreeInsertData(root, "07")
	// student.BTreeInsertData(root, "05")
	// student.BTreeInsertData(root, "12")
	// student.BTreeInsertData(root, "02")
	// student.BTreeInsertData(root, "10")
	// student.BTreeInsertData(root, "03")
	// node := student.BTreeSearchItem(root, "02")
	// fmt.Println("Before delete:")
	// student.BTreeApplyInorder(root, fmt.Println)
	// root = student.BTreeDeleteNode(root, node)
	// fmt.Println("After delete:")
	// student.BTreeApplyInorder(root, fmt.Println)

	// 01
	// 03
	// 05
	// 07
	// 10
	// 12

	// root := &student.TreeNode{Data: "33"}
	// student.BTreeInsertData(root, "5")
	// student.BTreeInsertData(root, "20")
	// student.BTreeInsertData(root, "31")
	// student.BTreeInsertData(root, "13")
	// student.BTreeInsertData(root, "11")
	// student.BTreeInsertData(root, "52")
	// node := student.BTreeSearchItem(root, "20")
	// fmt.Println("Before delete:")
	// student.BTreeApplyInorder(root, fmt.Println)
	// root = student.BTreeDeleteNode(root, node)
	// fmt.Println("After delete:")
	// student.BTreeApplyInorder(root, fmt.Println)

	// 11
	// 13
	// 31
	// 33
	// 5
	// 52

	// root := &student.TreeNode{Data: "03"}
	// student.BTreeInsertData(root, "39")
	// student.BTreeInsertData(root, "99")
	// student.BTreeInsertData(root, "44")
	// student.BTreeInsertData(root, "11")
	// student.BTreeInsertData(root, "14")
	// student.BTreeInsertData(root, "11")
	// node := student.BTreeSearchItem(root, "03")
	// fmt.Println("Before delete:")
	// student.BTreeApplyInorder(root, fmt.Println)
	// root = student.BTreeDeleteNode(root, node)
	// fmt.Println("After delete:")
	// student.BTreeApplyInorder(root, fmt.Println)

	// 11
	// 11
	// 14
	// 39
	// 44
	// 99

	// root := &student.TreeNode{Data: "03"}
	// student.BTreeInsertData(root, "03")
	// student.BTreeInsertData(root, "94")
	// student.BTreeInsertData(root, "19")
	// student.BTreeInsertData(root, "24")
	// student.BTreeInsertData(root, "111")
	// student.BTreeInsertData(root, "01")
	// node := student.BTreeSearchItem(root, "03")
	// fmt.Println("Before delete:")
	// student.BTreeApplyInorder(root, fmt.Println)
	// root = student.BTreeDeleteNode(root, node)
	// fmt.Println("After delete:")
	// student.BTreeApplyInorder(root, fmt.Println)

	// 01
	// 03
	// 111
	// 19
	// 24
	// 94
}

// func PrintList(l *student.NodeL) {
// 	it := l
// 	for it != nil {
// 		fmt.Print(it.Data, " -> ")
// 		it = it.Next
// 	}
// 	fmt.Print(nil, "\n")
// }
