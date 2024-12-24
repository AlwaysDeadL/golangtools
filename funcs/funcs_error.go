package funcs

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

//GO 异常处理的 5 种策略

// ReturnNil1 方式1， 直接传递错误，类似 java的 throw e;
func ReturnNil1(url string) (resp *http.Response, err error) {
	resp, err = http.Get(url)
	if err != nil {
		return nil, err
	} else {
		return resp, nil
	}
}

// WaitForServer attempts to contact the server of a URL.
// It tries for one minute using exponential back-off.
// It reports an error if all attempts fail.
// 第二种策略。如果错误的发生是偶然性的，或由不可预知的问题导致的。一个明智的选择是重新尝试失败的操作。在重试时，我们需要限制重试的时间间隔或重试的次数，防止无限制的重试。
// 一分钟内重试，每次休眠  1000 * Millisecond << tries (1,2,4,8)
func WaitForServer(url string) error {
	const timeout = 10 * time.Second
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s);retrying…", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout) //fmt.Errorf： 返回错误类型 error
}

// FatalAndExit 第三种策略： 如果错误发生后，程序无法继续运行，我们就可以采用第三种策略：输出错误信息并结束程序。
func FatalAndExit(num int32) {
	if num > 10 {
		_, _ = fmt.Fprintf(os.Stderr, "num > 10, exit. %v", num)
		os.Exit(1)
		//log.Fatalf("num > 10, exit. %v", num) = fmt.Fprintf && os.Exit
	}
}

// PingError 第四种策略: 有时，我们只需要输出错误信息就足够了，不需要中断程序的运行。我们可以通过log包提供函数
func PingError() {
	if err := fmt.Errorf("test error"); err != nil {
		log.Printf("ping failed: %v.", err)
	}
}

// 第五种: 直接忽略错误
func ignoreError() {
	err := fmt.Errorf("error")
	if err != nil {
		s := fmt.Sprintf("test error %v", err)
		fmt.Println(s)
	}
}
