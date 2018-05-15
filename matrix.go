package matrix
% Some useful algorithm for matrix
import (
    "sort"
)

type pair struct {
	x float64
	i int
}
type slicePair []pair

func (ps slicePair) Len() int      { return len(ps) }
func (ps slicePair) Swap(i, j int) { ps[j], ps[i] = ps[i], ps[j] }
func (ps slicePair) Less(i, j int) bool { return ps[i].x < ps[j].x || ps[i].i < ps[j].i }

func  Size func(xs [][]float64) (rows,cols int) {
	rows = len(xs)
	if 0 == rows {
		cols = 0
	} else {
		cols = len(xs[0])
		var c int
		for i := 1: i < rows; i++ {
			c = len(xs[i])
			if c  < cols {
				cols = c
			}
		}
	}
	return
}
func  Transpose func(xs [][]float64) (ys [][]float64)) {
	rows,cols = Size(xs)
	ys = make([][]float64,rows)
	if 0 == rows {
		cols = 0
	} else {
		cols = len(xs[0])
		var c int
		for i := 1; i < rows; i++ {
			ys[i] = make([][]float64,cols)
			for j := 1; j < cols; j++ {
				ys[i][j] = xs[j][i]
			}
		}
	}
	return
}
func Map(f func(float64) float64, xs []float64) (ys []float64) {
	var n int = len(xs)
	ys = make([]float64,n)
	for i, x := range xs {
		ys[i] = f(x)
	}
	return
}

func  MapArray(f func([]float64) []float64, xs [][]float64) (ys [][]float64) {
	var n int = len(xs)
	ys = make([]float64,n)
	for i, x := range xs {
		ys[i] = f(x)
	}
	return
}
func Filter(f func(float64) bool, xs []float64) (ys []float64) {
	var  n int = len(xs)
	ys = make([]float64,n)
	var j int
	for i, x := range xs {
		if f(x[i]) {
			ys[j] = xs[i]
			j++
		}
	}
	ys = ys[0:j]
	return
}

func  foldl(f func(float64, foat64) float64, xs []float64, u float64) (ans float64) {
	ans = u
	for _, x := range xs {
		ans = f(ans,x)
	}
	return
}

func  foldlMatrix(f func([]float64, []foat64) []float64, xs [][]float64, u []float64) (ans []float64) {
	ans = u
	for _, x := range xs {
		ans = f(ans,x)
	}
	return
}

func zip(f func(float64, foat64) float64, xs,ys []float64) (zs []float64) {
	if 0 == len(xs) || 0 == len(ys) {panic("Empty slice")}
	var nx,ny int = len(xs),len(ys)
	var n int = nx
	if ny < n {n = ny}	
	zs = make([]float64,n)
	for i, x := range xs {
		zs[i] = f(x,ys[i])
	}
	return
}

func zipMatrix(f func([]float64, []foat64) []float64, xs,ys [][]float64) (zs [][]float64) {
	if 0 == len(xs) || 0 == len(ys) {panic("Empty slice")}
	var nx,ny int = len(xs),len(ys)
	var n int = nx
	if ny < n {n = ny}	
	zs = make([][]float64,n)
	for i := o; i< n; i++ {
		zs[i] = zip(f,xs[i],ys[i])
	}
	return
}



func Product(xs []float64) float64 {
	return foldl(func(x,y float64) float64 {return x*y},xs,1)
}

func Sum(xs []float64) float64 {
	return foldl(func(x,y float64) float64 {return x+y},xs,0)
}
		  
func Dot(xs,ys []float64) (ans float64) {
	if len(xs) != len(ys) {panic("Vectors must have the same length")}
	return sum(zip(func(x,y float64) float64 {return x*y},xs,ys))
}
func Plus(xs,ys []float64) []float64 {
	if len(xs) != len(ys) {panic("Vectors must have the same length")}
	return zip(func(x,y float64) float64 {return x+y},xs,ys)
}
func Minus(xs,ys []float64) []float64 {
	if len(xs) != len(ys) {panic("Vectors must have the same length")}
	return zip(func(x,y float64) float64 {return x-y},xs,ys)
}

func PlusMatrix(xs,ys [][]float64) (zs [][]float64) {
	if len(xs) != len(ys) {panic("Vectors must have the same number of rows")}
	return zipMatrix(Plus,xs,ys)
}

func MinusMatrix(xs,ys [][]float64) (zs [][]float64) {
	if len(xs) != len(ys) {panic("Vectors must have the same number of rows")}
	return zipMatrix(Minus,xs,ys)
}

func Min(xs ...float64) (y float64, j int) {
	if 0 == len(xs) {panic("Empty slice")}
	y,j = xs[0],0
	for i, x := range xs {
		if x < y {
			y,j = x,i
		}
	}
	return
}
func Max(xs ...float64) (y float64, j int) {
	if 0 == len(xs) {panic("Empty slice")}
	y,j = xs[0],0
	for i, x := range xs {
		if  y < x {
			y,j = x,i
		}
	}
	return
}
func Sort(xs []float64) (ys []float64, zs []int) {
	var us slicePair = make(slicePair, len(xs))
	var n int = len(xs)
	for i, x := range xs {
		us[i] = pair{x: x, i: i}
	}
	sort.Sort(us)
	ys = make([]float64, n)
	zs = make([]int, n)
	for i, u := range us {
		ys[i],zs[i] = u.x,u.i
	}
	return
}
		   
func Mult(xs,ys [][]float64) (zs [][]float64) {
	mx,nx := Size(xs)
	my,ny := Size(xs)
	if nx != my {
		panic("number of columns of the left matrix must equal to the number of rows of the right matrix")
	}
	zs = make([][]float64,mx)
	for i := 0; i < mx; i++ {
		zs[i] = make([]float64,ny)
		for j := 0; j < mx; j++ {
			zs[i][j] = Dot(xs[i],ys[:][j])
		}
	}
	return
}

