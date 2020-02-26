package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

/*
Files
a_example.txt
b_read_on.txt
c_incunabula.txt
d_tough_choices.txt
e_so_many_books.txt
f_libraries_of_the_world.txt
*/

type lib struct {
	count   int
	books   [][]int
	signup  int
	canship int
	puan    float32
}

func main() {
	file := flag.Int("input", 0, "Mission File")
	files := [][]string{
		{"input/a_example.txt", "output/a_out.txt"},
		{"input/b_read_on.txt", "output/b_out.txt"},
		{"input/c_incunabula.txt", "output/c_out.txt"},
		{"input/d_tough_choices.txt", "output/d_out.txt"},
		{"input/e_so_many_books.txt", "output/e_out.txt"},
		{"input/f_libraries_of_the_world.txt", "output/f_out.txt"},
	}
	flag.Parse()
	f, err := os.Open(files[*file][0])
	check(err)
	defer f.Close()
	s := bufio.NewReader(f)
	var deadline int = 0
	var booklist [][]int
	var isLib bool = true
	var liblist []lib
	var templib lib

	data, _ := s.ReadString('\n')
	temp := strings.Split(string(data)[0:len(data)-1], " ")
	fmt.Println(temp[2])
	//books, _ = strconv.Atoi(temp[0])
	//libs, _ = strconv.Atoi(temp[1])
	deadline, _ = strconv.Atoi(string(temp[2]))

	fmt.Println(deadline)
	fmt.Println(`-----`)

	data, _ = s.ReadString('\n')
	temp = strings.Split(string(data)[0:len(data)-1], " ")
	for i := range temp {
		a, _ := strconv.Atoi(temp[i])
		tempa := []int{i, a}
		booklist = append(booklist, tempa)
	}
	var count int = 0
	for {
		data, err = s.ReadString('\n')
		if err != nil {
			break
		}
		temp := strings.Split(string(data)[0:len(data)-1], " ")
		if isLib {
			templib.signup, err = strconv.Atoi(temp[1])
			check(err)
			templib.canship, err = strconv.Atoi(temp[2])
			check(err)
			templib.count = count
			isLib = false
			count++
		} else {
			templist := [][]int{}
			for i := range temp {
				a, _ := strconv.Atoi(temp[i])
				templist = append(templist, booklist[a])
			}
			templib.books = Sorter(templist)
			liblist = append(liblist, templib)
			isLib = true
		}
	}
	for i := range liblist {
		liblist[i].puan = Yargic(liblist[i], deadline)
	}
	count = Derleme(quickSort(liblist), booklist, 1000)
	c, err := os.Create(files[*file][1])
	check(err)
	defer c.Close()
	dat, err := ioutil.ReadFile("temp.txt")
	check(err)
	c.WriteString(strconv.Itoa(count) + "\n" + string(dat))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeFile(library []lib, count int, out string) {
	f, err := os.Create(out)
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	countc := strconv.Itoa(count)
	_, err = f.WriteString(string(countc + "\n"))
	for i := range library {
		f.Sync()
		asa := strconv.Itoa(library[i].count)
		aba := strconv.Itoa(len(library[i].books))
		f.WriteString(asa + " " + aba + "\n")
		books := Sorter(library[i].books)
		var nbook []string
		for c := range books {
			aga := strconv.Itoa(books[len(books)-1-c][0])
			nbook = append(nbook, aga)
		}
		f.WriteString(string(strings.Join(nbook, " ") + "\n"))
		f.Sync()
	}
	w.Flush()
}

func Derleme(liblist []lib, booklist [][]int, deadline int) int {
	var top, count int
	f, err := os.Create("temp.txt")
	check(err)
	used := make(map[int]bool)
	defer f.Close()

	for _ = range liblist {
		// if used[liblist[len(liblist)-1].count] {
		// 	continue
		// }
		top += liblist[len(liblist)-1].signup
		if top > deadline {
			top -= liblist[len(liblist)-1].signup
			continue
		}

		asa := strconv.Itoa(liblist[len(liblist)-1].count)
		aba := strconv.Itoa(len(liblist[len(liblist)-1].books))
		f.WriteString(asa + " " + aba + "\n")
		books := Sorter(liblist[len(liblist)-1].books)
		var nbook []string
		for c := range books {
			aga := strconv.Itoa(books[len(books)-1-c][0])
			nbook = append(nbook, aga)
		}
		f.WriteString(string(strings.Join(nbook, " ") + "\n"))
		used[liblist[len(liblist)-1].count] = true
		for a := range liblist[len(liblist)-1].books {
			liblist[len(liblist)-1].books[a][1] = 0
		}

		for i := range liblist {
			liblist[i].puan = Yargic(liblist[i], deadline-top)
		}
		liblist = quickSort(liblist)
		count++
	}
	return count
}

func Yargic(library lib, signup int) float32 {
	var top, count float32
	for i := range library.books {
		top += float32(library.books[i][1])
		count++
	}
	ort := top / count
	book := (float32(len(library.books)) * (ort)) / float32(library.canship)
	return float32(book) * float32(math.Pow(float64(signup), (0-1/float64(signup))))
}

func Sorter(dizi [][]int) [][]int {
	if len(dizi) < 2 {
		return dizi
	}

	left, right := 0, len(dizi)-1

	rand.Seed(int64(len(dizi)))

	pivot := rand.Int() % len(dizi)

	dizi[pivot], dizi[right] = dizi[right], dizi[pivot]

	for i := range dizi {
		if dizi[i][1] < dizi[right][1] {
			dizi[left], dizi[i] = dizi[i], dizi[left]
			left++
		}
	}

	dizi[left], dizi[right] = dizi[right], dizi[left]

	Sorter(dizi[:left])
	Sorter(dizi[left+1:])

	return dizi
}

func quickSort(dizi []lib) []lib {
	if len(dizi) < 2 { // Eğer dizi 1 elemanlı ise diziyi dönüyor. Bu fonksiyonun recursive çalışması için önemli.
		return dizi
	}

	left, right := 0, len(dizi)-1 // Dizinin ilk ve son elemanlarının indexi alınıyor.

	rand.Seed(int64(len(dizi))) // Random için besleme yapılıyor. Dayanak seçimi için önemli.

	// Dayanak Seçimi
	pivot := rand.Int() % len(dizi)

	dizi[pivot], dizi[right] = dizi[right], dizi[pivot] // Dayanak ile son elemanın yerini değiştiriyoruz.

	for i := range dizi { // Dizi boyunda bir dönguye giriyoruz.
		if dizi[i].puan < dizi[right].puan { // Dizinin i elemanı dayanaktan küçük ise
			dizi[left], dizi[i] = dizi[i], dizi[left] // Onu ilk indexe alıyoruz
			left++                                    // Ve ilk indexi bir arttırıyoruz.
		}
	}

	dizi[left], dizi[right] = dizi[right], dizi[left] // sonrasında ilk index ile dayanağı yer değiştiriyoruz.

	quickSort(dizi[:left])   // ilk yarıyı tekrardan sıralamaya sokuyoruz.
	quickSort(dizi[left+1:]) // ikinci yarı da sıralamaya giriyor.

	// Bu işlemlerin tamamı dizi üzerinde yapıldığı için recursive sonunda dizi sıralanmış oluyor.
	// Böylece gereksiz bellek ve dizi kullanmadan en yüksek performansta tüm dizideki elemanları sıralıyoruz.
	return dizi
}

//100000
