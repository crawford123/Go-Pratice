package main


import "fmt"
import "os"

func main() {
	//located in e disk
	f := createFile("/tmp/test.txt")
	//Defer is used to ensure that a function call is performed later in a program’s execution 
    defer closeFile(f)
	writeFile(f)

}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
	//abort program
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	//write "data" to f
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %c\n", err)
		os.Exit(1)
	}
}