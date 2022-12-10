

# Array vs Linked List

## Array

연속된 메모리 블럭

[100] int = 8바이트 * 100 = 800바이트의 어레이청크를 생성함

### 특징

- 연속 메모리: 한꺼번에 할당되고 한꺼번에 해제(GC)된다.
- Random Access에 강하다.
  - 목적지: 배열의 시작주소 + (`INDEX`  * `Type Size`) => `O(1)`
- 삽입/삭제에 약하다 (배열의 처음/끝 에서는 괜찮다) => `O(n)`
- Cache Miss가 잘 일어나지 않는다.
- 요소가 사라질 때마다 GC되지 않는다.
  - 배열은 배열 전체가 참조되지 않을 때 GC가 발생.

게임의 렌더링(점의 색깔을 정해주는) 데이터를 다룰 때에는
렌더링 데이터를 배열로 때려놓고 캐시 히트를 여러번 일어나게 만들어줌

### 배열의 삽입과 삭제

```go
package main
import "fmt"

func main() {
  var a [10]int = [10]int {0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
  var b [11]int
  
  copy(b[0:], a[0:5])
  b[5] = 100
  copy(b[6:], a[5:])
  
  fmt.Println(a)
  fmt.Println(b)
}
```



## Linked List

- 포인터로 연결된 노드들로 구성
- 불연속(논리적) 메모리
- 그래프의 일종
- Single Linked List & Double Linked List
  - `only next` & `next & prev pointer`
- 

```go
package main

import "fmt"

type Node[T any] struct {
	next *Node[T]
	val  T
}

func main() {
	root := &Node[int]{nil, 10}
	root.next = &Node[int]{nil, 20}
	root.next.next = &Node[int]{nil, 30}

	for n := root; n != nil; n = n.next {
		fmt.Println(n.val)
	}

}

```

### 특징

- 불연속 메모리 - 그때그때 필요한 만큼의 메모리만 사용
- 삽입/삭제에 효율적 => `O(1)`
  - but RandomAccess = O(n)
- Cache Miss가 잘 일어남
- 요소가 사라질 때 GC가 일어남

### 삽입 삭제

```go
package main
import "fmt"

type Node[T any] struct {
  next *Node[T]
  val T
}

func Append[T any](node *Node[T], next *Node[T]) *Node[T] {
  node.next = next
  return next
}

func main() {
  root := &Node[int] {nil, 10}
  node := root
  for i := 0; i< 3; i++ {
    node := Append(root, &Node[int]{nil, (i+2)*10})  
  }
  
  for n := root; n != nil; n = n.next {
    fmt.Println("Val: ", n.val)
  }
  
  node = root.next
  originalNext := node.next
  node = Append(node, &Node[int]{nil, 100})
  node.next = originalNext
  
  for n := root; n != nil; n = n.next {
    fmt.Println("Val: ", n.val)
  }
  
  
}
```



- Array
  - Random Access = O(1)
  - Insert/Delete = O(n)
- Linked List
  - Random Access = O(n)
  - Insert/Delete = O(1)

## 결론

대부분의 경우에는 배열 승

단 요소수가 많고 맨 앞에서 삽입/삭제가 빈번한 큐는 링크드 리스트 사용 고려

멀티쓰레드 환경에서 Lockfree queue  등도 고려

=> 주로 대기열 큐에서 링크드 리스트 많이 씀

