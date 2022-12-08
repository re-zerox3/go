package hw08

import(
				"fmt"
				"math"
)


type Shape interface{
				Print()
}

type TwoD interface{
				Area()float64
				Perimeter()float64
}

type ThreeD interface{
				Volume()float64
				SurfaceArea()float64
}


type Rectangle struct{
				width float64
			 length float64
}

type Circle struct{
				 radius float64

}

type Sphere struct{
				 radius float64
}

type Box struct{
				 length float64
				 width float64
				 height float64
}



func(r Rectangle)Print(){
				fmt.Printf("Rectangle (%.3f,%.3f) : A = %.3f, P = %.3f\n",r.width,r.length,r.Area(),r.Perimeter())
}

func(c Circle)Print(){
				fmt.Printf("Circle (%.3f) : A = %.3f, P = %.3f\n",c.radius,c.Area(),c.Perimeter())
}

func(s Sphere)Print(){
				fmt.Printf("Sphere (%.3f) :V = %.3f, SA = %.3f\n",s.radius, s.Volume(),s.SurfaceArea())
}
func(b Box)Print(){
				fmt.Printf("Box (%.3f,%.3f,%.3f) : V = %.3f, SA = %.3f\n",b.length, b.width, b.height, b.Volume(),b.SurfaceArea())
}
func MakeRectangle(width float64,length float64)Rectangle{
				var r Rectangle
				r = Rectangle{width:width,length:length}
				return r
}

func MakeCircle(radius float64)Circle{
				var c Circle
				c = Circle{radius:radius}
				return c
}

func MakeSphere(radius float64)Sphere{
				var s Sphere
				s = Sphere{radius: radius}
				return s
}

func MakeBox(length float64, width float64, height float64)Box{
				var b Box
				b = Box{length:length,width:width,height:height}
				return b
}


func (r Rectangle)Area()float64{
				area := r.width * r.length
				return area
}

func (c Circle)Area()float64{
				area := math.Pi*math.Pow(c.radius,2)
				return area
}

func (r Rectangle)Perimeter()float64{
				perimeter := 2*(r.length + r.width)
				return perimeter
}

func (c Circle)Perimeter()float64{
				perimeter := 2*math.Pi*c.radius
				return perimeter
}


func (s Sphere)Volume()float64{
				volume := (4/3)*math.Pi*math.Pow(s.radius,3)
				return volume
}

func (s Sphere)SurfaceArea()float64{
				surfaceArea := 4*math.Pi*math.Pow(s.radius,2)
				return surfaceArea
}

func (b Box)Volume()float64{
				volume := b.length * b.width *b.height
				return volume
}

func (b Box)SurfaceArea()float64{
				surfaceArea := 2*(b.length*b.width +b.length*b.height + b.width*b.height)
				return surfaceArea
}










