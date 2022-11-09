package main

import (
	"fmt"
	"io"
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

	"golang.org/x/text/encoding/simplifiedchinese"

	"github.com/Trisia/randomness"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

var mainwin *ui.Window

var itemStr = [27]string{"单比特频数检测", "块内频数检测 m=10000", "扑克检测 m=4", "扑克检测 m=8", "重叠子序列检测 m=3 P1", "重叠子序列检测 m=3 P2", "重叠子序列检测 m=5 P1", "重叠子序列检测 m=5 P2", "游程总数检测", "游程分布检测", "块内最大1游程检测 m=10000", "块内最大0游程检测 m=10000", "二元推导检测 k=3", "二元推导检测 k=7", "自相关检测 d=1", "自相关检测 d=2", "自相关检测 d=8", "自相关检测 d=16", "矩阵秩检测", "累加和前向检测", "累加和后向检测", "近似熵检测 m=2", "近似熵检测 m=5", "线性复杂度检测 m=500", "线性复杂度检测 m=1000", "Maurer通用统计检测 L=7 Q=1280", "离散傅里叶检测"}

var selectArry [len(itemStr)]int
var selectslice []int
var cnt []int32

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
var buttonFile *ui.Button

var setcnt int

// Alpha 显著性水平α
// AlphaT 分布均匀性的显著性水平
// const AlphaT float64 = 0.0001
var Alpha float64
var entryAlpha *ui.Entry

var entryPath *ui.Entry
var entryOutPath *ui.Entry

var reportPath string
var inputPath string

var prosessbar *ui.ProgressBar

var seletctAll bool

var running int32

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
			//"块内频数检测 m=10000" 10  100 1000 10000 1000000 根据长度自动选择
			p, _ := randomness.FrequencyWithinBlockTest(bits)
			arr = append(arr, p)
		}

		if selectArry[2] == 1 {
			//"扑克检测 m=4"  至少8字节
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
			p, _ := randomness.RunsTest(bits)
			arr = append(arr, p)
		}

		if selectArry[9] == 1 {
			//"游程分布检测"
			p, _ := randomness.RunsDistributionTest(bits)
			arr = append(arr, p)
		}

		//至少128字节
		if selectArry[10] == 1 {
			//"块内最大1游程检测 m=10000"
			p, _ := randomness.LongestRunOfOnesInABlockTest(bits, true)
			arr = append(arr, p)
		}

		if selectArry[11] == 1 {
			//"块内最大0游程检测 m=10000"
			p, _ := randomness.LongestRunOfOnesInABlockTest(bits, false)
			arr = append(arr, p)
		}

		if selectArry[12] == 1 {
			//"二元推导检测 k=3" 至少7字节
			p, _ := randomness.BinaryDerivativeProto(bits, 3)
			arr = append(arr, p)
		}

		if selectArry[13] == 1 {
			//"二元推导检测 k=7"
			p, _ := randomness.BinaryDerivativeProto(bits, 7)
			arr = append(arr, p)
		}

		if selectArry[14] == 1 {
			//"自相关检测 d=1"
			//"自相关检测 d=2"
			//"自相关检测d=8"
			//"自相关检测 d=16" 至少16字节
			p, _ := randomness.AutocorrelationProto(bits, 1)
			arr = append(arr, p)
		}

		if selectArry[15] == 1 {
			p, _ := randomness.AutocorrelationProto(bits, 2)
			arr = append(arr, p)
		}

		if selectArry[16] == 1 {
			p, _ := randomness.AutocorrelationProto(bits, 8)
			arr = append(arr, p)
		}

		if selectArry[17] == 1 {
			p, _ := randomness.AutocorrelationProto(bits, 16)
			arr = append(arr, p)
		}

		if selectArry[18] == 1 {
			//"矩阵秩检测"
			p, _ := randomness.MatrixRankTest(bits)
			arr = append(arr, p)
		}

		if selectArry[19] == 1 {
			//"累加和检测"  前向
			p, _ := randomness.CumulativeTest(bits, true)
			arr = append(arr, p)
		}
		if selectArry[20] == 1 {
			//"累加和检测"  后向
			p, _ := randomness.CumulativeTest(bits, false)
			arr = append(arr, p)
		}

		if selectArry[21] == 1 {
			//"近似熵检测 m=2"
			//"近似熵检测 m=5"
			p, _ := randomness.ApproximateEntropyProto(bits, 2)
			arr = append(arr, p)
		}

		if selectArry[22] == 1 {
			p, _ := randomness.ApproximateEntropyProto(bits, 5)
			arr = append(arr, p)
		}

		if selectArry[23] == 1 {
			//"线性复杂度检测 m=500"
			//"线性复杂度检测 m=1000"
			p, _ := randomness.LinearComplexityProto(bits, 500)
			arr = append(arr, p)
		}

		if selectArry[24] == 1 {
			p, _ := randomness.LinearComplexityProto(bits, 1000)
			arr = append(arr, p)
		}

		if selectArry[25] == 1 {
			//"Maurer通用统计检测 L=7 Q=1280"
			p, _ := randomness.MaurerUniversalTest(bits)
			arr = append(arr, p)
		}

		if selectArry[26] == 1 {
			//"离散傅里叶检测"
			p, _ := randomness.DiscreteFourierTransformTest(bits)
			arr = append(arr, p)
		}

		go func(file string) {
			out <- &R{path.Base(file), arr}
		}(filename)
	}
}

// 结果集写入文件工作器
func resultWriter(in <-chan *R, w io.StringWriter, cnt []int32, wg *sync.WaitGroup) {
	for r := range in {
		_, _ = w.WriteString(r.Name)

		//lock.Lock()

		for j := 0; j < len(r.P); j++ {
			if r.P[j] >= 0.01 {
				atomic.AddInt32(&cnt[j], 1)
			} else {
				atomic.AddInt32(&failcnt[j], 1)
			}
			_, _ = w.WriteString(fmt.Sprintf(", %0.6f", r.P[j]))
		}
		_, _ = w.WriteString("\n")

		wg.Done()
	}
}

func StrToGBK(str string) string {
	sysType := runtime.GOOS
	if sysType == "windows" {
		//将utf-8编码的字符串转换为GBK编码
		ret, _ := simplifiedchinese.GBK.NewEncoder().String(str)
		return ret //如果转换失败返回空字符串
	} else {
		return str
	}

	//如果是[]byte格式的字符串，可以使用Bytes方法
	//b, err := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(str))
	//return string(b)
}

var startTime time.Time
var processfile int32
var percent int32

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
	vboxOp.Append(ui.NewLabel("建议使用1000*1000 或者1024*1024 bit大小的样本进行检测。"), false)
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

	buttonFile = ui.NewButton("选择样本文件")
	entryPath = ui.NewEntry()
	entryPath.SetReadOnly(true)
	buttonFile.OnClicked(func(*ui.Button) {
		setcnt = 0
		entryPath.SetText("")
		entryOutPath.SetText("")
		filename := ui.OpenFile(mainwin)
		if filename != "" {
			files, err := ioutil.ReadDir(filepath.Dir(filename))
			if err == nil {
				for _, file := range files {

					if strings.HasSuffix(file.Name(), ".bin") || strings.HasSuffix(file.Name(), ".dat") {
						setcnt++
					}
				}
			}
		}

		if setcnt == 0 {
			ui.MsgBoxError(mainwin, "随机数检测工具", "样本文件必须是 .dat 或 .bin文件")
			labelsetAlphaNum.SetText("")
			labelsetNum.SetText("")
		} else {

			buf, _ := os.ReadFile(filename)
			labelsetbit.SetText(fmt.Sprintf("%dbit %dbyte", len(buf)*8, len(buf)))

			if len(buf) < (7 * 1280) {
				//checkboxs[25].SetChecked(false)
				ui.MsgBoxError(mainwin, "随机数检测工具", fmt.Sprintf("Maurer通用统计检测 数据长度至少要满足 L*Q %dbyte", 7*1280))
				setcnt = 0
			} else {
				entryPath.SetText(filename)
				outpath := path.Join(filepath.Dir(filepath.Dir(filename)), "/RandomnessTestReport.csv")
				//outpath := filepath.Dir(filename) + "/RandomnessTestReport.csv"
				_ = os.MkdirAll(filepath.Dir(outpath), os.FileMode(0600))

				entryOutPath.SetText(outpath)
				var num float64 = (1 - Alpha - 3*math.Sqrt(float64((Alpha*(1-Alpha))/float64(setcnt)))) * float64(setcnt)
				//fmt.Printf("%f %d", num, int(math.Ceil(num)))
				labelsetAlphaNum.SetText(fmt.Sprintf("%d", int(math.Ceil(num))))
				labelsetNum.SetText(fmt.Sprintf("%d", setcnt))
			}
		}
	})
	hboxpath.Append(ui.NewLabel("样本路径:"), false)
	hboxpath.Append(entryPath, true)
	hboxpath.Append(buttonFile, false)

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
			buttonFile.Disable()

			for j := 0; j < len(itemStr); j++ {
				labels_succ[j].SetText("0")
				labels_fail[j].SetText("0")
			}
			prosessbar.SetValue(0)
			labeltime.SetText("")
			labeltimeend.SetText("")
			labeltimebegin.SetText("")

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

			inputPath = entryPath.Text()

			inputPath = filepath.Dir(inputPath)
			reportPath = entryOutPath.Text()

			startTime = time.Now()
			labeltimebegin.SetText(startTime.Format("2006.01.02 15:04:05"))

			//线程
			go func() {

				n := runtime.NumCPU()
				out := make(chan *R)
				jobs := make(chan string)

				//8核心一下的cpu使用一半的线程数
				if n <= 8 {
					n = n / 2
				}

				w, err := os.OpenFile(reportPath, os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.FileMode(0600))
				if err != nil {
					fmt.Printf(">> 无法打开写入报告文件")
					//ui.MsgBoxError(mainwin, "随机数检测工具", "无法打开写入文件 "+reportPath)
					running = 0
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

				reportheader = StrToGBK(reportheader)
				_, _ = w.WriteString(reportheader)

				processfile = 0
				if setcnt > 100 {
					percent = int32(setcnt / 100)
				} else {
					percent = int32(100 / setcnt)
				}

				files, err := ioutil.ReadDir(inputPath)
				if err != nil {
					//ui.MsgBoxError(mainwin, "随机数检测工具", "打开写入文件失败")
					running = 0
					return
				}

				var wg sync.WaitGroup
				cnt = make([]int32, len(selectslice))

				running = 1

				wg.Add(setcnt)

				// 启动数据写入消费者
				go resultWriter(out, w, cnt, &wg)
				// 检测工作器
				for i := 0; i < n; i++ {
					go worker(jobs, out)
				}

				// 结果工作器
				for _, file := range files {
					if strings.HasSuffix(file.Name(), ".bin") || strings.HasSuffix(file.Name(), ".dat") {

						atomic.AddInt32(&processfile, 1)

						jobs <- (path.Join(inputPath, "/", file.Name()))
					}
				}

				wg.Wait()

				running = 0

				reportheader = "总计"
				reportheader = StrToGBK(reportheader)
				_, _ = w.WriteString(reportheader)
				for i := 0; i < len(cnt); i++ {
					_, _ = w.WriteString(fmt.Sprintf(", %d", cnt[i]))
				}
				_, _ = w.WriteString("\n")

			}()
		} else {
			ui.MsgBoxError(mainwin, "随机数检测工具", "请选择样本文件 ")
			return
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

	grid.Append(ui.NewLabel("北京世纪龙脉科技有限公司 V1.4"),
		0, 2, 1, 1,
		true, ui.AlignCenter, true, ui.AlignCenter)

	//vbox.Append(grid, false)
	vboxOp.Append(grid, true)

	return hboxMain
}

func setupUI() {
	mainwin = ui.NewWindow("随机数检测工具-北京世纪龙脉科技有限公司", 640, 480, true)
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

	mainwin.Show()
	go counter()
}

func counter() {

	for {
		time.Sleep(1 * time.Second)

		ui.QueueMain(func() {

			if 1 == running {

				if setcnt > 100 {
					prosessbar.SetValue(int(processfile / percent))
				} else {
					prosessbar.SetValue(int(processfile * percent))
				}

				labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))
				for j := 0; j < len(selectslice); j++ {
					labels_succ[selectslice[j]].SetText(fmt.Sprintf("%d", cnt[j]))
					labels_fail[selectslice[j]].SetText(fmt.Sprintf("%d", failcnt[j]))
				}
			} else if 0 == running {

				prosessbar.SetValue(100)

				elapsedTime := time.Since(startTime) / time.Second // duration in s
				labeltime.SetText(fmt.Sprintf("%d 分钟 %d 秒", elapsedTime/60, elapsedTime%60))
				labeltimeend.SetText(time.Now().Format("2006.01.02 15:04:05"))

				buttonOK.Enable()
				buttonFile.Enable()

				for j := 0; j < len(selectslice); j++ {
					labels_succ[selectslice[j]].SetText(fmt.Sprintf("%d", cnt[j]))
					labels_fail[selectslice[j]].SetText(fmt.Sprintf("%d", failcnt[j]))
				}

				var errStr string
				var numCount int = 0
				//通过的组数
				var numpass int32 = int32(math.Ceil((1 - Alpha - 3*math.Sqrt(float64((Alpha*(1-Alpha))/float64(setcnt)))) * float64(setcnt)))
				for j := 0; j < len(selectslice); j++ {

					if cnt[j] < numpass {
						numCount++
						errStr += itemStr[selectslice[j]]
						errStr += ",\n"
					}
				}

				running = 3

				if 0 == numCount {
					ui.MsgBox(mainwin, "成功", "所有检测项全部通过")
				} else {
					ui.MsgBoxError(mainwin, "失败", fmt.Sprintf("有%d项检测未通过:%s\n", numCount, errStr))
				}

				processfile = 0
				percent = 1

			} else {

			}
		})
	}
}
func main() {
	running = 3
	percent = 1
	ui.Main(setupUI)
}

//go 大于1.8
//Debian、Ubuntu 等：sudo apt-get install libgtk-3-dev
//Red Hat / Fedora 等：sudo dnf install gtk3-devel
//go get github.com/andlabs/libui
//go get github.com/andlabs/ui
//go get golang.org/x/text
//go get github.com/Trisia/randomness
//go get github.com/josephspurrier/goversioninfo/cmd/goversioninfo

//apt-get install golang-go
//wget https://dl.google.com/go/go1.19.3.linux-amd64.tar.gz
//tar -zxvf go1.19.3.linux-amd64.tar.gz
//解压到/home/lm/Desktop/
//export GOROOT=/home/lm/Desktop/go
//export GOPATH=/home/lm/go
//export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
//go version
//go mod init RandomCheckTool

//隐藏命令行
//go build -ldflags="-H windowsgui"  RandomCheckTool.go

//图标
//go get github.com/akavel/rsrc
//main.rc
//IDI_ICON1 ICON "logo.ico"
//windres -o main.syso main.rc
//windres -i rc/main.rc -O coff -o main.syso
//go build -o RandomCheckTool.exe  -ldflags="-H windowsgui"

//其它
//go get github.com/josephspurrier/goversioninfo/cmd/goversioninfo

//编译andlabs/libui
//apt-get install ninja-build
//pip3 install --user meson

//export PATH=$PATH:/root/.local/bin

//cd /usr/local
//wget --no-check-certificate https://www.python.org/ftp/python/3.6.5/Python-3.6.5.tgz
//wget --no-check-certificate https://www.python.org/ftp/python/3.7.5/Python-3.7.5.tgz
//tar -xzvf Python-3.6.5.tgz
//cd Python-3.6.5
//./configure --prefix=/usr/local/python3.6.5
//make
//make install
//cp /usr/bin/python3 /usr/bin/python3_bak
//rm /usr/bin/python3
//ln -s /usr/local/Python-3.6.5/python /usr/bin/python3

//wget https://bootstrap.pypa.io/pip/3.6/get-pip.py
//python3 get-pip.py

//grep 取行，awk 按条件取指定列，cut 按分隔符取指定列。
//sed主要是用来将数据进行选取、替换、删除、新增的命令。可以放在管道符之后处理。

// ldd RandomCheckTool | cut -d ">" -f 2 |cut -d "(" -f 1
//cp $(ldd RandomCheckTool | cut -d ">" -f 2 |cut -d "(" -f 1 ) lib/
//ll $(ldd RandomCheckTool | cut -d ">" -f 2 |cut -d "(" -f 1 )
//ll $(ldd RandomCheckTool | cut -d ">" -f 2 |cut -d "(" -f 1 ) | awk '{print $11,$10,$9}' | cut -d '>' -f 2
