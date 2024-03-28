```go

func TestName(t *testing.T) {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}

	wg.Wait()

}

func TestWaitGroup(t *testing.T) {
	wg := sync.WaitGroup{}
	c := make(chan struct{})
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int, close <-chan struct{}) {
			defer wg.Done()
			<-close
			fmt.Println(num)
		}(i, c)
	}

	if WaitTimeout(&wg, time.Second*5) {
		close(c)
		fmt.Println("timeout exit")
	}
	time.Sleep(time.Second * 10)

}
func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	// 要求手写代码
	done := make(chan struct{})

	go func() {
		wg.Wait()
		done <- struct{}{}
	}()

	select {
	case <-done:
		fmt.Println("正常返回")
		return false
	case <-time.After(timeout):
		fmt.Println("超时返回")
		return true
	}

	// 要求sync.WaitGroup支持timeout功能
	// 如果timeout到了超时时间返回true
	// 如果WaitGroup自然结束返回false
}

var (
	ist  *instance
	once uint32
	mu   sync.Mutex
)

type instance struct{}

func GetInstance() *instance {
	if atomic.LoadUint32(&once) == 1 {
		return ist
	}
	// 否则先获取互斥锁
	mu.Lock()
	defer mu.Unlock()

	// 需要 double check
	if ist == nil {
		// 说明还是没有被复制过
		ist = &instance{}
		// 但是要标识对应的初始化
		atomic.StoreUint32(&once, 1)
	}

	return ist
}

func TestControl(t *testing.T) {

	// 带超时时间的 ctx
	ctx, cancel := context.WithTimeout(context.Background(), 800*time.Millisecond)

	// 确保退出时，ctx 关掉所有任务了
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go func(idx int) {
			defer wg.Done()
			// 执行任务
			task(ctx, idx)
		}(i)
	}

	// 模拟两秒后去取消 task 的执行
	time.Sleep(time.Second * 2)
	// 通过 ctx，去控制后台的任务，别跑了
	cancel()

	wg.Wait()
}

// 模拟 task
func task(ctx context.Context, idx int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("task 终止")
			return
		default:
			deadline, ok := ctx.Deadline()
			if ok && !time.Now().Before(deadline) {
				fmt.Println("Task 要退出了，不执行了")
				return
			}

			fmt.Printf("执行了 idx 号任务：%d\n", idx)
			// 模拟耗时
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func TestPrintNum(t *testing.T) {

	var wg sync.WaitGroup
	wg.Add(2)

	end := 20

	ch := make(chan struct{})
	defer close(ch)

	// 打印奇数
	go func() {
		defer wg.Done()
		for i := 1; i < end; i += 2 {
			// 打印奇数
			t.Log(i)

			// 打印完后，等待对方打印
			ch <- struct{}{}
			// 说明对方打印完成了，
			<-ch
		}

	}()

	// 打印偶数
	go func() {
		defer wg.Done()

		for i := 2; i <= end; i += 2 {
			// 说明对方打印了，自己可以打印了
			<-ch
			// 打印偶数
			t.Log(i)

			// 等待对方打印
			ch <- struct{}{}
		}
	}()

	wg.Wait()
}

func TestSingleFlight(t *testing.T) {

	var g singleflight.Group

	// 模拟的资源获取函数
	fetchData := func(key string) (interface{}, error) {
		// 假设这里是一个耗时的数据库查询或外部API请求
		time.Sleep(100 * time.Millisecond)
		return fmt.Sprintf("data for %s", key), nil
	}

	key := "my_resource"

	// 启动10个goroutine，模拟并发请求同一个资源
	for i := 0; i < 10; i++ {
		go func(id int) {
			// 使用Do方法请求资源，传入相同的键
			result, err, shared := g.Do(key, func() (interface{}, error) {
				return fetchData(key)
			})
			if err != nil {
				t.Logf("goroutine %d: error: %v\n", id, err)
				return
			}
			t.Logf("goroutine %d: got result: %v, shared: %t\n", id, result, shared)
		}(i)
	}

	// 等待足够长的时间以确保所有goroutine完成
	time.Sleep(1 * time.Second)

}

func TestName2(t *testing.T) {

	wg := sync.WaitGroup{}
	wg.Add(3)

	nums := []string{"A", "B", "C"}
	for i, v := range nums {
		go func() {
			defer wg.Done()

			t.Log(i)
			t.Log(v)
		}()
	}

	wg.Wait()
}

```