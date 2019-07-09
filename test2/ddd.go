package test2

type Rect struct{   //定义矩形类
	x,y float64       //类型只包含属性，并没有方法
	width,height float64
}
func (r *Rect) Area() float64{    //为Rect类型绑定Area的方法，*Rect为指针引用可以修改传入参数的值
	return r.width*r.height         //方法归属于类型，不归属于具体的对象，声明该类型的对象即可调用该类型的方法
}
func main ()  {

}
