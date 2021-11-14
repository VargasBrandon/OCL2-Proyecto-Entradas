package main;
import ( "fmt");

var t0, t1, t2, t3, t4, t5, t6, t7, t8 float64;
var P, H float64;
var stack [100000000]float64;
var heap  [100000000]float64;



func main(){
	// Regla 1
	t0 = 0;
	t0 = t1;
	t1 = t0;

	t2 = 10;
	t2 = t3;
	t3 = t2;
	fmt.Printf("%d", int(1));
	fmt.Printf("%d", int(1));
	fmt.Printf("%c", int(10));

	// Regla 2
	goto L1;
	fmt.Printf("%d", int(2));
	fmt.Printf("%d", int(2));
	fmt.Printf("%d", int(2));
	L1:
	fmt.Printf("%d", int(2));
	fmt.Printf("%d", int(2));
	fmt.Printf("%c", int(10));
	
	// Regla 3
	if 20 == 10 {goto L2;}
	goto L3;
	L2:
	fmt.Printf("%d", int(3));
	fmt.Printf("%d", int(3));
	fmt.Printf("%d", int(3));
	L3:
	fmt.Printf("%d", int(3));
	fmt.Printf("%d", int(3));
	fmt.Printf("%c", int(10));

	// Regla 4
	goto L4;
	fmt.Printf("%d", int(4));
	fmt.Printf("%d", int(4));
	fmt.Printf("%d", int(4));
	L4:
	goto L5;
	L5:
	fmt.Printf("%d", int(4));
	fmt.Printf("%d", int(4));
	fmt.Printf("%c", int(10));

	// Regla 5
	if 5 < 10 {goto L6;}
	fmt.Printf("%d", int(5));
	fmt.Printf("%d", int(5));
	fmt.Printf("%d", int(5));
	L6:
	goto L7;
	L7:
	fmt.Printf("%d", int(5));
	fmt.Printf("%d", int(5));
	fmt.Printf("%c", int(10));

	// Regla 6
	t4 = t4 + 0;
	t4 = t4 - 0;
	t4 = t4 * 1;
	t4 = t4 / 1;
	fmt.Printf("%d", int(6));
	fmt.Printf("%d", int(6));
	fmt.Printf("%c", int(10));

	// Regla 7
	t5 = t6 + 0;
	t5 = t6 - 0;
	t5 = t6 * 1;
	t5 = t6 / 1;
	fmt.Printf("%d", int(7));
	fmt.Printf("%d", int(7));
	fmt.Printf("%c", int(10));

	// Regla 8
	t7 = t8 * 2;
	t7 = t8 * 0;
	t7 = 0 / t8;
	fmt.Printf("%d", int(8));
	fmt.Printf("%d", int(8));
	fmt.Printf("%c", int(10));
	
}