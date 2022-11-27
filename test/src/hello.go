package main
// func main() {
// 	println("Hello")
// 	var a int
// 	a = 100
// 	println(a)
// 	a *= 3
// 	println(a)
// 	a >>= 2
// 	println(a)
// 	a |= 1
// 	println(a)

// 	names := []string{"홍길동", "이순신", "강감찬"}   //Go 컬렉션 타입은 아직..
 
// 	for index, name := range names {
//     	println(index, name)
// 	}
// }

// func main() {
//     msg := "Hello"
//     say(&msg)
//     println(msg) //변경된 메시지 출력
// }
 
// func say(msg *string) {
//     println(*msg)
//     *msg = "Changed" //메시지 변경
// }

// func main() {
// // 	var c string
// // 	var b int = 3
// // 	switch b {
// // 	case 1: c = "1반"
// // 	case 2: c = "2반"
// // 	default: c = "다른 반"	
// // 	}
// // 	println(c)
// 	var i int = 1
// 	for i < 10 {
// 		println(i)
// 		i++
// 	}
// }

// func main() {
//     var a = 1
//     for a < 15 {
//         if a == 5 {
//             a += a
//             continue // for루프 시작으로
//         }
//         a++
//         if a > 10 {
//             break  //루프 빠져나옴
//         }
// 		println(a)
//     }
//     if a == 11 {
//         goto END //goto 사용예
//     }
//     println(a)
 
// END:
//     println("End")
// }

// func main() {
//     var sum = 0
//     var i=0
//     var j=0
//     exit_Label:
//     for i = 1; i<10; i++ {
//         for j=1; j<10; j++ {
//             sum += i*j
//             if sum >=1000{
//                 break exit_Label;
//             }
//         }
//     }
//     println("1*1+1*2+...", i, "*", j, "=", sum) 
// }

// func main() {
// 	count, total := sum(1,2,3,4,5)
// 	println(count, total)
// }

// func sum(nums ...int) (count int, sum int) {
// 	sum = 0
// 	count = 0
// 	for _, num := range nums {
// 		sum += num
// 		count ++
// 	}
// 	return 
// }

	
import "fmt"

func main() {
    var a [3]int  //정수형 3개 요소를 갖는 배열 a 선언
    // a[0] = 1
    a[1] = 2
    a[2] = 3
    fmt.Println(a)
}