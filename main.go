package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

func main() {
  log.Println("로딩중...")

	const (
		// These paths will be different on your system.
		seleniumPath    = "selenium.jar"
		geckoDriverPath = "geckodriver"
		port            = 4444
	)

	opts := []selenium.ServiceOption{
		selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox
	}
	// selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}

	defer service.Stop()

	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}

	defer wd.Quit()
  log.Println("로딩 다함")

  wd.Get("https://nid.naver.com/nidlogin.login")
  log.Println("로그인해 게이야")
  for {
    url, err := wd.CurrentURL()
    if err != nil {
      continue
    }

    if url == "https://www.naver.com/" {
      log.Println("아구 잘해써~")
      break;  
    }
  }

  var elem selenium.WebElement

  for {
    log.Println("팔고있는가?")
    wd.Get(os.Args[1])

    elem, err = wd.FindElement(selenium.ByCSSSelector, "._2-uvQuRWK5")
    if err != nil {
      log.Println(err)
      log.Println("ㄴㄴ 없음")
      continue;
    }

    log.Println("ㅇㅇ 있음!")
    break
  }

  for {
    log.Println("페이지 바뀌었는지 확인")
    url, err := wd.CurrentURL()
    if err != nil {
      continue
    }

    if strings.HasPrefix(url, "https://order.pay.naver.com/orderSheet") {
      log.Println("결제 페이지로 이동됨")
      break
    }

    elem.Click()
  }

  for {
    url, err := wd.CurrentURL()
    if err != nil {
      continue
    }

    if strings.HasPrefix(url, "https://order.pay.naver.com/orderSheet/result") {
      break
    }

    log.Println("일반 결제 클릭 시도")
    elem, err = wd.FindElement(selenium.ByCSSSelector, "label[for=generalPaymentsRadio]")
    if err != nil {
      continue
    }

    err = elem.Click()
    if err != nil {
      continue
    }

    log.Println("나중에 결제 클릭 시도")
    elem, err = wd.FindElement(selenium.ByCSSSelector, "label[for=pay20]")
    if err != nil {
      continue
    }

    sel, err := elem.IsSelected()
    log.Println(sel, err)


    err = elem.Click()
    if err != nil {
      continue
    }

    log.Println("결제하기 버튼 클릭 시도")
    elem, err = wd.FindElement(selenium.ByCSSSelector, "._doPayButton")
    if err != nil {
      continue
    }

    elem.Click()
    time.Sleep(time.Second)
  }

  log.Println("됬다 게이야!!")
  for {}
}
