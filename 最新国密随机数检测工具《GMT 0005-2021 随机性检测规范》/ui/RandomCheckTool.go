package main

import (
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"math"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Trisia/randomness"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

//go get github.com/andlabs/libui
//go get github.com/andlabs/ui

var mainwin *ui.Window

var itemStr = [25]string{"单比特频数检测", "块内频数检测 m=10000", "扑克检测 m=4", "扑克检测 m=8", "重叠子序列检测 m=3 P1", "重叠子序列检测 m=3 P2", "重叠子序列检测 m=5 P1", "重叠子序列检测 m=5 P2", "游程总数检测", "游程分布检测", "块内最大游程检测 m=10000", "二元推导检测 k=3", "二元推导检测 k=7", "自相关检测 d=1", "自相关检测 d=2", "自相关检测d=8", "自相关检测 d=16", "矩阵秩检测", "累加和检测", "近似熵检测 m=2", "近似熵检测 m=5", "线性复杂度检测 m=500", "线性复杂度检测 m=1000", "Maurer通用统计检测 L=7 Q=1280", "离散傅里叶检测"}

var selectArry [25]int
var selectslice []int

var checkboxs [len(itemStr)]*ui.Checkbox
var labels_succ [len(itemStr)]*ui.Label
var labels_fail [len(itemStr)]*ui.Label

var failcnt = make([]int32, len(itemStr))

var labelsetbit *ui.Label
var labelsetNum *ui.Label
var labelsetAlphaNum *ui.Label
var labeltime *ui.Label
var labeltimebegin *ui.Label
var labeltimeend *ui.Label
var buttonOK *ui.Button
var setcnt int

// Alpha 显著性水平α
// AlphaT 分布均匀性的显著性水平
// const AlphaT float64 = 0.0001
var Alpha float64
var entryAlpha *ui.Entry

var entryPath *ui.Entry
var entryOutPath *ui.Entry

var prosessbar *ui.ProgressBar

var seletctAll bool

type R struct {
	Name string
	P    []float64
}

func worker(jobs <-chan string, out chan<- *R) {
	for filename := range jobs {
		buf, _ := os.ReadFile(filename)
		bits := randomness.B2bitArr(buf)
		buf = nil
		arr := make([]float64, 0, len(selectslice))


		if selectArry[0] == 1 {
			//"单比特频数检测"
			p, _ := randomness.MonoBitFrequencyTest(bits)
			arr = append(arr, p)
		}		

		if selectArry[1] == 1 {
			//"块内频数检测 m=10000"
			p, _ := randomness.FrequencyWithinBlockTest(bits)
			arr = append(arr, p)
		}		

		if selectArry[2] == 1 {
			//"扑克检测 m=4"
			//"扑克检测 m=8"
			p, _ := randomness.PokerProto(bits, 4)
			arr = append(arr, p)
		}		

		if selectArry[3] == 1 {
			p, _ := randomness.PokerProto(bits, 8)
			arr = append(arr, p)
		}		

		if selectArry[4] == 1 || selectArry[5] == 1 {
			//"重叠子序列检测 m=3 P1"
			//"重叠子序列检测 m=3 P2"
			//"重叠子序列检测 m=5 P1"
			//"重叠子序列检测 m=5 P2"
			p1, p2, _, _ := randomness.OverlappingTemplateMatchingProto(bits, 3)
			if selectArry[4] == 1 {
				arr = append(arr, p1)
			}
			if selectArry[5] == 1 {
				arr = append(arr, p2)
			}
		}

		if selectArry[6] == 1 || selectArry[7] == 1 {
			p1, p2, _, _ := randomness.OverlappingTemplateMatchingProto(bits, 5)
			if selectArry[6] == 1 {
				arr = append(arr, p1)
			}
			if selectArry[7] == 1 {
				arr = append(arr, p2)
			}
		}

		if selectArry[8] == 1 {
			//"游程总数检测"
			//"游程分布检测"
			//"块内最大游程检测 m=10000"
			p, _ := randomness.RunsTest(bits)
			arr = append(arr, p)
		}		

		if selectArry[9] == 1 {
			p, _ := randomness.RunsDistributionTest(bits)
			arr = append(arr, p)
		}		

		if selectArry[10] == 1 {
			p, _ := randomness.LongestRunOfOnesInABlockTest(bits, true)
			arr = append(arr, p)
		}		

		if selectArry[11] == 1 {
			//"二元推导检测 k=3"
			//"二元推导检测 k=7"
			p, _ := randomness.BinaryDerivativeProto(bits, 3)
			arr = append(arr, p)
		}		

		if selectArry[12] == 1 {
			p, _ := randomness.BinaryDerivativeProto(bits, 7)
			arr = append(arr, p)
		}		

		if selectArry[13] == 1 {
			//"自相关检测 d=1"
			//"自相关检测 d=2"
			//"自相关检测d=8"
			//"自相关检测 d=16"
			p, _ := randomness.AutocorrelationProto(bits, 1)
			arr = append(arr, p)
		}		

		if selectArry[14] == 1 {
			p, _ := randomness.AutocorrelationProto(bits, 2)
			arr = append(arr, p)
		}		

		if selectArry[15] == 1 {
			p, _ := randomness.AutocorrelationProto(bits, 8)
			arr = append(arr, p)
		}		

		if selectArry[16] == 1 {
			p, _ := randomness.AutocorrelationProto(bits, 16)
			arr = append(arr, p)
		}		

		if selectArry[17] == 1 {
			//"矩阵秩检测"
			p, _ := randomness.MatrixRankTest(bits)
			arr = append(arr, p)
		}		

		if selectArry[18] == 1 {
			//"累加和检测"
			p, _ := randomness.CumulativeTest(bits, true)
			arr = append(arr, p)
		}		

		if selectArry[19] == 1 {
			//"近似熵检测 m=2"
			//"近似熵检测 m=5"
			p, _ := randomness.ApproximateEntropyProto(bits, 2)
			arr = append(arr, p)
		}		

		if selectArry[20] == 1 {
			p, _ := randomness.ApproximateEntropyProto(bits, 5)
			arr = append(arr, p)
		}		

		if selectArry[21] == 1 {
			//"线性复杂度检测 m=500"
			//"线性复杂度检测 m=1000"
			p, _ := randomness.LinearComplexityProto(bits, 500)
			arr = append(arr, p)
		}		

		if selectArry[22] == 1 {
			p, _ := randomness.LinearComplexityProto(bits, 1000)
			arr = append(arr, p)
		}		

		if selectArry[23] == 1 {
			//"Maurer通用统计检测 L=7 Q=1280"
			p, _ := randomness.MaurerUniversalTest(bits)
			arr = append(arr, p)
		}		

		if selectArry[24] == 1 {
			//"离散傅里叶检测"
			p, _ := randomness.DiscreteFourierTransformTest(bits)
			arr = append(arr, p)
		}

		go func(file string) {
			out <- &R{path.Base(file), arr}
		}(filename)
	}
}

//var lock sync.Mutex

// 结果集写入文件工作器
func resultWriter(in <-chan *R, w io.StringWriter, cnt []int32, wg *sync.WaitGroup) {
	for r := range in {
		_, _ = w.WriteString(r.Name)

		//lock.Lock()

		for j := 0; j < len(r.P); j++ {
			if r.P[j] >= 0.01 {
				atomic.AddInt32(&cnt[j], 1)
				//崩溃
				//labels_succ[selectslice[j]].SetText(fmt.Sprintf("%d", cnt[j]))
			} else {
				atomic.AddInt32(&failcnt[j], 1)
				//labels_fail[selectslice[j]].SetText(fmt.Sprintf("%d", failcnt[j]))
			}
			_, _ = w.WriteString(fmt.Sprintf(", %0.6f", r.P[j]))
		}
		//lock.Unlock()
		_, _ = w.WriteString("\n")

		wg.Done()
	}
}

func makeBasicControlsPage() ui.Control {

	Alpha = 0.01

	//垂直布局
	hboxMain := ui.NewHorizontalBox()
	hboxMain.SetPadded(true)

	//水平布局
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	hboxMain.Append(vbox, true)

	//垂直布局
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	vbox.Append(hbox, true)

	//水平布局左
	vboxleft := ui.NewVerticalBox()
	vboxleft.SetPadded(false)
	hbox.Append(vboxleft, false)

	hbox.Append(ui.NewVerticalSeparator(), false)

	//水平布局右
	vboxright := ui.NewVerticalBox()
	vboxright.SetPadded(false)
	hbox.Append(vboxright, true)

	hbox.Append(ui.NewVerticalSeparator(), false)

	vboxstat := ui.NewVerticalBox()
	vboxstat.SetPadded(false)
	hbox.Append(vboxstat, true)

	vboxleft.Append(ui.NewLabel("检测项"), false)
	vboxright.Append(ui.NewLabel(">=显著水平样本数"), false)
	vboxstat.Append(ui.NewLabel("<显著水平样本数"), false)

	seletctAll = true
	for i := 0; i < len(itemStr); i++ {
		checkboxs[i] = ui.NewCheckbox(itemStr[i])
		checkboxs[i].SetChecked(seletctAll)
		labels_succ[i] = ui.NewLabel("0")
		labels_fail[i] = ui.NewLabel("0")

		vboxleft.Append(checkboxs[i], true)
		vboxleft.Append(ui.NewHorizontalSeparator(), false)
		vboxright.Append(labels_succ[i], true)
		vboxright.Append(ui.NewHorizontalSeparator(), false)
		vboxstat.Append(labels_fail[i], true)
		vboxstat.Append(ui.NewHorizontalSeparator(), false)
	}

	buttonSelectAll := ui.NewButton("")
	if true == seletctAll {
		buttonSelectAll.SetText("反选")
	} else if false == seletctAll {
		buttonSelectAll.SetText("全选")
	}

	buttonSelectAll.OnClicked(func(*ui.Button) {

		seletctAll = !seletctAll
		for i := 0; i < len(itemStr); i++ {
			checkboxs[i].SetChecked(seletctAll)
		}
		if true == seletctAll {
			buttonSelectAll.SetText("全选")
		} else if false == seletctAll {
			buttonSelectAll.SetText("反选")
		}
	})

	hboxAl := ui.NewHorizontalBox()
	hboxAl.SetPadded(true)
	vbox.Append(hboxAl, false)

	hboxAl.Append(buttonSelectAll, true)
	hboxAl.Append(ui.NewLabel("显著水平:"), false)
	entryAlpha = ui.NewEntry()
	entryAlpha.SetReadOnly(true)
	entryAlpha.SetText(fmt.Sprintf("%1.3f", Alpha))
	hboxAl.Append(entryAlpha, false)

	//右侧视图
	vboxOp := ui.NewVerticalBox()
	vboxOp.SetPadded(true)
	hboxMain.Append(ui.NewVerticalSeparator(), false)
	hboxMain.Append(vboxOp, true)

	//进度条
	prosessbar = ui.NewProgressBar()
	prosessbar.SetValue(0)
	//vbox.Append(prosessbar, false)
	vboxOp.Append(prosessbar, true)

	//其他
	hbox1 := ui.NewHorizontalBox()
	hbox1.SetPadded(true)
	hbox1.Append(ui.NewLabel("单样本大小:"), true)
	labelsetbit = ui.NewLabel("")
	hbox1.Append(labelsetbit, true)
	vboxOp.Append(hbox1, false)

	hbox1 = ui.NewHorizontalBox()
	hbox1.SetPadded(true)
	hbox1.Append(ui.NewLabel("样本数量:"), true)
	labelsetNum = ui.NewLabel("")
	hbox1.Append(labelsetNum, true)
	vboxOp.Append(hbox1, false)

	hbox1 = ui.NewHorizontalBox()
	hbox1.SetPadded(true)
	hbox1.Append(ui.NewLabel("成功通过检测项,需要>=显著水平的样本数量:"), true)
	labelsetAlphaNum = ui.NewLabel("")
	hbox1.Append(labelsetAlphaNum, true)
	vboxOp.Append(hbox1, false)

	//添加分割线
	//vbox.Append(ui.NewHorizontalSeparator(), false)
	vboxOp.Append(ui.NewHorizontalSeparator(), false)

	//样本路径
	hboxpath := ui.NewHorizontalBox()
	hboxpath.SetPadded(true)

	entryOutPath := ui.NewEntry()

	button := ui.NewButton("选择样本文件")
	entryPath = ui.NewEntry()
	entryPath.SetReadOnly(true)
	button.OnClicked(func(*ui.Button) {
		setcnt = 0

		filename := ui.OpenFile(mainwin)
		if filename != "" {
			/*
				filepath.Walk(filepath.Dir(filename), func(p string, _ fs.FileInfo, _ error) error {
					if strings.HasSuffix(p, ".bin") || strings.HasSuffix(p, ".dat") {
						setcnt++
					}
					return nil
				})
			*/
			files, err := ioutil.ReadDir(filepath.Dir(filename))
			if err == nil {
				for _, file := range files {

					if strings.HasSuffix(file.Name(), ".bin") || strings.HasSuffix(file.Name(), ".dat") {
						setcnt++
					}
					//fmt.Println(file.Name())
				}
			}
		}

		if setcnt == 0 {
			ui.MsgBoxError(mainwin, "国密随机数检测工具", "样本文件必须是 .dat 或 .bin文件")
			labelsetAlphaNum.SetText("")
			labelsetNum.SetText("")
		} else {

			entryPath.SetText(filename)
			buf, _ := os.ReadFile(filename)
			labelsetbit.SetText(fmt.Sprintf("%dbit %dbyte", len(buf)*8, len(buf)))

			outpath := path.Join(filepath.Dir(filename), "/RandomnessTestReport.csv")
			//outpath := filepath.Dir(filename) + "/RandomnessTestReport.csv"
			_ = os.MkdirAll(filepath.Dir(outpath), os.FileMode(0600))

			entryOutPath.SetText(outpath)
			var num float64 = (1 - Alpha - 3*math.Sqrt(float64((Alpha*(1-Alpha))/float64(setcnt)))) * float64(setcnt)
			//fmt.Printf("%f %d", num, int(math.Ceil(num)))
			labelsetAlphaNum.SetText(fmt.Sprintf("%d", int(math.Ceil(num))))
			labelsetNum.SetText(fmt.Sprintf("%d", setcnt))
		}
	})
	hboxpath.Append(ui.NewLabel("样本路径:"), false)
	hboxpath.Append(entryPath, true)
	hboxpath.Append(button, false)

	//vbox.Append(hboxpath, false)
	vboxOp.Append(hboxpath, false)

	//报告路径
	hboxpath = ui.NewHorizontalBox()
	hboxpath.SetPadded(true)

	buttonOK = ui.NewButton("开 始 检 测 ")
	buttonOK.OnClicked(func(*ui.Button) {

		filename := entryOutPath.Text()
		if filename != "" {

			buttonOK.Disable()
			button.Disable()
			_ = os.MkdirAll(filepath.Dir(filename), os.FileMode(0600))

			selectslice = selectslice[0:0]
			for i := 0; i < len(itemStr); i++ {
				if checkboxs[i].Checked() {
					selectslice = append(selectslice, i)
					selectArry[i] = 1
				} else {
					selectArry[i] = 0
				}
			}
			for i := 0; i < len(failcnt); i++ {
				failcnt[i] = 0
			}

			//线程
			go func() {

				startTime := time.Now()
				startTimestr := startTime.Format("2006.01.02 15:04:05")
				labeltimebegin.SetText(startTimestr)

				n := runtime.NumCPU()
				out := make(chan *R)
				jobs := make(chan string)

				inputPath := entryPath.Text()
				if filename == "" {
					fmt.Printf(">> 请选择样本文件")
					ui.MsgBoxError(mainwin, "国密随机数检测工具", "请选择样本文件 ")
					buttonOK.Enable()
					button.Enable()
					return
				}

				inputPath = filepath.Dir(inputPath)
				reportPath := entryOutPath.Text()
				w, err := os.OpenFile(reportPath, os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.FileMode(0600))
				if err != nil {
					fmt.Printf(">> 无法打开写入报告文件")
					ui.MsgBoxError(mainwin, "国密随机数检测工具", "无法打开写入文件 "+reportPath)
					buttonOK.Enable()
					button.Enable()
					return
				}
				defer w.Close()

				var reportheader string
				reportheader = "源数据,"

				for i := 0; i < len(selectslice); i++ {
					reportheader += itemStr[selectslice[i]]
					reportheader += ","
				}
				reportheader += "\n"
				_, _ = w.WriteString(reportheader)

				var wg sync.WaitGroup
				var cnt = make([]int32, len(selectslice))

				wg.Add(setcnt)

				// 启动数据写入消费者
				go resultWriter(out, w, cnt, &wg)
				// 检测工作器
				for i := 0; i < n; i++ {
					go worker(jobs, out)
				}

				process := 0
				var percent int
				if setcnt > 100 {
					percent = (setcnt / 100)
				} else {
					percent = (100 / setcnt)
				}

				// 结果工作器
				go filepath.Walk(inputPath, func(p string, _ fs.FileInfo, _ error) error {
					if strings.HasSuffix(p, ".bin") || strings.HasSuffix(p, ".dat") {
						process++
						//fmt.Printf("%d %d \n", process, percent)
						if setcnt > 100 {
							prosessbar.SetValue(process / percent)
						} else {
							prosessbar.SetValue(process * percent)
						}

						for j := 0; j < len(selectslice); j++ {
							labels_succ[selectslice[j]].SetText(fmt.Sprintf("%d", cnt[j]))
							labels_fail[selectslice[j]].SetText(fmt.Sprintf("%d", failcnt[j]))
						}

						labeltime.SetText(time.Now().Format("2006.01.02 15:04:05"))
						jobs <- p
					}
					return nil
				})

				wg.Wait()

				_, _ = w.WriteString("总计")
				for i := 0; i < len(cnt); i++ {
					_, _ = w.WriteString(fmt.Sprintf(", %d", cnt[i]))
				}
				_, _ = w.WriteString("\n")

				prosessbar.SetValue(100)

				elapsedTime := time.Since(startTime) / time.Second // duration in s
				labeltime.SetText(fmt.Sprintf("%d 分钟 %d 秒", elapsedTime/60, elapsedTime%60))
				labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))

				for j := 0; j < len(selectslice); j++ {
					labels_succ[selectslice[j]].SetText(fmt.Sprintf("%d", cnt[j]))
					labels_fail[selectslice[j]].SetText(fmt.Sprintf("%d", failcnt[j]))
				}

				var errStr string
				var numCount int = 0
				//通过的组数
				var numpass int32 = int32(math.Ceil((1 - Alpha - 3*math.Sqrt(float64((Alpha*(1-Alpha))/float64(setcnt)))) * float64(setcnt)))
				for j := 0; j < len(selectslice); j++ {

					//num, _ := strconv.Atoi(labels_succ[selectslice[j]].Text())
					if cnt[j] < numpass {
						numCount++
						errStr += itemStr[selectslice[j]]
						errStr += ",\n"
					}
				}

				if 0 == numCount {
					ui.MsgBox(mainwin, "成功", "所有检测项全部通过")
				} else {
					ui.MsgBoxError(mainwin, "失败", fmt.Sprintf("有%d项检测未通过:%s\n", numCount, errStr))
				}

				buttonOK.Enable()
				button.Enable()
			}()
		}
	})

	hboxpath.Append(ui.NewLabel("报告路径:"), false)
	hboxpath.Append(entryOutPath, true)
	hboxpath.Append(buttonOK, false)

	//vbox.Append(hboxpath, false)
	vboxOp.Append(hboxpath, false)

	hbox1 = ui.NewHorizontalBox()
	hbox1.SetPadded(true)
	hbox1.Append(ui.NewLabel("开始测试时间:"), true)
	labeltimebegin = ui.NewLabel("")
	hbox1.Append(labeltimebegin, true)
	vboxOp.Append(hbox1, false)

	hbox1 = ui.NewHorizontalBox()
	hbox1.SetPadded(true)
	hbox1.Append(ui.NewLabel("测试完成时间:"), true)
	labeltimeend = ui.NewLabel("")
	hbox1.Append(labeltimeend, true)
	vboxOp.Append(hbox1, false)

	hbox1 = ui.NewHorizontalBox()
	hbox1.SetPadded(true)
	hbox1.Append(ui.NewLabel("耗时:"), true)
	labeltime = ui.NewLabel("")
	hbox1.Append(labeltime, true)
	vboxOp.Append(hbox1, false)

	grid := ui.NewGrid()
	grid.SetPadded(false)

	grid.Append(ui.NewLabel("国密随机数质量检测工具,支持《GMT 0005-2021 随机性检测规范》。"),
		0, 0, 1, 1,
		true, ui.AlignCenter, true, ui.AlignCenter)
	grid.Append(ui.NewLabel("利用多核CPU进行检测,大大缩短检测时间。"),
		0, 1, 1, 1,
		true, ui.AlignCenter, false, ui.AlignCenter)

	grid.Append(ui.NewLabel("北京世纪龙脉科技有限公司 V1.1"),
		0, 2, 1, 1,
		true, ui.AlignCenter, true, ui.AlignCenter)

	//vbox.Append(grid, false)
	vboxOp.Append(grid, true)

	return hboxMain
}

func setupUI() {
	mainwin = ui.NewWindow("国密随机数检测工具-北京世纪龙脉科技有限公司", 640, 480, true)
	mainwin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	tab := ui.NewTab()
	mainwin.SetChild(tab)
	mainwin.SetMargined(true)

	tab.Append("GMT 0005-2021 随机性检测", makeBasicControlsPage())
	tab.SetMargined(0, true)

	//labels_succ[0].SetText("1")
	mainwin.Show()
}

func main() {
	ui.Main(setupUI)
}

//隐藏命令行
//go build -ldflags="-H windowsgui"  RandomCheckTool.go

//图标
//go get github.com/akavel/rsrc
//main.rc
//IDI_ICON1 ICON "logo.ico"
//windres -o main.syso main.rc
//go build -ldflags="-H windowsgui"
